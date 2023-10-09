package util

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return LogError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return LogError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

func LogErrorApp(app string, err error) error {
	if err != nil {
		log.Printf("%v: %v", app, err)
	}
	return err
}
