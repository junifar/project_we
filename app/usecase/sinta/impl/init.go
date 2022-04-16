package impl

import (
	sintarpi "project_we/app/repository/sinta"
	sintauc "project_we/app/usecase/sinta"
)

type sintaUsecase struct {
	sintaRP sintarpi.ISinta
}

func NewSintaUsecase(sintaRP sintarpi.ISinta) sintauc.ISinta {
	return &sintaUsecase{
		sintaRP: sintaRP,
	}
}
