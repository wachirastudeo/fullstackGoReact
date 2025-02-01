package main

import (
	"backend/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w, "Hello World from ", app.Domain)
	// json data
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go went Gone",
		Version: "1.0.0",
	}
	// out, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(out)
	_ = app.writeJSON(w, http.StatusOK, payload)

}
func (app *application) AllMovie(w http.ResponseWriter, r *http.Request) {
	movie, err := app.DB.AllMovies()
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movie)
}
func (app *application) About(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "About")
}

// Mockup data
func (app *application) AllMoviedemo(w http.ResponseWriter, r *http.Request) {

	rd, _ := time.Parse("2006-01-02", "2022-01-01")
	var movie []models.Movie

	handler := models.Movie{
		ID:          1,
		Title:       "handler",
		ReleaseDate: rd,
		Runtime:     200,
		MPAARating:  "PG-13",
		Description: "some description",
		Image:       "some image", Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	movie = append(movie, handler)
	rotla := models.Movie{
		ID:          2,
		Title:       "rotla",
		ReleaseDate: rd,
		Runtime:     100,
		MPAARating:  "PG-13",
		Description: "some description rotla",
		Image:       "some image", Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	movie = append(movie, rotla)

	out, err := json.Marshal(movie)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//read json payload (อ่านข้อมูลจาก json ที่ส่งมา)
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	// varidate user against database (ตรวจสอบข้อมูลผู้ใช้งาน)
	user, err := app.DB.GetUserByEmail(requestPayload.Email)

	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// 	check password against hash  (ตรวจสอบรหัสผ่าน)

	password, err := user.PasswordMatches(requestPayload.Password)
	if err != nil {
		app.errorJSON(w, errors.New("invalid password"), http.StatusBadRequest)
		return
	}
	if !password {
		app.errorJSON(w, errors.New("invalid password"), http.StatusBadRequest)
		return
	}

	// create a jwt user (สร้าง jwt user)

	u := jwtUser{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
	}
	//generate a token (สร้างโทเคน)
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		_ = app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, tokens)
}
