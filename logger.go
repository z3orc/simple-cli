package cli

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var infoLogger *log.Logger
var warnLogger *log.Logger
var errorLogger *log.Logger

func init() {

	infoText := color.New(color.FgWhite, color.Bold).SprintFunc()
	warnText := color.New(color.BgYellow, color.Bold).SprintFunc()
	errorText := color.New(color.BgHiRed, color.Bold).SprintFunc()

	infoLogger = log.New(os.Stdout, infoText("info:\t"), log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, warnText("warn:\t"), log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, errorText("error:\t"), log.Ldate|log.Ltime|log.Lshortfile)
}
