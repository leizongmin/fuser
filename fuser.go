package fuser

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func readAllPids() []int {
	list, err := os.ReadDir("/proc")
	if err != nil {
		return nil
	}

	pids := make([]int, 0, len(list))
	for _, entry := range list {
		pid, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}
		pids = append(pids, pid)
	}
	return pids
}

func readAllOpenFilesFromPid(pid int) []string {
	path := filepath.Join("/proc", strconv.Itoa(pid), "fd")
	list, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	files := make([]string, 0, len(list))
	for _, entry := range list {
		// files = append(files, filepath.Join(path, entry.Name()))
		rl, err := os.Readlink(filepath.Join(path, entry.Name()))
		if err != nil {
			continue
		}
		files = append(files, rl)
	}
	return files
}

// Options is a struct that contains options for BuildMap.
type Options struct {
	// Filter is a function that returns true if the file should be included in the map.
	Filter func(string) bool
}

// BuildMap returns a map of open files to the pids that have them open.
func BuildMap(options *Options) (map[string][]int, error) {
	if options == nil {
		options = &Options{}
	}

	ret := make(map[string][]int, 0)

	pids := readAllPids()
	for _, pid := range pids {
		files := readAllOpenFilesFromPid(pid)
		for _, file := range files {
			if options.Filter != nil && !options.Filter(file) {
				continue
			}

			if _, ok := ret[file]; !ok {
				ret[file] = make([]int, 0)
			}
			ret[file] = append(ret[file], pid)
		}
	}

	return ret, nil
}

var cacheMap map[string][]int

// Update updates the cache map.
func Update(options *Options) error {
	ret, err := BuildMap(options)
	if err != nil {
		return err
	}
	cacheMap = ret
	return nil
}

var notAPathRe = regexp.MustCompile(`/^\w+:\[\d+\]$`)

// GetPath returns the pids that have the given path open.
func GetPath(p string) []int {
	if cacheMap == nil {
		return nil
	}

	if !notAPathRe.MatchString(p) {
		s, err := filepath.Abs(p)
		if err != nil {
			log.Printf("fuser: failed to get absolute path: %v", err)
		}
		p = s
	}

	return cacheMap[p]
}
