// Code generated by qtc from "topic_add.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/ybs/topic_add.qtpl:1
package ybs

//line views/ybs/topic_add.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/ybs/topic_add.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/ybs/topic_add.qtpl:1
func (p *UserTopicAdd) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/ybs/topic_add.qtpl:1
	qw422016.N().S(`
<div class="write-box">
    <h1>`)
//line views/ybs/topic_add.qtpl:3
	qw422016.E().S(p.Title)
//line views/ybs/topic_add.qtpl:3
	qw422016.N().S(`</h1>
    <form class="pure-form" action="" method="post" onsubmit="form_post();return false;">
        <fieldset class="pure-group">
            <select id="select-nid">
                `)
//line views/ybs/topic_add.qtpl:7
	for _, item := range p.NodeLst {
//line views/ybs/topic_add.qtpl:7
		qw422016.N().S(`
                <option value="`)
//line views/ybs/topic_add.qtpl:8
		qw422016.N().DUL(item.ID)
//line views/ybs/topic_add.qtpl:8
		qw422016.N().S(`" `)
//line views/ybs/topic_add.qtpl:8
		if item.ID == p.DefaultNode.ID {
//line views/ybs/topic_add.qtpl:8
			qw422016.N().S(`selected="selected"`)
//line views/ybs/topic_add.qtpl:8
		}
//line views/ybs/topic_add.qtpl:8
		qw422016.N().S(`>`)
//line views/ybs/topic_add.qtpl:8
		qw422016.E().S(item.Name)
//line views/ybs/topic_add.qtpl:8
		qw422016.N().S(`</option>
                `)
//line views/ybs/topic_add.qtpl:9
	}
//line views/ybs/topic_add.qtpl:9
	qw422016.N().S(`
            </select>
            <input id="id-title" type="text" value="`)
//line views/ybs/topic_add.qtpl:11
	qw422016.E().S(p.DefaultTopic.Title)
//line views/ybs/topic_add.qtpl:11
	qw422016.N().S(`" class="pure-input-1" placeholder="* 标题 字符限制 `)
//line views/ybs/topic_add.qtpl:11
	qw422016.N().D(p.SiteCf.TitleMaxLen)
//line views/ybs/topic_add.qtpl:11
	qw422016.N().S(`" autocomplete="off" />
            <textarea id="id-content" class="pure-input-1 topic-con-input" placeholder="* 内容 字符限制 `)
//line views/ybs/topic_add.qtpl:12
	qw422016.N().D(p.SiteCf.TopicConMaxLen)
//line views/ybs/topic_add.qtpl:12
	qw422016.N().S(`">`)
//line views/ybs/topic_add.qtpl:12
	qw422016.N().S(p.DefaultTopic.Content)
//line views/ybs/topic_add.qtpl:12
	qw422016.N().S(`</textarea>
            `)
