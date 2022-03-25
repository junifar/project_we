package impl

import (
	dosenrp "project_we/app/repository/dosen"
	userrp "project_we/app/repository/user"
	useruc "project_we/app/usecase/user"
)

//UserUsecase -
type UserUsecase struct {
	User  userrp.IUser
	Dosen dosenrp.IDosen
}

func New(User userrp.IUser, Dosen dosenrp.IDosen) useruc.IUser {
	return &UserUsecase{
		User:  User,
		Dosen: Dosen,
	}
}
