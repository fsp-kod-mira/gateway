package controllers

import (
	"context"
	"fmt"
	"gateway/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type TemplatesService interface {
	GetAll(ctx context.Context) ([]*entity.TemplateStruct, error)
	GetFeatureByTemplate(ctx context.Context, id uint64) ([]*entity.FeatureStruct, error)
}

type TemplatesController struct {
	service TemplatesService
}

func NewTemplatesController(service TemplatesService) *TemplatesController {
	return &TemplatesController{
		service: service,
	}
}

func (h TemplatesController) GetAll() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		res, err := h.service.GetAll(ctx.Context())
		if err != nil {
			return internal(fmt.Sprintf("Unexpected error: %s", err.Error()))
		}

		return ok(ctx, res)
	}
}

func (h TemplatesController) GetFeaturesByTemplate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		templateId := ctx.QueryInt("templateId", 0)

		res, err := h.service.GetFeatureByTemplate(ctx.Context(), uint64(templateId))
		if err != nil {
			return internal(fmt.Sprintf("Unexpected error: %s", err.Error()))
		}

		return ok(ctx, res)
	}
}
