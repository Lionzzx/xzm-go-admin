package model

import "github.com/jinzhu/gorm"

type SysUser struct {
gorm.Model
Username    string       `json:"userName" gorm:"comment:'用户登录名'"`
Password    string       `json:"-"  gorm:"comment:'用户登录密码'"`
NickName    string       `json:"nickName" gorm:"default:'系统用户';comment:'用户昵称'" `
HeaderImg   string       `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
AuthorityId string       `json:"authorityId" gorm:"default:888;comment:'用户角色ID'"`
}
