package loggersentry

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func InitSentry() {

	sentry_dsn := os.Getenv("SENTRY_DSN")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentry_dsn,
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
