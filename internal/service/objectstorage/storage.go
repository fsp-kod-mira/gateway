package objectstorage

import (
	"gateway/api/objectstorage"
	"io"
	"log/slog"
	"math"

	"golang.org/x/net/context"
)

// var _ (controllers.ObjectStorage) = (*ObjectStorage)(nil)

type ObjectStorage struct {
	client objectstorage.ObjectStorageClient
}

func New(client objectstorage.ObjectStorageClient) *ObjectStorage {
	return &ObjectStorage{
		client: client,
	}
}

func (s *ObjectStorage) Upload(ctx context.Context, filename string, reader io.Reader, size uint64) (string, error) {
	stream, err := s.client.Upload(ctx)
	if err != nil {
		return "", err
	}

	log := slog.With(slog.String("filename", filename), slog.Uint64("size", size))

	log.Info("uploading in object storage")

	window := 1 << 20

	chunkCount := int(math.Ceil(float64(size) / float64(window)))

	log.Info("calculate chunk count", slog.Int("chunkCount", chunkCount))

	for i := 0; i < chunkCount; i++ {
		chunkSize := math.Min(float64(window), float64(size))
		chunk := make([]byte, int(chunkSize))
		size -= uint64(chunkSize)

		_, err := reader.Read(chunk)
		if err != nil {
			return "", err
		}

		if err := stream.Send(&objectstorage.File{
			Chunk: chunk,
			Meta: &objectstorage.Meta{
				Filename: filename,
			},
		}); err != nil {
			return "", err
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		return "", err
	}

	return reply.FileName, nil
}
