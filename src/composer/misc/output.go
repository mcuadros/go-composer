package misc

import (
	"fmt"
	"github.com/wsxiaoys/terminal"
	"strings"
)

const (
	EMERGENCY = 90
	ALERT     = 80
	CRITICAL  = 70
	ERROR     = 60
	WARNING   = 50
	NOTICE    = 40
	INFO      = 30
	DEBUG     = 20
	LOG       = 10
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

var instantiated *output = nil

func GetOutput() *output {
	if instantiated == nil {
		instantiated = new(output)
		instantiated.SetDecorated(true)
	}

	return instantiated
}

type output struct {
	verbosity int
	decorated bool
}

func (self *output) Emergency(message string, a ...interface{}) (result int, err error) {
	return self.print(message, EMERGENCY, false, true, a...)
}

func (self *output) Alert(message string, a ...interface{}) (result int, err error) {
	return self.print(message, ALERT, false, true, a...)
}

func (self *output) Critical(message string, a ...interface{}) (result int, err error) {
	return self.print(message, CRITICAL, false, true, a...)

}
func (self *output) Error(message string, a ...interface{}) (result int, err error) {
	return self.print(message, ERROR, false, true, a...)
}

func (self *output) Warning(message string, a ...interface{}) (result int, err error) {
	return self.print(message, WARNING, false, true, a...)
}

func (self *output) Notice(message string, a ...interface{}) (result int, err error) {
	return self.print(message, NOTICE, false, true, a...)
}

func (self *output) Info(message string, a ...interface{}) (result int, err error) {
	return self.print(message, INFO, false, true, a...)
}

func (self *output) Debug(message string, a ...interface{}) (result int, err error) {
	return self.print(message, DEBUG, false, true, a...)
}

func (self *output) Log(message string, a ...interface{}) (result int, err error) {
	return self.print(message, LOG, false, true, a...)
}

func (self *output) Write(message string, level int, a ...interface{}) (result int, err error) {
	return self.print(message, level, true, false, a...)
}

func (self *output) SetVerbosity(level int) {
	self.verbosity = level
}

func (self *output) GetVerbosity() (level int) {
	return self.verbosity
}

func (self *output) SetDecorated(descorated bool) {
	self.decorated = descorated
}

func (self *output) IsDecorated() (descorated bool) {
	return self.decorated
}

func (self *output) print(message string, level int, raw bool, newline bool, a ...interface{}) (result int, err error) {
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
