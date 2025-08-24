package dao

import (
	"fmt"

	"com.zhouhc/goim/model"
	"com.zhouhc/goim/utils"
)

// MySQL保存用户
func UserCreate(u *model.User) (*model.User, error) {
	db, err1 := utils.GetMySQLConnect()
	if err1 != nil {
		fmt.Printf("数据库连接失败,保存用户失败,%v\n", err1)
		return nil, err1
	}
	db.Create(u)
	return u, nil
}