//line views/ybs/topic_add.qtpl:13
	if p.CurrentUser.Flag >= 99 {
//line views/ybs/topic_add.qtpl:13
		qw422016.N().S(`
            <div class="pure-g">
                <select id="select-uid">
                    `)
//line views/ybs/topic_add.qtpl:16
		for _, item := range p.UserLst {
//line views/ybs/topic_add.qtpl:16
			qw422016.N().S(`
                    <option value="`)
//line views/ybs/topic_add.qtpl:17
			qw422016.N().DUL(item.ID)
//line views/ybs/topic_add.qtpl:17
			qw422016.N().S(`" `)
//line views/ybs/topic_add.qtpl:17
			if item.ID == p.DefaultUser.ID {
//line views/ybs/topic_add.qtpl:17
				qw422016.N().S(`selected="selected"`)
//line views/ybs/topic_add.qtpl:17
			}
//line views/ybs/topic_add.qtpl:17
			qw422016.N().S(`>`)
//line views/ybs/topic_add.qtpl:17
			qw422016.E().S(item.Name)
//line views/ybs/topic_add.qtpl:17
			qw422016.N().S(`</option>
                    `)
//line views/ybs/topic_add.qtpl:18
		}
//line views/ybs/topic_add.qtpl:18
		qw422016.N().S(`
                </select>
                <input id="id-addtime" type="text" value="`)
//line views/ybs/topic_add.qtpl:20
		qw422016.N().DL(p.DefaultTopic.AddTime)
//line views/ybs/topic_add.qtpl:20
		qw422016.N().S(`" class="pure-u-1-6" placeholder="发表的时间戳" />
            </div>
            `)
//line views/ybs/topic_add.qtpl:22
	} else {
//line views/ybs/topic_add.qtpl:22
		qw422016.N().S(`
            <input type="hidden" id="select-uid" value="`)
//line views/ybs/topic_add.qtpl:23
		qw422016.N().DUL(p.DefaultTopic.UserId)
//line views/ybs/topic_add.qtpl:23
		qw422016.N().S(`">
            <input type="hidden" id="id-addtime" value="`)
//line views/ybs/topic_add.qtpl:24
		qw422016.N().DL(p.DefaultTopic.AddTime)
//line views/ybs/topic_add.qtpl:24
		qw422016.N().S(`">
            `)
//line views/ybs/topic_add.qtpl:25
	}
//line views/ybs/topic_add.qtpl:25
	qw422016.N().S(`
        </fieldset>
        <fieldset>
            <label for="ReadAuthed">
                <input type="checkbox" id="ReadAuthed" `)
//line views/ybs/topic_add.qtpl:29
	if p.DefaultTopic.ReadAuthed {
//line views/ybs/topic_add.qtpl:29
		qw422016.N().S(`checked`)
//line views/ybs/topic_add.qtpl:29
	}
//line views/ybs/topic_add.qtpl:29
	qw422016.N().S(`> 登录可浏览
            </label>
            <label for="ReadReply">
                <input type="checkbox" id="ReadReply" `)
//line views/ybs/topic_add.qtpl:32
	if p.DefaultTopic.ReadReply {
//line views/ybs/topic_add.qtpl:32
		qw422016.N().S(`checked`)
//line views/ybs/topic_add.qtpl:32
	}
//line views/ybs/topic_add.qtpl:32
	qw422016.N().S(`> 回复可浏览
            </label>
            <button id="insert-break" type="button" class="pure-button">插入分割线</button>
            `)
//line views/ybs/topic_add.qtpl:35
	if !p.SiteCf.UploadLimit || (p.SiteCf.UploadLimit && p.CurrentUser.Flag >= 99) {
//line views/ybs/topic_add.qtpl:35
		qw422016.N().S(`
            <input id="fileUpload" type="file" accept="image/*,video/*,audio/*" onChange="uploadFile()" class="pure-button" name="fileUpload" style="font-size: .8334em;width: 100px;" />
            `)
//line views/ybs/topic_add.qtpl:37
	}
//line views/ybs/topic_add.qtpl:37
	qw422016.N().S(`
        </fieldset>
        <div id="id-msg"></div>
        <div class="left pure-button-group">
            <input id="btn-preview" type="button" value="预览" name="submit" class="pure-button button-success" />
            <input id="btn-submit" type="submit" value="发表" name="submit" class="pure-button pure-button-primary" />
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
        let readAuthedEle = document.getElementById("ReadAuthed");
        let readReplyEle = document.getElementById("ReadReply");

        document.getElementById("insert-break").addEventListener('click', function (event) {
            let img_url = "\n`)
//line views/ybs/topic_add.qtpl:65
	qw422016.N().S(p.ReadMoreBreak)
