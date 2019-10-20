
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// User is a github user information
type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	// HTTP call
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode JSON
	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)


	//github
	// user, err := userInfo("tebeka")
	// if err != nil {
	// 	log.Fatalf("error: %s", err)
	// }

	// fmt.Printf("%+v\n", user)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
// [END all]
