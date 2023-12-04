package myapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gouProcess "github.com/yaoapp/gou/process"
)

func TestProcessHello(t *testing.T) {
	args := []interface{}{"hello world", "foo", "bar", 0.618}
	res, err := gouProcess.New("yao.myapi.hello", args...).Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []interface{}{"hello world", "foo", "bar", 0.618}, res.(map[string]interface{})["args"])
}
