package output

import (
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

func (self *Terminal) Emergency(message string) (result int, err error) {
	return self.print(message, EMERGENCY, false, true)
}

func (self *Terminal) Alert(message string) (result int, err error) {
	return self.print(message, ALERT, false, true)
}

func (self *Terminal) Critical(message string) (result int, err error) {
	return self.print(message, CRITICAL, false, true)

}
func (self *Terminal) Error(message string) (result int, err error) {
	return self.print(message, ERROR, false, true)
}

func (self *Terminal) Warning(message string) (result int, err error) {
	return self.print(message, WARNING, false, true)
}

func (self *Terminal) Notice(message string) (result int, err error) {
	return self.print(message, NOTICE, false, true)
}

func (self *Terminal) Info(message string) (result int, err error) {
	return self.print(message, INFO, false, true)
}

func (self *Terminal) Debug(message string) (result int, err error) {
	return self.print(message, DEBUG, false, true)
}

func (self *Terminal) Log(message string) (result int, err error) {
	return self.print(message, LOG, false, true)
}

func (self *Terminal) Write(message string, level int) (result int, err error) {
	return self.print(message, level, true, false)
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

func (self *Terminal) print(message string, level int, raw bool, newline bool) (result int, err error) {
	if self.verbosity > level {
		return 0, nil
	}

	if raw == false {
		message = strings.Replace(message, "@", "@@", -1)
	}

	format := "%d: %s"

	channel := terminal.Stderr
	if level > 70 {
		channel = terminal.Stdout
	}

	if self.decorated {
		channel.Color(colors[level])
	}

	channel.Colorf(format, level, message)

	if newline == true {
		channel.Nl()
	}

	return 0, nil
}
