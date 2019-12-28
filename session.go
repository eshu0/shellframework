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
	//interactive bool
	shell       sfinterfaces.IShell
	idmethod func(ss sfinterfaces.ISession) string
	interactiveMethod func(ss sfinterfaces.ISession) bool
}

func DefaultBuildIDMethod(ss sfinterfaces.ISession) {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	ss.SetID(string(b))
}

func DefaultNonInteractiveMethod(ss sfinterfaces.ISession) bool {
//	ss.interactive = false
//	return ss.interactive
	return false
}

func DefaultInteractiveMethod(ss sfinterfaces.ISession) bool {
//	ss.interactive = true
//	return ss.interactive
	return true
}

// this is a very simple random string
// this is just to uniquely identify each session
// this is meant to be over written if needs be
func NewSession(idmethod func(ss sfinterfaces.ISession) string,	interactiveMethod func(ss sfinterfaces.ISession) bool ) sfinterfaces.ISession {
	ss := new(Session)
	//ss.interactive = false
	ss.SetInteractiveMethod(interactiveMethod)
	ss.SetIDMethod(idmethod)
	return ss
}

func NewDefaultInteractiveSession(idmethod func(ss sfinterfaces.ISession) string) sfinterfaces.ISession {
	ss := NewSession(idmethod, DefaultInteractiveMethod)
	return ss
}

func NewDefaultNonInteractiveSession(idmethod func(ss sfinterfaces.ISession) string) sfinterfaces.ISession {
	ss := NewSession(idmethod, DefaultNonInteractiveMethod)
	return ss
}

func (session *Session) SetShell(shell sfinterfaces.IShell) {
	session.shell = shell
}

func (session *Session) GetShell() sfinterfaces.IShell {
	return session.shell
}

func (session *Session) SetInteractiveMethod(interactiveMethod  func(ss sfinterfaces.ISession) bool) {
	session.interactiveMethod = interactiveMethod
}

func (session *Session) GetInteractiveMethod() func(ss sfinterfaces.ISession) bool {
	return session.interactiveMethod
}

func (session *Session) SetBuildIDMethod(builidmethod  func(ss sfinterfaces.ISession)) {
	session.builidmethod = builidmethod
}

func (session *Session) ID() string {
	return session.id
}
func (session *Session) SetID(id string) string {
	session.id = id
}
