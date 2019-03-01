package relations

import "blogserver/persistence"

/*
	generator.go 根据结构体(Models) 统一创建数据库关系
 */

func InitRelations() {
	db := persistence.GetOrm()
	// 判断存不存在表, 不存在就新建, 否则就是自动迁移(其他修改)
	if !db.HasTable("users") {
	 	db.CreateTable(&models.User{}, &models.UserInfo{}, &models.Category{}, &models.CategoryItem{},
			&models.Article{}, models.FriendlyLink{})
	}
}
