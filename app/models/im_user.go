package models

type IMUser struct {
	Id       int32  `xorm:"int notnull autoincr pk 'id'"`
	Name     string `xorm:"varchar(63) default 'IMER' 'er_name'"`
	Username string `xorm:"varchar(63) notnull 'login_usr'"`
	Salt     string `xorm:"varchar(63) notnull 'salt'"`
	Password string `xorm:"varchar(63) notnull 'login_pwd'"`
}
