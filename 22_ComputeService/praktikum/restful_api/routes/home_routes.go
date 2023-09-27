package routes

import (
	"net/http"
	"path"
)

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path.Join(".", "index.html"))
}
