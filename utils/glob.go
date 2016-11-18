package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MediaFile struct {
	BaseName string
	AbsPath  string
}

func (m MediaFile) String() string {
	return fmt.Sprintf("%s \n", m.CleanName())
}

func (m MediaFile) CleanName() string {
	baseName := strings.Replace(m.BaseName, ".", " ", -1)
	baseName = strings.ToLower(baseName)
	regex, _ := regexp.Compile("(.+)s([0-9]+)e([0-9]+)")
	matches := regex.FindStringSubmatch(baseName)
	if len(matches) > 2 {
		return matches[1]
	}

	return m.BaseName
}

var mediaFiles []MediaFile

func visit(path string, f os.FileInfo, err error) error {
	basePath := filepath.Base(path)
	absPath, _ := filepath.Abs(path)

	if strings.HasSuffix(basePath, ".torrent") {
		mediaFile := MediaFile{BaseName: basePath, AbsPath: absPath}
		mediaFiles = append(mediaFiles, mediaFile)
	}

	return nil
}

func GetTorrents(path string) []MediaFile {
	mediaFiles = make([]MediaFile, 0)
	_ = filepath.Walk(path, visit)
	return mediaFiles

}
