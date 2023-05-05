package golog

import (
	"testing"
)

func TestLogger(t *testing.T) {
	// simple logger
	l := NewSimpleLogger()
	l.SetFormatFunc(DefaultFormatFunc)
	l.SetDebug(false)
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
	l.Debug("Debug")
	l.Fatal("Fatal")
	l.Panic("Panic")
	l.Infof("Infof")
	l.Warnf("Warnf")
	l.Errorf("Errorf")
	l.Debugf("Debugf")
	l.Fatalf("Fatalf")
	l.Panicf("Panicf")

	// tag logger
	tl := NewTagLogger(l, "tag")
	tl.Info("Info")
	tl.Warn("Warn")
	tl.Error("Error")
	tl.Debug("Debug")
	tl.Fatal("Fatal")
	tl.Panic("Panic")
	tl.Infof("Infof")
	tl.Warnf("Warnf")
	tl.Errorf("Errorf")
	tl.Debugf("Debugf")
	tl.Fatalf("Fatalf")
	tl.Panicf("Panicf")

	// tag logger x2
	tl2 := NewTagLogger(tl, "tag2")
	tl2.Info("Info")
	tl2.Warn("Warn")
	tl2.Error("Error")
	tl2.Debug("Debug")
	tl2.Fatal("Fatal")
	tl2.Panic("Panic")
	tl2.Infof("Infof")
	tl2.Warnf("Warnf")
	tl2.Errorf("Errorf")
	tl2.Debugf("Debugf")
	tl2.Fatalf("Fatalf")
	tl2.Panicf("Panicf")

	// logger writer
	lw := NewLogWriter(tl2, Info)
	lw.Write([]byte("Write"))
}
