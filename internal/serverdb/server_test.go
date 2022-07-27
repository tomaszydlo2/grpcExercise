package serverdb

import (
	"context"
	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grpcExercise/internal/db/mocks"
	"grpcExercise/internal/users"
	"testing"
)

func MockServerInit(t *testing.T) (mocks.MockStorage, *Server) {

	controller := NewController(t)
	defer controller.Finish()

	storage := mocks.NewMockStorage(controller)

	return *storage, &Server{Database: storage}
}

func TestServer_CreateUser(t *testing.T) {

	storage, ts := MockServerInit(t)

	usr := users.User{
		Id:       1,
		Username: "test",
		Name:     "test",
		Surname:  "test",
	}

	storage.EXPECT().CreateUser(&usr)

	res, err := ts.CreateUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Created user "+usr.GetUsername(), res.GetMessage())
}

func TestServer_UpdateUser(t *testing.T) {
	storage, ts := MockServerInit(t)

	usr := users.User{
		Id:       1,
		Username: "test",
		Name:     "test",
		Surname:  "test",
	}

	storage.EXPECT().UpdateUser(&usr)

	res, err := ts.UpdateUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())
}

func TestServer_DeleteUser(t *testing.T) {
	storage, ts := MockServerInit(t)

	usr := users.Id{
		Id: 1,
	}

	storage.EXPECT().DeleteUser(&usr)

	res, err := ts.DeleteUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())
}

func TestServer_ReadUser(t *testing.T) {
	storage, ts := MockServerInit(t)

	usr := users.Id{
		Id: 1,
	}

	storage.EXPECT().ReadUser(&usr).Return("Couldn't find user of id " + string(usr.GetId()))

	res, err := ts.ReadUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())
}

func TestServer_ReadUsers(t *testing.T) {
	storage, ts := MockServerInit(t)

	storage.EXPECT().ReadUsers()

	res, err := ts.ReadUsers(context.Background(), &users.Empty{})

	assert.NoError(t, err)
	assert.Equal(t, "", res.GetMessage())
}
