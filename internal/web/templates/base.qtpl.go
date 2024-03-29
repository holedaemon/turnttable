// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type Page interface {
	Title() string
	StreamTitle(qw422016 *qt422016.Writer)
	WriteTitle(qq422016 qtio422016.Writer)
	Body() string
	StreamBody(qw422016 *qt422016.Writer)
	WriteBody(qq422016 qtio422016.Writer)
}

func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
	qw422016.N().S(`
<!DOCTYPE html>
<html class="has-background-dark has-text-light">
    <head>
        <title>`)
	p.StreamTitle(qw422016)
	qw422016.N().S(` &bull; turnttable</title>

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
        <meta property="og:image" content="/static/turnttable.png">

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
                        <a class="navbar-item" href="/admin">
                            Admin
                        </a>
                    </div>
                </div>
            </nav>

            `)
	p.StreamBody(qw422016)
	qw422016.N().S(`
        </div>
    </body>
</html>
`)
}

func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamPageTemplate(qw422016, p)
	qt422016.ReleaseWriter(qw422016)
}

func PageTemplate(p Page) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WritePageTemplate(qb422016, p)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

type BasePage struct{}

func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(` `)
}

func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) Title() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(` `)
}

func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *BasePage) Body() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WriteBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
