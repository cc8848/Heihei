package routers

import (
	"Heihei/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goat/indexer", &controllers.BookfulController{}, "get:Indexer")
	beego.Router("/goat/searcher/:query", &controllers.BookfulController{}, "get:Searcher")
}
