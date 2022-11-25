package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/best4tires/kit/env"
	"github.com/best4tires/kit/errs"
	"github.com/best4tires/kit/log"
	"github.com/best4tires/kit/maps"
	"github.com/best4tires/kit/srv"
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
			Name: fmt.Sprintf("Foo#%d", i+1),
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

func (r *Repository) InsertOrUpdate(foo Foo) {
	r.foos[foo.ID] = foo
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

func (s *Service) Route(router *srv.PrefixRouter) {
	router.GET("foos/", s.handleGETFoos)
	router.POST("foos/", s.handlePOSTFoos)
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
	srv.WriteJSON(w, http.StatusOK, foos)
}

func (s *Service) handlePOSTFoos(w http.ResponseWriter, r *http.Request) {
	var foo Foo
	err := json.NewDecoder(r.Body).Decode(&foo)
	if err != nil {
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}
	s.repo.InsertOrUpdate(foo)
	foos, err := s.repo.FindAll()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	srv.WriteJSON(w, http.StatusOK, foos)
}

func (s *Service) handleGETFoo(w http.ResponseWriter, r *http.Request) {
	id := srv.Var(r, "id")
	v, err := s.repo.Find(id)
	switch {
	case err == nil:
		srv.WriteJSON(w, http.StatusOK, v)
	case errors.Is(err, errs.NotFound()):
		http.Error(w, fmt.Sprintf("no such foo %q", id), http.StatusNotFound)
	default:
		log.Errorf("find %q: %v", id, err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
