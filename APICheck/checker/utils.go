package main

import (
	"fmt"
	"strings"

	apiclient "github.com/BoyerDamien/42APIClient"
)

// CheckProjectStatus compares the validation status between the old and new project data
func CheckProjectStatus(username string, p1, p2 *Project) error {
	if p1.FinalMark != p2.FinalMark && p2.Validated {
		return fmt.Errorf("%s a validé le projet %s à %2.f%%! Félicitation à toi!", username, p2.Name, p2.FinalMark)
	}
	if p1.FinalMark != p2.FinalMark && !p2.Validated {
		if strings.Contains(p2.Name, "Exam") {
			return nil
		}
		return fmt.Errorf("%s a raté le projet %s à %2.f%%! Boooh!", username, p2.Name, p2.FinalMark)
	}
	if p1.FinalMark != p2.FinalMark && p1.FinalMark == -42 {
		return fmt.Errorf("@everyone %s a triché au projet %s! Quelle honte!", username, p2.Name)
	}
	return nil
}

// CheckProjectSubscribed checks whether the user has subscribed to a new project or not
func CheckProjectSubscribed(dbUser, apiUser *apiclient.User) error {
	dbUserProjectsLen := len(dbUser.ProjectsUsers)
	apiUserProjectsLen := len(apiUser.ProjectsUsers)
	if dbUserProjectsLen > 0 && dbUserProjectsLen < apiUserProjectsLen {
		project := BuildProject(apiUser.ProjectsUsers[0])
		return fmt.Errorf("%s s'est inscrit au projet %s! Bon courage à toi!", apiUser.Login, project.Name)
	}
	return nil
}

// CheckUserLocation checks whether a user is login in a cluster or not
func CheckUserLocation(dbUser, apiUser *apiclient.User) error {
	if apiUser.Location != dbUser.Location {
		if apiUser.Location == nil {
			return fmt.Errorf("%s s'est déconnecté", apiUser.Login)
		}
		return fmt.Errorf("%s s'est connecté en %s", apiUser.Login, apiUser.Location)
	}
	return nil
}

func CheckUserStudent(dbUser, apiUser *apiclient.User) error {
	for _, cursus := range dbUser.CursusUsers {
		if cursus.Cursus.Slug != "c-piscine" {
			return nil
		}
	}
	for _, cursus := range apiUser.CursusUsers {
		if cursus.Cursus.Slug != "c-piscine" {
			return fmt.Errorf("%s est maintenant student! Bienvenue à 42", apiUser.Login)
		}
	}
	return nil
}
