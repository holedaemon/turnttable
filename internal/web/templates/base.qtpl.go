// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
package templates

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
type Page interface {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	Title() string
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	StreamTitle(qw422016 *qt422016.Writer)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	WriteTitle(qq422016 qtio422016.Writer)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	Body() string
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	StreamBody(qw422016 *qt422016.Writer)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
	WriteBody(qq422016 qtio422016.Writer)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:1
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:7
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:7
	qw422016.N().S(`
<!DOCTYPE html>
<html class="has-background-dark has-text-light">
    <head>
        <title>turnttable &bull; `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:11
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:11
	qw422016.N().S(`</title>

        <link rel="apple-touch-icon" sizes="180x180" href="/static/apple-touch-icon.png">
        <link rel="icon" type="image/png" sizes="32x32" href="/static/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="16x16" href="/static/favicon-16x16.png">
        <link rel="manifest" href="/static/site.webmanifest">
        <link rel="mask-icon" href="/static/safari-pinned-tab.svg" color="#5bbad5">
        <link rel="shortcut icon" href="/static/favicon.ico">
        <meta name="msapplication-TileColor" content="#da532c">
        <meta name="msapplication-config" content="/static/browserconfig.xml">
        <meta name="theme-color" content="#ffffff">

        <meta name="twitter:card" content="summary">
        <meta name="twitter:site" content="@holedaemon">
        <meta name="twitter:creator" content="@holedaemon">
        <meta name="twitter:description" content="A place for all the ladies to ogle my music collection">
        <meta name="twitter:title" content="turnttable">

        <meta property="og:title" content="turnttable">
        <meta property="og:description" content="A place for all the ladies to ogle my music collection">
        <meta property="og:type" content="website">
        <meta property="og:image" content="/static/hsuhao.png">

        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="/static/turnttable.css">
    </head>
    <body>
        <div class="container">
             <nav class="navbar is-dark mb-5" role="navigation" aria-label="main navigation">
                <div class="navbar-brand">
                    <a class="navbar-item is-size-4" href="/">
                        turnttable
                    </a>
                </div>
                
                <div class="navbar-menu">
                    <div class="navbar-end">
                        <a class="navbar-item" href="/about">
                            About
                        </a>
                        <a class="navbar-item" href="/admin">
                            Admin
                        </a>
                    </div>
                </div>
            </nav>

            `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:58
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:58
	qw422016.N().S(`
        </div>
    </body>
</html>
`)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	StreamPageTemplate(qw422016, p)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
func PageTemplate(p Page) string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	WritePageTemplate(qb422016, p)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:62
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:64
type BasePage struct{}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qw422016.N().S(` `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	p.StreamTitle(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
func (p *BasePage) Title() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	p.WriteTitle(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:65
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qw422016.N().S(` `)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	p.StreamBody(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
}

//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
func (p *BasePage) Body() string {
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	p.WriteBody(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qs422016 := string(qb422016.B)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
	return qs422016
//line /home/max/git/holedaemon/turnttable/internal/web/templates/base.qtpl:66
}
