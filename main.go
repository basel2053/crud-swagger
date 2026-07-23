package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/basel2053/crud-swagger/api"
)

type Server struct {
	port int
	mux  *http.ServeMux
}

var _ api.ServerInterface = (*Server)(nil)

var serverWrapper Server = Server{
	port: 3000,
	mux:  http.DefaultServeMux,
}

const swaggerUIHTML = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>User API Docs</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.onload = function () {
        SwaggerUIBundle({
          url: "/openapi.yaml",
          dom_id: "#swagger-ui"
        });
      };
    </script>
  </body>
</html>`

func (sw *Server) GetV1GetUserByID(w http.ResponseWriter, r *http.Request, id string) {
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		fmt.Errorf("Error while converting id to number")
	}
	user := api.User{
		Id:   idNumber,
		Name: "Bassel",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	server := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(swaggerUIHTML))
	})
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, "./openapi.yaml")
	})
	handler := api.HandlerFromMux(server, mux)
	fmt.Println("Server is up and running on port 3000")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatal("Couldn't start the server")
	}
}
