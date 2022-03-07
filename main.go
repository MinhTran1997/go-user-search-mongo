package main

import (
	"context"
	"fmt"
	"github.com/core-go/config"
	sv "github.com/core-go/service"
	"github.com/gorilla/mux"
	"net/http"

	"search-users/internal/app"
)

func main() {
	var conf app.Root
	er1 := config.Load(&conf, "configs/config")
	if er1 != nil {
		panic(er1)
	}

	r := mux.NewRouter()

	er2 := app.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}

	fmt.Println(sv.ServerInfo(conf.Server))
	er3 := http.ListenAndServe(sv.Addr(conf.Server.Port), r)
	if er3 != nil {
		fmt.Println(er3.Error())
	}
}