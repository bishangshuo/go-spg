package utils

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	LOG_LEVEL_INFO    = 0x00000001
	LOG_LEVEL_WARNING = 0x00000002
	LOG_LEVEL_ERROR   = 0x00000003
)

type Loger struct {
	logFile *os.File
}

var gInstance *Loger
var once sync.Once
var chanOutput chan bytes.Buffer

func getLogerInstance() *Loger {
	once.Do(func() {
		debugFile := GetDataDir() + "\\debug.log"
		fd, err := os.OpenFile(debugFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err == nil {
			gInstance = &Loger{logFile: fd}
			chanOutput = make(chan bytes.Buffer)
			go ThreadOutput(gInstance.logFile)
		} else {
			panic("Can not open log file")
		}
	})
	return gInstance
}

func ThreadOutput(file *os.File) {
	for ch := range chanOutput {
		fmt.Println(ch.String())
		file.Write(ch.Bytes())
	}
}

func (loger *Loger) printAndWriteLog(level int, ctxFormat string, ctx ...interface{}) {
	currentTime := time.Now()
	var strLevel string
	switch level {
	case LOG_LEVEL_INFO:
		strLevel = "INFO"
	case LOG_LEVEL_WARNING:
		strLevel = "WARNING"
	case LOG_LEVEL_ERROR:
		strLevel = "ERROR"
	default:
		strLevel = "INFO"
	}

	content := fmt.Sprintf(ctxFormat, ctx...)

	var byteContent bytes.Buffer
	byteContent.WriteString(currentTime.String())
	byteContent.WriteString(" ")
	byteContent.WriteString(strLevel)
	byteContent.WriteString(" ")
	byteContent.WriteString(content)
	byteContent.WriteString("\r\n")

	chanOutput <- byteContent
}

func PrintInfo(ctxFormat string, ctx ...interface{}) {
	getLogerInstance().printAndWriteLog(LOG_LEVEL_INFO, ctxFormat, ctx...)
}

func PrintWarning(ctxFormat string, ctx ...interface{}) {
	getLogerInstance().printAndWriteLog(LOG_LEVEL_WARNING, ctxFormat, ctx...)
}

func PrintError(ctxFormat string, ctx ...interface{}) {
	getLogerInstance().printAndWriteLog(LOG_LEVEL_ERROR, ctxFormat, ctx...)
}
