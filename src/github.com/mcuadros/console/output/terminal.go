package output

import (
	"fmt"
	"github.com/wsxiaoys/terminal"
	"strings"
)

var colors = map[int]string{
	EMERGENCY: "!wR",
	ALERT:     "!rW",
	CRITICAL:  "!r",
	ERROR:     "r",
	WARNING:   "!y",
	NOTICE:    "g",
	INFO:      "",
	DEBUG:     "m",
	LOG:       "/",
}

type Terminal struct {
	verbosity int
	decorated bool
}

func NewTerminal() *Terminal {
	terminal := new(Terminal)
	terminal.SetDecorated(true)

	return terminal
}

func (self *Terminal) Emergency(message string, a ...interface{}) (result int, err error) {
	return self.print(message, EMERGENCY, false, true, a...)
}

func (self *Terminal) Alert(message string, a ...interface{}) (result int, err error) {
	return self.print(message, ALERT, false, true, a...)
}

func (self *Terminal) Critical(message string, a ...interface{}) (result int, err error) {
	return self.print(message, CRITICAL, false, true, a...)

}
func (self *Terminal) Error(message string, a ...interface{}) (result int, err error) {
	return self.print(message, ERROR, false, true, a...)
}

func (self *Terminal) Warning(message string, a ...interface{}) (result int, err error) {
	return self.print(message, WARNING, false, true, a...)
}

func (self *Terminal) Notice(message string, a ...interface{}) (result int, err error) {
	return self.print(message, NOTICE, false, true, a...)
}

func (self *Terminal) Info(message string, a ...interface{}) (result int, err error) {
	return self.print(message, INFO, false, true, a...)
}

func (self *Terminal) Debug(message string, a ...interface{}) (result int, err error) {
	return self.print(message, DEBUG, false, true, a...)
}

func (self *Terminal) Log(message string, a ...interface{}) (result int, err error) {
	return self.print(message, LOG, false, true, a...)
}

func (self *Terminal) Write(message string, level int, a ...interface{}) (result int, err error) {
	return self.print(message, level, true, false, a...)
}

func (self *Terminal) SetVerbosity(level int) {
	self.verbosity = level
}

func (self *Terminal) GetVerbosity() (level int) {
	return self.verbosity
}

func (self *Terminal) SetDecorated(descorated bool) {
	self.decorated = descorated
}

func (self *Terminal) IsDecorated() (descorated bool) {
	return self.decorated
}

func (self *Terminal) print(message string, level int, raw bool, newline bool, a ...interface{}) (result int, err error) {
	if self.verbosity > level {
		return 0, nil
	}

	message = fmt.Sprintf(message, a...)

	if raw == false {
		message = strings.Replace(message, "@", "@@", -1)
	}

	channel := terminal.Stderr
	if level > 70 {
		channel = terminal.Stdout
	}

	if self.decorated {
		channel.Color(colors[level])
	}

	channel.Colorf(message).Reset()

	if newline == true {
		channel.Nl()
	}

	return 0, nil
}
