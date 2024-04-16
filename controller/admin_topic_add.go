package controller

import (
	"fmt"
	"github.com/ego008/goutils/json"
	"github.com/ego008/sdb"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/valyala/fasthttp"
	"goyoubbs/model"
	"goyoubbs/util"
	"goyoubbs/views/admin"
	"log"
	"strconv"
	"strings"
)

func (h *BaseHandler) AdminTopicAddPage(ctx *fasthttp.RequestCtx) {
	curUser, _ := h.CurrentUser(ctx)
	if curUser.Flag < model.FlagAuthor {
		ctx.Redirect(h.App.Cf.Site.MainDomain+"/login", 302)
		return
	}

	db := h.App.Db
	scf := h.App.Cf.Site

	evn := &admin.TopicAdd{}
	evn.CurrentUser = curUser
	evn.SiteCf = scf
	evn.Title = "发表文章"
	evn.PageName = "topic_input"

	evn.ReadMoreBreak = model.ReadMoreBreak
	evn.DefaultTopic = model.Topic{
		NodeId: 1,
		UserId: curUser.ID,
	}
	evn.DefaultUser = curUser
	evn.HasMsg = model.MsgCheckHasOne(db, curUser.ID)
	// evn.DefaultNode, _ = model.NodeGetById(h.App.Db, evn.DefaultTopic.NodeId)
	// evn.DefaultNode = model.Node{}
	evn.NodeLst = model.NodeGetAll(h.App.Mc, db)
	evn.HasMsg = model.MsgCheckHasOne(db, curUser.ID)
	if curUser.Flag >= model.FlagAdmin {
		evn.UserLst = model.UserGetAllAdmin(db)
		evn.HasTopicReview = model.CheckHasTopic2Review(db)
		evn.HasReplyReview = model.CheckHasComment2Review(db)
	}

	ctx.SetContentType("text/html; charset=utf-8")
	admin.WritePageTemplate(ctx, evn)
}

