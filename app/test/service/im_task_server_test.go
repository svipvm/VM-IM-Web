package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/services"
)

func TestAddOneTask(t *testing.T) {
	task_service := services.BuildIMTaskServer()
	task := &models.IMTask{
		UserId: 23,
		Name:   "finally_one",
		Conf:   "/home/svipvm/conf/files",
	}
	err := task_service.AddOneTask(task)
	if err != nil {
		t.Log(err)
	}
}
