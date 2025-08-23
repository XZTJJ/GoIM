package modules

import (
	"time"

	"gorm.io/gorm"
)

// 用户相关信息类
type User struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    uint32
	LoginTime     time.Time
	HeartbeatTime uint64
	LoginOutTime  time.Time `grom:"column:login_outtime"`
	IsLogout      bool
	DeviceInfo    string
}

// 数据库对应的表名
func (u *User) TableName() string {
	return "go_user_info"
}
