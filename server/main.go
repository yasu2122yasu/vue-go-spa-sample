package main

import (
	"github.com/aocm/vue-go-spa-sample/server/handler"
	"github.com/aocm/vue-go-spa-sample/server/infra/accessor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	IntiDb()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/yamabiko", handler.YamabikoAPI())
	e.GET("/yamabiko", handler.GetHistory())
	e.OPTIONS("/yamabiko", handler.OptionsCheck())

	e.Start(":8080")
}

func InitDb() {
	accessor.AccessDB(accessor.MysqlAccessor{})
}

/*
package main

パッケージをimportする
import (
	"github.com/aocm/vue-go-spa-sample/server/handler"
	"github.com/aocm/vue-go-spa-sample/server/infra/accessor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

フレームワーク名.New()を変数eに代入。
initdbはPostgreSQLのデータベースクラスタを新しく作成します。
 データベースクラスタとは、1つのサーバインスタンスで管理されるデータベースの集合です
func main() {
	e := echo.New()
	IntiDb()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/yamabiko", handler.YamabikoAPI())
	e.GET("/yamabiko", handler.GetHistory())
	e.OPTIONS("/yamabiko", handler.OptionsCheck())

	e.Start(":8080")
}

func InitDb() {
	accessor.AccessDB(accessor.MysqlAccessor{})
}


*/
