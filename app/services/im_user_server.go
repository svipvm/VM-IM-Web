package services

import (
	"github.com/svipvm/VM-IM-Web/app/daos"
	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/utils"
)

type imUserService struct {
	dao *daos.IMUserDao
}

type IMUserService interface {
	AddOneUser(user *models.IMUser) error
}

func BuildIMUserServer() IMUserService {
	return &imUserService{
		dao: daos.BuildIMUserDao(utils.GetEngine()),
	}
}

func (s *imUserService) AddOneUser(user *models.IMUser) error {
	_, err := s.dao.Insert(user)
	return err
}
