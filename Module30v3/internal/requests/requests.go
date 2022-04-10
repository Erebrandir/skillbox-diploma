package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"skillbox/internal/friends"
	"skillbox/internal/update"
	"skillbox/internal/user"
	"strconv"
)

const UserStorageFile = "userStorage.json"

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var makeFriend friends.SelectUser
	makeFriend.AddFriends(user.Db, w, *r)
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(user.Db)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u user.User
	_ = json.NewDecoder(r.Body).Decode(&u)

	_, err := strconv.Atoi(u.Age)
	if err != nil {
		_, err := w.Write([]byte("Возраст должен быть типа int"))
		if err != nil {
			return
		}
		return
	}
	u.Id = u.AddUserId(user.Db)
	user.Db = user.Storage.UpdateStorage(user.Db, u)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u.Id)
}

func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	user.Storage.GetFriends(user.Db, params, w)
}

func UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateAge update.Update
	(*update.Update).UpdateAge(&updateAge, user.Db, w, *r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var del friends.SelectUser
	del.DeleteUser(&user.Db, w, *r)

}
