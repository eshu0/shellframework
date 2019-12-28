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
	builidmethod func(ss sfinterfaces.ISession, shell sfinterfaces.IShell)
	interactiveMethod func(ss sfinterfaces.ISession) bool
}

func DefaultBuildIDMethod(ss sfinterfaces.ISession, shell sfinterfaces.IShell) {
	if(ss.ID() == ""){
		rand.Seed(time.Now().UnixNano())
		b := make([]rune, 10)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		ss.SetID(string(b))
	}
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
func NewSession(buildidmethod func(ss sfinterfaces.ISession, shell sfinterfaces.IShell),	interactiveMethod func(ss sfinterfaces.ISession) bool ) sfinterfaces.ISession {
	ss := new(Session)
	//ss.interactive = false
	ss.SetInteractiveMethod(interactiveMethod)
	ss.SetBuildIDMethod(buildidmethod)
	return ss
}

func NewDefaultInteractiveSession(buildidmethod func(ss sfinterfaces.ISession, shell sfinterfaces.IShell)) sfinterfaces.ISession {
	ss := NewSession(buildidmethod, DefaultInteractiveMethod)
	return ss
}

func NewDefaultNonInteractiveSession(buildidmethod func(ss sfinterfaces.ISession, shell sfinterfaces.IShell)) sfinterfaces.ISession {
	ss := NewSession(buildidmethod, DefaultNonInteractiveMethod)
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

func (session *Session) SetBuildIDMethod(builidmethod  func(ss sfinterfaces.ISession, shell sfinterfaces.IShell)) {
	session.builidmethod = builidmethod
}

func (session *Session) CallBuildIDMethod(shell sfinterfaces.IShell) {
	session.builidmethod(session, shell)
}

func (session *Session) ID() string {
	return session.id
}
func (session *Session) SetID(id string) {
	session.id = id
}
