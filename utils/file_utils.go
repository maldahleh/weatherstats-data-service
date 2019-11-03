package utils

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func DeleteFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		log.Error("failed to delete file", filepath, "error", err)
	}
}
