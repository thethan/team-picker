package team_picker

import (
	"github.com/thethan/team_picker/core"
	"github.com/thethan/team_picker/entities"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetTeamFromYaml(dependencies *core.Dependencies) (*entities.Team, error) {
	yamlfile, err := ioutil.ReadFile(dependencies.Config.DataLocation +"nfl-team.yaml")
	if err != nil {
		return nil, err
	}
	team := entities.Team{}
	err2 := yaml.Unmarshal(yamlfile, &team)
	if err2 != nil {
		return nil, err2
	}
	return &team, nil
}

func GetTeamsFromYaml(dependencies *core.Dependencies) (*entities.Teams ,error) {

	yamlfile, err := ioutil.ReadFile(dependencies.Config.DataLocation +"nfl-teams.yaml")
	if err != nil {
		return nil, err
	}
	teamSlice := make([]*entities.Team, 32)
	teams := entities.Teams{Teams: teamSlice}
	err2 := yaml.Unmarshal(yamlfile, &teams)
	if err2 != nil {
		return nil, err2
	}
	return &teams, nil
}
