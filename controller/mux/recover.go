package mux

import (
	"GraphNeo4jGO/DTO"
	"log"
	"net/http"
)

func recoverHttp(w http.ResponseWriter) {
	if err := recover(); err != nil {
		switch e := err.(type) {
		case error:
			writeJson(w, http.StatusInternalServerError, DTO.Error{Status: DTO.StatusError, Err: e.Error()})
			//r.logger.Error(logger.LogData{Message: e.Error(), Section: "unknown"})
		case string:
			writeJson(w, http.StatusInternalServerError, DTO.Error{Status: DTO.StatusError, Err: e})
			//r.logger.Error(logger.LogData{Message: e, Section: "unknown"})
		case *Error:
			handle(*e, w)
		case Error:
			handle(e, w)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("err: %#v\n", e)
		}
	}
}

type LogLevel uint8

const (
	InfoLevel LogLevel = iota
	WarnLevel
	ErrorLevel
)

type Error struct {
	Err        error
	Section    string
	StatusCode int
	LogLevel   LogLevel
}

func handle(e Error, w http.ResponseWriter) {
	//switch e.LogLevel {
	//case InfoLevel:
	//	writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: e.Err.Error()})
	//case WarnLevel:
	//	writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: e.Err.Error()})
	//case ErrorLevel:
	//	writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: e.Err.Error()})
	//default:
	//	writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: e.Err.Error()})
	//}

	writeJson(w, e.StatusCode, DTO.Error{Status: DTO.StatusError, Err: e.Err.Error()})
}
