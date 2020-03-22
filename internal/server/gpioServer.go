package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ameyer8/omicron/internal/data"
	"github.com/ameyer8/omicron/internal/model"
	"github.com/gorilla/mux"
)

func (s *Server) gpioRoutes(r *mux.Router) {

	r.HandleFunc("/healthz", gpioHealthHandler())
	r.HandleFunc("/{portId}/direction", gpioGetDirectionHandler()).Methods("GET")
	r.HandleFunc("/{portId}/direction", gpioSetDirectionHandler()).Methods("POST")
	r.HandleFunc("/{portId}/value", gpioGetValueHandler()).Methods("GET")
	r.HandleFunc("/{portId}/value", gpioSetValueHandler()).Methods("POST")
	r.HandleFunc("/{portId:[0-9]+(?:\\/)?}", gpioGetStatusHandler()).Methods("GET")
	r.HandleFunc("/{portId}", gpioSetStatusHandler()).Methods("POST")

}

func gpioHealthHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))

	}
}

func gpioGetDirectionHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		portID, _ := strconv.Atoi(mux.Vars(r)["portId"])

		dirRet := data.DirectionReturn{
			Direction: model.GetGPIODirection(portID),
		}
		jsonData, _ := json.Marshal(dirRet)
		w.Write(jsonData)
	}
}

func gpioSetDirectionHandler() http.HandlerFunc {

	type directionReturn struct {
		Direction string `json:"direction"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		portID, _ := strconv.Atoi(mux.Vars(r)["portId"])

		var dirRet directionReturn
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("[gpioSetDirectionHandler] Could not read from Request body")
		}
		err = json.Unmarshal(data, &dirRet)
		if err != nil {
			log.Println("[gpioSetDirectionHandler] Could not Unmarshall data")
		}
		err = model.SetGPIODirection(portID, dirRet.Direction)
		if err != nil {
			log.Println("[gpioSetDirectionHandler] Could not Set Port Direction")
		}

		dirRet.Direction = model.GetGPIODirection(portID)
		jsonData, _ := json.Marshal(dirRet)
		w.Write(jsonData)
	}
}

func gpioGetValueHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		portID, _ := strconv.Atoi(mux.Vars(r)["portId"])

		valueRet := data.ValueReturn{
			Value: model.GetGPIOValue(portID),
		}
		jsonData, _ := json.Marshal(valueRet)
		w.Write(jsonData)
	}
}
func gpioSetValueHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		portID, _ := strconv.Atoi(mux.Vars(r)["portId"])
		query := r.URL.Query()

		var valueRet data.ValueReturn

		if query.Get("toggle") == "true" {
			err := model.ToggleGPIOValue(portID)
			if err != nil {
				log.Println("[gpioSetDirectionHandler] Could not Set Port Value")
				w.WriteHeader(400)
				w.Write([]byte("Could not set port value"))
			}
		} else {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("[gpioSetValueHandler] Could not read from Request body")
			}
			err = json.Unmarshal(data, &valueRet)
			if err != nil {
				log.Println("[gpioSetDirectionHandler] Could not Unmarshall data")
			}

			err = model.SetGPIOValue(portID, valueRet.Value)
			if err != nil {
				log.Println("[gpioSetDirectionHandler] Could not Set Port Value")
				w.WriteHeader(400)
				w.Write([]byte("Could not set port value"))
			}

		}

		valueRet.Value = model.GetGPIOValue(portID)
		jsonData, _ := json.Marshal(valueRet)
		w.Write(jsonData)

	}
}

func gpioGetStatusHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		portID, _ := strconv.Atoi(mux.Vars(r)["portId"])

		statusRet := data.StatusReturn{
			Value:     model.GetGPIOValue(portID),
			Direction: model.GetGPIODirection(portID),
		}
		jsonData, _ := json.Marshal(statusRet)
		w.Write(jsonData)

	}
}
func gpioSetStatusHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))

	}
}
