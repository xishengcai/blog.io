package main

import (
	"blog.io/app"
	"blog.io/relations"
)

func main() {
	// 判断表是否存在 存在就自动迁移模式
	relations.InitRelations()

	// 启动http 服务
	app := app.NewApp()
	defer app.Destory()
	app.Launch()
}
