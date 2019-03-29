package tool

import (
	"crypto/md5"
	"errors"
	"strconv"
	"time"
)

func GenerateCode(id string, tm time.Time) (string, error) {
	return GenerateWithDigit(id, tm, 6)
}

func GenerateWithDigit(id string, tm time.Time, digit int) (string, error) {
	if digit < 1 {
		return "", errors.New("Invalid Digit !")
	}
	hs, err := hash(id + timeToStr(tm))
	if err != nil {
		return "", err
	}
	return byteToCode(hs, digit), nil
}

func timeToStr(tm time.Time) string {
	str := tm.Format("200601021504")
	if perSecond(tm.Second(), 30) == 1 {
		str += "1"
	} else {
		str += "2"
	}
	return str
}

func perSecond(second, per int) int {
	return (second + per) / per
}

func byteToCode(data []byte, digit int) string {
	str := ""
	for i := 0; i < digit; i++ {
		num := int(data[i])
		str += strconv.Itoa(num % 10)
	}
	return str
}

func hash(str string) ([]byte, error) {
	h := md5.New()
	_, err := h.Write([]byte(str))
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
