package app

import (
	"context"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, ctx context.Context, root Root) error {
	app, err := NewApp(ctx, root)
	if err != nil {
		return err
	}

	userPath := "/users"
	r.HandleFunc(userPath, app.UserHandler.All).Methods("GET")
	r.HandleFunc(userPath+"/search", app.UserHandler.Search).Methods("POST")

	return nil
}
