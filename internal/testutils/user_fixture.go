package testutils

import "github.com/ChristopherVennemann/Go-AcademyDay/internal/model"

type UserComparable struct {
	Username string
	Email    string
}

func NewUser(username string, email string) *model.User {
	return &model.User{
		Username: username,
		Email:    email,
	}
}

func SavedUser(user *model.User, id int, createdAt string) *model.User {
	saved := *user
	saved.ID = id
	saved.CreatedAt = createdAt
	return &saved
}

func ToUserComparable(users []*model.User) []UserComparable {
	res := make([]UserComparable, len(users))
	for i, u := range users {
		res[i] = UserComparable{Username: u.Username, Email: u.Email}
	}
	return res
}
