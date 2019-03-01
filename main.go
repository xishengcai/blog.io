package main

import "blog.io/relations"

func main() {
	// 判断表是否存在 存在就自动迁移模式
	relations.InitRelations()

}
