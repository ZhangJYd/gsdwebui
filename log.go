package gsdwebui

import log "github.com/sirupsen/logrus"

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(
		&log.TextFormatter{
			ForceQuote:      true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
}
