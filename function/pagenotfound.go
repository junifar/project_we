package function

import "net/http"

// PageNotFound is function to handle if page not found
func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("the page are you looking for does not exist"))
}
