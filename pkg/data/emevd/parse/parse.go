package parse

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/params/id"
)

// Event parsing regexps
const (
	eventLine1Regexp        = `Event\((\d+), (\w+), function\(\) {`
	eventFunctionBodyRegexp = `\s{4}(\w+)\((\w+)\);`
	eventLastLine           = `});`
)

// parseEvent parses a single event from an EMEVD file
//
// @param ev: A slice of strings, where each string is a line from the EMEVD file
//
// @return: An event
func parseEvent(ev []string) Event {
	result := Event{
		ID:         id.ID(0),
		Statements: []Statement{},
	}

	// Validate event format
	//
	// Event(<ID>, Default, function() {
	//     <function name>(<args>);
	//     <function name>(<args>);
	//     <function name>(<args>);
	//     ...
	// });

	re := regexp.MustCompile(eventLine1Regexp)
	match := re.FindStringSubmatch(ev[0])

	if match == nil {
		panic("invalid event header format")
	}

	// Extract ID
	ID, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}

	result.ID = id.ID(ID)

	// Extract statements
	re = regexp.MustCompile(eventFunctionBodyRegexp)
	for i := 1; i < len(ev)-1; i++ {
		match = re.FindStringSubmatch(ev[i])

		if match == nil {
			panic("invalid event body format")
		}

		result.Statements = append(result.Statements, Statement{
			Name: match[0],
			Args: strings.Split(match[1], ", "),
		})
	}

	// Validate last line
	matched, err := regexp.MatchString(eventLastLine, ev[len(ev)-1])
	if err != nil || !matched {
		panic(err)
	}

	return result
}

func parseFile(path string) []Event {
	result := []Event{}
	lines := []string{}

	// Open file
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	f.Close()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	currentEv := []string{}
	for i := 0; i < len(lines); i++ {
		// The JS-converted EMEVD files start with several lines of comments.
		// Skip them.
		if strings.HasPrefix(lines[i], "//") {
			continue
		}

		// Events have empty lines between them
		if lines[i] == "" {
			// Add previous event to result
			if len(currentEv) > 0 {
				result = append(result, parseEvent(currentEv))
			}

			// End of previous event
			currentEv = []string{}

			continue
		}

		// Add line to current event
		currentEv = append(currentEv, lines[i])
	}

	return result
}

func Parse() ([]Event, error) {
	result := []Event{}

	// for each EMEVD file
	//   result = append(result, parseFile(<path to EMEVD file>))

	return result, nil
}
