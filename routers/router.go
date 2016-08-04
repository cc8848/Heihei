package routers

import (
	"Heihei/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/heihei/indexer", &controllers.BookfulController{}, "get:Indexer")
	beego.Router("/heihei/searcher/:query", &controllers.BookfulController{}, "get:Searcher")
  beego.Router("/heihei/bookful/:query", &controllers.BookfulController{}, "get:SearchBookful")
}
