package agent

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Load(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open user-agent file: %s", path)
	}
	defer file.Close()

	var uas []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			uas = append(uas, line)
		}
	}
	if len(uas) == 0 {
		return nil, fmt.Errorf("List of user-agent is empty")
	}
	return uas, nil
}

func Random(uas []string) string {
	rand.Seed(time.Now().UnixNano())
	return uas[rand.Intn(len(uas))]
}