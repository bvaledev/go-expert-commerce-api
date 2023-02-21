package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_ShouldCreateNewUserWithCorrectValues(t *testing.T) {
	user, err := NewUser("Hermanoteu", "hermanoteu@gmail.com", "123456789")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Hermanoteu", user.Name)
	assert.Equal(t, "hermanoteu@gmail.com", user.Email)
	assert.NotEqual(t, "123456789", user.Password)
}

func TestUser_ShouldValidatePassword(t *testing.T) {
	user, err := NewUser("Hermanoteu", "hermanoteu@gmail.com", "123456789")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	isValidPassword := user.ValidatePassword("123456789")

	assert.True(t, isValidPassword)

	invalidPassword := user.ValidatePassword("invalid_password")

	assert.False(t, invalidPassword)

}
