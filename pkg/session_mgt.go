package pkg

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type SessionData struct {
	Email     string
	Role      string
	ExpiresAt time.Time
}

type SessionManager struct {
	activeSessions map[string]SessionData
	l              *sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		activeSessions: make(map[string]SessionData),
		l:              &sync.Mutex{},
	}
}

func (s *SessionManager) CreateSession(data SessionData) string {
	s.l.Lock()
	id := uuid.New().String()
	s.activeSessions[id] = data
	s.l.Unlock()
	return id
}

func (s *SessionManager) GetSession(id string) (SessionData, error) {
	s.l.Lock()
	val, ok := s.activeSessions[id]
	if !ok {
		s.l.Unlock()
		return SessionData{}, errors.New("session not found")
	}
	s.l.Unlock()
	return val, nil

}

func (s *SessionManager) DeleteSession(id string) {
	s.l.Lock()
	delete(s.activeSessions, id)
	s.l.Unlock()
}
