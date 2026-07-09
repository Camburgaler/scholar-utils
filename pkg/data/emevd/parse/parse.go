package parse

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Camburgaler/scholar-utils/pkg/data/params/id"
)

// Event parsing regexps
const (
	eventLine1Regexp        = `Event\((\d+),\sDefault,\sfunction\s?\(\)\s{`
	eventFunctionBodyRegexp = `\s{4}(\w+)\s*\(([^)]*)\);`
	eventLastLine           = `}\);`
)

// File names
const (
	EMEVDFileArmor  = "SpEffectArmor"
	EMEVDFileRing   = "SpEffectRing"
	EMEVDFileWeapon = "SpEffectWeapon"
)

var EMEVDFiles = []string{
	EMEVDFileArmor,
	EMEVDFileRing,
	EMEVDFileWeapon,
}

// parseEvent parses a single event from an EMEVD file
//
// @param ev: A slice of strings, where each string is a line from the EMEVD file
//
// @return: An event
func parseEvent(ev []string) Events {
	result := Events{}

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
		errMsgPrefix := "invalid event header format:"
		spacePadding := strings.Repeat(" ", len(errMsgPrefix)-1)
		caratUnderline := strings.Repeat("^", len(ev[0]))
		panic(fmt.Errorf("%s%s\n%s%s", errMsgPrefix, ev[0], spacePadding, caratUnderline))
	}

	// Extract ID
	matchZero, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}

	ID := id.ID(matchZero)

	result[ID] = []Statement{}

	// Extract statements
	re = regexp.MustCompile(eventFunctionBodyRegexp)
	for i := 1; i < len(ev)-1; i++ {
		match = re.FindStringSubmatch(ev[i])

		if match == nil {
			errMsgPrefix := "invalid event statement format:"
			spacePadding := strings.Repeat(" ", len(errMsgPrefix)-1)
			caratUnderline := strings.Repeat("^", len(ev[i]))
			panic(fmt.Errorf("%s%s\n%s%s", errMsgPrefix, ev[i], spacePadding, caratUnderline))
		}

		result[ID] = append(result[ID], Statement{
			Name: match[1],
			Args: strings.Split(match[2], ", "),
		})
	}

	// Validate last line
	matched, err := regexp.MatchString(eventLastLine, ev[len(ev)-1])
	if err != nil || !matched {
		panic(err)
	}

	return result
}

func parseFile(path string) []Events {
	result := []Events{}
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

	fmt.Println("Parsed", len(result), "events")

	return result
}

func Parse() ([]Events, error) {
	result := []Events{}

	for _, file := range EMEVDFiles {
		fmt.Printf("\nParsing %s.emevd.js...\n", file)

		result = append(result, parseFile(fmt.Sprintf("inputs/%s.emevd.js", file))...)
	}

	return result, nil
}
