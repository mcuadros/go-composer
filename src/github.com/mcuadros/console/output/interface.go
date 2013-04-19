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
	Emergency(message string) (result int, err error)
	Alert(message string) (result int, err error)
	Critical(message string) (result int, err error)
	Error(message string) (result int, err error)
	Warning(message string) (result int, err error)
	Notice(message string) (result int, err error)
	Info(message string) (result int, err error)
	Debug(message string) (result int, err error)
	Log(message string) (result int, err error)
	Write(message string, level int) (result int, err error)
	SetVerbosity(level int)
	GetVerbosity() (level int)
	SetDecorated(descorated bool)
	IsDecorated() (descorated bool)
}
