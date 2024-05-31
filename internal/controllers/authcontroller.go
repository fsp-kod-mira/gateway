package controllers

import (
	"context"
	"errors"
	"gateway/api/auth"
	"gateway/internal/entity"
	"gateway/internal/entity/dto"
	"gateway/internal/service/authservice"
	"gateway/pkg/middleware"
	"log/slog"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	SignIn(ctx context.Context, credentials *entity.UserCredentials) (*entity.Tokens, error)
	SignUp(ctx context.Context, dto *dto.RegisterUser) (*entity.Tokens, error)
	SignOut(ctx context.Context, accessToken string) error
	Authenticate(ctx context.Context, accessToken string, role []auth.Role) (*entity.UserClaims, error)
	Refresh(ctx context.Context, refreshToken string) (*entity.Tokens, error)
}

type UserService interface {
	FindById(ctx context.Context, id string) (*entity.User, error)
}

type AuthController struct {
	service     AuthService
	userService UserService
	validator   *validator.Validate
}

func NewAuthController(service AuthService, userService UserService) *AuthController {
	return &AuthController{
		service:     service,
		validator:   validator.New(),
		userService: userService,
	}
}

func (a *AuthController) SignIn() fiber.Handler {
	type request struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	return func(ctx *fiber.Ctx) error {
		var req request

		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "signIn")

		if err := ctx.BodyParser(&req); err != nil {
			logger.Error("failed to parse request", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		if err := a.validator.Struct(req); err != nil {
			logger.Debug("failed to validate request", slog.String("err", err.Error()))
			return bad(err.Error())
		}

		credentials := &entity.UserCredentials{
			Email:    req.Email,
			Password: req.Password,
		}

		tokens, err := a.service.SignIn(ctx.Context(), credentials)
		if err != nil {
			logger.Error("failed to sign in", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) SignUp() fiber.Handler {
	type request struct {
		Email      string `json:"email" validate:"required"`
		Password   string `json:"password" validate:"required"`
		LastName   string `json:"lastName" validate:"required"`
		FirstName  string `json:"firstName" validate:"required"`
		MiddleName string `json:"middleName" validate:"required"`
		Role       string `json:"role" validate:"required"`
	}

	return func(ctx *fiber.Ctx) error {
		var req request

		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "signUp")

		if err := ctx.BodyParser(&req); err != nil {
			logger.Error("failed to parse request", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		if err := a.validator.Struct(req); err != nil {
			logger.Debug("failed to validate request", slog.String("err", err.Error()))
			return bad(err.Error())
		}

		logger.Info("sign up user", slog.Any("request", req))

		u := &dto.RegisterUser{
			User: entity.User{
				Email:      req.Email,
				LastName:   req.LastName,
				FirstName:  req.FirstName,
				MiddleName: req.MiddleName,
				Role:       entity.Role(req.Role),
			},
			Password: req.Password,
		}

		logger.Info("sign up user", slog.Any("dto", u))

		tokens, err := a.service.SignUp(ctx.Context(), u)
		if err != nil {

			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.AlreadyExists:
					return bad(e.Message())
				}
			}

			logger.Error("failed to sign up", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) SignOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "signOut")

		accessToken, k := ctx.Locals("accessToken").(string)
		if !k {
			logger.Error("missing access token")
			return internal("missing access token")
		}

		if err := a.service.SignOut(ctx.Context(), accessToken); err != nil {
			logger.Error("failed to sign out", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		return ok(ctx)
	}
}

func (a *AuthController) Refresh() fiber.Handler {
	type req struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	return func(ctx *fiber.Ctx) error {
		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "refresh")

		var r req

		if err := ctx.BodyParser(&r); err != nil {
			logger.Error("failed to parse request", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		if err := a.validator.Struct(r); err != nil {
			logger.Debug("failed to validate request", slog.String("err", err.Error()))
			return bad(err.Error())
		}

		tokens, err := a.service.Refresh(ctx.Context(), r.RefreshToken)
		if err != nil {
			logger.Error("failed to refresh", slog.String("err", err.Error()))
			return internal(err.Error())
		}

		return ok(ctx, tokens)
	}
}

func (a *AuthController) AuthRequired(role ...auth.Role) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "authRequired")

		authorization := ctx.Get("Authorization")
		logger.Debug("authorization", slog.String("authorization", authorization))

		s := strings.Split(authorization, " ")
		if len(s) < 2 {
			logger.Debug("failed to parse authorization header")
			return bad("failed to parse authorization header")
		}

		accessToken := s[1]

		u, err := a.service.Authenticate(ctx.Context(), accessToken, role)
		if err != nil {

			if errors.Is(err, authservice.ErrUnauthorized) {
				return unauthorized(err.Error())
			}

			if errors.Is(err, authservice.ErrForbidden) {
				return forbidden(err.Error())
			}

			if errors.Is(err, authservice.ErrNotFound) {
				return unauthorized(err.Error())
			}

			if errors.Is(err, authservice.ErrInvalidRequest) {
				return bad(err.Error())
			}

			logger.Error("failed to authenticate", slog.String("err", err.Error()))

			return internal(err.Error())
		}

		ctx.Locals("accessToken", accessToken)
		ctx.Locals("user", u)

		return ctx.Next()
	}
}

func (a *AuthController) Profile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := ctx.Context().Value(middleware.LOGGER).(*slog.Logger).With("controller", "auth").With("method", "profile")

		u, k := ctx.Locals("user").(*entity.UserClaims)
		if !k {
			logger.Error("missing user")
			return internal("missing user")
		}

		user, err := a.userService.FindById(ctx.Context(), u.Id)
		if err != nil {
			return err
		}

		return ok(ctx, user)
	}
}
