package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"skillbox/internal/requests"
	"skillbox/internal/user"
)

func main() {
	rawData, _ := ioutil.ReadFile(requests.UserStorageFile)
	_ = json.Unmarshal(rawData, &user.Db)

	router := chi.NewRouter()
	router.MethodFunc("GET", "/users", requests.GetUsers)
	router.MethodFunc("POST", "/create", requests.CreateUser)
	router.MethodFunc("GET", "/friends/{id}", requests.GetUserFriends)
	router.MethodFunc("PUT", "/{id}", requests.UpdateUserAge)
	router.MethodFunc("DELETE", "/user", requests.DeleteUser)
	router.MethodFunc("POST", "/make_friends", requests.MakeFriends)

	log.Fatal(http.ListenAndServe(":8080", router))
}
