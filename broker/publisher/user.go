package publisher

import (
	"github.com/pkg/errors"

	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

const (
	channelNameUserCreated = "userCreated"
)

func (p *Publisher) UserCreated(user *proto.User) error {
	userCreatedChannel, err := p.Rabbit.GetSender(channelNameUserCreated)
	if err != nil {
		return errors.Wrapf(err, "failed to get a sender channel")
	}
	return p.Publish(userCreatedChannel, "int.bootstrap.debug", "new", user) // TODO: use some constants or variables
}

func (p *Publisher) initUserCreated() error {
	userCreatedChannel, err := p.Rabbit.GetSender(channelNameUserCreated)
	if err != nil {
		return errors.Wrapf(err, "failed to get a sender channel")
	}

	err = userCreatedChannel.ExchangeDeclare(
		"int.bootstrap.debug", // name // TODO: use constant or variable
		"direct",              // type // TODO: change me
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	return errors.Wrap(err, "failed to create an exchange")
}
