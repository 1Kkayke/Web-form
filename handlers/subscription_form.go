package handlers

import (
	"MyPackages/controller"
	"fmt"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
		}

		err := controller.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "Error saving data: %v", err)
		}
	}
	http.ServeFile(w, r, "templates/form.html")

}
