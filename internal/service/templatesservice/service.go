package templatesservice

import (
	"context"
	"gateway/api/templates"
	"gateway/internal/entity"
	"log/slog"

	lop "github.com/samber/lo/parallel"
)

// var _ (controllers.TemplatesService) = (*Service)(nil)

type Service struct {
	client templates.TemplatesClient
	logger *slog.Logger
}

func New(client templates.TemplatesClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}

func (s Service) GetAll(ctx context.Context) ([]*entity.TemplateStruct, error) {
	resp, err := s.client.GetAllTemplates(ctx, &templates.Empty{})
	if err != nil {
		return nil, err
	}

	templates := lop.Map(resp.Items, func(item *templates.TemplateStruct, index int) *entity.TemplateStruct {
		return &entity.TemplateStruct{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
		}
	})

	return templates, nil
}

func (s *Service) GetFeatureByTemplate(ctx context.Context, id uint64) ([]*entity.FeatureStruct, error) {
	resp, err := s.client.GetFeaturesByTemplateId(ctx, &templates.IdStruct{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	features := lop.Map(resp.Items, func(item *templates.FeatureLinkTemplate, index int) *entity.FeatureStruct {
		return &entity.FeatureStruct{
			Id:          item.Link.Id,
			FeatureId:   item.Feature.Id,
			TemplateId:  item.Link.TemplateId,
			Value:       item.Link.Value,
			Name:        item.Feature.Name,
			FeatureType: entity.FeatureType(item.Feature.FeatureType.String()),
		}
	})

	return features, nil
}
