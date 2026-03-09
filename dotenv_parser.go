package config

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	envRegex = regexp.MustCompile(`\$\{(.+?)\}`)
)

func parseLineToEntry(line string) (string, Entry) {
	idx := strings.Index(line, "=")
	if idx == -1 {
		return "", Entry{}
	}
	key := strings.TrimSpace(line[:idx])
	value := strings.TrimSpace(line[idx+1:])

	if envRegex.MatchString(value) {
		matches := envRegex.FindAllStringSubmatch(value, -1)

		for _, match := range matches {
			envKey := match[1]
			envValue := os.Getenv(envKey)
			value = strings.ReplaceAll(value, match[0], envValue)
		}
	}

	if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return key, NewInt(intValue)
	}
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return key, NewFloat(floatValue)
	}
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return key, NewBool(boolValue)
	}
	return key, NewString(value)
}

func readDotenvFile(path string) (result map[string]Entry, err error) {
	// #nosec G304 -- dotenv paths are intentionally provided by caller options.
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := file.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()

	scanner := bufio.NewScanner(file)
	result = make(map[string]Entry)

	for scanner.Scan() {
		line := scanner.Text()
		key, entry := parseLineToEntry(line)
		if key != "" {
			result[key] = entry
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
