package store

import(
	"fmt"
	"errors"
)

//对扩展开放，对修改关闭
//数据库接口
type DBStore interface {
	Save(data string) error
	// Load() string
}

//业务逻辑只依赖接口
type DBService struct {
	db DBStore//为了解耦合，这里不使用具体数据库，高层依赖不会直接调用低层依赖
}

//开启具体业务
func (service *DBService)createUserDB(id int, name string) error {
	//关键逻辑，关键执行哪个方法，取决于db里装的什么
	err := service.db.Save()
	if err != nil {
		return fmt.Errorf("保存用户信息失败：%w", error)
	}
	fmt.Println("用户创建成功")
	return nil
}



//具体数据库
type MySQLDB struct {

}

func (db MySQLDB)Save(data string) error {
	fmt.Println("数据保存中...")
}

//测试数据库或者后期替代数据库
type TestDB struct{}
func (db TestDB)Save(data string) error{}


//如果就这样结束的话，我们大概率会这么使用service，如果多个地方运用到了保存数据，
//那么又会产生我们担心的事情，高度产生具体数据库的初始化，牵一发动全身
//某一个需要调用此方法的地方
func useSaveFunction() {
	productService := DBService{db : MySQLDB{}}
	productService.createUserDB(1001, "tome")

	testService := DBService(db: TestDB{})
	testService.createUserDB(10001, "tom")
}

//为了防止多个地方使用service，造成的多处MySQLDB数据库初始化造成的耦合，
//我们做一个工厂函数，再也不需要改动一项任务要更改多处代码

package factory

//一种很好的方式，是使用配置文件来初始化数据库，用工厂管理service
//TODO:
//pkg/app/container.go
package app
type container struct {
	dbService *dbService
	testService *testService
	//还有很多其它服务，这里进行统一管理
	//...
}

var global *container

func Init(config *Config) {
	global = &Container{
		dbService: &DBService{db: createDB(config)}
	}
}

//如果未来需要修改数据库的类型，只需要改这里就好
func createDB(config: Config) *DBStore {
	if config.EVN == "sql" {
		return DBStore{}//里面可以传一些配置信息，name, host等等
	} 
	return TestDB{}
}

func Get() *Container {
    if global == nil {
        panic("app not initialized")
    }
    return global
}

package main
// main.go
func main() {
    config := loadConfig()
    app.Init(config)
    // 使用全局容器获取服务
    service := app.Get().dbService
    // 启动服务
}

