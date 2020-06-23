package service

import (
	"ginEssential/global"
	"ginEssential/model"
	"ginEssential/model/request"
	"time"
)

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser

	err = db.Count(&total).Error
	err = db.Find(&userList).Limit(limit).Offset(offset).Error
	return err, userList, total
}

func CreateUser(user model.SysUser) (err error) {
	db := global.GVA_DB.Model(&model.SysUser{})
	err = db.Create(&user).Error
	return err
}
func Test(user model.SysUser) (err error, key string) {
	db := global.GVA_DB.Model(&model.SysUser{})
	err = db.Model(user).Update("username",user.Username).Error
	global.GVA_REDIS.Set("test","我修改了",time.Hour).Result()
	key,_ = global.GVA_REDIS.Get("test").Result()
	return err, key
}