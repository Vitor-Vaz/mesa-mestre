package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {

		r.Post("/owner", CreateOwnerHandler)

	})
}
func CreateOwnerHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic to create an owner

	fmt.Println("to criando um dono")
}
