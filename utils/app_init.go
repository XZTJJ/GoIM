package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化MySQL
var config *viper.Viper

// 初始化方式
func init() {
	v, err1 := initConfig()
	if err1 != nil {
		fmt.Printf("app.yml文件初始化失败,请检查工作目录的config目录下是否有app.yml文件, %v\n", err1)
	}
	config = v
}

// 初始化配置文件
func initConfig() (*viper.Viper, error) {
	v := viper.New()
	//可执行文件目录，当前目录
	exeDir, _ := os.Executable()
	pwdDir, _ := os.Getwd()
	//处理后缀文件,去掉多余的后缀
	separator := string(filepath.Separator)
	exeDirAll := strings.Split(exeDir, separator)
	pwdDirAll := strings.Split(pwdDir, separator)
	//viper会自动在这些目录中查找
	v.AddConfigPath(exeDir)
	v.AddConfigPath(filepath.Join(exeDir, "config"))

	exeParentDir := strings.Join(exeDirAll[0:len(exeDirAll)-1], separator)
	v.AddConfigPath(exeParentDir)
	v.AddConfigPath(filepath.Join(exeParentDir, "config"))

	pwdParentDir := strings.Join(pwdDirAll[0:len(pwdDirAll)-1], separator)
	v.AddConfigPath(pwdParentDir)
	v.AddConfigPath(filepath.Join(pwdParentDir, "config"))

	v.AddConfigPath(pwdDir)
	v.AddConfigPath(filepath.Join(pwdDir, "config"))

	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	//设置文件名和后缀
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	err1 := v.ReadInConfig()
	return v, err1
}

// 初始化MySQl连接
func GetMySQLConnect() (*gorm.DB, error) {
	user := config.GetString("db.mysql.user")
	password := config.GetString("db.mysql.password")
	database := config.GetString("db.mysql.database")
	host := config.GetString("db.mysql.host")
	port := config.GetString("db.mysql.port")
	metaUrl := config.GetString("db.mysql.metaUrl")
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, database, metaUrl)
	return gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{})
}
