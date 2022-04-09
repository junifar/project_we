package impl

import (
	dosenrp "project_we/app/repository/dosen"
	institusirp "project_we/app/repository/institusi"
	jabatanfungsionalrp "project_we/app/repository/jabatanfungsional"
	jenjangpendidikanrp "project_we/app/repository/jenjangpendidikan"
	programstudirp "project_we/app/repository/programstudi"
	userrp "project_we/app/repository/user"
	useruc "project_we/app/usecase/user"
)

//UserUsecase -
type UserUsecase struct {
	User              userrp.IUser
	Dosen             dosenrp.IDosen
	Institusi         institusirp.IInstitusi
	JenjangPendidikan jenjangpendidikanrp.IJenjangPendidikan
	ProgramStudi      programstudirp.IProgramStudi
	JabatanFungsional jabatanfungsionalrp.IJabatanFungsional
}

func New(User userrp.IUser,
	Dosen dosenrp.IDosen,
	Institusi institusirp.IInstitusi,
	JenjangPendidikan jenjangpendidikanrp.IJenjangPendidikan,
	ProgramStudi programstudirp.IProgramStudi,
	JabatanFungsional jabatanfungsionalrp.IJabatanFungsional,
) useruc.IUser {
	return &UserUsecase{
		User:              User,
		Dosen:             Dosen,
		Institusi:         Institusi,
		JenjangPendidikan: JenjangPendidikan,
		ProgramStudi:      ProgramStudi,
		JabatanFungsional: JabatanFungsional,
	}
}
