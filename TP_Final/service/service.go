package service

import "time"

type User struct {
	ID       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname"`
}

type UserService interface {
	Save(User) (User, error)
	Find() []User
}

type userServices struct {
	data map[int64]User
}

func NewUserServices() UserService {
	d := map[int64]User{}
	return &userServices{d}
}

func (s *userServices) Save(u User) (User, error) {
	u.ID = time.Now().Unix()
	s.data[u.ID] = u
	return u, nil
}

func (s *userServices) Find() []User {
	list := []User{}
	for _, v := range s.data {
		list = append(list, v)
	}
	return list
}
