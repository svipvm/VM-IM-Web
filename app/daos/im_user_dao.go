package daos

import (
	"github.com/svipvm/VM-IM-Web/app/models"

	"github.com/go-xorm/xorm"
)

const TABLE_USER string = "im_user"

type IMUserDao struct {
	engine *xorm.Engine
}

func BuildIMUserDao(engine *xorm.Engine) *IMUserDao {
	return &IMUserDao{engine: engine}
}

func (dao *IMUserDao) Insert(data *models.IMUser) error {
	_, err := dao.engine.Table(TABLE_USER).Insert(data)
	return err
}

func (dao *IMUserDao) Delete(username string) error {
	data := &models.IMUser{Username: username}
	_, err := dao.engine.Table(TABLE_USER).Delete(data)
	return err
}

func (dao *IMUserDao) Query(username string) *models.IMUser {
	data := &models.IMUser{Username: username}
	ok, err := dao.engine.Table(TABLE_USER).Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (dao *IMUserDao) Update(data *models.IMUser) error {
	// must need this id
	_, err := dao.engine.Table(TABLE_USER).ID(data.Id).Update(data)
	return err
}
