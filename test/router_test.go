package test

import (
	"testing"

	"com.zhouhc/goim/router"
)

// 运行的命令: go test -run ^TestGinRouter$ ./test -v
func TestGinRouter(t *testing.T) {
	r := router.Router()
	r.Run(":8083")
}
