package logs

type Logger interface {
	Print(args ...any)
}

type AppLogger struct {
	eventsBuf chan event
	outputs   []Logger
}

type event struct {
	args []any
}

func NewAppLogger(outputs ...Logger) *AppLogger {
	al := AppLogger{
		eventsBuf: make(chan event, 100),
		outputs:   make([]Logger, 0),
	}

	al.outputs = append(al.outputs, outputs...)

	go al.start()

	return &al
}

func (l *AppLogger) start() {
	for {
		e := <-l.eventsBuf

		for _, o := range l.outputs {
			o.Print(e.args...)
		}
	}
}

func (l *AppLogger) Event(args ...any) {
	l.eventsBuf <- event{
		args: args,
	}
}
