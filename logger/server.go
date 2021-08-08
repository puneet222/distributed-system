package logger

import (
	"distributed"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// this service will instantiate the logger with log file
// and create required http handlers to run

var _logger *log.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (n int, err error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	return f.Write(data)
}

func NewLoggerService(destination string) distributed.Service {
	// initializing logger
	_logger = log.New(fileLog(destination), "", log.LstdFlags)
	return distributed.Service{
		Name:                "Logger Service",
		Host:                "",
		Port:                "4000",
		HttpHandleFunctions: handleHttp,
	}
}

func handleHttp() {
	http.HandleFunc("/log", func(w http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		write(string(data))
	})
}

func write(msg string) {
	_logger.Printf("%v\n", msg)
}
