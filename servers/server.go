package servers

import (
	"fmt"
	"net/http"
)

func Start(port string, id string) {
	//so that each 
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", id)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	fmt.Printf("%s running on :%s\n", id, port)

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Println("server error:", err)
	}

}
