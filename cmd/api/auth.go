package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/Ay-afk-stack/gopher-socials/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
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

	tokenExp, err := time.ParseDuration(app.config.mail.exp)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.store.Users.CreateAndInvite(r.Context(), user, hashToken, tokenExp); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	userWithToken := UserWithToken{
		User: user,
		Token: plainToken,
	}

	activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, plainToken)

	if err := app.mailer.Send(user.Username, user.Email, activationURL); err != nil {
		app.logger.Errorw("error sending welcome email", "error", err)

		if err := app.store.Users.Delete(r.Context(), user.ID); err != nil {
			app.logger.Errorw("error deleting user", "error", err)
			return
		}

		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}