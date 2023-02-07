// Code generated by qtc from "admin_alter.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:1
package templates

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:1
import "github.com/holedaemon/turnttable/internal/db/models"

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:4
type AdminAlterPage struct {
	BasePage
	Rows models.RecordSlice
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
func (p *AdminAlterPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qw422016.N().S(`Alter Record`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
func (p *AdminAlterPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
func (p *AdminAlterPage) Title() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	p.WriteTitle(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:10
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:12
func (p *AdminAlterPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:12
	qw422016.N().S(`
<div id="buttons" class="field is-grouped is-grouped-right">
    <p class="control">
        <a class="button is-primary" href="?filter=all">All</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="?filter=cd">CDs</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="?filter=vinyl">Vinyl</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="?filter=cassette">Tape</a>
    </p>
</div>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:27
	if len(p.Rows) == 0 {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:27
		qw422016.N().S(`
    <nav class="level">
        <div class="level-item has-text-centered">
            <p class="title has-text-white is-size-3">No rows... wtf!!!</p>
        </div>
    </nav>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:33
	} else {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:33
		qw422016.N().S(`
    <div class="table-container">
        <table id="turnt-table" class="table is-fullwidth">
            <thead>
                <tr>
                    <th><abbr title="Position">Pos</abbr></th>
                    <th>Title</th>
                    <th>Artist</th>
                    <th>Label</th>
                    <th><abbr title="Catalogue Number">CN</abbr></th>
                    <th>Genre</th>
                    <th>Medium</th>
                    <th>Released</th>
                    <th>Purchased</th>
                    <th>Edit</th>
                    <th>Delete</th>
                </tr>
            </thead>
            <tbody>
                `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:52
		for i, row := range p.Rows {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:52
			qw422016.N().S(`
                    <tr>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:54
			qw422016.N().D(i + 1)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:54
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:55
			qw422016.E().S(row.Title)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:55
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:56
			qw422016.E().S(row.Artist)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:56
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:57
			qw422016.E().S(row.Label)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:57
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:58
			qw422016.E().S(row.CN)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:58
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:59
			qw422016.E().S(row.Genre)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:59
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:60
			qw422016.N().S(formatMedium(row.Medium))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:60
			qw422016.N().S(`</td>
                        <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:61
			qw422016.E().S(row.Released.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:61
			qw422016.N().S(`</td>

                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:63
			if row.Purchased.Time.IsZero() {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:63
				qw422016.N().S(`
                            <td class="alter-td">N/A</td>
                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:65
			} else {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:65
				qw422016.N().S(`
                            <td class="alter-td">`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:66
				qw422016.E().S(row.Purchased.Time.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:66
				qw422016.N().S(`</td>
                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:67
			}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:67
			qw422016.N().S(`


                        <td class="alter-td"><a class="button is-primary" href="/admin/edit/`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:70
			qw422016.N().DL(row.ID)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:70
			qw422016.N().S(`">Edit</a></td>
                        <td class="alter-td"><a class="button is-danger" href="/admin/delete/`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:71
			qw422016.N().DL(row.ID)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:71
			qw422016.N().S(`">Delete</a></td>
                    </tr>
                `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:73
		}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:73
		qw422016.N().S(`
            </tbody>
        </table>
    </div>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:77
	}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:77
	qw422016.N().S(`
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
func (p *AdminAlterPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
func (p *AdminAlterPage) Body() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	p.WriteBody(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_alter.qtpl:78
}
