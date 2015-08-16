package sessions

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"net/http"
	"strings"
)

// Session represents a web session
type Session struct {
	id     string
	name   string
	Values map[interface{}]interface{}
	s      Store
}

// NewSession is called by Stores to generate a new session
func NewSession(store Store, name string) *Session {
	return &Session{
		name:   name,
		Values: make(map[interface{}]interface{}),
		s:      store,
	}
}

// Set sets key to value in the session
func (s *Session) Set(key interface{}, value interface{}) {
	s.Values[key] = value
}

// Get returns the value of key and if it exists
func (s *Session) Get(key interface{}) (interface{}, bool) {
	v, ok := s.Values[key]
	return v, ok
}

// GetWithDefault returns the value of key if it exists or def
func (s *Session) GetWithDefault(key interface{}, def interface{}) interface{} {
	if v, ok := s.Get(key); ok {
		return v
	}
	return def
}

// Save is a wrapper for the Store.Save method attached to the session
func (s *Session) Save(r *http.Request, w http.ResponseWriter) error {
	return s.s.Save(r, w, s)
}

// Name returns the name of the session
func (s *Session) Name() string {
	return s.name
}

func (s *Session) RegenerateID() {
	s.id = strings.TrimRight(
		base32.StdEncoding.EncodeToString(
			GenerateRandomKey(32)), "=")
}

// Store returns the underlying Store for the session
func (s *Session) Store() Store {
	return s.s
}

// A Store is what saves and gets a session to/from storage
type Store interface {
	Get(r *http.Request, name string) *Session
	New(name string) *Session
	Save(r *http.Request, w http.ResponseWriter, s *Session) error
}

func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}

func saveIDCookie(w http.ResponseWriter, name string, id string) {
	http.SetCookie(w, &http.Cookie{
		Name:   name,
		Value:  id,
		Path:   "/",
		MaxAge: 86400 * 30,
	})
}

func getSessionID(r *http.Request, name string) string {
	c, e := r.Cookie(name)
	if e != nil {
		return ""
	}
	return c.Value
}
