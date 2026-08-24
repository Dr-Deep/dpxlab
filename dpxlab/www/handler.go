package www

import "net/http"

type WebHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
