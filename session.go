package shellframework

import (
	"math/rand"
	"time"

	"github.com/eshu0/shellframework/interfaces"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Session struct {
	sfinterfaces.ISession
	id          string
	interactive bool
	shell       sfinterfaces.IShell
	idmethod func() string
	interactiveMethod func() bool
}

// this is a very simple random string
// this is just to uniquely identify each session
// this is meant to be over written if needs be
func NewSession() sfinterfaces.ISession {
	ss := new(Session)

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	ss.id = string(b)
	ss.SetInteractive(false)
	return ss
}

func NewInteractiveSession() sfinterfaces.ISession {
	ss := NewSession()
	ss.SetInteractive(true)
	return ss
}

func (session *Session) SetShell(shell sfinterfaces.IShell) {
	session.shell = shell
}

func (session *Session) GetShell() sfinterfaces.IShell {
	return session.shell
}

func (session *Session) SetInteractiveMethod(interactiveMethod  func() bool) {
	session.interactiveMethod = interactiveMethod
}

func (session *Session) GetInteractiveMethod()  func() bool {
	return session.interactiveMethod
}


func (session *Session) SetIDMethod(idmethod  func() string) {
	session.idmethod = idmethod
}

func (session *Session) GetIDMethod()  func() string {
	return session.idmethod
}
