package main

// Project model
type Project struct {
	Name      string
	Status    string
	Slug      string
	FinalMark float64
	Validated bool
}

// BuildProject builds a project structure form project raw data
func BuildProject(descriptions map[string]interface{}) Project {
	rawProject := descriptions["project"].(map[string]interface{})
	name := rawProject["name"].(string)
	status := descriptions["status"].(string)
	slug := rawProject["slug"].(string)

	var finalMark float64
	var validated bool
	if descriptions["validated?"] == nil {
		validated = false
		finalMark = 0
	} else {
		validated = descriptions["validated?"].(bool)
		finalMark = descriptions["final_mark"].(float64)
	}
	return Project{Name: name, Status: status, FinalMark: finalMark, Validated: validated, Slug: slug}
}
