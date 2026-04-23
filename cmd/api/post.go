package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Tags []string `json:"tags"`
}

func (app *application) createPostHandlers(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	userID := 1

	post := &store.Post{
		Title: payload.Title,
		Content: payload.Content,
		Tags: payload.Tags,
		UserID: int64(userID),
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err:= WriteJSON(w, http.StatusCreated, post); err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			WriteJSONError(w, http.StatusNotFound, err.Error())
		default:
			WriteJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err = WriteJSON(w, http.StatusOK, post); err != nil {
		WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}