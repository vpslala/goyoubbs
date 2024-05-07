package controller

import (
	"github.com/VictoriaMetrics/fastcache"
	"github.com/ego008/goutils/json"
	"github.com/gorilla/securecookie"
	"github.com/valyala/fasthttp"
	"goyoubbs/model"
	"goyoubbs/util"
	"goyoubbs/views/admin"
	"strconv"
	"strings"
	"time"
)

func (h *BaseHandler) AdminSiteConfigPage(ctx *fasthttp.RequestCtx) {
	curUser, _ := h.CurrentUser(ctx)
	if curUser.Flag < model.FlagAdmin {
		ctx.Redirect(h.App.Cf.Site.MainDomain+"/login", 302)
		return
	}

	scf := h.App.Cf.Site

	evn := &admin.SiteConfig{}
	evn.CurrentUser = *curUser
	evn.SiteCf = scf
	evn.Title = "网站设置"
	evn.PageName = "admin_site_setting"

	evn.NodeLst = model.NodeGetAll(h.App.Mc, h.App.Db)

	//
	siteConf := model.SiteConf{}
	model.SiteConfLoad(&siteConf, h.App.Db)
	evn.SiteConf = siteConf

	evn.HasMsg = model.MsgCheckHasOne(h.App.Db, curUser.ID)
	evn.HasTopicReview = model.CheckHasTopic2Review(h.App.Db)
	evn.HasReplyReview = model.CheckHasComment2Review(h.App.Db)

	ctx.SetContentType("text/html; charset=utf-8")
	admin.WritePageTemplate(ctx, evn)
}

func b2int(b []byte, df int) int {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		return df
	}
	return i
}

func b2bool(b []byte, df bool) bool {
	i, err := strconv.ParseBool(string(b))
	if err != nil {
		return df
	}
	return i
}

