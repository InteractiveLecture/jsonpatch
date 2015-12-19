package jsonpatch

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatch(t *testing.T) {

	patchStr := `{"version" : 1,"operations": [{"op": "ADD", "path": "/biscuits/1", "value": {"name": "Ginger Nut"}}, {"op": "REMOVE", "path": "/biscuits"}]}`

	reader := strings.NewReader(patchStr)
	patch, err := Decode(reader)
	assert.Nil(t, err)
	assert.Equal(t, ADD, patch.Operations[0].Type)
	assert.Equal(t, 1, patch.Version)
	assert.Equal(t, "/biscuits/1", patch.Operations[0].Path)
	assert.Equal(t, 2, len(patch.Operations))

}
