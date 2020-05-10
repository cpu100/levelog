package levelog

import (
    "fmt"
    "strings"
    "time"
)

var Level = 50

const (
    LevelError   = 10
    LevelWarn    = 20
    LevelHttp    = 30
    LevelInfo    = 40
    LevelVerbose = 50
    LevelSilly   = 60
)

func log(prefix interface{}, a ...interface{}) {

    length := len(a)

    if length == 0 {
        fmt.Println(prefix)
        return
    }

    now := time.Now().Format("2006-01-02 15:04:05")

    if length == 1 {
        fmt.Printf("[%s] %v %v \n", now, prefix, a[0])
    } else {
        // Type switches
        // Type assertions
        // https://tour.golang.org/methods/15
        // format := a[0].(string) // 数字会出错 todo 数字、interface{} 转字符串的最佳方式
        // 真很捉急
        // .(Type) 和 .(*Type) 又不一样
        format := fmt.Sprintf("%v", a[0])
        var msgs []interface{}
        if strings.Index(format, "%") < 0 {
            format = strings.Repeat("%v ", length)
            msgs = a
        } else {
            msgs = a[1:]
        }
        fmt.Printf(fmt.Sprintf("[%s] %v %s \n", now, prefix, format), msgs...)
    }
}

func Log(level int, prefix interface{}, a ...interface{}) {
    if level <= Level {
        log(prefix, a...)
    }
}

func Error(prefix interface{}, a ...interface{}) {
    if LevelError <= Level {
        log(fmt.Sprintf("\x1B[31m%v\x1B[0m", prefix), a...)
    }
}

func Warn(prefix interface{}, a ...interface{}) {
    if LevelWarn <= Level {
        log(fmt.Sprintf("\x1B[33m%v\x1B[0m", prefix), a...)
    }
}

func Http(prefix interface{}, a ...interface{}) {
    if LevelHttp <= Level {
        log(prefix, a...)
    }
}

func Info(prefix interface{}, a ...interface{}) {
    if LevelInfo <= Level {
        log(fmt.Sprintf("\x1B[32m%v\x1B[0m", prefix), a...)
    }
}

func Verbose(prefix interface{}, a ...interface{}) {
    if LevelVerbose <= Level {
        log(prefix, a...)
    }
}

func Silly(prefix interface{}, a ...interface{}) {
    if LevelSilly <= Level {
        log(prefix, a...)
    }
}
