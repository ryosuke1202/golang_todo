package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 0666はパーミッション
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)   // logの標準出力先を指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // logのフォーマットを指定
	log.SetOutput(multiLogFile)                          // logの出力先を指定
}
