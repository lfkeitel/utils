package sessions

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type FileSystemStore struct {
	path string
}

func NewFileStore(path string) *FileSystemStore {
	if path == "" {
		path = os.TempDir()
	}
	fs := &FileSystemStore{
		path: path,
	}
	return fs
}

func (f *FileSystemStore) Get(r *http.Request, name string) *Session {
	id := getSessionID(r, name)
	if id == "" {
		return NewSession(f, name)
	}
	s, err := f.load(id)
	if err != nil {
		return NewSession(f, name)
	}
	s.s = f
	s.name = name
	return s
}

func (f *FileSystemStore) New(name string) *Session {
	return NewSession(f, name)
}

func (f *FileSystemStore) Save(r *http.Request, w http.ResponseWriter, s *Session) error {
	if s.id == "" {
		s.RegenerateID()
	}
	if err := f.save(s); err != nil {
		fmt.Println(err.Error())
		return err
	}
	saveIDCookie(w, s.name, s.id)
	return nil
}

func (f *FileSystemStore) save(s *Session) error {
	filename := filepath.Join(f.path, "session_"+s.id)
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	if err := enc.Encode(s); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data.Bytes(), 0600)
}

func (f *FileSystemStore) load(id string) (*Session, error) {
	filename := filepath.Join(f.path, "session_"+id)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	dec := gob.NewDecoder(file)
	var s Session
	if err = dec.Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}
