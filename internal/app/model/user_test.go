package model_test

import (
	"testing"

	"github.com/Jereyji/Wep-app.git/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name string
		u func() *model.User
		isValid bool
	} {
		{
			name: "valid",
			u: func()*model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Email = ""

				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty name",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Name = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short name",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Name = "S"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Password = "1234"

				return u
			},
			isValid: false,
		},
		{
			name: "with encrypt password",
			u: func()*model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted password"
				return u
			},
			isValid: true,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}