package controllers

import (
	"github.com/astaxie/beego"
  "Heihei/models"
  "Heihei/services"
)

type BookfulController struct {
	beego.Controller
}


func (this *BookfulController) Indexer() {
	var rt models.Result

	err, rtv := services.Indexer()
	if err != nil {
		rt.Msg = "o_o"
		beego.Info(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		rt.Msg = "^_^"
		rt.Data = make([]models.Recs, 1)
		rt.Data[0] = rtv
	}

	this.Data["json"] = &rt
	this.ServeJSON()
}

func (this *BookfulController) Searcher() {
  var rt models.Result
  var query = this.GetString(":query")

  err, rtv := services.Searcher(query)
	if err != nil {
		rt.Msg = "o_o"
		beego.Info(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		rt.Msg = "^_^"
		rt.Data = make([]models.Recs, 1)
		rt.Data[0] = rtv
	}

  this.Data["json"] = &rt
  this.ServeJSON()
}

func (this *BookfulController) SearchBookful() {
  var rt models.Result
  var query = this.GetString(":query")

  err, rtv := services.SearchBookful(query)
  if err != nil {
    rt.Msg = "o_o"
    beego.Info(err)
    this.Ctx.ResponseWriter.WriteHeader(500)
  } else {
    rt.Msg = "^_^"
    rt.Data = make([]models.Recs, 1)
    rt.Data[0] = rtv
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
