package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

var Rebirth = log.WithFields(logrus.Fields{
	"source":  "rebirth",
	"command": "run",
})

var Server = log.WithFields(logrus.Fields{
	"source": "server",
})

var Extension = log.WithFields(logrus.Fields{
	"source": "extension",
})

func init() {
	log.SetOutput(os.Stdout)
}

func SetFormatter(f string) {
	if f == "text" {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
}

// Enable sets the REBIRTH_DEBUG env var to true
// and makes the logger to log at debug level.
func EnableDebug() {
	_ = os.Setenv("REBIRTH_DEBUG", "1")
	log.SetLevel(logrus.DebugLevel)
}

// IsEnabled checks whether the debug flag is set or not.
func IsEnabled() bool {
	return os.Getenv("REBIRTH_DEBUG") != ""
}