package storage

import (
	"encoding/binary"
	"io"
	"os"
	"sync"
)

type Log struct {
	file   *os.File
	mu     sync.Mutex
	offset int64
}

func NewLog(logPath string) (*Log, error) {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	info, _ := file.Stat()

	return &Log{
		file:   file,
		offset: info.Size(),
	}, err
}

func (l *Log) Append(msg []byte) (int64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	length := int64(len(msg))
	if err := binary.Write(l.file, binary.BigEndian, length); err != nil {
		return -1, err
	}

	n, err := l.file.Write(msg)
	if err != nil {
		return -1, err
	}

	offset := l.offset
	l.offset += int64(8 + n)

	// Returning original offset
	// Know where it starts writing from
	return offset, l.file.Sync()
}

func (l *Log) ReadAll() ([][]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := l.file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var res [][]byte
	for {
		var length int64
		if err := binary.Read(l.file, binary.BigEndian, &length); err != nil {

			if err == io.EOF {
				break
			}

			return nil, err
		}

		buf := make([]byte, length)
		_, err := l.file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		res = append(res, buf)
	}

	return res, err
}

func (l *Log) ReadFrom(offset int64) ([][]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := l.file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	var res [][]byte
	for {
		var length int64
		if err := binary.Read(l.file, binary.BigEndian, &length); err != nil {

			if err == io.EOF {
				break
			}

			return nil, err
		}

		buf := make([]byte, length)
		_, err := l.file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		res = append(res, buf)
	}

	return res, err
}
