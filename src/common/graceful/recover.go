package graceful

import (
	"common/logs"
	"fmt"
	"runtime"
	"strings"
)

var logger = logs.New("panic-recover")

func Recover(fn ...func(msg string)) {
	if err := recover(); err != nil {
		if len(fn) > 0 {
			fn[0](fmt.Sprint(err))
		}
		logger.Errorf("%s\n%s", err, GetStacks())
	}
}

func GetStacks() string {
	var stack []string
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, fmt.Sprintf("\tat %s:%d", file, line))
	}
	joinStr := "\n"
	return strings.Join(stack, joinStr)
}