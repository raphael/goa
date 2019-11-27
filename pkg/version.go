package goa

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	// Major version number
	Major = 3
	// Minor version number
	Minor = 0
	// Build number
	Build = 8
	// Suffix - set to empty string in release tag commits.
	Suffix = ""
)

var (
	// Version format
	versionFormat = regexp.MustCompile(`v(\d+?)\.(\d+?)\.(\d+?)(?:-.+)?`)
)

// Version returns the complete version number.
func Version() string {
	if Suffix != "" {
		return fmt.Sprintf("v%d.%d.%d-%s", Major, Minor, Build, Suffix)
	}
	return fmt.Sprintf("v%d.%d.%d", Major, Minor, Build)
}

// Compatible returns true if Major matches the major version of the given version string.
// It returns an error if the given string is not a valid version string.
func Compatible(v string) (bool, error) {
	matches := versionFormat.FindStringSubmatch(v)
	if len(matches) != 4 {
		return false, fmt.Errorf("invalid version string format %#v, %+v", v, matches)
	}
	mj, err := strconv.Atoi(matches[1])
	if err != nil {
		return false, fmt.Errorf("invalid major version number %#v, must be number, %v", matches[1], err)
	}
	return mj == Major, nil
}
