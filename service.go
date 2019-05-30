package therundown

import (
	"os"

	blaster "github.com/joelhill/go-rest-http-blaster"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// CompressionHeaderGzip - type of compression header
	CompressionHeaderGzip = "gzip"
)

// IService - all functions that implement the My Sports Feed Service Interface.
type IService interface {
	NewEventBySportOptions() *EventBySportOptions
	EventBySport(options *EventBySportOptions) (TheRunDownAPILines, int, error)

	NewEventBySportByDateOptions() *EventBySportByDateOptions
	EventBySportByDate(options *EventBySportByDateOptions) (TheRunDownAPILines, int, error)
}

// Service - dependencies for the my sports feed service
type Service struct {
	Config *Config
	Logger *logrus.Entry
}

// NewService -
func NewService(config *Config) *Service {

	// JSON formatter for log output if not running in a TTY
	// because Loggly likes JSON but humans like colors
	if !terminal.IsTerminal(int(os.Stderr.Fd())) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stderr)
	}

	logLevel, err := logrus.ParseLevel("debug")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error_message": err.Error(),
		}).Error("log level error")
	}
	logrus.SetLevel(logLevel)

	blaster.SetDefaults(&blaster.Defaults{
		ServiceName:    "the-rundown-api-go",
		RequireHeaders: false,
		StatsdRate:     0.0,
	})

	l := logrus.New()
	lg := logrus.NewEntry(l)

	return &Service{
		Config: config,
		Logger: lg,
	}
}
