package test

import (
	"fmt"
	"testing"

	"com.zhouhc/goim/modules"
	"com.zhouhc/goim/utils"
)

// 运行命令 go test -run ^TestGorm$ ./test -v
func TestGorm(t *testing.T) {
	db, err1 := utils.GetMySQLConnect()
	if err1 != nil {
		fmt.Printf("数据库初始化失败, %v\n", err1)
		panic("测试失败")
	}
	user := modules.User{}
	db.Find(&user)
	fmt.Println(user)
	if user.Name == "" {
		panic("测试失败")
	}
}
