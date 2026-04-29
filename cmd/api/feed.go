package main

import (
	"net/http"

	"github.com/Ay-afk-stack/gopher-socials/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	fs := store.PaginationFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fs.Parse(r)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := validate.Struct(fq); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()

	feeds, err := app.store.Posts.GetFeed(ctx, int64(2), &fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feeds); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}