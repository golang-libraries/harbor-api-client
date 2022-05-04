package helpers

import (
	harborErr "github.com/golang-libraries/harbor-api-client/pkg/harbor/errors"
	"strconv"
	"strings"
)

const (
	BadID = -1
)

func ParseResourceLocation(location string) (int64, error) {
	splited := strings.Split(location, "/")
	if len(splited) == 0 {
		return -1, harborErr.ErrFailedToParseResourceLocation
	}
	strID := splited[len(splited)-1]
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func BooToStr(b bool) string {
	if b {
		return "True"
	}
	return "False"
}
