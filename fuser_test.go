package fuser

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAllPids(t *testing.T) {
	ret := readAllPids()
	assert.NotNil(t, ret)
	assert.Greater(t, len(ret), 0)
	assert.Contains(t, ret, os.Getpid())
}

func TestReadAllOpenFilesFromPid(t *testing.T) {
	tmp, err := os.MkdirTemp(t.TempDir(), "*")
	assert.NoError(t, err)
	tmpFile, err := os.Open(tmp)
	assert.NoError(t, err)
	defer tmpFile.Close()

	ret := readAllOpenFilesFromPid(os.Getpid())
	assert.NotNil(t, ret)
	assert.Greater(t, len(ret), 0)
	assert.Contains(t, ret, tmp)
}

func TestBuildMap(t *testing.T) {
	ret, err := BuildMap(nil)
	assert.NoError(t, err)
	assert.NotNil(t, ret)
	assert.Greater(t, len(ret), 0)
}

func TestUpdateAndGetPath(t *testing.T) {
	tmp, err := os.MkdirTemp(t.TempDir(), "*")
	assert.NoError(t, err)
	tmpFile, err := os.Open(tmp)
	assert.NoError(t, err)
	defer tmpFile.Close()

	{
		ret := GetPath(tmp)
		assert.Nil(t, ret)
	}
	{
		err := Update(nil)
		assert.NoError(t, err)
		ret := GetPath(tmp)
		assert.NotNil(t, ret)
		assert.Greater(t, len(ret), 0)
		ret2 := GetPath(tmp + "2")
		assert.Nil(t, ret2)
	}
	{
		assert.NoError(t, tmpFile.Close())
		err := Update(nil)
		assert.NoError(t, err)
		ret := GetPath(tmp)
		assert.Nil(t, ret)
	}
}

func TestBuildMapWithFilter(t *testing.T) {
	tmp, err := os.MkdirTemp(t.TempDir(), "*")
	assert.NoError(t, err)
	tmpFile, err := os.Open(tmp)
	assert.NoError(t, err)
	defer tmpFile.Close()

	{
		files := make(map[string]bool)
		ret, err := BuildMap(&Options{
			Filter: func(s string) bool {
				files[s] = true
				return true
			},
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, ret)
		assert.Contains(t, ret, tmp)
		assert.Contains(t, files, tmp)
		keys1 := make([]string, 0, len(ret))
		for k := range ret {
			keys1 = append(keys1, k)
		}
		keys2 := make([]string, 0, len(files))
		for k := range files {
			keys2 = append(keys2, k)
		}
		assert.Equal(t, len(keys1), len(keys2))
		sort.Strings(keys1)
		sort.Strings(keys2)
		assert.Equal(t, keys1, keys2)
	}
	{
		files := make(map[string]bool)
		ret, err := BuildMap(&Options{
			Filter: func(s string) bool {
				files[s] = true
				return false
			},
		})
		assert.NoError(t, err)
		assert.Empty(t, ret)
		assert.Contains(t, files, tmp)
	}
}
