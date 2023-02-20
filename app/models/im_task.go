package models

import "time"

type IMTask struct {
	Id         int32     `xorm:"int notnull autoincr pk 'id'"`
	UserId     int32     `xorm:"int notnull 'user_id'"`
	Name       string    `xorm:"varchar(63) notnull 'task_name'"`
	Conf       string    `xorm:"varchar(255) notnull 'task_conf'"`
	State      int8      `xorm:"int(8) default 0 'task_state'"`
	CreateData time.Time `xorm:"datatime created 'create_data'"`
	ModifyData time.Time `xorm:"datatime created 'modify_data'"`
}
