package logger

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://696fb708f8ab46d6834ba32a9ca4dfd0@o4504869668388864.ingest.sentry.io/4504870783680512",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

}

func CaptureErrorMessage(_string string) {
	InitSentry()
	sentry.CaptureMessage(_string)
	defer sentry.Flush(2 * time.Second)
}

func CaptureErrorException(string string) {
	InitSentry()
	sentry.CaptureException(fmt.Errorf(string))
	defer sentry.Flush(2 * time.Second)

}
