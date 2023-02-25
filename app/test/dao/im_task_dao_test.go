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
		UserId: 23,
		Name:   "test_five_hah",
		Conf:   "/home/vmice/conf/one",
	}
	affected, err := task_dao.Insert(task)
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}

func TestUpdateOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	task := &models.IMTask{
		Id:    13,
		Name:  "new_trhee",
		State: 1,
	}
	affected, err := task_dao.Update(task)
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}

func TestDeleteOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	affected, err := task_dao.Delete(12)
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}

func TestQueryOneTask(t *testing.T) {
	task_dao := daos.BuildIMTaskDao(utils.GetEngine())
	data := task_dao.Query(13)
	t.Log(data)
}
