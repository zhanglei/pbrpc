package pbrpc

import (
	"bytes"
	"io/ioutil"

	gzip "compress/gzip"
)

func GZIP(b []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	w, _ := gzip.NewWriterLevel(buf, gzip.NoCompression)
	defer w.Close()

	_, err := w.Write(b)
	w.Flush()

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GUNZIP(b []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(b)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)

	return undatas, nil
}
