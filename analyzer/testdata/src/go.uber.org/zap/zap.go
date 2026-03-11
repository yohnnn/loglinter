package zap

type Logger struct{}
type SugaredLogger struct{}

func NewNop() *Logger                          { return &Logger{} }
func L() *Logger                               { return &Logger{} }
func (l *Logger) Sugar() *SugaredLogger        { return &SugaredLogger{} }
func (l *Logger) Info(msg string, fields ...any)  {}
func (l *Logger) Error(msg string, fields ...any) {}
func (l *Logger) Warn(msg string, fields ...any)  {}
func (l *Logger) Debug(msg string, fields ...any) {}
func (l *Logger) Fatal(msg string, fields ...any) {}
func (l *Logger) Panic(msg string, fields ...any) {}
