package dotenv

import (
	"bufio"
	"github.com/Oleexo/config-go"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	envRegex = regexp.MustCompile(`\$\{(.+?)\}`)
)

func parseLineToEntry(line string) (string, config.Entry, error) {
	idx := strings.Index(line, "=")
	if idx == -1 {
		return "", config.Entry{}, nil
	}
	key := strings.TrimSpace(line[:idx])
	value := strings.TrimSpace(line[idx+1:])

	if envRegex.MatchString(value) {
		matches := envRegex.FindAllStringSubmatch(value, -1)

		for _, match := range matches {
			envKey := match[1]
			envValue := os.Getenv(envKey)
			value = strings.Replace(value, match[0], envValue, -1)
		}
	}

	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return key, config.NewEntryInt(intValue), nil
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return key, config.NewEntryFloat(floatValue), nil
	}
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return key, config.NewEntryBool(boolValue), nil
	}
	return key, config.NewEntryString(value), nil
}

func readFile(path string) (map[string]config.Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	result := make(map[string]config.Entry)

	for scanner.Scan() {
		line := scanner.Text()
		var key string
		var entry config.Entry
		key, entry, err = parseLineToEntry(line)
		if err != nil {
			return nil, err
		}
		if key != "" {
			result[key] = entry
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
