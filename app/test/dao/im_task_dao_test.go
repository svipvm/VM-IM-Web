package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/daos"
	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/utils"
)

func TestInsertOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	task := &models.IMTask{
		UserId: 2,
		Name:   "test_five",
		Conf:   "/home/vmice/conf/one",
	}
	err := task_dao.Insert(task)
	if err != nil {
		t.Log(err)
	}
}

func TestUpdateOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	task := &models.IMTask{
		Id:    3,
		Name:  "new_trhee",
		State: 1,
	}
	err := task_dao.Update(task)
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	err := task_dao.Delete(4)
	if err != nil {
		t.Log(err)
	}
}

func TestQueryOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	data := task_dao.Query(5)
	// t.Log(data)
	if data.Id != 0 {
		t.Log(data)
	}
}
