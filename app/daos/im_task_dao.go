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

func (dao *IMTaskDao) Insert(data *models.IMTask) (bool, error) {
	affected, err := dao.engine.Table(TABLE_TASK).Insert(data)
	return affected > 0, err
}

func (dao *IMTaskDao) Delete(id int) (bool, error) {
	data := &models.IMTask{Id: int32(id)}
	affected, err := dao.engine.Table(TABLE_TASK).Delete(data)
	return affected > 0, err
}

func (dao *IMTaskDao) Query(id int) *models.IMTask {
	data := &models.IMTask{Id: int32(id)}
	ok, err := dao.engine.Table(TABLE_TASK).Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = -1
		return data
	}
}

func (dao *IMTaskDao) Update(data *models.IMTask) (bool, error) {
	// must need this id
	affected, err := dao.engine.Table(TABLE_TASK).ID(data.Id).Update(data)
	return affected > 0, err
}
