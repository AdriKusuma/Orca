package output

import (
	"fmt"
	"os"
	"sync"
)

type Writer struct {
	file *os.File
	mu   sync.Mutex
}

func New(path string) (*Writer, error) {
	if path == "" {
		return &Writer{}, nil
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &Writer{file: f}, nil
}

func (w *Writer) Write(line string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.file != nil {
		fmt.Fprintln(w.file, line)
	}
	fmt.Println(line) 
}

func (w *Writer) Close() {
	if w.file != nil {
		w.file.Close()
	}
}