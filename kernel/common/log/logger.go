package log

import (
	"bytes"
	"fmt"
	"io"
	"kernel/common/file"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

const (
	ExitCodeReadOnlyDatabase = 20 // 数据库文件被锁
	ExitCodeUnavailablePort  = 21 // 端口不可用
	ExitCodeWorkspaceLocked  = 24 // 工作空间已被锁定
	ExitCodeInitWorkspaceErr = 25 // 初始化工作空间失败
	ExitCodeFileSysErr       = 26 // 文件系统错误
	ExitCodeOk               = 0  // 正常退出
	ExitCodeFatal            = 1  // 致命错误
)

var (
	logger  *Logger
	logFile *os.File
	LogPath string
)

// init 初始化日志输出位置及文件名
func init() {
	dir, err := os.Getwd()
	if nil != err {
		log.Printf("get current dir failed: %s", err)
	}
	LogPath = filepath.Join(dir, "log", "logging.log")
}

func SetLogPath(path string) {
	LogPath = path
}

func Trace(format string, v ...interface{}) {
	defer closeLogger()
	openLogger()

	if !logger.IsTraceEnabled() {
		return
	}
	logger.LogTrace(format, v...)
}

func Debug(format string, v ...interface{}) {
	defer closeLogger()
	openLogger()

	if !logger.IsDebugEnabled() {
		return
	}
	logger.LogDebug(format, v...)
}

func Info(format string, v ...interface{}) {
	defer closeLogger()
	openLogger()
	logger.LogInfo(format, v...)
}

func Error(format string, v ...interface{}) {
	defer closeLogger()
	openLogger()
	logger.LogError(format, v...)
}

func Warn(format string, v ...interface{}) {
	defer closeLogger()
	openLogger()

	if !logger.IsWarnEnabled() {
		return
	}
	logger.LogWarn(format, v...)
}

func Fatal(exitCode int, format string, v ...interface{}) {
	openLogger()
	logger.LogFatal(exitCode, format, v...)
}

var lock = sync.Mutex{}

func openLogger() {
	lock.Lock()

	// Todo 临时解决日志文件过大的问题
	if file.IsExist(LogPath) {
		if size := file.GetFileSize(LogPath); 1024*1024*32 <= size {
			// 日志文件大于 32M 的话删了重建
			_ = os.Remove(LogPath)
		}
	}

	dir, _ := filepath.Split(LogPath)
	if !file.IsExist(dir) {
		if err := os.MkdirAll(dir, 0755); nil != err {
			log.Printf("create log dir [%s] failed: %s", dir, err)
		}
	}

	var err error
	logFile, err = os.OpenFile(LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		log.Printf("create log file [%s] failed: %s", LogPath, err)
	}
	logger = NewLogger(io.MultiWriter(os.Stdout, logFile))
}

func closeLogger() {
	_ = logFile.Close()
	lock.Unlock()
}

func Recover() {
	if e := recover(); nil != e {
		stack := stack()
		msg := fmt.Sprintf("%s PANIC RECOVERED: %v\n%s\n", e, stack)
		Error(msg)
	}
}

func RecoverError(e any) {
	if nil != e {
		stack := stack()
		msg := fmt.Sprintf("%s PANIC RECOVERED: %v\n%s\n", e, stack)
		Error(msg)
	}
}

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// stack implements Stack, skipping 2 frames.
func stack() []byte {
	buf := &bytes.Buffer{} // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := 2; ; i++ { // Caller we care about is the user, 2 frames up
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := os.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		line-- // in stack trace, lines are 1-indexed but our array is 0-indexed
		_, _ = fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.Trim(lines[n], " \t")
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Since the package path might contains dots (e.g. code.google.com/...),
	// we first remove the path prefix if there is one.
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

// Logging level.
const (
	OFF = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

// the global default logging level, it will be used for creating logger.
var logLevel = DEBUG

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level  int
	logger *log.Logger
}

// NewLogger creates a logger.
func NewLogger(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile)}
	return ret
}

// SetLogLevel sets the logging level of all loggers.
func SetLogLevel(level string) {
	logLevel = getLevel(level)
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return OFF
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return INFO
	}
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level string) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= TRACE
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= DEBUG
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= WARN
}

// Trace prints trace level message with format.
func (l *Logger) LogTrace(format string, v ...interface{}) {
	if TRACE < l.level {
		return
	}

	l.logger.SetPrefix("TRACE ")
	_ = l.logger.Output(3, fmt.Sprintf(format, v...))
}

// Debug prints debug level message with format.
func (l *Logger) LogDebug(format string, v ...interface{}) {
	if DEBUG < l.level {
		return
	}

	l.logger.SetPrefix("DEBUG ")
	_ = l.logger.Output(3, fmt.Sprintf(format, v...))
}

// Info prints info level message with format.
func (l *Logger) LogInfo(format string, v ...interface{}) {
	if INFO < l.level {
		return
	}

	l.logger.SetPrefix("INFO  ")
	_ = l.logger.Output(3, fmt.Sprintf(format, v...))
}

// Warn prints warning level message with format.
func (l *Logger) LogWarn(format string, v ...interface{}) {
	if WARN < l.level {
		return
	}

	l.logger.SetPrefix("WARN  ")
	msg := fmt.Sprintf(format, v...)
	_ = l.logger.Output(3, msg)
}

// Error prints error level message with format.
func (l *Logger) LogError(format string, v ...interface{}) {
	if ERROR < l.level {
		return
	}

	l.logger.SetPrefix("ERROR ")
	msg := fmt.Sprintf(format, v...)
	_ = l.logger.Output(3, msg)
	//sentry.CaptureMessage(msg)
}

// Fatal prints fatal level message with format and exit process with code 1.
func (l *Logger) LogFatal(exitCode int, format string, v ...interface{}) {
	if FATAL < l.level {
		return
	}

	l.logger.SetPrefix("FATAL ")
	format += "\n%s"
	v = append(v, shortStack())
	msg := fmt.Sprintf(format, v...)
	_ = l.logger.Output(3, msg)
	//sentry.CaptureMessage(msg)
	closeLogger()
	os.Exit(exitCode)
}

func shortStack() string {
	output := string(debug.Stack())
	lines := strings.Split(output, "\n")
	if 11 < len(lines) {
		lines = lines[11:]
	}
	buf := bytes.Buffer{}
	for _, l := range lines {
		if strings.Contains(l, "gin-gonic") {
			break
		}
		buf.WriteString("    ")
		buf.WriteString(l)
		buf.WriteByte('\n')
	}
	return buf.String()
}
