package main

import (
	"fmt"
	"net/http"

	"github.com/best4tires/kit/req"
)

type Foo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Bars int    `json:"bars"`
}

func main() {
	clt := &http.Client{}
	foos, err := req.PostJSON[[]Foo](clt, "http://127.0.0.1:8080/api/fooapi/foos/", Foo{
		ID:   "new-id",
		Name: "new-name",
		Bars: 43,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(foos)
}
