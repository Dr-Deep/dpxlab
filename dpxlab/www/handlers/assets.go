package handlers

import (
	"embed"
	"net/http"
)

//go:embed assets
var assets embed.FS

type AssetsHandler struct{}

func NewAssetsHandler() *AssetsHandler {
	return &AssetsHandler{}
}

func (h *AssetsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Error(w, http.StatusMethodNotAllowed)
		return
	}

	// valid path?
	buf, err := assets.ReadFile(r.URL.Path)
	if err != nil {
		//! log
		Error(w, http.StatusNotFound)
		return
	}

	w.Write(buf)
}
