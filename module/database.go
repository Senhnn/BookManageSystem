package module

import (
	"bookmanagesystem/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func InitDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// 禁用默认事务
		SkipDefaultTransaction: false,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",
			// 使用单数命名，迁移操作之后User表名是user，而不是users
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		FullSaveAssociations: false,
		Logger:               nil,
		NowFunc:              nil,
		DryRun:               false,
		PrepareStmt:          false,
		DisableAutomaticPing: false,
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		fmt.Println("数据库连接失败: ", err)
	}

	// 迁移数据库
	db.AutoMigrate(&User{})

	sqlDB, _ := db.DB()
	// 设置连接池中最大的限制连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置数据库最大的连接数量
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	sqlDB.SetConnMaxLifetime(5 * time.Second)
}