// AdminTopicAddPost 发表、审核、编辑 公用
func (h *BaseHandler) AdminTopicAddPost(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=UTF-8")

	curUser, _ := h.CurrentUser(ctx)

	if curUser.Flag < model.FlagAuthor {
		_, _ = ctx.WriteString(`{"Code":401,"Msg":"请先登录"}`)
		return
	}

	var rec model.TopicRecForm
	err := util.Bind(ctx, util.JSON, &rec)
	if err != nil {
		_, _ = ctx.WriteString(`{"Code":400,"Msg":"unable to read body"}`)
		return
	}

	db := h.App.Db
	scf := h.App.Cf.Site

	rec.Title = strings.TrimSpace(rec.Title)
	rec.Content = strings.TrimSpace(rec.Content)

	titleLen := len(rec.Title)
	if titleLen == 0 {
		_, _ = ctx.WriteString(`{"Code":400,"Msg":"文章标题不能为空"}`)
		return
	} else if titleLen > scf.TitleMaxLen {
		msg := fmt.Sprintf(`{"Code":400,"Msg":"文章标题太长 %d > %d "}`, titleLen, scf.TitleMaxLen)
		_, _ = ctx.WriteString(msg)
		return
	}

	contentLen := len(rec.Content)
	if curUser.Flag < model.FlagAdmin && contentLen > scf.TopicConMaxLen {
		msg := fmt.Sprintf(`{"Code":400,"Msg":"文章内容太长 %d > %d "}`, contentLen, scf.TopicConMaxLen)
		_, _ = ctx.WriteString(msg)
		return
	}

	// edit
	var oldTopic model.Topic
	var isEdit bool
	if rec.ID > 0 {
		isEdit = true
	}

	// check title
	titleMd5 := fnv1a.HashString64(rec.Title)
	if rs := db.Hget("title_fnv1a", sdb.I2b(titleMd5)); rs.OK() {
		if rec.ID != sdb.B2i(rs.Bytes()) {
			_, _ = ctx.WriteString(`{"Code":400,"Msg":"相同的文章标题已存在，请修改"}`)
			return
		}
	}

	var topic model.Topic
	if isEdit {
		topic = model.TopicGetById(db, rec.ID)
		if topic.ID == 0 {
			_, _ = ctx.WriteString(`{"Code":400,"Msg":"该 id 帖子不存在"}`)
			return
		}
		oldTopic = topic
	}

	topic.NodeId = rec.NodeId
	topic.UserId = rec.UserId // curUser.ID
	topic.Title = rec.Title
	topic.Content = rec.Content
	topic.AddTime = rec.AddTime // util.GetCNTM()
	topic.ReadAuthed = rec.ReadAuthed
	topic.ReadReply = rec.ReadReply

	// may fix
	if topic.UserId == 0 {
		topic.UserId = curUser.ID
	}
	if topic.AddTime == 0 {
		topic.AddTime = util.GetCNTM(model.TimeOffSet)
	}
	topic.EditTime = topic.AddTime

	// 审核发帖删掉信息
	if !isEdit {
		if rec.AddTime > 0 {
			// 删掉管理员列表
			_ = db.Hdel(model.TopicReviewTbName, sdb.I2b(uint64(rec.AddTime)))
			// 删掉个人待审核列表
			_ = db.Hdel("review_topic:"+strconv.FormatUint(rec.UserId, 10), sdb.I2b(uint64(rec.AddTime)))
		}
	}

	type response struct {
		model.NormalRsp
		Tid uint64
	}

	rsp := response{}
	rsp.Code = 200

	// 编辑
	if isEdit {
		if curUser.Flag < model.FlagAdmin {
			_, _ = ctx.WriteString(`{"Code":403,"Msg":"权限限制"}`)
			return
		}
		if oldTopic.Title != topic.Title || oldTopic.Content != topic.Content {
			topic.EditTime = util.GetCNTM(model.TimeOffSet)
		}
		// 直接更新
		model.TopicSet(db, topic)
		// 分类、title 变化
		if oldTopic.NodeId != topic.NodeId {
			_ = db.Zset("topic_update:"+strconv.FormatUint(topic.NodeId, 10), sdb.I2b(topic.ID), uint64(topic.AddTime))
			_ = db.Zdel("topic_update:"+strconv.FormatUint(oldTopic.NodeId, 10), sdb.I2b(topic.ID))
			// 该分类的文章数
			_, _ = db.Hincr(model.NodeTopicNumTbName, sdb.I2b(topic.NodeId), 1)
			_, _ = db.Hincr(model.NodeTopicNumTbName, sdb.I2b(oldTopic.NodeId), -1)
			// 删除分类缓存
			h.App.Mc.Del([]byte("NodeGetAll"))
		}
		if oldTopic.Title != topic.Title {
			_ = db.Hdel("title_fnv1a", sdb.I2b(fnv1a.HashString64(oldTopic.Title)))
			_ = db.Hset("title_fnv1a", sdb.I2b(titleMd5), sdb.I2b(topic.ID))
			// 自动从标题里提取标签
			log.Println("scf.GetTagApi2", scf.GetTagApi, "h.App.Cf.Site2", h.App.Cf.Site.GetTagApi)
			if len(scf.GetTagApi) > 0 {
				_ = db.Hset("task_to_get_tag", sdb.I2b(topic.ID), sdb.S2b(topic.Title))
			}
		}
		rsp.Tid = topic.ID
		_ = json.NewEncoder(ctx).Encode(rsp)

		// 删除缓存
		h.App.Mc.Del([]byte("ContentFmt:" + strconv.FormatUint(topic.ID, 10)))

		return
	}

	// 以下是添加或审核
	if curUser.Flag < model.FlagAdmin && scf.PostReview {
		// 非管理员+开启审核
		// 把jb 内容暂存到审核列表
		jb, _ := json.Marshal(topic)
		// 给管理员看
		_ = db.Hset(model.TopicReviewTbName, sdb.I2b(uint64(topic.AddTime)), jb)
		// 把key 放到个人的列表
		_ = db.Hset("review_topic:"+strconv.FormatUint(topic.UserId, 10), sdb.I2b(uint64(topic.AddTime)), nil)
		rsp.Code = 200 // 201
		rsp.Msg = "* 您的帖子已经提交，系统开启了发帖审核，请耐心等管理员审核"
		_ = json.NewEncoder(ctx).Encode(rsp)
		return
	}

	// 直接保存
	topic = model.TopicAdd(h.App.Mc, db, topic)
	rsp.Tid = topic.ID

	// 自动从标题里提取标签
	log.Println("scf.GetTagApi", scf.GetTagApi, "h.App.Cf.Site", h.App.Cf.Site.GetTagApi)
	if len(scf.GetTagApi) > 0 {
		_ = db.Hset("task_to_get_tag", sdb.I2b(topic.ID), sdb.S2b(topic.Title))
	}

	// 记录标题md5
	_ = db.Hset("title_fnv1a", sdb.I2b(titleMd5), sdb.I2b(topic.ID))

	_ = json.NewEncoder(ctx).Encode(rsp)
}
