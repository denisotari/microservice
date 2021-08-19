package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	"gitlab.com/faemproject/backend/faem/pkg/logs"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

type GreeterPublisher interface {
	UserCreated(*proto.User) error
}

func (h *Handler) Hello(ctx context.Context, in *proto.User) (*proto.HelloOut, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	in.Name = strings.Title(in.Name)
	logs.LoggerForContext(ctx).WithFields(logrus.Fields{
		"name": in.Name,
	}).Info("hello called") // just a logging sample

	// Publish a message to the broker
	go func() {
		if err := h.Pub.UserCreated(in); err != nil {
			logs.Eloger.Error(err) // just a sample error handling
		}
	}()
	// Return the result if all is ok
	return &proto.HelloOut{Message: fmt.Sprintf("Hello, %s", in.Name)}, nil
}
