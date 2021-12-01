package utils

import "os"

// CreateDirIfNotExist Check if dir exists, if not create it
func CreateDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

// CreateFileIfNotExist Check if file exists, if not create it
func CreateFileIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			LogError(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}
