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
	user.Username = "svipvm"
	user.Password = "admins"
	user.Salt = "fjalsjdo2"
	err := user_dao.Insert(user)
	if err != nil {
		t.Log(err)
	}
}

func TestDeleteOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	err := user_dao.Delete("vmice")
	if err != nil {
		t.Log(err)
	}
}

func TestQueryOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	data := user_dao.Query("vmices")
	// t.Log(data)
	if data.Id != 0 {
		t.Log(data)
	}
}

func TestUpdateOneUser(t *testing.T) {
	user_dao := daos.BuildIMUserDao(utils.GetEngine())
	user := new(models.IMUser)
	user.Id = 2
	user.Name = "WHQ3"
	user.Username = "vmicesss"
	user.Password = "admins"
	user.Salt = "aa32fd3"
	err := user_dao.Update(user)
	if err != nil {
		t.Log(err)
	}
}
