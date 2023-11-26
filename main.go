package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
)

var (
	reverse bool
)

func main() {
	getCliFlags()
	logger := log.New(os.Stderr, "", log.Lshortfile)
	versions, err := readStdin()
	if err != nil {
		logger.Fatal(err)
	}
	printSortedVersions(os.Stdout, versions, reverse)
}

func readStdin() ([]*semver.Version, error) {
	var parsedVersions []*semver.Version
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawVersion := scanner.Text()
		v, err := semver.NewVersion(rawVersion)
		if err != nil {
			return nil, fmt.Errorf("Error parsing version %s: %s", rawVersion, err)
		}
		parsedVersions = append(parsedVersions, v)
	}
	return parsedVersions, nil
}

func printSortedVersions(output io.Writer, versions []*semver.Version, reverse bool) {
	semver.Sort(versions)

	if reverse {
		for i := len(versions) - 1; i >= 0; i-- {
			fmt.Fprintln(output, versions[i])
		}
		return
	}

	for _, v := range versions {
		fmt.Fprintln(output, v)
	}
}

func getCliFlags() {
	var help bool
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nReads versions from stdin and prints them to stdout in sorted order.")
	}
	flag.BoolVar(&reverse, "reverse", false, "Reverse the order of the sorted versions")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
}
