package team_picker

import (
	"github.com/stretchr/testify/assert"
	"github.com/thethan/team_picker/core"
	"github.com/thethan/team_picker/entities"
	"testing"
)


func Test_GetTeamFromYaml(t *testing.T) {
	config := core.SetupTestConfig()
	dep := core.SetupTestDependencies()
	dep.Config = config
	result, err := GetTeamFromYaml(dep)
	assert.Nil(t, err)
	assert.Equal(t, &entities.Team{Name: "Greenbay Packers", Conference:"NFC", Division:"North"}, result)

}

func Test_GetTeamsFromYaml(t *testing.T) {
	config := core.SetupTestConfig()
	dep := core.SetupTestDependencies()
	dep.Config = config
	result, err := GetTeamsFromYaml(dep)
	assert.Nil(t, err)
	assert.IsType(t, &entities.Teams{}, result)

	assert.Equal(t, 32, len(result.Teams))
}