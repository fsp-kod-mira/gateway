package controllers

import (
	"context"
	"fmt"
	"gateway/internal/entity"
	"io"

	"github.com/gofiber/fiber/v2"
)

type CvService interface {
	Upload(ctx context.Context, fileInfo *entity.CVInfo) error
	Get(ctx context.Context, id string) (*entity.CV, error)
	GetAll(context.Context, *entity.Pagination) ([]*entity.CV, error)
}

type CvController struct {
	service  CvService
	uploader ObjectStorage
}

type ObjectStorage interface {
	Upload(ctx context.Context, filename string, reader io.Reader, size uint64) (string, error)
}

func NewCvController(service CvService, uploader ObjectStorage) *CvController {
	return &CvController{
		service:  service,
		uploader: uploader,
	}
}

func (c *CvController) Upload() fiber.Handler {
	type response struct {
		Filename string `json:"filename"`
	}

	return func(ctx *fiber.Ctx) error {
		f, err := ctx.FormFile("cv")
		if err != nil {
			return internal(err.Error())
		}

		u := ctx.Locals("user").(*entity.UserClaims)

		reader, err := f.Open()
		if err != nil {
			return internal(err.Error())
		}

		filename, err := c.uploader.Upload(ctx.Context(), f.Filename, reader, uint64(f.Size))
		if err != nil {
			return internal(err.Error())
		}

		c.service.Upload(ctx.Context(), &entity.CVInfo{
			FileId:     filename,
			UploadedBy: u.Id,
		})

		return ok(ctx, &response{
			Filename: filename,
		})
	}
}

func (c *CvController) Get() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id", "")
		if id == "" {
			return bad("missing id")
		}

		cv, err := c.service.Get(ctx.Context(), id)
		if err != nil {
			return internal(fmt.Sprintf("Unexpected error: %s", err.Error()))
		}

		return ok(ctx, cv)
	}
}

func (cc *CvController) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 0)
		limit := c.QueryInt("limit", 0)

		if page == 0 {
			return bad("page is missing or equal zero")
		}

		if limit == 0 {
			return bad("limit is missing or equal zero")
		}

		cvs, err := cc.service.GetAll(c.Context(), &entity.Pagination{
			Limit:  limit,
			Offset: (page - 1) * limit,
		})
		if err != nil {
			return internal(fmt.Sprintf("Unexpected error: %s", err.Error()))
		}

		return ok(c, cvs)
	}
}
