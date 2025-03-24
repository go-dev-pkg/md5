package md5

import (
	"crypto/md5"
	"encoding/hex"
)

type length int

const (
	Length16 length = 16
	Length32 length = 32
)

type m struct {
	length length
}

type Option func(*m)

func WithLength(l length) Option {
	return func(m *m) {
		m.length = l
	}
}

func Encrypt(str string, opts ...Option) string {
	_m := m{
		length: Length32,
	}
	for _, opt := range opts {
		opt(&_m)
	}

	md5New := md5.New()
	md5New.Write([]byte(str))
	str = hex.EncodeToString(md5New.Sum(nil))

	if _m.length == Length16 {
		str = str[8:24]
	}

	return str
}
