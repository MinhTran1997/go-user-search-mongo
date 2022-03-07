package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler interface {
	All(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService	UserService
}

func NewUserHandler(service UserService) UserHandler {
	return	&userHandler{userService: service}
}

func (h *userHandler) All(w http.ResponseWriter, r *http.Request) {
	res, err := h.userService.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(w, res)
}

func (h *userHandler) Search(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := filterScan(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.userService.Search(r.Context(), bodyRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respond(w, res)
}

func filterScan(w http.ResponseWriter, r *http.Request) (*UserFilter, error) {
	var res UserFilter
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &res, err
	}
	return &res, nil
}

func respond(w http.ResponseWriter, result interface{}) {
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
