package user

import (
	"encoding/json"
	"github.com/KyriakosMilad/go-rest-cms/helpers"
	"github.com/KyriakosMilad/valdn"
	"net/http"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.HandleError(http.StatusMethodNotAllowed, "method not allowed", w)
		return
	}

	rules := valdn.Rules{
		"firstName": {"required", "kind:string", "minLen:3", "maxLen:13"},
		"lastName":  {"required", "kind:string", "minLen:3", "maxLen:13"},
		"email":     {"required", "email", "maxLen:254"},
		"password":  {"required", "kind:string", "minLen:6", "maxLen:64"},
	}

	errors := valdn.ValidateRequest(r, rules)
	if len(errors) > 0 {
		helpers.HandleValidationErrors(errors, w)
		return
	}

	u := User{}

	// parse request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		helpers.HandleError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	password, err := helpers.HashPassword(u.Password)
	if err != nil {
		helpers.HandleError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	u.Password = password

	err = u.Create()
	if err != nil {
		helpers.HandleError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "user created successfully with id: " + helpers.ToString(u.ID)})
}
