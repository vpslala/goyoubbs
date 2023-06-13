// Code generated by qtc from "user.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/admin/user.qtpl:1
package admin

//line views/admin/user.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/admin/user.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/admin/user.qtpl:1
func (p *User) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/admin/user.qtpl:1
	qw422016.N().S(`
<div class="index">
    <div class="markdown-body entry-content">
    <h1>`)
//line views/admin/user.qtpl:4
	qw422016.E().S(p.Title)
//line views/admin/user.qtpl:4
	qw422016.N().S(`</h1>
    <form action="" method="get" class="pure-form">
        <fieldset>
            <input type="text" name="q" class="pure-input-rounded" placeholder="搜索用户 id/name" />
        </fieldset>
    </form>
    <form action="" method="post" class="pure-form pure-form-stacked">
        <fieldset>
            <legend>`)
//line views/admin/user.qtpl:12
	qw422016.E().S(p.Act)
//line views/admin/user.qtpl:12
	qw422016.N().S(` 用户</legend>
            `)
//line views/admin/user.qtpl:13
	if p.User.ID > 0 {
//line views/admin/user.qtpl:13
		qw422016.N().S(`
            <img id="img" src="/static/avatar/`)
//line views/admin/user.qtpl:14
		qw422016.N().DUL(p.User.ID)
//line views/admin/user.qtpl:14
		qw422016.N().S(`.jpg" alt="`)
//line views/admin/user.qtpl:14
		qw422016.E().S(p.User.Name)
//line views/admin/user.qtpl:14
		qw422016.N().S(` avatar" onclick="document.getElementById('file-input').click();" title="点击更换头像" style="cursor: pointer;height: 119px;width: 119px;">
            <input id="file-input" type="file" name="file" accept="image/*" style="display: none;" />
            <script>
                const reader = new FileReader();
                const fileInput = document.getElementById("file-input");
                const img = document.getElementById("img");
                let file;

                reader.onload = e => {
                    img.src = e.target.result;
                }

                fileInput.addEventListener('change', e => {
                    const f = e.target.files[0];

                    let formData = new FormData();
                    formData.append("UserId", "`)
//line views/admin/user.qtpl:30
		qw422016.N().DUL(p.User.ID)
//line views/admin/user.qtpl:30
		qw422016.N().S(`");
                    formData.append("file", f);

                    postAjax("/user/avatar/upload", formData, function(data){
                        var obj = JSON.parse(data)
                        console.log(obj);
                        if(obj.Code === 200) {
                            reader.readAsDataURL(f);
                        }
                    });
                })
            </script>
            `)
//line views/admin/user.qtpl:42
	}
//line views/admin/user.qtpl:42
	qw422016.N().S(`

            <p>Flag:
            `)
//line views/admin/user.qtpl:45
	for _, item := range p.FlagLst {
//line views/admin/user.qtpl:45
		qw422016.N().S(`
            <a href="?flag=`)
//line views/admin/user.qtpl:46
		qw422016.N().D(item.Flag)
//line views/admin/user.qtpl:46
		qw422016.N().S(`">`)
//line views/admin/user.qtpl:46
		qw422016.E().S(item.Name)
//line views/admin/user.qtpl:46
		qw422016.N().S(`</a>,
            `)
//line views/admin/user.qtpl:47
	}
//line views/admin/user.qtpl:47
	qw422016.N().S(`
            </p>

            <div>
                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Name">登录名： </label>
                    <input id="Name" name="Name" class="pure-u-23-24" type="text" value="`)
//line views/admin/user.qtpl:53
	qw422016.E().S(p.User.Name)
//line views/admin/user.qtpl:53
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Password">密　码： </label>
                    <input id="Password" name="Password" class="pure-u-23-24" type="text" value="">
                </div>

                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="select-nid">权　限： </label>
                    <select id="select-nid" name="Flag">
                        `)
//line views/admin/user.qtpl:64
	for _, item := range p.FlagLst {
//line views/admin/user.qtpl:64
		qw422016.N().S(`
                        <option value="`)
//line views/admin/user.qtpl:65
		qw422016.N().D(item.Flag)
//line views/admin/user.qtpl:65
		qw422016.N().S(`" `)
//line views/admin/user.qtpl:65
		if item.Flag == p.User.Flag {
//line views/admin/user.qtpl:65
			qw422016.N().S(`selected="selected"`)
//line views/admin/user.qtpl:65
		}
//line views/admin/user.qtpl:65
		qw422016.N().S(`>`)
//line views/admin/user.qtpl:65
		qw422016.E().S(item.Name)
//line views/admin/user.qtpl:65
		qw422016.N().S(`</option>
                        `)
//line views/admin/user.qtpl:66
	}
//line views/admin/user.qtpl:66
	qw422016.N().S(`
                    </select>
                </div>

            </div>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-1">
                    <input name="Url" type="text" value="`)
//line views/admin/user.qtpl:74
	qw422016.E().S(p.User.Url)
//line views/admin/user.qtpl:74
	qw422016.N().S(`" class="pure-input-1" placeholder="URL http(s)://example.com" />
                    <textarea name="About" class="pure-input-1" placeholder="About...">`)
//line views/admin/user.qtpl:75
	qw422016.N().S(p.User.About)
//line views/admin/user.qtpl:75
	qw422016.N().S(`</textarea>
                </div>
            </div>


            <button type="submit" class="pure-button pure-button-primary">提交</button>
        </fieldset>
    </form>

    <h2>列表</h2>
    <ul>
        <li class="bot-line">
            ID - Name - Flag - Url - About
        </li>
        `)
//line views/admin/user.qtpl:89
	for _, v := range p.UserLst {
//line views/admin/user.qtpl:89
		qw422016.N().S(`
        <li class="bot-line">
            `)
//line views/admin/user.qtpl:91
		qw422016.N().DUL(v.ID)
//line views/admin/user.qtpl:91
		qw422016.N().S(` - <a href="/admin/user?id=`)
//line views/admin/user.qtpl:91
		qw422016.N().DUL(v.ID)
//line views/admin/user.qtpl:91
		qw422016.N().S(`">`)
//line views/admin/user.qtpl:91
		qw422016.E().S(v.Name)
//line views/admin/user.qtpl:91
		qw422016.N().S(`</a> - `)
//line views/admin/user.qtpl:91
		qw422016.N().D(v.Flag)
//line views/admin/user.qtpl:91
		qw422016.N().S(` - `)
//line views/admin/user.qtpl:91
		qw422016.E().S(v.Url)
//line views/admin/user.qtpl:91
		qw422016.N().S(` - `)
//line views/admin/user.qtpl:91
		qw422016.E().S(v.About)
//line views/admin/user.qtpl:91
		qw422016.N().S(`
        </li>
        `)
//line views/admin/user.qtpl:93
	}
//line views/admin/user.qtpl:93
	qw422016.N().S(`
    </ul>
</div>
</div>

`)
//line views/admin/user.qtpl:98
}

//line views/admin/user.qtpl:98
func (p *User) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/admin/user.qtpl:98
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin/user.qtpl:98
	p.StreamMainBody(qw422016)
//line views/admin/user.qtpl:98
	qt422016.ReleaseWriter(qw422016)
//line views/admin/user.qtpl:98
}

//line views/admin/user.qtpl:98
func (p *User) MainBody() string {
//line views/admin/user.qtpl:98
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin/user.qtpl:98
	p.WriteMainBody(qb422016)
//line views/admin/user.qtpl:98
	qs422016 := string(qb422016.B)
//line views/admin/user.qtpl:98
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin/user.qtpl:98
	return qs422016
//line views/admin/user.qtpl:98
}
