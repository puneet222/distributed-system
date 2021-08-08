package register

import (
	"bytes"
	"distributed"
	"encoding/json"
	"fmt"
	"net/http"
)

var url = "http://localhost:5000/register"

func Add(service distributed.Service) error {
	buf, err := getBufferedData(service)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service as status recieved : %v", res.StatusCode)
	}
	return err
}

func Remove(service distributed.Service) error {
	buf, err := getBufferedData(service)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, url, buf)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service as status recieved : %v", res.StatusCode)
	}
	return err
}

func getBufferedData(service distributed.Service) (*bytes.Buffer, error) {
	data := Registration{
		ServiceName: service.Name,
		ServiceUrl:  service.Host + ":" + service.Port,
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(&data)
	return buf, err
}