package main

import "fmt"

type Project struct {
	Name      string
	Status    string
	Validated bool
}

func BuildProject(descriptions []map[string]interface{}) Project {
	/*rawProject := descriptions[2]["Value"].([]map[string]interface{})
	name := rawProject[3]["Value"].(string)
	status := descriptions[6]["Value"].(string)
	validated := descriptions[7]["Value"].(bool)*/
	fmt.Println(len(descriptions))
	return Project{Name: "test", Status: "ok", Validated: true}
}
