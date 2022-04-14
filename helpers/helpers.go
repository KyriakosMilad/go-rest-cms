package helpers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type response struct {
	message string
	code    int
	errors  map[string]string
	success bool
}

func ToString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HandleError(code int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{
		message: message,
		code:    code,
		success: false,
	})
}

func HandleValidationErrors(errors map[string]string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(response{
		message: "inputs are not valid!",
		errors:  errors,
		code:    http.StatusUnprocessableEntity,
		success: false,
	})
}

// HandleMany not ready!!
func HandleMany(originalHandler http.Handler, middles ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middles) - 1; i > -1; i-- {
		originalHandler = middles[i](originalHandler)
	}
	return originalHandler
}
