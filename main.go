package main

import "com.zhouhc/goim/router"

func main() {
	r := router.Router()
	r.Run(":8080")
}
