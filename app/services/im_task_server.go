package services

import (
	"github.com/svipvm/VM-IM-Web/app/daos"
	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/utils"
)

type imTaskServer struct {
	dao *daos.IMTaskDao
}

type IMTaskServer interface {
	AddOneTask(task *models.IMTask) error
}

func BuildIMTaskServer() IMTaskServer {
	return &imTaskServer{
		dao: daos.BuildIMTaskDao(utils.GetEngine()),
	}
}

func (s *imTaskServer) AddOneTask(task *models.IMTask) error {
	return s.dao.Insert(task)
}
