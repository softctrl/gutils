// Package sendmail allow user to send emails to a smpt server.
package sendmail

import (
	"fmt"
	"net/smtp"

	scg "github.com/softctrl/scgotils"
)

const (
	MESSAGE_FORMAT = "From: %s\nTo: %s\nSubject: %s\n\n%s"
)

type SMPT struct {
	Server string
	Port   int
}

type SendMail struct {
	From     string
	Password string
	Smpt     SMPT
}

// SendMail
func (_obj *SendMail) Addr() string {
	return fmt.Sprintf("%s:%d", _obj.Smpt.Server, _obj.Smpt.Port)
}

// NewSendMail
func NewSendMail(_user, _password, _smpt_server string, _smtp_port int) *SendMail {
	return &SendMail{
		From:     _user,
		Password: _password,
		Smpt:     SMPT{Server: _smpt_server, Port: _smtp_port},
	}
}

// SendFromTo
func (_obj *SendMail) SendFromTo(_from string, _to []string, _subject, _message string) error {

	msg := fmt.Sprintf(MESSAGE_FORMAT, _from, scg.Join(_to, ", "), _subject, _message)
	return smtp.SendMail(_obj.Addr(),
		smtp.PlainAuth("", _obj.From, _obj.Password, _obj.Smpt.Server),
		_from,
		_to,
		[]byte(msg))

}

// SendTo
func (_obj *SendMail) SendTo(_to []string, _subject, _message string) error {

	return _obj.SendFromTo(_obj.From, _to, _subject, _message)

}
