package app

import (
	"blog.io/config"
	"blog.io/controllers"
	"blog.io/persistence"
	"blog.io/routers"
	"bufio"
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"os"
)

// 博客应用服务器
type App struct {
	DB     *gorm.DB
	RPool  *redis.Pool
	Conf   *config.BlogConfig
	Server *dotweb.DotWeb
}

func NewApp() *App {
	return &App{}
}

// 启动服务器
func (app *App) Launch() error {
	app.Conf = config.Config()
	app.initDB()

}

// 打印 佛陀log
func (app *App) bless() {
	if !app.Conf.EnvProd {
		return
	}

	file, err := os.Open("bless.txt")
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 关闭操作
func (app *App) Destory() {
	if app.DB != nil {
		app.DB.Close()
	}

	if app.RPool != nil {
		app.RPool.Close()
	}

	if app.Server != nil {
		app.Server.Close()
	}
}

// 根据配置文件初始化数据库
func (app *App) initDB() {
	app.DB = persistence.GetOrm()
}

//根据配置文件初始化Redis
func (app *App) initRedis() {
	app.RPool = persistence.GetRedisPool()
}

func (app *App) initError() {
	ec := controllers.NewErrorController()
	app.Server.SetNotFoundHandle(ec.NotFound)
	app.Server.SetExceptionHandle(ec.Internal)
	app.Server.SetMethodNotAllowedHandle(ec.MethodNotAllowed)
}

// 根据配置初始化服务器
func (app *App) initServer() {
	app.Server = dotweb.New()

	// 配置Log
	app.Server.SetEnabledLog(app.Conf.LogEnable)
	app.Server.SetLogPath(app.Conf.LogPath)

	// 配置字定义error
	app.initError()

	// 配置环境模式
	if app.Conf.EnvProd {
		app.Server.SetProductionMode()
	} else {
		app.Server.SetDevelopmentMode()
	}

	// 开启Gzip压缩
	app.Server.HttpServer.SetEnabledGzip(true)

}

// 初始化路由器配置
func (app *App) initRoute(){
	r := routers.NewApiRouter(app.Server.HttpServer)

	// /api/v1/ version interface
	r.V1()

	// backend interface
	r.Admin()
}

