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

func (dao *IMUserDao) Insert(data *models.IMUser) (bool, error) {
	affected, err := dao.engine.Table(TABLE_USER).Insert(data)
	return affected > 0, err
}

func (dao *IMUserDao) Delete(username string) (bool, error) {
	data := &models.IMUser{Username: username}
	affected, err := dao.engine.Table(TABLE_USER).Delete(data)
	return affected > 0, err
}

func (dao *IMUserDao) Update(data *models.IMUser) (bool, error) {
	// must need this id
	affected, err := dao.engine.Table(TABLE_USER).ID(data.Id).Update(data)
	return affected > 0, err
}

func (dao *IMUserDao) Query(username string) *models.IMUser {
	data := &models.IMUser{Username: username}
	has, err := dao.engine.Table(TABLE_USER).Get(data)
	if has && err == nil {
		return data
	} else {
		data.Id = -1
		return data
	}
}
