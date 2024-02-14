package web

import "database/sql"

// Option is used to configure a server.
type Option func(*Server)

// WithAddr sets a server's address.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithDB sets a server's DB client.
func WithDB(db *sql.DB) Option {
	return func(s *Server) {
		s.db = db
	}
}

// WithAdmins sets a server's admins.
func WithAdmins(admins map[string]string) Option {
	return func(s *Server) {
		s.admins = admins
	}
}
