package models

//CREATE TABLE `users` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
//`name` varchar(64) NOT NULL DEFAULT '' COMMENT '姓名',
//`email` varchar(128) NOT NULL COMMENT '邮箱',
//`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
//`birthday` date DEFAULT '0000-00-00' COMMENT '生日',
//`sex` enum('MAN','WOMAN','UNKNOWN') NOT NULL DEFAULT 'UNKNOWN' COMMENT '性别',
//`passwd` char(32) NOT NULL COMMENT '密码',
//`salt` char(4) NOT NULL COMMENT '盐',
//`status` enum('INIT','ON','OFF','DELETE') NOT NULL DEFAULT 'INIT' COMMENT '用户状态:初始化，正常，禁用，删除',
//`remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注字段',
//`extra` varchar(20) NOT NULL DEFAULT '' COMMENT '附加字段',
//`ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//PRIMARY KEY (`id`),
//KEY `mobile` (`mobile`)
//) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

type User struct {
	Id       int
	Name     string
	Email    string
	Mobile   string
	Password string `gorm:"column:passwd"`
	Salt     string
	Status   string
	Ctime    string
	Utime    string
}

func NewUser() *User {
	return &User{}
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "users"
}

func (u *User) GetUserByMobile(userName string) (User, error) {
	var user User
	err := db.Where("user_name = ?", userName).First(&user).Error
	return user, err
}
