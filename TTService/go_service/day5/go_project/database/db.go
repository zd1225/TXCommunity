package database

import (
	"fmt"
	"go_project/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.AppConfig
	//fmt.Sprintf()是生成一串字符串
	//root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local
	//是给 MySQL 驱动用的连接信息。
	//把这段连接字符串传给驱动，去连接数据库。
	//只不过你用了配置文件 + 格式化拼接，更灵活、更规范。
	//mysql.Open(dsn) = 用 DSN 连接串连接 MySQL
	//charset=utf8mb4：支持表情符号、所有中文（必须）
	//parseTime=True：让 Go 能正确处理时间（超级重要）
	//Go 的 time.Time 类型，默认不支持从 MySQL 直接读取。
	/*
		如果你不写 parseTime=True：
		数据库里是 datetime / timestamp
		Go 读出来会报错
		或者读成字符串，无法比较、格式化
	*/

	/*
			为什么必须写 loc=Local？
		不写 = 时间错 8 小时！
		MySQL 默认时区是 UTC
		你的电脑 / 服务器是 东八区（UTC+8）
		如果不写 loc=Local：
		存进去：2025-01-01 10:00:00
		读出来：2025-01-01 02:00:00 ❌ 直接少 8 小时
		2024-01-01 10:00:00 +0800 → UTC 时间 = 2024-01-01 02:00:00 +0000

		数据库里实际存的数值是：2024-01-01 02:00:00(MySQL 认为这就是一个“无时区标记的时间字面量”）
		如果你 错误地 把这个 02:00 UTC 直接当成北京时间来用（比如不做任何转换就显示给用户），那么用户看到的是 凌晨 2 点，比你想要的上午 10 点 少了 8 小时。

		这就是文档里常说的“少了 8 小时” —— 是 显示上的偏差，不是时间点绝对值错误。


	*/
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	var err error
	//下面这句是：？？？？
	//gorm.Open(驱动, 配置)
	//GORM.Open() :创建数据库连接的核心函数
	//DB：数据库连接对象，后续所有增删改查都用它

	//mysql.Open(dsn):MySQL 驱动打开连接
	//dsn = 数据源名称

	//&gorm.Config{}:GORM 的配置项(你也可以在这里开启日志、禁用事务等：)
	/*
			&gorm.Config{
		  	SkipDefaultTransaction: true,  // 禁用默认事务
	*/
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	//自动迁移
	err = DB.AutoMigrate(&models.User())
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connected successfully")
}
