package internal

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){

		t := template.Must(template.ParseFiles("templates/search.html"))

		if err := t.ExecuteTemplate(w, "search.html", ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
