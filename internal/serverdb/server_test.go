package serverdb

import (
	"context"
	"errors"
	. "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"grpcExercise/internal/db/mocks"
	"grpcExercise/internal/users"
	"testing"
)

func MockServerInit(t *testing.T) (mocks.MockStorage, *Server, *Controller) {
	controller := NewController(t)

	storage := mocks.NewMockStorage(controller)

	return *storage, &Server{Database: storage}, controller
}

func TestServer_CreateUser(t *testing.T) {

	storage, ts, controller := MockServerInit(t)
	defer controller.Finish()

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
	storage, ts, controller := MockServerInit(t)
	defer controller.Finish()

	usr := users.User{
		Id:       1,
		Username: "test",
		Name:     "test",
		Surname:  "test",
	}

	// Fail
	storage.EXPECT().UpdateUser(&usr).Return(errors.New("id not found in database"))

	res, err := ts.UpdateUser(context.Background(), &usr)

	assert.Error(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())

	// Pass
	storage.EXPECT().UpdateUser(&usr)

	res, err = ts.UpdateUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Successfully updated user "+usr.GetUsername(), res.GetMessage())
}

func TestServer_DeleteUser(t *testing.T) {
	storage, ts, controller := MockServerInit(t)
	defer controller.Finish()

	usr := users.Id{
		Id: 1,
	}

	storage.EXPECT().DeleteUser(&usr)

	res, err := ts.DeleteUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())
}

func TestServer_ReadUser(t *testing.T) {
	storage, ts, controller := MockServerInit(t)
	defer controller.Finish()

	usr := users.Id{
		Id: 1,
	}

	storage.EXPECT().ReadUser(&usr).Return("Couldn't find user of id " + string(usr.GetId()))

	res, err := ts.ReadUser(context.Background(), &usr)

	assert.NoError(t, err)
	assert.Equal(t, "Couldn't find user of id "+string(usr.GetId()), res.GetMessage())
}

func TestServer_ReadUsers(t *testing.T) {
	storage, ts, controller := MockServerInit(t)
	defer controller.Finish()

	storage.EXPECT().ReadUsers()

	res, err := ts.ReadUsers(context.Background(), &users.Empty{})

	assert.NoError(t, err)
	assert.Equal(t, "", res.GetMessage())
}
