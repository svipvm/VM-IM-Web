package daos

import (
	"github.com/go-xorm/xorm"
	"github.com/svipvm/VM-IM-Web/app/models"
)

type IMUserDao struct {
	engine *xorm.Engine
}

func (dao *IMUserDao) Insert(data *models.IMUser) error {
	_, err := dao.engine.Table("im_user").Insert(data)
	return err
}
