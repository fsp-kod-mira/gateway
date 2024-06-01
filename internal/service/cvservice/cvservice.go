package cvservice

import (
	"context"
	"gateway/api/cv"
	"gateway/internal/entity"
	"log/slog"

	lop "github.com/samber/lo/parallel"
)

// var _ (controllers.CvService) = (*Service)(nil)

type Service struct {
	client cv.CvServiceClient
	logger *slog.Logger
}

func New(client cv.CvServiceClient, logger *slog.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}

// Get implements controllers.CvService.
func (s *Service) Get(ctx context.Context, id string) (*entity.CV, error) {
	cv, err := s.client.GetOne(ctx, &cv.Id{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &entity.CV{
		Id:         cv.Id,
		Status:     cv.Status,
		FileId:     cv.FileId,
		UploadedBy: cv.UploadedById,
	}, nil
}

// GetAll implements controllers.CvService.
func (s *Service) GetAll(ctx context.Context, pagination *entity.Pagination) ([]*entity.CV, error) {
	cvs, err := s.client.GetAll(ctx, &cv.Paggination{
		Limit:  int32(pagination.Limit),
		Offset: int32(pagination.Offset),
	})
	if err != nil {
		// if e, ok := status.FromError(err); ok {
		// }
		return nil, err
	}

	cvcv := lop.Map(cvs.Cvs, func(cv *cv.CV, index int) *entity.CV {
		return &entity.CV{
			Id:         cv.Id,
			Status:     cv.Status,
			FileId:     cv.FileId,
			UploadedBy: cv.UploadedById,
		}
	})

	return cvcv, nil
}

// Upload implements controllers.CvService.
func (s *Service) Upload(ctx context.Context, info *entity.CVInfo) error {
	slog.Info("uploading cv", slog.Any("cv-info", info))

	_, err := s.client.Upload(ctx, &cv.UploadRequest{
		FileId:       info.FileId,
		UploadedById: info.UploadedBy,
	})
	if err != nil {
		return err
	}

	return nil
}
