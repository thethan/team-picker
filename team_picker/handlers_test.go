package team_picker_test

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/thethan/team_picker/core"
	"github.com/thethan/team_picker/team_picker"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"testing"
)


func TestAPI_GetTeams(t *testing.T) {
	gameId := rand.Int()
	api := getMockAPI()

	mockEcho := getMockEcho()
	mockEcho.SetParamNames("gameId")
	mockEcho.SetParamValues(strconv.Itoa(gameId))
	mockEcho.Set("JWT", core.GetMockJWTToken(gameId))

	err := api.GetTeams(mockEcho)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, mockEcho.Assertable.HttpCode)


	var rgx = regexp.MustCompile(`getenv\('(.*?)', true\)`)


	s := `getenv('WASHER', true), Ethan => getenv('DRYER', true)`
	rs := rgx.FindAllStringSubmatch(s, -1)

	for i := range rs {
		fmt.Println(rs[i][1])
	}

	fmt.Println(rs[1])
	assert.Equal(t, "WASHER",rs[0][1])

}





func getMockEcho() *core.EchoContextMock {

	mockable := core.EchoMockable{ContextData: map[string]interface{}{}}
	mockable.ContextData["JWT"] = ""
	assertable := core.EchoAssertable{}

	mockEcho := core.EchoContextMock{
		Mockable: mockable,
		Assertable: assertable,
	}

	return &mockEcho
}

func getMockAPI() *team_picker.API {
	return &team_picker.API{
		Dependencies: &core.Dependencies{
			Config: core.SetupTestConfig(),
			Logger: log.New(),

		},
	}
}
