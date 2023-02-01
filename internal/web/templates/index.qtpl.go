// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/web/templates/index.qtpl:1
package templates

//line internal/web/templates/index.qtpl:1
import "github.com/holedaemon/turnttable/internal/db/models"

//line internal/web/templates/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/web/templates/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/web/templates/index.qtpl:4
type IndexPage struct {
	BasePage
	Rows models.RecordSlice
}

//line internal/web/templates/index.qtpl:10
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
//line internal/web/templates/index.qtpl:10
	qw422016.N().S(`Home`)
//line internal/web/templates/index.qtpl:10
}

//line internal/web/templates/index.qtpl:10
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
//line internal/web/templates/index.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/index.qtpl:10
	p.StreamTitle(qw422016)
//line internal/web/templates/index.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/index.qtpl:10
}

//line internal/web/templates/index.qtpl:10
func (p *IndexPage) Title() string {
//line internal/web/templates/index.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/index.qtpl:10
	p.WriteTitle(qb422016)
//line internal/web/templates/index.qtpl:10
	qs422016 := string(qb422016.B)
//line internal/web/templates/index.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/index.qtpl:10
	return qs422016
//line internal/web/templates/index.qtpl:10
}

//line internal/web/templates/index.qtpl:12
func (p *IndexPage) StreamMeta(qw422016 *qt422016.Writer) {
//line internal/web/templates/index.qtpl:12
}

//line internal/web/templates/index.qtpl:12
func (p *IndexPage) WriteMeta(qq422016 qtio422016.Writer) {
//line internal/web/templates/index.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/index.qtpl:12
	p.StreamMeta(qw422016)
//line internal/web/templates/index.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/index.qtpl:12
}

//line internal/web/templates/index.qtpl:12
func (p *IndexPage) Meta() string {
//line internal/web/templates/index.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/index.qtpl:12
	p.WriteMeta(qb422016)
//line internal/web/templates/index.qtpl:12
	qs422016 := string(qb422016.B)
//line internal/web/templates/index.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/index.qtpl:12
	return qs422016
//line internal/web/templates/index.qtpl:12
}

//line internal/web/templates/index.qtpl:14
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
//line internal/web/templates/index.qtpl:14
	qw422016.N().S(`
<table id="turnt-table" class="table is-fullwidth">
                <thead>
                    <tr>
                        <th><abbr title="Position">Pos</abbr></th>
                        <th>Title</th>
                        <th>Artist</th>
                        <th>Label</th>
                        <th><abbr title="Catalogue Number">CN</abbr></th>
                        <th>Genre</th>
                        <th>Released</th>
                        <th>Purchased</th>
                    </tr>
                </thead>
                <tbody>
                    `)
//line internal/web/templates/index.qtpl:29
	for i, row := range p.Rows {
//line internal/web/templates/index.qtpl:29
		qw422016.N().S(`
                        <tr>
                            <td>`)
//line internal/web/templates/index.qtpl:31
		qw422016.N().D(i)
//line internal/web/templates/index.qtpl:31
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:32
		qw422016.E().S(row.Title)
//line internal/web/templates/index.qtpl:32
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:33
		qw422016.E().S(row.Artist)
//line internal/web/templates/index.qtpl:33
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:34
		qw422016.E().S(row.Label)
//line internal/web/templates/index.qtpl:34
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:35
		qw422016.E().S(row.CN)
//line internal/web/templates/index.qtpl:35
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:36
		qw422016.E().S(row.Genre)
//line internal/web/templates/index.qtpl:36
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:37
		qw422016.E().S(row.Released.Format("2006-01-02"))
//line internal/web/templates/index.qtpl:37
		qw422016.N().S(`</td>
                            <td>`)
//line internal/web/templates/index.qtpl:38
		qw422016.E().S(row.Purchased.Time.Format("2006-01-02"))
//line internal/web/templates/index.qtpl:38
		qw422016.N().S(`</td>
                        </tr>
                    `)
//line internal/web/templates/index.qtpl:40
	}
//line internal/web/templates/index.qtpl:40
	qw422016.N().S(`
                </tbody>
            </table>
`)
//line internal/web/templates/index.qtpl:43
}

//line internal/web/templates/index.qtpl:43
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
//line internal/web/templates/index.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/index.qtpl:43
	p.StreamBody(qw422016)
//line internal/web/templates/index.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/index.qtpl:43
}

//line internal/web/templates/index.qtpl:43
func (p *IndexPage) Body() string {
//line internal/web/templates/index.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/index.qtpl:43
	p.WriteBody(qb422016)
//line internal/web/templates/index.qtpl:43
	qs422016 := string(qb422016.B)
//line internal/web/templates/index.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/index.qtpl:43
	return qs422016
//line internal/web/templates/index.qtpl:43
}
