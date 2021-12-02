package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "deficonnect",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          true,
		CookieLifeTime:  3600,
	})
	go globalSessions.GC()
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Controller.Prepare()
	c.Layout = "layout.html"
	c.LayoutSections = map[string]string{}
}

func (c *MainController) Get() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "index"
	c.LayoutSections["Scripts"] = "index_scripts.html"
	c.TplName = "index-new.html"
}

func (c *MainController) Books() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "books"
	c.TplName = "books.html"
}

func (c *MainController) Contact() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "contact"
	c.TplName = "contact.html"
}

func (c *MainController) Author() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)


	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "author"
	c.TplName = "author.html"
}

func (c *MainController) Buy() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)


	c.Data["Debug"] = c.GetString("debug")
	c.Data["JsController"] = "buy"
	c.TplName = "buy.html"
}

