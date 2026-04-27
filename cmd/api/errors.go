package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusInternalServerError, "something went wrong")
}

func (app *application) badRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error:%s path: %s error:%s", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictsError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error:%s path: %s error:%s", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusConflict, "already exists")
}

func (app *application) notFoundError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusNotFound, "not found")
}