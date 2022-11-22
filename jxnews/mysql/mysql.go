package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"jxnew/mode"
	"time"
)

var DB *gorm.DB

func init() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:root@tcp(127.0.0.1:3307)/jxnews?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 191,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           // 最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 容纳的链接数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 链接的最大可复用时间
	DB = db
}

//func main() {
//	add(mode.TM{
//		Question: "要完善以______为核心的中国特色社会主义法律体系，加强宪法实施和监督，加强重点领域、新兴领域、涉外领域立法，推进科学立法、民主立法、依法立法。",
//		Answer:   "1",
//	})
//}

// Query 查询全部
func Query() (tms []mode.TM) {
	DB.Find(&tms)
	return tms
}

// QueryOne 单个查询
func QueryOne(question string) (tmQuery mode.TM) {
	err := DB.Model(tmQuery).Where("question = ?", question).First(&tmQuery).Error
	if err != nil {
		return mode.TM{}
	}
	return tmQuery
}

// Add 添加单个
func Add(tmAdd mode.TM) {
	one := QueryOne(tmAdd.Question)
	if one.Question == "" && one.Answer == "" {
		DB.Create(&tmAdd) // 添加
	}
}
