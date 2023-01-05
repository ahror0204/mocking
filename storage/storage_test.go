package storage

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *User {
	u, err := strg.CreateUser(&User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, u)

	return u
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestGetUser(t *testing.T) {
	u := createUser(t)

	u, err := strg.GetUser(u.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
}
