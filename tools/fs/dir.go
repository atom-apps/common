package fs

import (
	"errors"
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirs(dirs ...string) (err error) {
	var exist bool
	for _, v := range dirs {
		exist, err = PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetFilePaths(dir string) ([]string, error) {
	paths := make([]string, 0)

	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		return paths, err
	}

	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return paths, err
	}

	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			subPaths, err := GetFilePaths(filePath)
			if err != nil {
				return paths, err
			}
			paths = append(paths, subPaths...)
		} else {
			paths = append(paths, filePath)
		}
	}
	return paths, nil
}
