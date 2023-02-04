// Code generated by qtc from "about.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/web/templates/about.qtpl:1
package templates

//line internal/web/templates/about.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/web/templates/about.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/web/templates/about.qtpl:2
type Aboutpage struct {
	BasePage
}

//line internal/web/templates/about.qtpl:7
func (p *Aboutpage) StreamTitle(qw422016 *qt422016.Writer) {
//line internal/web/templates/about.qtpl:7
	qw422016.N().S(`Home`)
//line internal/web/templates/about.qtpl:7
}

//line internal/web/templates/about.qtpl:7
func (p *Aboutpage) WriteTitle(qq422016 qtio422016.Writer) {
//line internal/web/templates/about.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/about.qtpl:7
	p.StreamTitle(qw422016)
//line internal/web/templates/about.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/about.qtpl:7
}

//line internal/web/templates/about.qtpl:7
func (p *Aboutpage) Title() string {
//line internal/web/templates/about.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/about.qtpl:7
	p.WriteTitle(qb422016)
//line internal/web/templates/about.qtpl:7
	qs422016 := string(qb422016.B)
//line internal/web/templates/about.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/about.qtpl:7
	return qs422016
//line internal/web/templates/about.qtpl:7
}

//line internal/web/templates/about.qtpl:9
func (p *Aboutpage) StreamBody(qw422016 *qt422016.Writer) {
//line internal/web/templates/about.qtpl:9
	qw422016.N().S(`
 <div class="columns is-centered is-multiline is-mobile is-variable is-0-mobile">
        <div class="column has-text-right-desktop has-text-centered-touch is-full-mobile">
            <img id="img-profile" src="/static/hsuhao.png">
        </div>
        <div class="column has-text-centered-mobile has-text-left-desktop is-full-mobile">
            <h1 class="title has-text-white is-size-4">TURNTTABLE ENTERPRISES, LLC.</h1>
            <hr id="about-hr">
            <p class="has-text-white mb-3">
                Founded in the year 1920 by the beautiful man seen on the left, TURNTTABLE ENTERPRISES LLC has grown to become the lead supplier in high-quality, harmful microplastics, made specifically for human consumption. It's the hope, dream, and promise of TURNTTABLE ENTERPRISES LLC founder Bald Hsu Hao that each and every customer of TURNTTABLE ENTERPRISES LLC experience any kind of health complications as result of consuming TURNTTABLE ENTERPRISES LLC microplastics.
            </p>
            <p class="has-text-white">
                In reality, turnttable is a website made by <a class="has-text-primary" href="https://holedaemon.net">holedaemon</a> to track the physical music <span id="about-pronouns">they</span> <span id="about-owns">own</span>. The source code is on <a class="has-text-primary" href="https://github.com/holedaemon/turnttable">GitHub</a>.
            </p>
        </div>
    </div>
</div>
<script>
    const pronouns = ["they", "he", "she"];
    const span = document.getElementById("about-pronouns");
    const owns = document.getElementById("about-owns")

    const selection = pronouns[Math.floor(Math.random() * pronouns.length)]

    if (selection !== "they") {
        owns.innerText = "owns"
    }

    span.innerText = selection
</script>
`)
//line internal/web/templates/about.qtpl:39
}

//line internal/web/templates/about.qtpl:39
func (p *Aboutpage) WriteBody(qq422016 qtio422016.Writer) {
//line internal/web/templates/about.qtpl:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/web/templates/about.qtpl:39
	p.StreamBody(qw422016)
//line internal/web/templates/about.qtpl:39
	qt422016.ReleaseWriter(qw422016)
//line internal/web/templates/about.qtpl:39
}

//line internal/web/templates/about.qtpl:39
func (p *Aboutpage) Body() string {
//line internal/web/templates/about.qtpl:39
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/web/templates/about.qtpl:39
	p.WriteBody(qb422016)
//line internal/web/templates/about.qtpl:39
	qs422016 := string(qb422016.B)
//line internal/web/templates/about.qtpl:39
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/web/templates/about.qtpl:39
	return qs422016
//line internal/web/templates/about.qtpl:39
}
