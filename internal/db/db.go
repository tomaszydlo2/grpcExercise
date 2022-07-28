package db

import (
	"errors"
	"grpcExercise/internal/users"
)

type Db struct {
	usersList []users.User
}

type Storage interface {
	CreateUser(*users.User)
	UpdateUser(*users.User) error
	DeleteUser(*users.Id) bool
	ReadUser(*users.Id) string
	ReadUsers() string
}

func (db *Db) CreateUser(in *users.User) {
	db.usersList = append(db.usersList, users.User{
		Id:       in.GetId(),
		Username: in.GetUsername(),
		Name:     in.GetName(),
		Surname:  in.GetSurname(),
	})
}

func (db *Db) UpdateUser(in *users.User) error {
	for i, v := range db.usersList {
		if v.Id == in.GetId() {
			db.usersList[i] = users.User{
				Id:       in.GetId(),
				Username: in.GetUsername(),
				Name:     in.GetName(),
				Surname:  in.GetSurname(),
			}
			return nil
		}
	}
	return errors.New("id not found in database")
}

func (db *Db) DeleteUser(in *users.Id) bool {
	for i, v := range db.usersList {
		if v.Id == in.GetId() {
			db.usersList = append(db.usersList[:i], db.usersList[i+1:]...)
			return true
		}
	}
	return false
}

func (db *Db) ReadUser(in *users.Id) string {
	for _, v := range db.usersList {
		if v.Id == in.GetId() {
			return "Username: " + v.GetUsername() + " Name: " + v.GetName() + " Surname: " + v.GetSurname()
		}
	}
	return "Couldn't find user of id " + string(in.GetId())
}

func (db *Db) ReadUsers() string {
	var message string
	for _, v := range db.usersList {
		message += "Username: " + v.GetUsername() + " Name: " + v.GetName() + " Surname: " + v.GetSurname() + "\n"
	}
	return message
}
