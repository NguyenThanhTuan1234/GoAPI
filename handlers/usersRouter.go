package handlers

import (
	"fmt"
	"net/http"
)

//UserRouter handles the users route
func UsersRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
