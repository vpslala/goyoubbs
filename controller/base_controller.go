package controller

import (
	"errors"
	"github.com/ego008/goutils/json"
	"github.com/ego008/sdb"
	"github.com/valyala/fasthttp"
	"goyoubbs/model"
	"goyoubbs/util"
	"net"
	"strconv"
	"strings"
	"time"
)

var rangeTopicLst []model.TopicLi // 边栏显示最近被浏览的文章

type (
	BaseHandler struct {
		App *model.Application
	}
)

func ReadUserIP(ctx *fasthttp.RequestCtx) string {
	clientIp := string(ctx.Request.Header.Peek("X-Real-Ip"))
	netIP := net.ParseIP(clientIp)
	if netIP != nil {
		return clientIp
	}

	ips := string(ctx.Request.Header.Peek(fasthttp.HeaderXForwardedFor))
	splitIps := util.StringSplit(ips, ",")
	for _, ip := range splitIps {
		netIP = net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	ips = string(ctx.Request.Header.Peek("X-FORWARDED-FOR"))
	splitIps = util.StringSplit(ips, ",")
	for _, ip := range splitIps {
		netIP = net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String())
	if err != nil {
		return ""
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	return ""
}

func (h *BaseHandler) CurrentUser(ctx *fasthttp.RequestCtx) (*model.User, error) {
	user := &model.User{}
	ssValue := h.GetCookie(ctx, "SessionID")
	if len(ssValue) == 0 {
		return user, errors.New("SessionID cookie not found ")
	}
	index := strings.Index(ssValue, ":")
	if index == -1 {
		return user, errors.New("SessionID cookie not found ")
	}
	uId, _ := strconv.ParseUint(ssValue[:index], 10, 64)
	if uId == 0 {
		return user, errors.New("UserID is 0 ")
	}
	var ok bool
	model.UserMapMux.RLock()
	user, ok = model.UserMap[uId]
	model.UserMapMux.RUnlock()
	if !ok {
		rs := h.App.Db.Hget(model.UserTbName, sdb.I2b(uId))
		if rs.OK() {
			_ = json.Unmarshal(rs.Data[0], &user)
			if user.ID > 0 {
				model.UserMapMux.Lock()
				model.UserMap[uId] = user
				model.UserMapMux.Unlock()
			}
		}
	}
	if user != nil {
		_ = h.SetCookie(ctx, "SessionID", ssValue, 365)
		return user, nil
	}

	return &model.User{}, errors.New("user not found")
}

func (h *BaseHandler) SetCookie(ctx *fasthttp.RequestCtx, name, value string, days int) error {
	name = name + h.App.Cf.Main.Addr
	encoded, err := h.App.Sc.Encode(name, value)
	if err != nil {
		return err
	}

	var c fasthttp.Cookie
	c.SetKey(name)
	c.SetValue(encoded)
	c.SetPath("/")
	c.SetSecure(false)
	c.SetHTTPOnly(true)
	c.SetSameSite(fasthttp.CookieSameSiteStrictMode)
	c.SetExpire(time.Now().UTC().AddDate(0, 0, days))
	ctx.Response.Header.SetCookie(&c)

	return err
}

func (h *BaseHandler) GetCookie(ctx *fasthttp.RequestCtx, name string) string {
	name = name + h.App.Cf.Main.Addr
	if cookieByte := ctx.Request.Header.Cookie(name); len(cookieByte) > 0 {
		var value string
		if err := h.App.Sc.Decode(name, sdb.B2s(cookieByte), &value); err == nil {
			return value
		}
	}
	return ""
}

func (h *BaseHandler) DelCookie(ctx *fasthttp.RequestCtx, name string) {
	if len(name) > 0 {
		name = name + h.App.Cf.Main.Addr
		ctx.Response.Header.DelClientCookie(name)
	}
}
