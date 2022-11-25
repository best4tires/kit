package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/best4tires/kit/env"
	"github.com/best4tires/kit/httpsrv"
	"github.com/best4tires/kit/log"
	"github.com/best4tires/kit/svc"
)

func main() {
	svc.NewRuntimeEnvironment("bartask").Run(NewService())
}

// Service definition
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Route(router *httpsrv.PrefixRouter) {}

func (s *Service) Shutdown() {}

func (s *Service) RunCtx(ctx context.Context, env env.Env) error {
	timer := time.NewTimer(0)
	for {
		select {
		case <-ctx.Done():
			log.Infof("context is done ... exit loop")
			return nil
		case <-timer.C:
			s.executeCtx(ctx)
			timer.Reset(2 * time.Second)
		}
	}
	return nil
}

func (s *Service) executeCtx(ctx context.Context) {
	dur := time.Duration(1000+rand.Intn(2000)) * time.Millisecond
	log.Infof("executing for %s ...", dur)
	select {
	case <-time.After(dur):
	case <-ctx.Done():
	}
	log.Infof("executing for %s ... done", dur)
}
