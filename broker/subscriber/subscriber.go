package subscriber

import (
	"sync"
	"time"

	"gitlab.com/faemproject/backend/faem/pkg/logs"
	"gitlab.com/faemproject/backend/faem/pkg/rabbit"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/broker"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/handler"
)

type Subscriber struct {
	Rabbit  *rabbit.Rabbit
	Encoder broker.Encoder
	Handler *handler.Handler

	wg     sync.WaitGroup
	closed chan struct{}
}

func (s *Subscriber) Init() error {
	s.closed = make(chan struct{})

	// call all the initializers here, multierr package might be useful
	return s.initUserCreated()
}

func (s *Subscriber) Wait(shutdownTimeout time.Duration) {
	// try to shutdown the listener gracefully
	stoppedGracefully := make(chan struct{}, 1)
	go func() {
		// Notify subscribers about exit, wait for their work to be finished
		close(s.closed)
		s.wg.Wait()
		stoppedGracefully <- struct{}{}
	}()

	// wait for a graceful shutdown and then stop forcibly
	select {
	case <-stoppedGracefully:
		logs.Eloger.Info("subscriber stopped gracefully")
	case <-time.After(shutdownTimeout):
		logs.Eloger.Info("subscriber stopped forcibly")
	}
}
