package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	u "skillbox/Module30/pkg/User"
	"strconv"
)

type request struct {
	TargetID int `json:"target_id"`
	SourceID int `json:"source_id"`
	Age      int `json:"new age"`
}

type service struct {
	store map[int]*u.User
}

func NewService() *service {
	return &service{make(map[int]*u.User)}
}

func (s *service) Contains(u *u.User) bool {
	for _, i := range s.store {
		if i == u {
			return true
		}
	}
	return false
}

func (s *service) newId() int {
	var id int
	for s.store[id+1] != nil {
		id += 1
	}
	return id + 1
}

func (s *service) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for id, user := range s.store {
		w.Write([]byte("id: " + strconv.Itoa(int(id)) + "\nПользователь: " + user.ToString() + "\n"))
	}
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Не удается прочитать данные из запроса")
		w.Write([]byte(err.Error()))
		return
	}

	tmpUser := u.NewUser("", 0)

	if err := json.Unmarshal(content, &tmpUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Error("Не удается обработать данные из json")
		return
	}

	id := s.newId()
	s.store[id] = &tmpUser

	func(u []*u.User) {
		for _, man := range u {
			if !s.Contains(man) {
				newId := s.newId()
				s.store[newId] = man
			}
		}
	}(s.store[id].GetFriends())

	log.Info("Новый пользователь: ", id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("\nПользователь " + tmpUser.Name + " был создан\nid:" + strconv.Itoa(id) + "\n"))
}

func (s *service) ChangeAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req request
	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Error("Неправильный запрос на изменение возраста")
		return
	}

	if err := json.Unmarshal(content, &req); err != nil {
		log.Error("Не удается обработать запрос на изменение возраста")
		return
	}

	//tmp, _ := strconv.Atoi(r.URL.Query().Get("id"))
	//req.TargetID = tmp
	req.TargetID = 1

	if _, ok := s.store[req.TargetID]; !ok {
		w.Write([]byte("Нет такого пользователя"))
		return
	}
	s.store[req.TargetID].SetAge(req.Age)
	w.Write([]byte("Возраст пользователя был обновлен\n"))
	log.Info("Пользователь ", req.TargetID, " возраст был изменен на ", req.Age)
}

func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//tmp, _ := strconv.Atoi(r.URL.Query().Get("id"))
	//id = tmp
	id := 1

	if _, ok := s.store[id]; !ok {
		w.Write([]byte("Нет такого пользователя"))
		return
	}

	answer := func(u []*u.User) string {
		if len(u) == 0 {
			return "У пользователя нет друзей\n"
		}
		result := "Друзья " + s.store[id].Name + ": "
		for _, man := range u {
			result += "\n" + man.ToString()
		}
		return result + "\n"
	}(s.store[id].GetFriends())

	w.Write([]byte(answer))
}

func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Не удается обработать данные запроса")
		w.Write([]byte(err.Error()))
		return
	}
	var data request
	if err := json.Unmarshal(content, &data); err != nil {
		log.Error("Не удается обработать данные для создания друзей")
		w.Write([]byte(err.Error()))
		return
	}

	if data.TargetID == 0 || data.SourceID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Вам необходимо указать идентификаторы пользователей"))
		return
	}

	if data.TargetID == data.SourceID {
		w.Write([]byte("Тот же пользователь"))
		return
	}

	var tmp u.User
	tmp.Name = s.store[data.TargetID].GetName()
	tmp.Age = s.store[data.TargetID].GetAge()
	tmp.Friends = s.store[data.TargetID].GetFriends()

	if s.store[data.TargetID].AddFriend(s.store[data.SourceID]) {
		s.store[data.SourceID].AddFriend(&tmp)
	} else {
		w.Write([]byte("Пользователи уже являются друзьями\n"))
		return
	}

	w.Write([]byte("Пользователи " + s.store[data.TargetID].GetName() + " и " +
		s.store[data.SourceID].GetName() + " теперь друзья\n"))
	log.Info("Пользователи ", s.store[data.TargetID].GetName()+" и "+
		s.store[data.SourceID].GetName(), " теперь друзья")
}

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Неправильный запрос на удаление пользователя")
		return
	}
	var data request
	if err := json.Unmarshal(content, &data); err != nil {
		log.Error("Не удается обработать данные для удаления пользователя")
		return
	}
	if _, ok := s.store[data.TargetID]; !ok {
		w.Write([]byte("Нет такого пользователя"))
		return
	}

	for _, man := range s.store[data.TargetID].GetFriends() {
		man.RemoveFriend(*s.store[data.TargetID])
	}

	w.Write([]byte("Пользователь " + s.store[data.TargetID].GetName() + " был удален\n"))
	log.Info("Пользователь " + s.store[data.TargetID].GetName() + " был удален")
	delete(s.store, data.TargetID)
}
