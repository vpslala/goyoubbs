// Code generated by qtc from "link.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/admin/link.qtpl:1
package admin

//line views/admin/link.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/admin/link.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/admin/link.qtpl:1
func (p *Link) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/admin/link.qtpl:1
	qw422016.N().S(`
<div class="index">
    <div class="markdown-body entry-content">
    <h1>`)
//line views/admin/link.qtpl:4
	qw422016.E().S(p.Title)
//line views/admin/link.qtpl:4
	qw422016.N().S(`</h1>

    <form action="" method="post" class="pure-form pure-form-stacked">
        <fieldset>
            <legend>`)
//line views/admin/link.qtpl:8
	qw422016.E().S(p.Act)
//line views/admin/link.qtpl:8
	qw422016.N().S(` 链接</legend>

            <p>Score: 显示排序，大排前</p>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-5">
                    <label for="Name">Name</label>
                    <input id="Name" name="Name" class="pure-u-23-24" type="text" value="`)
//line views/admin/link.qtpl:15
	qw422016.E().S(p.Link.Name)
//line views/admin/link.qtpl:15
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-1-5">
                    <label for="Score">Score</label>
                    <input id="Score" name="Score" class="pure-u-23-24" type="number" value="`)
//line views/admin/link.qtpl:20
	qw422016.N().D(p.Link.Score)
//line views/admin/link.qtpl:20
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-2-5">
                    <label for="Url">Url</label>
                    <input id="Url" name="Url" class="pure-u-23-24" type="text" value="`)
//line views/admin/link.qtpl:25
	qw422016.E().S(p.Link.Url)
//line views/admin/link.qtpl:25
	qw422016.N().S(`" required>
                </div>

            </div>

            <button type="submit" class="pure-button pure-button-primary">提交</button>
        </fieldset>
    </form>

    <h2>链接列表</h2>

        <table class="pure-table">
            <thead>
                <tr>
                    <th>#ID</th>
                    <th>Name</th>
                    <th>Score</th>
                    <th>URL</th>
                </tr>
            </thead>
            <tbody>
                `)
//line views/admin/link.qtpl:46
	for i, v := range p.LinkLst {
//line views/admin/link.qtpl:46
		qw422016.N().S(`
                <tr `)
//line views/admin/link.qtpl:47
		if i%2 == 0 {
//line views/admin/link.qtpl:47
			qw422016.N().S(`class="pure-table-odd"`)
//line views/admin/link.qtpl:47
		}
//line views/admin/link.qtpl:47
		qw422016.N().S(`>
                    <td>`)
//line views/admin/link.qtpl:48
		qw422016.N().DUL(v.ID)
//line views/admin/link.qtpl:48
		qw422016.N().S(`</td>
                    <td><a href="/admin/link?id=`)
//line views/admin/link.qtpl:49
		qw422016.N().DUL(v.ID)
//line views/admin/link.qtpl:49
		qw422016.N().S(`">`)
//line views/admin/link.qtpl:49
		qw422016.E().S(v.Name)
//line views/admin/link.qtpl:49
		qw422016.N().S(`</a></td>
                    <td>`)
//line views/admin/link.qtpl:50
		qw422016.N().D(v.Score)
//line views/admin/link.qtpl:50
		qw422016.N().S(`</td>
                    <td>`)
//line views/admin/link.qtpl:51
		qw422016.E().S(v.Url)
//line views/admin/link.qtpl:51
		qw422016.N().S(`</td>
                </tr>
                `)
//line views/admin/link.qtpl:53
	}
//line views/admin/link.qtpl:53
	qw422016.N().S(`
            </tbody>
        </table>

    </div>
</div>

`)
//line views/admin/link.qtpl:60
}

//line views/admin/link.qtpl:60
func (p *Link) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/admin/link.qtpl:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin/link.qtpl:60
	p.StreamMainBody(qw422016)
//line views/admin/link.qtpl:60
	qt422016.ReleaseWriter(qw422016)
//line views/admin/link.qtpl:60
}

//line views/admin/link.qtpl:60
func (p *Link) MainBody() string {
//line views/admin/link.qtpl:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin/link.qtpl:60
	p.WriteMainBody(qb422016)
//line views/admin/link.qtpl:60
	qs422016 := string(qb422016.B)
//line views/admin/link.qtpl:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin/link.qtpl:60
	return qs422016
//line views/admin/link.qtpl:60
}
