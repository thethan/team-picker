package entities

type Teams struct {
	Teams []*Team `yaml:"teams",json:""`
}

type Team struct {
	Name string `json:"name",yaml:"name"`
	Conference string `json:"conference",yaml:"conference"`
	Division string `json:"division",yaml:"division"`
}

