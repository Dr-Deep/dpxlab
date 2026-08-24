package www

import (
	"net/http"

	"github.com/Dr-Deep/dpxlab/dpxlab/www/handlers"
)

type WebSrv struct {
	routes map[string]http.HandlerFunc
	mux    *http.ServeMux
}

func NewHttpListener() *WebSrv {
	www := &WebSrv{
		routes: map[string]http.HandlerFunc{
			"/":       handlers.NewIndexHandler().Handle,
			"/assets": handlers.NewAssetsHandler().Handle,
		},
		mux: http.NewServeMux(),
	}

	return www
}

func (www *WebSrv) RegisterHandler(route string, h WebHandler) {
	www.routes[route] = h.Handle
}

func (www *WebSrv) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if h, oke := www.routes[r.URL.Path]; oke {
		h(w, r)
		return
	}

	handlers.Error(w, http.StatusNotFound)
}

func (www *WebSrv) Start(socket string) error {
	www.mux.HandleFunc("/", www.HandleRequest)
	return http.ListenAndServe(socket, www.mux) //!
}

func (www *WebSrv) Stop() {

}
