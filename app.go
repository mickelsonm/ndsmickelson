package ndsmickelson

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/data", data)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, HomeTemplate)
}

func data(w http.ResponseWriter, r *http.Request) {
	peeps := GetTestPeople()
	js, err := peeps.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(js))
}
