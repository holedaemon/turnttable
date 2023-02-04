// Code generated by qtc from "admin.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/web/templates/admin.qtpl:1
package templates

//line internal/web/templates/admin.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/web/templates/admin.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/web/templates/admin.qtpl:2
type AdminPage struct {
	BasePage
}

//line internal/web/templates/admin.qtpl:7
func (p *AdminPage) StreamTitle(qw422016 *qt422016.Writer) {
//line internal/web/templates/admin.qtpl:7
	qw422016.N().S(`Admin`)
//line internal/web/templates/admin.qtpl:7
}

//line internal/web/templates/admin.qtpl:7
func (p *AdminPage) WriteTitle(qq422016 qtio422016.Writer) {
//line internal/web/templates/admin.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/admin.qtpl:7
	p.StreamTitle(qw422016)
//line internal/web/templates/admin.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/admin.qtpl:7
}

//line internal/web/templates/admin.qtpl:7
func (p *AdminPage) Title() string {
//line internal/web/templates/admin.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/admin.qtpl:7
	p.WriteTitle(qb422016)
//line internal/web/templates/admin.qtpl:7
	qs422016 := string(qb422016.B)
//line internal/web/templates/admin.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/admin.qtpl:7
	return qs422016
//line internal/web/templates/admin.qtpl:7
}

//line internal/web/templates/admin.qtpl:9
func (p *AdminPage) StreamBody(qw422016 *qt422016.Writer) {
//line internal/web/templates/admin.qtpl:9
	qw422016.N().S(`
<form method="post">
    <div class="columns is-multiline mx-1">
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Release Title</label>
                <div class="control">
                    <input class="input" type="text" name="title" placeholder="Release Title" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Artist</label>
                <div class="control">
                    <input class="input" type="text" name="artist" placeholder="Artist" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Label</label>
                <div class="control">
                    <input class="input" type="text" name="label" placeholder="Label" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Catalogue Number</label>
                <div class="control">
                    <input class="input" type="text" name="cn" placeholder="Catalogue Number" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Genre</label>
                <div class="control">
                    <input class="input" type="text" name="genre" placeholder="Genre" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Released</label>
                <div class="control">
                    <input class="input" type="date" name="release" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Purchased</label>
                <div class="control">
                    <input class="input" type="date" name="purchased" required>
                </div>
            </div>
        </div>
        <div class="column is-one-quarter">
            <div class="field">
                <label class="label has-text-white-ter">Medium</label>
                <div class="control">
                    <div class="select">
                        <select name="medium">
                            <option value="cd">CD</option>
                            <option value="vinyl">Vinyl</option>
                            <option value="cassette">Cassette</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="columns mx-1">
        <div class="column is-one-quarter">
            <div class="field">
                <div class="control">
                    <button class="button is-primary">Submit</button>
                </div>
            </div>
        </div>
    </div>
</form>
`)
//line internal/web/templates/admin.qtpl:93
}

//line internal/web/templates/admin.qtpl:93
func (p *AdminPage) WriteBody(qq422016 qtio422016.Writer) {
//line internal/web/templates/admin.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/admin.qtpl:93
	p.StreamBody(qw422016)
//line internal/web/templates/admin.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/admin.qtpl:93
}

//line internal/web/templates/admin.qtpl:93
func (p *AdminPage) Body() string {
//line internal/web/templates/admin.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/admin.qtpl:93
	p.WriteBody(qb422016)
//line internal/web/templates/admin.qtpl:93
	qs422016 := string(qb422016.B)
//line internal/web/templates/admin.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/admin.qtpl:93
	return qs422016
//line internal/web/templates/admin.qtpl:93
}
