package handlers

import "net/http"

/*
# dpxlab.de nginx index
* Poudriere build status
* pkg repos (Cloudflare buckets)
* radicle http mirror
* The Heart
* Impressum (about me)

*/

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	// write index.html
	buf, err := assets.ReadFile("index.html")
	if err != nil {
		//! log
		Error(w, http.StatusInternalServerError)
		return
	}

	w.Write(buf)
}
