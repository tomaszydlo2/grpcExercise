package db

import (
	. "github.com/golang/mock/gomock"
	"grpcExercise/internal/db/mocks"
	"grpcExercise/internal/users"
	"testing"
)

func Test_db(t *testing.T) {
	controller := NewController(t)
	defer controller.Finish()

	storage := mocks.NewMockStorage(controller)

	//mockUser := users.User{
	//	Id:       1,
	//	Username: "test",
	//	Name:     "test",
	//	Surname:  "test",
	//}
	mockUsrId := users.Id{Id: 3}

	storage.EXPECT().ReadUser(&mockUsrId).Return("No user of id ")

}
