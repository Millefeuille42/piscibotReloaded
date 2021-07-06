package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/debug"

	apiclient "github.com/BoyerDamien/42APIClient"
)

// CheckEnvVariables tests the existence of required env variables
func CheckEnvVariables() {
	envVariables := []string{"UID", "SECRET", "PORT"}
	for _, val := range envVariables {
		if os.Getenv(val) == "" {
			log.Fatalf("Missing %s env variable", val)
		}
		log.Printf("%s env variable [OK]", val)
	}
}

// logError Prints error + StackTrace to stderr if error
func logError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err, string(debug.Stack()))
	}
}

// createDirIfNotExist Check if dir exists, if not create it
func createDirIfNotExist(path string) error {
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

// createFileIfNotExist Check if file exists, if not create it
func createFileIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			logError(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func writeUserData(data apiclient.User) error {
	dataBytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		logError(err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("./data/%s.json", data.Login), dataBytes, 0677)
	if err != nil {
		logError(err)
		return err
	}
	return nil
}

func readUserData(login string) (apiclient.User, error) {
	target := apiclient.User{}

	fileData, err := ioutil.ReadFile(fmt.Sprintf("./data/targets/%s.json", id))
	if err != nil {
		logError(err)
		return apiclient.User{}, err
	}

	err = json.Unmarshal(fileData, &target)
	if err != nil {
		logError(err)
		return apiclient.User{}, err
	}

	return target, nil
}
