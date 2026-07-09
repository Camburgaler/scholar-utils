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
func (e *Events) parseEvent(ev []string) error {
	statements := []Statement{}

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
		return err
	}

	ID := id.ID(matchZero)

	// Extract statements
	re = regexp.MustCompile(eventFunctionBodyRegexp)
	for i := 1; i < len(ev)-1; i++ {
		match = re.FindStringSubmatch(ev[i])

		if match == nil {
			errMsgPrefix := "invalid event statement format:"
			spacePadding := strings.Repeat(" ", len(errMsgPrefix)-1)
			caratUnderline := strings.Repeat("^", len(ev[i]))
			return fmt.Errorf("%s%s\n%s%s", errMsgPrefix, ev[i], spacePadding, caratUnderline)
		}

		statements = append(statements, Statement{
			Name: match[1],
			Args: strings.Split(match[2], ", "),
		})
	}

	// Validate last line
	matched, err := regexp.MatchString(eventLastLine, ev[len(ev)-1])
	if err != nil || !matched {
		return err
	}

	(*e)[ID] = statements
	return nil
}

func (e *DS2EMEVD) parseFile(path string) error {
	result := Events{}
	lines := []string{}

	// Open file
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	f.Close()

	if err := scanner.Err(); err != nil {
		return err
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
				err := result.parseEvent(currentEv)
				if err != nil {
					return err
				}
			}

			// End of previous event
			currentEv = []string{}

			continue
		}

		// Add line to current event
		currentEv = append(currentEv, lines[i])
	}

	fmt.Println("Parsed", len(result), "events")

	switch path {
	case EMEVDFileArmor:
		e.SpEffectArmor = result
	}

	return nil
}

func Parse() (DS2EMEVD, error) {
	result := DS2EMEVD{}

	// For each file, populate the struct
	for _, file := range EMEVDFiles {
		fmt.Printf("\nParsing %s.emevd.js...\n", file)

		// Parse
		result.parseFile(fmt.Sprintf("inputs/%s.emevd.js", file))
	}

	return result, nil
}
