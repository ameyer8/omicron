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

import "github.com/gorilla/mux"

//Server is the base data type for this module
type Server struct {
	Router *mux.Router
	Port   int
}

func (s *Server) routes() {

	s.Router.Handle("/", s.Router)

}

//Start server functionality
func (s *Server) Start() {

}

//TurnDownServer preps for shutdown
func (s *Server) TurnDownServer() {

}
