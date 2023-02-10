package user

import (
	"context"

	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) func(context.Context, string) (User, error) {
	return func(c context.Context, userName string) (User, error) {

		var user User
		err := db.Find(&user).Error

		return user, err
	}
}
