package service

import (
	"io/ioutil"
	"log"
)

var _ Manager = &Storage{}

type Manager interface {
	Store(file *File) error
}

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) Store(file *File) error {
	log.Printf("Store: %v", "start>>>")
	if err := ioutil.WriteFile(s.dir+file.name, file.buffer.Bytes(), 0644); err != nil {
		log.Printf("Store: %v", err)
		return err
	}

	return nil
}
