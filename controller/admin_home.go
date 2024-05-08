package controller

import (
	"github.com/valyala/fasthttp"
	"goyoubbs/model"
)

func (h *BaseHandler) AdminHomePage(ctx *fasthttp.RequestCtx) {
	curUser, _ := h.CurrentUser(ctx)
	if curUser.Flag < model.FlagAdmin {
		ctx.Redirect(h.App.Cf.Site.MainDomain+"/login", 302)
		return
	}

	ctx.Redirect(h.App.Cf.Site.MainDomain+"/admin/topic/add", 302)
}
