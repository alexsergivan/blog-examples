package main

import (
	"math/rand"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", homePageHandler)
	r.Mount("/admin", adminRouter())
	http.ListenAndServe(":3000", r)
}

func adminRouter() http.Handler {
	r := chi.NewRouter()
	//Middleware with access rules for router.
	r.Use(AdminOnly)
	r.Get("/", adminPageHandler)

	return r
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // If user is admin, allows access.
		if IsLoggedInAdmin(r) {
			next.ServeHTTP(w, r)
		} else {
			// Otherwise, 403.
			http.Error(w, http.StatusText(403), 403)
			return
		}

		return
	})
}

func IsLoggedInAdmin(r *http.Request) bool {
	return rand.Float32() < 0.5
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}

func adminPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is admin page"))
}
