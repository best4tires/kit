package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/best4tires/kit/env"
	"github.com/best4tires/kit/errs"
	"github.com/best4tires/kit/httpsrv"
	"github.com/best4tires/kit/log"
	"github.com/best4tires/kit/maps"
	"github.com/best4tires/kit/svc"
)

func main() {
	svc.NewRuntimeEnvironment("fooapi").Run(NewService(10))
}

// Model
type Foo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Bars int    `json:"bars"`
}

type Repository struct {
	foos map[string]Foo
}

func NewRepository(count int) *Repository {
	r := &Repository{
		foos: make(map[string]Foo),
	}
	for i := 0; i < count; i++ {
		id := fmt.Sprintf("%03d", i+1)
		r.foos[id] = Foo{
			ID:   id,
			Name: fmt.Sprintf("Foo#%d", i),
			Bars: i*2 + 10,
		}
	}
	return r
}

func (r *Repository) Close() {
	r.foos = nil
}

func (r *Repository) FindAll() ([]Foo, error) {
	return maps.OrderedValues(r.foos, func(v1, v2 Foo) bool {
		return v1.ID < v2.ID
	}), nil
}

func (r *Repository) Find(id string) (Foo, error) {
	if id == "003" {
		return Foo{}, fmt.Errorf("you can't pass")
	}

	if v, ok := r.foos[id]; ok {
		return v, nil
	}
	return Foo{}, errs.NotFound()
}

// Service definition
type Service struct {
	repo *Repository
}

func NewService(fooCount int) *Service {
	return &Service{
		repo: NewRepository(fooCount),
	}
}

func (s *Service) Route(router *httpsrv.PrefixRouter) {
	router.GET("foos/", s.handleGETFoos)
	router.GET("foos/{id}", s.handleGETFoo)
}

func (s *Service) RunCtx(ctx context.Context, env env.Env) error {
	return nil
}

func (s *Service) Shutdown() {
	s.repo.Close()
}

// handler
func (s *Service) handleGETFoos(w http.ResponseWriter, r *http.Request) {
	foos, err := s.repo.FindAll()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	httpsrv.WriteJSON(w, http.StatusOK, foos)
}

func (s *Service) handleGETFoo(w http.ResponseWriter, r *http.Request) {
	id := httpsrv.Var(r, "id")
	v, err := s.repo.Find(id)
	switch {
	case err == nil:
		httpsrv.WriteJSON(w, http.StatusOK, v)
	case errors.Is(err, errs.NotFound()):
		http.Error(w, fmt.Sprintf("no such foo %q", id), http.StatusNotFound)
	default:
		log.Errorf("find %q: %v", id, err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
