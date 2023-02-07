// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:1
package templates

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:1
import "github.com/holedaemon/turnttable/internal/db/models"

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:3
func streamformatMedium(qw422016 *qt422016.Writer, m models.Medium) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:3
	qw422016.N().S(`
    `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:4
	if m == models.MediumCD {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:4
		qw422016.N().S(`
        CD
    `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:6
	} else if m == models.MediumVinyl {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:6
		qw422016.N().S(`
        Vinyl
    `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:8
	} else {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:8
		qw422016.N().S(`
        Cassette
    `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:10
	}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:10
	qw422016.N().S(`
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
func writeformatMedium(qq422016 qtio422016.Writer, m models.Medium) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	streamformatMedium(qw422016, m)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
func formatMedium(m models.Medium) string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	writeformatMedium(qb422016, m)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:11
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:14
type IndexPage struct {
	BasePage
	Rows models.RecordSlice
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qw422016.N().S(`Home`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
func (p *IndexPage) Title() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	p.WriteTitle(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:20
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:22
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:22
	qw422016.N().S(`
<div id="buttons" class="field is-grouped is-grouped-right">
    <p class="control">
        <a class="button is-primary" href="/?filter=all">All</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="/?filter=cd">CDs</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="/?filter=vinyl">Vinyl</a>
    </p>
    <p class="control">
        <a class="button is-primary" href="/?filter=cassette">Tape</a>
    </p>
</div>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:37
	if len(p.Rows) == 0 {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:37
		qw422016.N().S(`
    <nav class="level">
        <div class="level-item has-text-centered">
            <p class="title has-text-white is-size-3">No rows... wtf!!!</p>
        </div>
    </nav>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:43
	} else {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:43
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
                </tr>
            </thead>
            <tbody>
                `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:60
		for i, row := range p.Rows {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:60
			qw422016.N().S(`
                    <tr>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:62
			qw422016.N().D(i + 1)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:62
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:63
			qw422016.E().S(row.Title)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:63
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:64
			qw422016.E().S(row.Artist)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:64
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:65
			qw422016.E().S(row.Label)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:65
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:66
			qw422016.E().S(row.CN)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:66
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:67
			qw422016.E().S(row.Genre)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:67
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:68
			qw422016.N().S(formatMedium(row.Medium))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:68
			qw422016.N().S(`</td>
                        <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:69
			qw422016.E().S(row.Released.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:69
			qw422016.N().S(`</td>

                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:71
			if row.Purchased.Time.IsZero() {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:71
				qw422016.N().S(`
                            <td>N/A</td>
                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:73
			} else {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:73
				qw422016.N().S(`
                            <td>`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:74
				qw422016.E().S(row.Purchased.Time.Format("2006-01-02"))
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:74
				qw422016.N().S(`</td>
                        `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:75
			}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:75
			qw422016.N().S(`
                    </tr>
                `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:77
		}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:77
		qw422016.N().S(`
            </tbody>
        </table>
    </div>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:81
	}
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:81
	qw422016.N().S(`
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
func (p *IndexPage) Body() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	p.WriteBody(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/index.qtpl:82
}
