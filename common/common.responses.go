package common

// import (
// 	"fmt"
// 	"net/http"
// )

// type Handler struct{}

// var NewUrl string = "/trade/gem/GEM"

// func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	uri := r.URL.Path
// 	if uri == "/" {
// 		http.Redirect(w, r, NewUrl, http.StatusSeeOther)
// 	}
// 	fmt.Fprintf(w, uri)
// 	return
// }

// func main() {
// 	handler := new(Handler)
// 	http.ListenAndServe(":9000", handler)
// }
