package impl

import (
	dosenrp "project_we/app/repository/dosen"
	institusirp "project_we/app/repository/institusi"
	userrp "project_we/app/repository/user"
	useruc "project_we/app/usecase/user"
)

//UserUsecase -
type UserUsecase struct {
	User      userrp.IUser
	Dosen     dosenrp.IDosen
	Institusi institusirp.IInstitusi
}

func New(User userrp.IUser, Dosen dosenrp.IDosen, Institusi institusirp.IInstitusi) useruc.IUser {
	return &UserUsecase{
		User:      User,
		Dosen:     Dosen,
		Institusi: Institusi,
	}
}
