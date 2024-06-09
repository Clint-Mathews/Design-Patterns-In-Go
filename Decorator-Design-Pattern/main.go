package main

import (
	"fmt"
	"time"
)

// The Decorator Design Pattern in Go is a structural pattern that allows you to dynamically add behavior to an object without affecting the behavior of other objects from the same type

// Component Interface
type Logger interface {
	log(message string)
}

// Concrete Component
type SimpleLogger struct{}

func (s *SimpleLogger) log(message string) {
	fmt.Printf("Log: %s\n", message)
}

// Decorator
// LoggerDecorator maintains a reference to a Logger object and implements the Logger interface.
type LoggerDecorator struct {
	Logger
}

// Concrete Decorators

// TimeStamped concrete decorator
type TimeStampedLogger struct {
	LoggerDecorator
}

func (l *TimeStampedLogger) log(message string) {
	timeUTC := time.Now().UTC()
	timestamp := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d.%03dZ", timeUTC.Year(), timeUTC.Month(), timeUTC.Day(), timeUTC.Hour(), timeUTC.Minute(), timeUTC.Second(), timeUTC.Nanosecond()/1000000)
	l.LoggerDecorator.log(fmt.Sprintf("[%s] %s", timestamp, message))
}

// Prefixed concrete decorator
type PrefixedLogger struct {
	LoggerDecorator
	Prefix string
}

func (p *PrefixedLogger) log(message string) {
	p.LoggerDecorator.log(fmt.Sprintf("[%s] %s", p.Prefix, message))
}

// Client using logger
func main() {
	var simpleLogger Logger = &SimpleLogger{}
	simpleLogger.log("Message")

	timeStampedLogger := &TimeStampedLogger{LoggerDecorator{Logger: simpleLogger}}
	timeStampedLogger.log("Message")

	prefixedLogger := &PrefixedLogger{LoggerDecorator: LoggerDecorator{Logger: timeStampedLogger}, Prefix: "Debug"}
	prefixedLogger.log("Message")
}
