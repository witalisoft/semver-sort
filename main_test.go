package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/coreos/go-semver/semver"
)

func TestPrintSortedVersions(t *testing.T) {
	buf := new(bytes.Buffer)

	testCases := []struct {
		name     string
		versions []*semver.Version
		reverse  bool
		expected []string
	}{
		{
			name:     "Ascending order",
			versions: []*semver.Version{semver.New("1.0.0"), semver.New("2.0.0"), semver.New("1.1.0"), semver.New("1.0.0-rc.1")},
			reverse:  false,
			expected: []string{"1.0.0-rc.1", "1.0.0", "1.1.0", "2.0.0"},
		},
		{
			name:     "Descending order",
			versions: []*semver.Version{semver.New("1.0.0"), semver.New("2.0.0"), semver.New("1.1.0"), semver.New("1.0.0-rc.1")},
			reverse:  true,
			expected: []string{"2.0.0", "1.1.0", "1.0.0", "1.0.0-rc.1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()

			printSortedVersions(buf, tc.versions, tc.reverse)

			output := strings.Split(strings.TrimSpace(buf.String()), "\n")

			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("unexpected output:\nexpected: %v\ngot: %v", tc.expected, output)
			}
		})
	}
}
