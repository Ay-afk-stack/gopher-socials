package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}




func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	if err := ReadJSON(w,r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email: payload.Email,
	}

	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(r.Context(), user, hashToken, app.config.mail.exp); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}