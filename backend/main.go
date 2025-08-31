package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

type application struct {
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}

	app.logger.Printf("Starting tennis match logger server on :8080")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Basic routes for tennis match logging
	router.HandlerFunc(http.MethodGet, "/", app.healthCheck)
	router.HandlerFunc(http.MethodGet, "/v1/health", app.healthCheck)

	// Tennis match routes (to be implemented)
	router.HandlerFunc(http.MethodGet, "/v1/matches", app.getMatches)
	router.HandlerFunc(http.MethodPost, "/v1/matches", app.createMatch)
	router.HandlerFunc(http.MethodGet, "/v1/matches/:id", app.getMatch)
	router.HandlerFunc(http.MethodPut, "/v1/matches/:id", app.updateMatch)

	// Player routes (to be implemented)
	router.HandlerFunc(http.MethodGet, "/v1/players", app.getPlayers)
	router.HandlerFunc(http.MethodPost, "/v1/players", app.createPlayer)

	standard := alice.New(app.recoverPanic, app.logRequest)

	return standard.Then(router)
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tennis Match Logger API is running")
}

// Placeholder handlers - to be implemented
func (app *application) getMatches(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get matches - to be implemented")
}

func (app *application) createMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create match - to be implemented")
}

func (app *application) getMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get match - to be implemented")
}

func (app *application) updateMatch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update match - to be implemented")
}

func (app *application) getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get players - to be implemented")
}

func (app *application) createPlayer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create player - to be implemented")
}

// Middleware
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				http.Error(w, "Internal Server Error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}