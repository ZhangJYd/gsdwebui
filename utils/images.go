package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type ImageInfo struct {
	Base64 string
	Width  int
	Height int
}

func ImageToBase64(img io.Reader) (string, error) {
	m, _, err := image.Decode(img)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, m); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func GetImageInfo(img io.Reader) (*ImageInfo, error) {
	m, _, err := image.Decode(img)
	if err != nil {
		m, err = jpeg.Decode(img)
		if err != nil {
			return nil, err
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, m); err != nil {
		return nil, err
	}

	b := m.Bounds()
	ret := &ImageInfo{}
	ret.Width = b.Max.X
	ret.Height = b.Max.Y
	ret.Base64 = base64.StdEncoding.EncodeToString(buf.Bytes())

	return ret, nil
}

func Base64ToImage(ba string) ([]byte, error) {
	bs, err := base64.StdEncoding.DecodeString(ba) //成图片文件并把文件写入到buffer
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(bs)
	m, _, err := image.Decode(buf) // 图片文件解码
	if err != nil {
		return nil, err
	}
	var ret bytes.Buffer

	if err := png.Encode(&ret, m); err != nil {
		return nil, err
	}
	return ret.Bytes(), nil
}
