package common

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	nested "github.com/antonfisher/nested-logrus-formatter"
)

var (
	Log        *logrus.Logger
)

func CreateLog() {
	if Log == nil {
		Log = logrus.New()
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&nested.Formatter{
			HideKeys:    false,
			FieldsOrder: []string{"handler", "issue"},
			NoColors:    true,
		})

		pathErr := os.MkdirAll("/var/log/cryposerver", 0755)
		if pathErr != nil {
			panic(fmt.Errorf("cannot create %q", "/var/log/cryposerver"))
		}
		f, err := os.OpenFile("/var/log/cryposerver/handlers.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Errorf("cannot open log file"))
		}
		Log.SetOutput(f)
	}
}