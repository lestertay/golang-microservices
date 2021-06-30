package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoParam(t *testing.T) {
	req := CreateRepoParam{
		Name:        "golang introduction",
		Description: "a golang introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(req)
	assert.Nil(t, err)

	var target CreateRepoParam
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.Equal(t, target, req)
}
