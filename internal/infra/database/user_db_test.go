package database

import (
	"testing"

	"github.com/bvaledev/go-expert-commerce-api/internal/domain/entity"
	"github.com/bvaledev/go-expert-commerce-api/pkg/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestUserDb_ShouldCreateUser(t *testing.T) {
	db := testhelper.SetupDBTest(&entity.User{})
	user, err := entity.NewUser("Hermanoteu", "hermanoteu@gmail.com", "123456789")
	if err != nil {
		t.Error(err)
	}
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.NotEmpty(t, userFound.ID)
	assert.NotEmpty(t, userFound.Password)
	assert.Equal(t, "Hermanoteu", userFound.Name)
	assert.Equal(t, "hermanoteu@gmail.com", userFound.Email)
	assert.NotEqual(t, "123456789", userFound.Password)
}

func TestUserDb_ShouldFindUserByEmail(t *testing.T) {
	emailToFind := "hermanoteu@gmail.com"
	db := testhelper.SetupDBTest(&entity.User{})
	user, err := entity.NewUser("Hermanoteu", emailToFind, "123456789")
	if err != nil {
		t.Error(err)
	}
	db.Create(user)
	userDB := NewUserDB(db)

	userFound, err := userDB.FindByEmail(emailToFind)
	assert.Nil(t, err)
	assert.NotNil(t, userFound)

	assert.NotEmpty(t, userFound.ID)
	assert.NotEmpty(t, userFound.Password)
	assert.Equal(t, "Hermanoteu", userFound.Name)
	assert.Equal(t, "hermanoteu@gmail.com", userFound.Email)
	assert.NotEqual(t, "123456789", userFound.Password)

	userFound, err = userDB.FindByEmail("any_email@email.com")
	assert.Nil(t, userFound)
	assert.NotNil(t, err)
}
