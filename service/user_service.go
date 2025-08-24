package service

import (
	"net/http"

	"com.zhouhc/goim/dao"
	"com.zhouhc/goim/model"
	"github.com/gin-gonic/gin"
)

// 创建用户
func UserCreate(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//数据库保存
	if _, err1 := dao.UserCreate(&u); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	//返回数据
	c.JSON(http.StatusOK, u)
}
