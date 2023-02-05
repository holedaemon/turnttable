package web

import (
	"net/http"

	"github.com/holedaemon/turnttable/internal/web/templates"
)

func (s *Server) statusPage(w http.ResponseWriter, title, header, subtitle string) {
	templates.WritePageTemplate(w, &templates.StatusPage{
		PageTitle: title,
		Header:    header,
		Subtitle:  subtitle,
	})
}

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	s.statusPage(w, "uh oh!!!!", "Huh what?", "The requested resource ain't here")
}

func (s *Server) unauthorized(w http.ResponseWriter, header bool) {
	if header {
		w.Header().Add("WWW-Authenticate", `Basic realm="turnttable"`)
	}

	w.WriteHeader(http.StatusUnauthorized)
	s.statusPage(w, "Oh Uh", "LOL HOW ARE YOU SO SMALL???", "You need special pants to open this door (Unauthorized)")
}

func (s *Server) internalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	s.statusPage(w, "uh oh sisters", "idk man, I think it's your catalytic converter", "An internal server error has occurred, try again later")
}

func (s *Server) badRequest(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	s.statusPage(w, "uh oh sisters!!!", "lol ur request sucks bro", message)
}
