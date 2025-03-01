package users

import (
	"GoRecordurbate/modules/db"
	"GoRecordurbate/modules/file"
	"GoRecordurbate/modules/handlers/cookies"
	"GoRecordurbate/modules/handlers/login"
	"GoRecordurbate/modules/handlers/status"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if !cookies.Session.IsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(db.Users.Users)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	if !cookies.Session.IsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	type RequestData struct {
		OldUsername string `json:"oldUsername"`
		NewUsername string `json:"newUsername"`
		NewPassword string `json:"newPassword"`
	}
	var reqData RequestData
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	modified := false
	for i, user := range db.Users.Users {
		if user.Name == reqData.OldUsername {
			db.Users.Users[i].Name = reqData.NewUsername
			db.Users.Users[i].Key = string(login.HashedPassword(reqData.NewPassword))
			modified = true
			break
		}
	}
	if modified {
		db.Update(file.Users_json_path, db.Users)
	}

	resp := status.Response{
		Message: "User modified!",
	}
	for _, u := range db.Users.Users {
		cookies.UserStore[u.Name] = u.Key
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if !cookies.Session.IsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	type RequestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var reqData RequestData
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	for _, user := range db.Users.Users {
		if user.Name == reqData.Username {
			resp := status.Response{
				Message: "User already exists!",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}

	db.Users.Users = append(db.Users.Users, db.Login{Name: reqData.Username, Key: string(login.HashedPassword(reqData.Password))})
	db.Update(file.Users_json_path, db.Users)

	resp := status.Response{
		Message: reqData.Username + " added!",
	}
	for _, u := range db.Users.Users {
		cookies.UserStore[u.Name] = u.Key
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