//line views/ybs/topic_add.qtpl:65
	qw422016.N().S(`\n";
            let pos = conEle.selectionStart;
            let con = conEle.value;
            conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
        }, false);

        btnReviewEle.addEventListener('click', function (event) {
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
                btnReviewEle.removeAttribute('disabled')
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
            postAjax("/topic/add", JSON.stringify({"Act": "submit", "NodeId": parseInt(nodeEle.value, 10), "Title": title, "Content": con, "UserId": parseInt(userIdEle.value, 10), "AddTime": parseInt(addTimeEle.value, 10), "ReadAuthed": readAuthedEle.checked, "ReadReply": readReplyEle.checked}), function(data){
                var obj = JSON.parse(data)
                //console.log(obj);
                if(obj.Code === 200) {
                    msgEle.style.display = "none";
                    if(obj.Tid > 0){
                        window.location.href = "/t/"+obj.Tid;
                    }else{
                        window.location.href = "/my/topic";
                    }
                } else if(obj.Code === 201){
                    msgEle.style.display = "block";
                    msgEle.innerText = obj.Msg;
                    titleEle.value = "";
                    conEle.value = "";

                    window.location.href = "/member/`)
//line views/ybs/topic_add.qtpl:131
	qw422016.N().DUL(p.CurrentUser.ID)
//line views/ybs/topic_add.qtpl:131
	qw422016.N().S(`";
                    return;
                }else{
                    msgEle.style.display = "block";
                    msgEle.innerText = obj.Msg;
                }
                submitEle.removeAttribute('disabled');
            });

            return false;
        }

        `)
//line views/ybs/topic_add.qtpl:143
	if !p.SiteCf.UploadLimit || (p.SiteCf.UploadLimit && p.CurrentUser.Flag >= 99) {
//line views/ybs/topic_add.qtpl:143
		qw422016.N().S(`
        document.addEventListener('paste', function (evt) {
            var url = "/file/upload";
            var items = evt.clipboardData && evt.clipboardData.items;
            var file = null;
            if(items && items.length) {
                for(var i=0; i!==items.length; i++) {
                    var iType = items[i].type;
                    if(iType.indexOf('image') !== -1 || iType.indexOf('video') !== -1 || iType.indexOf('audio') !== -1) {
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
                                let img_url = "\n" + s2tag(obj.Url, `)
//line views/ybs/topic_add.qtpl:165
		qw422016.E().V(p.SiteCf.AutoDecodeMp4)
//line views/ybs/topic_add.qtpl:165
		qw422016.N().S(`) + "\n";
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
                    let img_url = "\n" + s2tag(obj.Url, `)
//line views/ybs/topic_add.qtpl:184
		qw422016.E().V(p.SiteCf.AutoDecodeMp4)
//line views/ybs/topic_add.qtpl:184
		qw422016.N().S(`) + "\n";
                    let pos = conEle.selectionStart;
                    let con = conEle.value;
                    conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                }else{
                    console.warn(obj.Msg);
                }
            });
        }
        `)
//line views/ybs/topic_add.qtpl:193
	}
//line views/ybs/topic_add.qtpl:193
	qw422016.N().S(`

    </script>

</div>

`)
//line views/ybs/topic_add.qtpl:199
}

//line views/ybs/topic_add.qtpl:199
func (p *UserTopicAdd) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/ybs/topic_add.qtpl:199
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/ybs/topic_add.qtpl:199
	p.StreamMainBody(qw422016)
//line views/ybs/topic_add.qtpl:199
	qt422016.ReleaseWriter(qw422016)
//line views/ybs/topic_add.qtpl:199
}

//line views/ybs/topic_add.qtpl:199
func (p *UserTopicAdd) MainBody() string {
//line views/ybs/topic_add.qtpl:199
	qb422016 := qt422016.AcquireByteBuffer()
//line views/ybs/topic_add.qtpl:199
	p.WriteMainBody(qb422016)
//line views/ybs/topic_add.qtpl:199
	qs422016 := string(qb422016.B)
//line views/ybs/topic_add.qtpl:199
	qt422016.ReleaseByteBuffer(qb422016)
//line views/ybs/topic_add.qtpl:199
	return qs422016
//line views/ybs/topic_add.qtpl:199
}
