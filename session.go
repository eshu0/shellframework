package simpleshell

import (
	"math/rand"
	"time"

	"github.com/eshu0/shellframework"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type SimpleSession struct {
	id string
}

// this is a very simple random string
// this is just to uniquely identify each session
// this is meant to be over written if needs be
func NewSimpleSession() shellframework.ISession {
	ss := new(SimpleSession)

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	ss.id = string(b)
	return ss
}

// returns the string ID for the session
func (ss *SimpleSession) ID() string {
	return ss.id
}
