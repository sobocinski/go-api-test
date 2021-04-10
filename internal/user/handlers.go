package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service *Service
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	user, err := h.Service.GetUser(UserIdFromString(vars["id"]))
	if err != nil {
		if err == pg.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)

		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	users, err := h.Service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	user := User{
		UpdatedAt: time.Now(),
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}
	newUser, err := h.Service.RegisterUser(user)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		panic(err)
	}
}
