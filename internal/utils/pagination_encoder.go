package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

// DecodeCursor func decode cursor
func DecodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	res, err = time.Parse(time.RFC3339, arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}

// EncodeCursor func encode cursor
func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
