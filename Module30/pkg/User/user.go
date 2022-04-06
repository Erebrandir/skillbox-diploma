package User

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

func (u *User) ToString() string {
	result := fmt.Sprintf("Имя %s, Возраст %d", u.Name, u.Age)
	if len(u.Friends) > 0 {
		if len(u.Friends) > 1 {
			result += ", Друзья ["
			for i, man := range u.Friends {
				result += "{"
				result += man.GetName() + " "
				result += strconv.Itoa(man.GetAge()) + "}"
				if i != len(u.Friends)-1 {
					result += ", "
				}
			}
			result += "]\n"
		} else {
			result += " Друзья {"
			result += u.Friends[0].GetName() + " "
			result += strconv.Itoa(u.Friends[0].GetAge()) + "}\n"
		}
	}
	return result
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetFriends() []*User {
	return u.Friends
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) isFriend(friend *User) bool {
	for _, i := range u.Friends {
		if (i.Name == friend.Name) && (i.Age == friend.Age) {
			for j := range i.Friends {
				if reflect.DeepEqual(friend, i.Friends[j]) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (u *User) AddFriend(friend *User) bool {
	if u.isFriend(friend) {
		return false
	}
	u.Friends = append(u.Friends, friend)
	return true
}

func (u *User) ClearFriends() {
	u.Friends = nil
}

func (u *User) RemoveFriend(friend User) {
	for i := range u.Friends {
		if reflect.DeepEqual(u.Friends[i], friend) {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
			return
		}
	}
}

func (u *User) RemoveFriends(friends ...User) {
	for _, man := range friends {
		u.RemoveFriend(man)
	}
}

func NewUser(name string, age int) User {
	user := User{
		Name: name,
		Age:  age,
	}
	return user
}
