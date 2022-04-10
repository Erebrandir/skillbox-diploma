package update

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"skillbox/internal/user"
	"strconv"
)

type Update struct {
	NewAge  string `json:"new_age"`
	NewName string `json:"new_name"`
}

func (newUserAge *Update) UpdateAge(userStorage user.Storage, w http.ResponseWriter, r http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&newUserAge)
	params := chi.URLParam(&r, "id")
	_, err := strconv.Atoi(newUserAge.NewAge)

	if err != nil {
		_, err := w.Write([]byte("Возраст должен быть типа int"))
		if err != nil {
			return
		}
		return
	}
	for index, item := range userStorage.Users {
		if item.Id == params {
			userStorage.Users[index].Age = newUserAge.NewAge

			_, err = w.Write([]byte("Пользователь " + item.Name + ". Возраст изменен! Статус: " + strconv.Itoa(http.StatusOK)))
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
