package servers

import (
	"fmt"
	"net/http"
)

func Start() {

	//when a http req comes to "/" call handler
	http.HandleFunc("/", handler)

	fmt.Println("Server running on :8081")

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("server error:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from backend server")
}
