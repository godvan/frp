package server

import (
	"net"
	"sync"
)

type Server struct {
	listener net.Listener
	sessions map[string]*Session
	mutex    sync.Mutex
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		sessions: make(map[string]*Session),
	}, nil
}

func (s *Server) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// 处理错误
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	// 处理客户端连接、会话管理、隧道请求等逻辑
}
