package server

type Server struct {
	users map[string]string
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
	}
}

func (s *Server) AddUser(user string) {
	s.users[user] = user
}
