package handleers

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"revision.go/models"
)

func TestValidateData(t *testing.T) {
	var err error
	users := `{"id":1,"name": "Bola","gender":"female"}`

	var use models.User
	_ = json.Unmarshal([]byte(users), &use)

	var res *models.User
	res, err = ValidateData(use)

	require.NoError(t, err)
	require.NotEmpty(t, res)

}

func TestValidateParameter(t *testing.T) {
	var err error
	var n string
	n = "name"
	n, err = ValidateParameter(n)
	require.NoError(t, err)
	require.NotEmpty(t, n)

}

func TestGetUsersResults(t *testing.T) {
	var err error
	v := `[{"id":1,"name":"kunle","gender":"female"},{"id":2,"name":"folake","gender":"female"}]`
	var users []models.User
	if err = json.Unmarshal([]byte(v), &users); err != nil {
		log.Fatal(err)
	}

	var us []models.User
	us, err = ValidateUsersResults(users)
	require.NoError(t, err)
	require.NotEmpty(t, us)
}
