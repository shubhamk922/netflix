package logger

type logger struct {
}

func NewLogger() *logger {
	return &logger{}
}

func (l *logger) Info(msg string, args ...any) {
	return
}

func (l *logger) Error(msg string, args ...any) {
	return
}

func (l *logger) Debug(msg string, args ...any) {
	return
}

func (l *logger) Warn(msg string, args ...any) {
	return
}
