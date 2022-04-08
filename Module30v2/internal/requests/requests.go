package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"skillbox/internal/actionsJson"
	"skillbox/internal/user"
	"sort"
	"strconv"
)

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var makeFriend actionsJson.MakeFriends
	_ = json.NewDecoder(r.Body).Decode(&makeFriend)
	var name1 string
	var name2 string
	_, err := strconv.Atoi(makeFriend.TargetId)
	if err != nil {
		_, err := w.Write([]byte("ID должен быть типа int!"))
		if err != nil {
			return
		}
		return
	}
	_, err = strconv.Atoi(makeFriend.SourceId)
	if err != nil {
		_, err := w.Write([]byte("ID должен быть типа int!"))
		if err != nil {
			return
		}
		return
	}
	for _, u := range user.Users {
		if u.Id == makeFriend.TargetId {
			name1 = u.Name
		}
		if u.Id == makeFriend.SourceId {
			name2 = u.Name
		}
	}
	if name1 == "" || name2 == "" {
		_, err := w.Write([]byte("Пользователи не найдены."))
		if err != nil {
			return
		}
		return
	}
	for index, u := range user.Users {
		if u.Id == makeFriend.TargetId {
			user.Users[index].Friends = append(user.Users[index].Friends, makeFriend.SourceId)
		}
		if u.Id == makeFriend.SourceId {
			user.Users[index].Friends = append(user.Users[index].Friends, makeFriend.TargetId)
		}
	}

	_, err = w.Write([]byte("Пользователь " + name1 + " и пользователь " + name2 + " теперь друзья! Статус: " + strconv.Itoa(http.StatusOK)))
	if err != nil {
		return
	}
}
func GetUsers(w http.ResponseWriter, _ *http.Request) {

	var response string
	for _, u := range user.Users {
		response += u.ToString() + "\n"
	}
	_, err := w.Write([]byte(response))
	if err != nil {
		return
	}

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u user.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	_, err := strconv.Atoi(u.Age)
	if err != nil {
		_, err := w.Write([]byte("Возраст должен быть типа int."))
		if err != nil {
			return
		}
		return
	}
	u.Id = strconv.Itoa(len(user.Users) + 1)
	for i, u := range user.Users {
		if u.Id != strconv.Itoa(i+1) {
			u.Id = strconv.Itoa(i + 1)
			break
		}
		if u.Id == u.Id {
			id, _ := strconv.Atoi(u.Id)
			u.Id = strconv.Itoa(id + 1)
		}
	}
	user.Users = append(user.Users, u)
	sort.SliceStable(user.Users, func(i, j int) bool {
		return user.Users[i].Id < user.Users[j].Id
	})
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Имя: " + u.Name + "\nID пользователя: " + u.Id + "\nСтатус:" + strconv.Itoa(http.StatusCreated)))
	if err != nil {
		return
	}
}

func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := chi.URLParam(r, "id")
	for _, u := range user.Users {
		if u.Id == params {
			_, err := w.Write([]byte("Пользователь: " + u.Name + "\nДрузья в списке: " + u.FriendsToString()))
			if err != nil {
				return
			}
			return
		}
	}
	_, err := w.Write([]byte("Пользователь не найден."))
	if err != nil {
		return
	}
}

func UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateAge actionsJson.UpdateUser
	_ = json.NewDecoder(r.Body).Decode(&updateAge)
	params := chi.URLParam(r, "id")
	_, err := strconv.Atoi(updateAge.NewAge)

	if err != nil {
		_, err := w.Write([]byte("Возраст должен быть типа int."))
		if err != nil {
			return
		}
		return
	}
	for index, item := range user.Users {
		if item.Id == params {
			user.Users[index].Age = updateAge.NewAge
			_, err := w.Write([]byte("Пользователь " + item.Name + ". Возраст изменен. Статус: " + strconv.Itoa(http.StatusOK)))
			if err != nil {
				return
			}
			return
		}
	}
	_, err = w.Write([]byte("Пользователь не найден."))
	if err != nil {
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var makeFriend actionsJson.MakeFriends
	_ = json.NewDecoder(r.Body).Decode(&makeFriend)

	for i, u := range user.Users {
		for j, f := range u.Friends {
			if f == makeFriend.TargetId {
				user.Users[i].Friends = append(u.Friends[:j], u.Friends[j+1:]...)
			}
		}
	}
	for index, u := range user.Users {
		if u.Id == makeFriend.TargetId {
			user.Users = append(user.Users[:index], user.Users[index+1:]...)
			_, err := w.Write([]byte("Пользователь " + u.Name + " удалён. Статус: " + strconv.Itoa(http.StatusOK)))
			if err != nil {
				return
			}
			return
		}
	}
	_, err := w.Write([]byte("Пользователь не найден."))
	if err != nil {
		return
	}
}
