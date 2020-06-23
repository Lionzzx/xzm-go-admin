package v1

import (
	"fmt"
	"ginEssential/global/response"
	"ginEssential/model/request"
	"ginEssential/model"
	resp "ginEssential/model/response"
	"ginEssential/service"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
)

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)

	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

func CreateUser(c *gin.Context) {
	var user model.SysUser
	_ = c.ShouldBindJSON(&user)

	err := service.CreateUser(user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建数据失败，%v", err), c)
	} else {
		response.OkWithData(true, c)
	}
}
func Test(c *gin.Context) {
	file, header, _ := c.Request.FormFile("file")
	fmt.Println(file)
	err, filePath, _ := utils.Upload(header)
	fmt.Println(err)
	response.OkWithData(filePath, c)
}