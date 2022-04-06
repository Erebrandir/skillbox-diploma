package main

import (
	s "awesomeProject/Module30/pkg/service"
	"net/http"
)

func main() {
	srv := s.NewService()
	router := http.NewServeMux()

	router.HandleFunc("/get", srv.GetAllUsers)
	router.HandleFunc("/create", srv.Create)
	router.HandleFunc("/make_friends", srv.MakeFriends)
	router.HandleFunc("/user", srv.DeleteUser)
	router.HandleFunc("/age", srv.ChangeAge)
	router.HandleFunc("/friends", srv.GetFriends)

	http.Handle("/", router)
	http.ListenAndServe("localhost:8080", nil)
}
