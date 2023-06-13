// Code generated by qtc from "topic_add.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/admin/topic_add.qtpl:1
package admin

//line views/admin/topic_add.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/admin/topic_add.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/admin/topic_add.qtpl:1
func (p *TopicAdd) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/admin/topic_add.qtpl:1
	qw422016.N().S(`

<div class="index">
    `)
//line views/admin/topic_add.qtpl:4
	if p.PageName == "admin_topic_review" && len(p.DefaultTopic.Title) == 0 {
//line views/admin/topic_add.qtpl:4
		qw422016.N().S(`
    <h1>没有 `)
//line views/admin/topic_add.qtpl:5
		qw422016.E().S(p.Title)
//line views/admin/topic_add.qtpl:5
		qw422016.N().S(`</h1>
    `)
//line views/admin/topic_add.qtpl:6
	} else {
//line views/admin/topic_add.qtpl:6
		qw422016.N().S(`
    <h1>
    `)
//line views/admin/topic_add.qtpl:8
		qw422016.E().S(p.Title)
//line views/admin/topic_add.qtpl:8
		qw422016.N().S(`
    `)
//line views/admin/topic_add.qtpl:9
		if len(p.DefaultTopic.Title) > 0 && p.PageName == "admin_topic_edit" {
//line views/admin/topic_add.qtpl:9
			qw422016.N().S(`
    <a href="/admin/topic/edit?id=`)
//line views/admin/topic_add.qtpl:10
			qw422016.N().DUL(p.DefaultTopic.ID)
//line views/admin/topic_add.qtpl:10
			qw422016.N().S(`&del=1">删除</a>
    `)
//line views/admin/topic_add.qtpl:11
		}
//line views/admin/topic_add.qtpl:11
		qw422016.N().S(`
    </h1>
    <p>`)
//line views/admin/topic_add.qtpl:13
		qw422016.E().S(p.DefaultTopic.Tags)
//line views/admin/topic_add.qtpl:13
		qw422016.N().S(`</p>
    <form class="pure-form" action="" method="post" onsubmit="form_post();return false;">
        <fieldset class="pure-group">
            <select id="select-nid">
                `)
//line views/admin/topic_add.qtpl:17
		for _, item := range p.NodeLst {
//line views/admin/topic_add.qtpl:17
			qw422016.N().S(`
                <option value="`)
//line views/admin/topic_add.qtpl:18
			qw422016.N().DUL(item.ID)
//line views/admin/topic_add.qtpl:18
			qw422016.N().S(`" `)
//line views/admin/topic_add.qtpl:18
			if item.ID == p.DefaultTopic.NodeId {
//line views/admin/topic_add.qtpl:18
				qw422016.N().S(`selected="selected"`)
//line views/admin/topic_add.qtpl:18
			}
//line views/admin/topic_add.qtpl:18
			qw422016.N().S(`>`)
//line views/admin/topic_add.qtpl:18
			qw422016.E().S(item.Name)
//line views/admin/topic_add.qtpl:18
			qw422016.N().S(`</option>
                `)
//line views/admin/topic_add.qtpl:19
		}
//line views/admin/topic_add.qtpl:19
		qw422016.N().S(`
            </select>
            <input id="id-title" type="text" value="`)
//line views/admin/topic_add.qtpl:21
		qw422016.E().S(p.DefaultTopic.Title)
//line views/admin/topic_add.qtpl:21
		qw422016.N().S(`" class="pure-input-1" placeholder="* 标题 MaxLen `)
//line views/admin/topic_add.qtpl:21
		qw422016.N().D(p.SiteCf.TitleMaxLen)
//line views/admin/topic_add.qtpl:21
		qw422016.N().S(`" autocomplete="off" />
            <textarea id="id-content" class="pure-input-1 topic-con-input" placeholder="* 内容 MaxLen `)
//line views/admin/topic_add.qtpl:22
		qw422016.N().D(p.SiteCf.TopicConMaxLen)
//line views/admin/topic_add.qtpl:22
		qw422016.N().S(`">`)
//line views/admin/topic_add.qtpl:22
		qw422016.N().S(p.DefaultTopic.Content)
//line views/admin/topic_add.qtpl:22
		qw422016.N().S(`</textarea>
            `)
//line views/admin/topic_add.qtpl:23
		if p.CurrentUser.Flag >= 99 {
//line views/admin/topic_add.qtpl:23
			qw422016.N().S(`
            <select id="select-uid">
                `)
//line views/admin/topic_add.qtpl:25
			for _, item := range p.UserLst {
//line views/admin/topic_add.qtpl:25
				qw422016.N().S(`
                <option value="`)
//line views/admin/topic_add.qtpl:26
				qw422016.N().DUL(item.ID)
//line views/admin/topic_add.qtpl:26
				qw422016.N().S(`" `)
//line views/admin/topic_add.qtpl:26
				if item.ID == p.DefaultUser.ID {
//line views/admin/topic_add.qtpl:26
					qw422016.N().S(`selected="selected"`)
//line views/admin/topic_add.qtpl:26
				}
//line views/admin/topic_add.qtpl:26
				qw422016.N().S(`>`)
//line views/admin/topic_add.qtpl:26
				qw422016.E().S(item.Name)
//line views/admin/topic_add.qtpl:26
				qw422016.N().S(`</option>
                `)
//line views/admin/topic_add.qtpl:27
			}
//line views/admin/topic_add.qtpl:27
			qw422016.N().S(`
            </select>
            <input id="id-addtime" type="text" value="`)
//line views/admin/topic_add.qtpl:29
			qw422016.N().DL(p.DefaultTopic.AddTime)
//line views/admin/topic_add.qtpl:29
			qw422016.N().S(`" class="pure-input-1" placeholder="发表的时间戳" />
            `)
//line views/admin/topic_add.qtpl:30
		} else {
//line views/admin/topic_add.qtpl:30
			qw422016.N().S(`
            <input type="hidden" id="select-uid" value="`)
//line views/admin/topic_add.qtpl:31
			qw422016.N().DUL(p.DefaultTopic.UserId)
//line views/admin/topic_add.qtpl:31
			qw422016.N().S(`">
            <input type="hidden" id="id-addtime" value="`)
//line views/admin/topic_add.qtpl:32
			qw422016.N().DL(p.DefaultTopic.AddTime)
//line views/admin/topic_add.qtpl:32
			qw422016.N().S(`">
            `)
//line views/admin/topic_add.qtpl:33
		}
//line views/admin/topic_add.qtpl:33
		qw422016.N().S(`
        </fieldset>
        <div id="id-msg"></div>
        <div class="fleft pure-button-group">
            <input id="btn-preview" type="button" value="预览" name="submit" class="pure-button button-success" />
            <input id="btn-submit" type="submit" value="发表" name="submit" class="pure-button pure-button-primary" />
            <input id="fileUpload" type="file" accept="image/*" onChange="uploadFile()" class="pure-button" name="fileUpload" style="font-size: .8334em;width: 95px;" />
            `)
//line views/admin/topic_add.qtpl:40
		if p.PageName == "admin_topic_review" {
//line views/admin/topic_add.qtpl:40
			qw422016.N().S(`
            <a href="?act=del" class="pure-button button-warning fr">直接删除</a>
            `)
//line views/admin/topic_add.qtpl:42
		}
//line views/admin/topic_add.qtpl:42
		qw422016.N().S(`
        </div>
        <div class="c"></div>

        <div id="id-preview" class="topic-content markdown-body"></div>
    </form>

    <script>
        let nodeEle = document.getElementById("select-nid");
        let titleEle = document.getElementById("id-title");
        let conEle = document.getElementById("id-content");
        let btnReviewEle = document.getElementById("btn-preview");
        let submitEle = document.getElementById("btn-submit");
        let msgEle = document.getElementById("id-msg");
        let addTimeEle = document.getElementById("id-addtime");
        let userIdEle = document.getElementById("select-uid");
        let reviewEle = document.getElementById("id-preview");

        btnReviewEle.addEventListener("click", function(){
            let con = conEle.value.trim();
            let title = titleEle.value.trim();
            if (con === "") {
                conEle.focus();
                return
            }

            btnReviewEle.setAttribute("disabled", "disabled");

            postAjax("/content/preview", JSON.stringify({Act: "topicPreview", Title: title, Content: con}), function(data){
                var obj = JSON.parse(data)
                //console.log(obj);
                if(obj.Code === 200) {
                    msgEle.style.display = "none";
                    reviewEle.innerHTML = obj.Html;
                    reviewEle.style.display = "block";
                }else{
                    reviewEle.innerHTML = "";
                    reviewEle.style.display = "none";
                    msgEle.innerText = obj.Msg;
                }
                btnReviewEle.removeAttribute('disabled');
            });
        }, false);

        function form_post(){
            let title = titleEle.value.trim();
            let con = conEle.value.trim();

            if (title === "") {
                titleEle.focus();
                return false;
            }

            if (con === "") {
                conEle.focus();
                return false;
            }

            reviewEle.innerHTML = "";
            reviewEle.style.display = "none";



            submitEle.setAttribute("disabled", "disabled");
            postAjax("/admin/topic/add", JSON.stringify({"Act": "submit", "ID": `)
//line views/admin/topic_add.qtpl:106
		qw422016.N().DUL(p.DefaultTopic.ID)
//line views/admin/topic_add.qtpl:106
		qw422016.N().S(`, "NodeId": parseInt(nodeEle.value, 10), "Title": title, "Content": con, "UserId": parseInt(userIdEle.value, 10), "AddTime": parseInt(addTimeEle.value.trim(), 10)}), function(data){
                var obj = JSON.parse(data)
                //console.log(obj);
                if(obj.Code === 200) {
                    msgEle.style.display = "none";
                    `)
//line views/admin/topic_add.qtpl:111
		if p.GoBack {
//line views/admin/topic_add.qtpl:111
			qw422016.N().S(`
                    window.location.href = "/t/`)
//line views/admin/topic_add.qtpl:112
			qw422016.N().DUL(p.DefaultTopic.ID)
//line views/admin/topic_add.qtpl:112
			qw422016.N().S(`";
                    return
                    `)
//line views/admin/topic_add.qtpl:114
		}
//line views/admin/topic_add.qtpl:114
		qw422016.N().S(`
                    `)
//line views/admin/topic_add.qtpl:115
		if p.PageName == "admin_topic_review" {
//line views/admin/topic_add.qtpl:115
			qw422016.N().S(`
                    window.location.href = "/admin/topic/review";
                    `)
//line views/admin/topic_add.qtpl:117
		} else {
//line views/admin/topic_add.qtpl:117
			qw422016.N().S(`
                    if(data.Tid > 0){
                        window.location.href = "/t/"+data.Tid;
                    }else{
                        window.location.href = "/admin/my/topic";
                    }
                    `)
//line views/admin/topic_add.qtpl:123
		}
//line views/admin/topic_add.qtpl:123
		qw422016.N().S(`
                    return false;
                } else if(obj.Code === 201){
                    msgEle.style.display = "block";
                    msgEle.innerText = obj.Msg;
                    titleEle.value = "";
                    conEle.value = "";

                    window.location.href = "/member/`)
//line views/admin/topic_add.qtpl:131
		qw422016.N().DUL(p.CurrentUser.ID)
//line views/admin/topic_add.qtpl:131
		qw422016.N().S(`";
                    return false;
                }else{
                    msgEle.style.display = "block";
                    msgEle.innerText = obj.Msg;
                }
                submitEle.removeAttribute('disabled');
            });

            return false;
        }

        document.addEventListener('paste', function (evt) {
            var url = "/file/upload";
            var items = evt.clipboardData && evt.clipboardData.items;
            var file = null;
            if(items && items.length) {
                for(var i=0; i!==items.length; i++) {
                    if(items[i].type.indexOf('image') !== -1) {
                        file = items[i].getAsFile();
                        if(!!!file) {
                            continue;
                        }

                        // upload file object.
                        var form = new FormData();
                        form.append('file', file);

                        postAjax("/file/upload", form, function(data){
                            let obj = JSON.parse(data)
                            //console.log(obj);
                            if(obj.Code === 200) {
                                let img_url = "\n![](" + obj.Url + ")\n";
                                let pos = conEle.selectionStart;
                                let con = conEle.value;
                                conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                            }else{
                                console.warn(obj.Msg);
                            }
                        });
                    }
                }
            }
        });
        function uploadFile() {
            let form = new FormData();
            form.append("file", fileUpload.files[0]);
            postAjax("/file/upload", form, function(data){
                let obj = JSON.parse(data)
                if(obj.Code === 200) {
                    let img_url = "\n![](" + obj.Url + ")\n";
                    let pos = conEle.selectionStart;
                    let con = conEle.value;
                    conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                }else{
                    console.warn(obj.Msg);
                }
            });
        }
    </script>

    `)
//line views/admin/topic_add.qtpl:192
	}
//line views/admin/topic_add.qtpl:192
	qw422016.N().S(`

</div>

`)
//line views/admin/topic_add.qtpl:196
}

//line views/admin/topic_add.qtpl:196
func (p *TopicAdd) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/admin/topic_add.qtpl:196
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin/topic_add.qtpl:196
	p.StreamMainBody(qw422016)
//line views/admin/topic_add.qtpl:196
	qt422016.ReleaseWriter(qw422016)
//line views/admin/topic_add.qtpl:196
}

//line views/admin/topic_add.qtpl:196
func (p *TopicAdd) MainBody() string {
//line views/admin/topic_add.qtpl:196
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin/topic_add.qtpl:196
	p.WriteMainBody(qb422016)
//line views/admin/topic_add.qtpl:196
	qs422016 := string(qb422016.B)
//line views/admin/topic_add.qtpl:196
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin/topic_add.qtpl:196
	return qs422016
//line views/admin/topic_add.qtpl:196
}
