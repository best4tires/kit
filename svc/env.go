package svc

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/best4tires/kit/env"
	"github.com/best4tires/kit/httpsrv"
	"github.com/best4tires/kit/log"
)

const (
	envKeyHttpPort   = "http.port"
	envKeyHttpPrefix = "http.prefix"
)

type Service interface {
	Route(router *httpsrv.PrefixRouter)
	RunCtx(ctx context.Context, env env.Env) error
}

func NewRuntimeEnvironment(name string) *RuntimeEnvironment {
	e := &RuntimeEnvironment{
		name: name,
	}
	return e
}

type RuntimeEnvironment struct {
	name string
}

func (e *RuntimeEnvironment) Run(svcs ...Service) {
	err := e.run(svcs...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "run finished with error: %v\n", err)
		os.Exit(1)
	}
}

func (e *RuntimeEnvironment) run(svcs ...Service) (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = fmt.Errorf("runtime-env: recovered: %v", perr)
			log.Errorf("%v", err)
			log.DebugStack()
		}
	}()
	env := env.Load()

	//http params
	httpPort := env.StringWithTagOrDefault(envKeyHttpPort, e.name, "0")
	httpPrefix := env.StringWithTagOrDefault(envKeyHttpPrefix, e.name, fmt.Sprintf("/api/%s/", e.name))

	//router
	router := httpsrv.NewRouter().WithPrefix(httpPrefix)
	for _, svc := range svcs {
		svc.Route(router)
	}

	//server
	bind := fmt.Sprintf(":%s", httpPort)
	server, err := httpsrv.NewServer(bind)
	if err != nil {
		return fmt.Errorf("new-server on %q: %w", bind, err)
	}
	handler := router.Handler(
		httpsrv.GZIP(),
		httpsrv.Recovery(),
		httpsrv.Logging(),
	)
	go server.Run(handler)

	// run services
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	wg := sync.WaitGroup{}
	for _, svc := range svcs {
		wg.Add(1)
		go func(s Service) {
			defer wg.Done()
			s.RunCtx(ctx, env)
		}(svc)
	}

	// wait until done
	<-ctx.Done()

	// shutdown server and wait for services in parallel
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Shutdown(3 * time.Second)
	}()

	waitC := make(chan struct{})
	go func() {
		defer close(waitC)
		wg.Wait()
	}()
	select {
	case <-waitC:
	case <-time.After(5 * time.Second):
	}
	return nil
}
