package register

import (
	"distributed"
	"encoding/json"
	"log"
	"net/http"
)

type Registration struct {
	ServiceName string `json:"service_name"`
	ServiceUrl string `json:"service_url"`
}

var ServiceRegistrations []Registration

func NewRegistrationService() distributed.Service {
	return distributed.Service{
		Name:                "Registration Service",
		Host:                "",
		Port:                "5000",
		HttpHandleFunctions: HandleFunctions,
	}
}

func add(reg Registration) {
	ServiceRegistrations = append(ServiceRegistrations, reg)
	log.Printf("%v Registered", reg.ServiceName)
}

func remove(serviceUrl string) {
	for i, reg := range ServiceRegistrations {
		if reg.ServiceUrl == serviceUrl {
			ServiceRegistrations = append(ServiceRegistrations[:i], ServiceRegistrations[i+1:]...)
			log.Printf("%v Unregistered", reg.ServiceName)
			return
		}
	}
	log.Printf("Service not found with URL %v", serviceUrl)
}

func HandleFunctions() {
	http.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		// POST for registering service
		var reg Registration
		err := json.NewDecoder(req.Body).Decode(&reg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		switch req.Method {
		case http.MethodPost:
			add(reg)
		case http.MethodDelete:
			remove(reg.ServiceUrl)
		}
	})
}
