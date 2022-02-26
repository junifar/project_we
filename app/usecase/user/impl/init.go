package impl

import (
	userrp "project_we/app/repository/user"
	useruc "project_we/app/usecase/user"
)

//UserUsecase -
type UserUsecase struct {
	User userrp.IUser
}

func New(User userrp.IUser) useruc.IUser {
	return &UserUsecase{
		User: User,
	}
}
