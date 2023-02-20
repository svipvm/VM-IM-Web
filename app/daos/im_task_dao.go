package daos

import (
	"github.com/svipvm/VM-IM-Web/app/models"

	"github.com/go-xorm/xorm"
)

const TABLE_TASK string = "im_task"

type IMTaskDao struct {
	engine *xorm.Engine
}

func BuildIMTaskDao(engine *xorm.Engine) *IMTaskDao {
	return &IMTaskDao{engine: engine}
}

func (dao *IMTaskDao) Insert(data *models.IMTask) error {
	_, err := dao.engine.Table(TABLE_TASK).Insert(data)
	return err
}

func (dao *IMTaskDao) Delete(id int) error {
	data := &models.IMTask{Id: int32(id)}
	_, err := dao.engine.Table(TABLE_TASK).Delete(data)
	return err
}

func (dao *IMTaskDao) Query(id int) *models.IMTask {
	data := &models.IMTask{Id: int32(id)}
	ok, err := dao.engine.Table(TABLE_TASK).Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (dao *IMTaskDao) Update(data *models.IMTask) error {
	// must need this id
	_, err := dao.engine.Table(TABLE_TASK).ID(data.Id).Update(data)
	return err
}
