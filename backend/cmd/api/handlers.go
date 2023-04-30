package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	var payload = struct {
		Status      string   `json:"status"`
		Message     string   `json:"message"`
		Version     string   `json:"version"`
		Collections []string `josn:"collections"`
	}{
		Status:  "active",
		Message: "Welcome to subwatch!",
		Version: "1.0.0",
	}

	db, err := app.connectToDB()
	if err != nil {
		log.Println(err)
	}

	// get all collections in you database
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
	} else {

		for _, coll := range collections {
			payload.Collections = append(payload.Collections, coll)
		}
		_ = app.writeJSON(w, http.StatusOK, payload)
	}

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	// app.Database is alredy initalized to mongodb connection
	movies, err := app.Database.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	movieId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(movieId)
	if err != nil {
		app.errorJSON(w, err)
	}
	oneMovie, err := app.Database.GetOneMovie(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, oneMovie)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload submitted by the end user
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.Database.GetUserByEmail(requestPayload.Email)
	// user, err := app.Database.GetUserByEmail("anubhav11697@gmail.com")
	if err != nil {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password (match the hash)
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.Token)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)

}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	// range through all the cookie received

	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// now parse the token to get the claims

			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("Unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the user id form the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("Unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.Database.GetUserById(userID)
			if err != nil {
				app.errorJSON(w, errors.New("Unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			// generate new token pairs

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("errror generating tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

func (app *application) MovieCatalog(w http.ResponseWriter, r *http.Request) {
	// app.Database is alredy initalized to mongodb connection
	movies, err := app.Database.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies)
}
