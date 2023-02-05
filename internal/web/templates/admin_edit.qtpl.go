// Code generated by qtc from "admin_edit.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:1
package templates

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:1
import (
	"github.com/holedaemon/turnttable/internal/db/models"
	"strings"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:7
type AdminEditPage struct {
	BasePage
	Record *models.Record
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
func (p *AdminEditPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qw422016.N().S(`Edit Record`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
func (p *AdminEditPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
func (p *AdminEditPage) Title() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	p.WriteTitle(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:13
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:15
func (p *AdminEditPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:15
	qw422016.N().S(`
<form method="post">
    <div class="columns is-multiline mx-1">
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Release Title</label>
                <div class="control">
                    <input class="input" type="text" name="title" placeholder="Release Title" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:22
	qw422016.E().S(p.Record.Title)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:22
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Artist</label>
                <div class="control">
                    <input class="input" type="text" name="artist" placeholder="Artist" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:30
	qw422016.E().S(p.Record.Artist)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:30
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Label</label>
                <div class="control">
                    <input class="input" type="text" name="label" placeholder="Label" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:38
	qw422016.E().S(p.Record.Label)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:38
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Catalogue Number</label>
                <div class="control">
                    <input class="input" type="text" name="cn" placeholder="Catalogue Number" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:46
	qw422016.E().S(p.Record.CN)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:46
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Genre</label>
                <div class="control">
                    <input class="input" type="text" name="genre" placeholder="Genre" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:54
	qw422016.E().S(p.Record.Genre)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:54
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Released</label>
                <div class="control">
                    <input class="input" type="date" name="release" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:62
	qw422016.E().S(p.Record.Released.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:62
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-half">
            <div class="field">
                <label class="label has-text-white-ter">Purchased</label>
                <div class="control">
                    <input class="input" type="date" name="purchased" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:70
	qw422016.E().S(p.Record.Purchased.Time.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:70
	qw422016.N().S(`" required>
                </div>
            </div>
        </div>
        <div class="column is-one-quarter">
            <div class="field">
                <label class="label has-text-white-ter">Medium</label>
                <div class="control">
                    <div class="select">
                        <select name="medium" value="`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:79
	qw422016.N().S(strings.ToLower(formatMedium(p.Record.Medium)))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:79
	qw422016.N().S(`">
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
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
func (p *AdminEditPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
func (p *AdminEditPage) Body() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	p.WriteBody(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_edit.qtpl:99
}
