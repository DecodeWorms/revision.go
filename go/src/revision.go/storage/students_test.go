package storage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"revision.go/types"
)

var s Students

func TestSave(t *testing.T) {
	var err error
	users := `{"id":1,"name": "Bola","gender":"female"}`

	var use types.User
	if err = json.Unmarshal([]byte(users), &use); err != nil {
		panic(err)
	}

	var res *types.User
	res, err = SaveData(use)

	require.NoError(t, err)
	require.NotEmpty(t, res)

}

func TestCheckNameLength(t *testing.T) {
	var err error
	var n string
	n = "Biolas"
	n, err = CheckNameLength(n)
	require.NoError(t, err)
	require.NotEmpty(t, n)

}

func TestGetSlice(t *testing.T) {
	var err error
	v := `[{"id":1,"name":"kunle","gender":"female"},{"id":2,"name":"folake","gender":"female"}]`
	var users []types.User
	if err = json.Unmarshal([]byte(v), &users); err != nil {
		panic(err)
	}

	var us []types.User
	us, err = GetSlice(users)
	require.NoError(t, err)
	require.NotEmpty(t, us)
}

func TestCheckForTwoStrings(t *testing.T) {
	var err error

	old, new := "Deola", "Yinka"
	var o, n string

	o, n, err = CheckForTwoStrings(old, new)
	require.NoError(t, err)
	require.NotEmpty(t, o, n)

}
