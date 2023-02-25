package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/daos"
	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/utils"
)

func TestInsertOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	user := new(models.IMUser)
	user.Name = "WHQ4"
	user.Username = "svipvm6"
	user.Password = "admins"
	user.Salt = "fjalsjdo2"
	affected, err := user_dao.Insert(user)
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}

func TestDeleteOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	affected, err := user_dao.Delete("svipvms")
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}

func TestQueryOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	data := user_dao.Query("svipvms")
	t.Log(data)
}

func TestUpdateOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	user := new(models.IMUser)
	user.Id = 1
	user.Name = "WHQ3"
	user.Username = "vmicesss"
	user.Password = "admins"
	user.Salt = "aa32fd3"
	affected, err := user_dao.Update(user)
	if err != nil {
		t.Log(err)
	} else if affected == false {
		t.Log("No something.")
	}
}
