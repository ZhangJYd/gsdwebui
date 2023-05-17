package utils

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestImageToBase64(t *testing.T) {
	f1, err := os.Open("../tmp/1.jpg")
	if err != nil {
		log.Error(err)
		return
	}
	defer f1.Close()

	img, err := GetImageInfo(f1)
	if err != nil {
		log.Error(err)
		return
	}
	data, _ := Base64ToImage(img.Base64)

	out, err := os.Create("../tmp/1.png")
	if err != nil {
		log.Error(err)
		return
	}
	defer out.Close()
	out.Write(data)

	log.Info(img, err)
}
