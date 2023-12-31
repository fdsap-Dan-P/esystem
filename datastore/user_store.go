package datastore

import (
	u "simplebank/user"
	"sync"
)

// UserStore is an interface to store users
type UserStore interface {
	// Save saves a user to the store
	Save(user *u.User) error
	// Find finds a user by username
	Find(username string) (*u.User, error)
}

// InMemoryUserStore stores users in memory
type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*u.User
}

// NewInMemoryUserStore returns a new in-memory user store
func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*u.User),
	}
}

// Save saves a user to the store
func (store *InMemoryUserStore) Save(user *u.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return ErrAlreadyExists
	}

	store.users[user.Username] = user.Clone()
	return nil
}

// Find finds a user by username
func (store *InMemoryUserStore) Find(username string) (*u.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
