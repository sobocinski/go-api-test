package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sobocinski/go-api-test/db"
	"github.com/sobocinski/go-api-test/internal/user"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.Println("Starting app...")

	ctx := context.Background()
	db, _ := db.NewDatabase(ctx)
	defer db.Close()

	router := mux.NewRouter()
	userRepository := user.UserRepositoryDb(db)
	userService := user.NewService(userRepository)
	userHanlder := user.NewHandler(userService)

	router.HandleFunc("/api/user", userHanlder.Register).Methods("POST")
	router.HandleFunc("/api/user", userHanlder.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", userHanlder.GetUser).Methods("GET")

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(struct{ Message string }{Message: "healthy"}); err != nil {
			panic(err)
		}
	})

	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(loggingMiddleware)
	log.Println("Listean at :8080")
	http.ListenAndServe(":8080", router)
	return nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, mux.Vars(r))
		next.ServeHTTP(w, r)
	})
}

func main() {
	app := App{"API", "0.1"}
	app.Run()
}
