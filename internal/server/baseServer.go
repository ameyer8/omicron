/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	onion "github.com/ameyer8/omicron/external/omega2gpio"

	"github.com/gorilla/mux"
)

//Server is the base data type for this module
type Server struct {
	router *mux.Router
	Port   int
}

func (s *Server) routes() {
	s.router.HandleFunc("/api/v1/healthz", healthHandler())

	s.gpioRoutes(s.router.PathPrefix("/api/v1/gpio").Subrouter())

	http.Handle("/", s.router)
}

func healthHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("health"))
	}
}

//Start server functionality
func (s *Server) Start() {

	s.router = mux.NewRouter()
	s.routes()
	addr := fmt.Sprintf(":%d", s.Port)
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	onion.Setup()
	log.Fatal(srv.ListenAndServe())
}

//TurnDownServer preps for shutdown
func (s *Server) TurnDownServer() {

}
