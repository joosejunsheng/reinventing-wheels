package storage

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"testing"
)

func TestNewLog(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "log-*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	log, err := NewLog(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create log: %v", err)
	}

	if log.offset != 0 {
		t.Fatalf("Invalid offset for new file: %v", err)
	}
}

func TestAppend(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "log-*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	log, err := NewLog(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create log: %v", err)
	}

	msg := []byte("Hello")

	_, err = log.Append(msg)
	if err != nil {
		t.Fatalf("Failed to append message 1: %v", err)
	}

	stat, err := tmpFile.Stat()
	if err != nil {
		t.Fatalf("Failed to get file stat: %v", err)
	}

	expectedSize := 8 + int64(len(msg))
	if stat.Size() != expectedSize {
		t.Fatalf("Expected file size %d after first append, but got %d", expectedSize, stat.Size())
	}

	_, err = tmpFile.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("Failed to seek to the beginning of the file: %v", err)
	}

	var length int64
	if err := binary.Read(tmpFile, binary.BigEndian, &length); err != nil {
		t.Fatalf("Failed to read length prefix: %v", err)
	}

	buf := make([]byte, length)
	_, err = tmpFile.Read(buf)
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	t.Logf("File content: %s", buf)

	if !bytes.Equal(buf, msg) {
		t.Fatalf("Expected message %s, but got %s", msg, buf)
	}
}
