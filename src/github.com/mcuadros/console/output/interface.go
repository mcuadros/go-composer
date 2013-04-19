package output

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

type Output interface {
	Emergency(message string, a ...interface{}) (result int, err error)
	Alert(message string, a ...interface{}) (result int, err error)
	Critical(message string, a ...interface{}) (result int, err error)
	Error(message string, a ...interface{}) (result int, err error)
	Warning(message string, a ...interface{}) (result int, err error)
	Notice(message string, a ...interface{}) (result int, err error)
	Info(message string, a ...interface{}) (result int, err error)
	Debug(message string, a ...interface{}) (result int, err error)
	Log(message string, a ...interface{}) (result int, err error)
	Write(message string, level int, a ...interface{}) (result int, err error)
	SetVerbosity(level int)
	GetVerbosity() (level int)
	SetDecorated(descorated bool)
	IsDecorated() (descorated bool)
}