func (h *BaseHandler) AdminSiteConfigPost(ctx *fasthttp.RequestCtx) {
	curUser, _ := h.CurrentUser(ctx)
	if curUser.Flag < model.FlagAdmin {
		ctx.Redirect(h.App.Cf.Site.MainDomain+"/login", 302)
		return
	}

	obj := model.SiteConf{}
	model.SiteConfLoad(&obj, h.App.Db)

	obj.Name = string(ctx.FormValue("Name"))
	obj.Desc = string(ctx.FormValue("Desc"))
	obj.MainDomain = strings.TrimSuffix(string(ctx.FormValue("MainDomain")), "/")
	obj.HeaderPartCon = string(ctx.FormValue("HeaderPartCon"))
	obj.GoogleAutoAdJs = string(ctx.FormValue("GoogleAutoAdJs"))
	obj.FooterPartHtml = string(ctx.FormValue("FooterPartHtml"))

	obj.TimeZone = b2int(ctx.FormValue("TimeZone"), 8)
	if obj.TimeZone < -12 || obj.TimeZone > 12 {
		obj.TimeZone = 8
	}
	model.TimeOffSet = time.Duration(obj.TimeZone) * time.Hour

	obj.PageShowNum = b2int(ctx.FormValue("PageShowNum"), 32)
	obj.TopRateNum = b2int(ctx.FormValue("TopRateNum"), 10)
	obj.RecentCommentNum = b2int(ctx.FormValue("RecentCommentNum"), 10)
	obj.TitleMaxLen = b2int(ctx.FormValue("TitleMaxLen"), 110)
	obj.TopicConMaxLen = b2int(ctx.FormValue("TopicConMaxLen"), 12000)
	obj.CommentConMaxLen = b2int(ctx.FormValue("CommentConMaxLen"), 5000)

	obj.AutoDataBackup = b2bool(ctx.FormValue("AutoDataBackup"), false)
	obj.DataBackupDir = strings.TrimSuffix(string(ctx.FormValue("DataBackupDir")), "/")
	if obj.UploadDir == "" {
		obj.UploadDir = "data_backup"
	}

	obj.Authorized = b2bool(ctx.FormValue("Authorized"), false)
	obj.AllowNameReg = b2bool(ctx.FormValue("AllowNameReg"), true)
	obj.RegReview = b2bool(ctx.FormValue("RegReview"), false)
	obj.CloseReg = b2bool(ctx.FormValue("CloseReg"), false)
	obj.CloseReply = b2bool(ctx.FormValue("CloseReply"), false)
	obj.PostReview = b2bool(ctx.FormValue("PostReview"), false)

	obj.ResetCookieKey = b2bool(ctx.FormValue("ResetCookieKey"), false)
	if obj.ResetCookieKey {
		hashKey := securecookie.GenerateRandomKey(64)
		blockKey := securecookie.GenerateRandomKey(32)
		_ = h.App.Db.Hmset(model.KeyValueTb, []byte("hashKey"), hashKey, []byte("blockKey"), blockKey)
		h.App.Sc = securecookie.New(hashKey, blockKey)
	}

	obj.AutoDecodeMp4 = b2bool(ctx.FormValue("AutoDecodeMp4"), false)
	// check ffmpeg exist
	if obj.AutoDecodeMp4 {
		obj.AutoDecodeMp4 = util.CmdExists("ffmpeg")
	}

	obj.GetTagApi = string(ctx.FormValue("GetTagApi"))

	obj.UploadLimit = b2bool(ctx.FormValue("UploadLimit"), false)

	var reloadRouter bool
	oldUploadDir := obj.UploadDir
	obj.UploadDir = strings.TrimSuffix(string(ctx.FormValue("UploadDir")), "/")
	if obj.UploadDir == "" {
		obj.UploadDir = "upload"
	}
	if obj.UploadDir != oldUploadDir {
		reloadRouter = true
	}

	obj.UploadMaxSize = b2int(ctx.FormValue("UploadMaxSize"), 20)
	if obj.UploadMaxSize < 1 {
		obj.UploadMaxSize = 1
	}
	obj.UploadMaxSizeByte = int64(obj.UploadMaxSize) << 20

	oldCachedSize := obj.CachedSize
	obj.CachedSize = b2int(ctx.FormValue("CachedSize"), 1)
	if obj.CachedSize < 1 {
		obj.CachedSize = 1
	}

	obj.RateLimitDay = b2int(ctx.FormValue("RateLimitDay"), 0)
	model.RateLimitDay = obj.RateLimitDay
	obj.RateLimitHour = b2int(ctx.FormValue("RateLimitHour"), 0)
	model.RateLimitHour = obj.RateLimitHour

	oldSaveTopicIcon := obj.SaveTopicIcon
	obj.SaveTopicIcon = b2bool(ctx.FormValue("SaveTopicIcon"), false)

	obj.SaveImg2db = b2bool(ctx.FormValue("SaveImg2db"), false)
	obj.RemotePostPw = string(ctx.FormValue("RemotePostPw"))
	obj.QQClientID = string(ctx.FormValue("QQClientID"))
	obj.QQClientSecret = string(ctx.FormValue("QQClientSecret"))
	obj.WeiboClientID = string(ctx.FormValue("WeiboClientID"))
	obj.WeiboClientSecret = string(ctx.FormValue("WeiboClientSecret"))
	obj.GithubClientID = string(ctx.FormValue("GithubClientID"))
	obj.GithubClientSecret = string(ctx.FormValue("GithubClientSecret"))
	obj.SendEmail = b2bool(ctx.FormValue("SendEmail"), false)
	obj.SmtpHost = string(ctx.FormValue("SmtpHost"))
	obj.SmtpPort = b2int(ctx.FormValue("SmtpPort"), 465)
	obj.SmtpEmail = string(ctx.FormValue("SmtpEmail"))
	obj.SmtpPassword = string(ctx.FormValue("SmtpPassword"))
	obj.SendToEmail = string(ctx.FormValue("SendToEmail"))

	jb, _ := json.Marshal(obj)
	_ = h.App.Db.Hset(model.KeyValueTb, []byte("site_config"), jb)

	// in old conf
	obj.IsDevMod = h.App.Cf.Site.IsDevMod
	obj.SelfHash = h.App.Cf.Site.SelfHash

	h.App.Cf.Site = &obj

	if reloadRouter {
		// 在保存后 reload
		RouterReload(h.App)
	}

	// 清空缓存
	h.App.Mc.Reset()
	if obj.CachedSize != oldCachedSize {
		h.App.Mc = fastcache.New(obj.CachedSize * 1024 * 1024)
	}

	// 清除帖子九宫格图片
	if (oldSaveTopicIcon != obj.SaveTopicIcon) && !obj.SaveTopicIcon {
		if h.App.Db.Hscan("topic_icon", nil, 1).OK() {
			_ = h.App.Db.HdelBucket("topic_icon")
		}
	}

	ctx.Redirect(h.App.Cf.Site.MainDomain+"/admin/site/conf", 302)
}
