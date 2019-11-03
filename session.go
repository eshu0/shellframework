package shellframework

import (
	"math/rand"
	"time"

	"github.com/eshu0/shellframework/interfaces"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Session struct {
	id          string
	interactive bool
	shell       sfinterfaces.IShell
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

func (command *Session) SetShell(shell sfinterfaces.IShell) {
	command.shell = shell
}

func (command *Session) GetShell() sfinterfaces.IShell {
	return command.shell
}

// returns the string ID for the session
func (ss Session) ID() string {
	return ss.id
}

// Session is Interactive?
func (session Session) GetInteractive() bool {
	return session.interactive
}

func (session Session) SetInteractive(interactive bool) {
	session.interactive = interactive
}
