package mlog

import (
	"io"
	"log"
	"os"
)

type logger struct {
	*log.Logger
}

func (l logger) Errorln(v ...any) {
	if Debug {
		l.Logger.Fatalln(v...)
	} else {
		l.Logger.Println("Error:")
		l.Logger.Println(v...)
	}
}

func (l logger) Alarm(v ...any) {
	if Debug {
		l.Logger.Fatalln(v...)
	} else {
		// TODO: alarm
	}
}

func (l logger) PrintStruct(v any) {
	l.Logger.Printf("%+v", v)
}

var Log logger

func init() {
	//f, err := os.OpenFile("./output.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.FileMode(0666))
	//if err != nil {
	//	panic(err.Error())
	//}

	var w io.Writer
	//if Debug {
	//	//w = io.MultiWriter(f, os.Stdout)
	//	w = os.Stdout
	//} else {
	//	f, err := os.OpenFile("./output.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.FileMode(0666))
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	w = io.MultiWriter(f, os.Stdout)
	//}
	w = os.Stdout
	Log = logger{Logger: log.New(w, "", log.LstdFlags)}
}

func AddLogOutputFile(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.FileMode(0666))
	if err != nil {
		Log.Errorln("AddLogOutputFile:", err)
	}
	Log.Logger.SetOutput(io.MultiWriter(Log.Logger.Writer(), f))
}
