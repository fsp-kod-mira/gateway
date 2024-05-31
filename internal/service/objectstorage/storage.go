package objectstorage

import (
	"gateway/api/objectstorage"
	"io"
	"math"

	"golang.org/x/net/context"
)

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

	chunkCount := int(math.Ceil(float64(size / (1 << 20))))

	for i := 0; i < chunkCount; i++ {
		chunk := make([]byte, 1<<20)

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
