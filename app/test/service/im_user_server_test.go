package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/models"
	"github.com/svipvm/VM-IM-Web/app/services"
)

func TestAddOneUser(t *testing.T) {
	user_service := services.BuildIMUserServer()
	user := &models.IMUser{
		Name:     "wwwwhq",
		Username: "adminss",
		Salt:     "dfadf21",
		Password: "rootadmin",
	}
	err := user_service.AddOneUser(user)
	if err != nil {
		t.Log(err)
	}
}
