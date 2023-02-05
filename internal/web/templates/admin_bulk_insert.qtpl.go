// Code generated by qtc from "admin_bulk_insert.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:1
package templates

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:2
type AdminBulkInsertPage struct {
	BasePage
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
func (p *AdminBulkInsertPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qw422016.N().S(`Bulk Insert`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
func (p *AdminBulkInsertPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
func (p *AdminBulkInsertPage) Title() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	p.WriteTitle(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:7
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:9
func (p *AdminBulkInsertPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:9
	qw422016.N().S(`
<form method="post" enctype="multipart/form-data">
    <div class="columns is-multiline is-centered">
        <div class="column is-narrow">
            <div class="file is-boxed is-large is-link is-centered has-name mb-3">
                <label class="file-label">
                    <input id="file-selector" class="file-input" type="file" name="records" accept="text/csv">
                    <span class="file-cta">
                        <span class="file-label">
                            Choose a file...
                        </span>
                    </span>
                    <span id="file-name" class="file-name has-text-centered">
                        No file uploaded
                    </span>
                </label>
            </div>
        </div>
        <div class="column is-full">
            <div class="field">
                <div class="control has-text-centered">
                    <button class="button is-primary">Submit</button>
                </div>
            </div>
        </div>
    </div>
</form>
<script>
    const fileInput = document.getElementById("file-selector");
    fileInput.onchange = () => {
        if (fileInput.files.length > 0) {
            const fileName = document.getElementById("file-name");
            fileName.innerText = fileInput.files[0].name;
        }
    }
</script>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
func (p *AdminBulkInsertPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
func (p *AdminBulkInsertPage) Body() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	p.WriteBody(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/admin_bulk_insert.qtpl:45
}
