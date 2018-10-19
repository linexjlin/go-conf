package conf

import (
	"testing"
)

var conf = New()

func TestInit(t *testing.T) {
	if e := conf.Init("sample/*"); e != nil {
		t.Log(e)
		t.Fail()
	}
}

func TestRenderAsString(t *testing.T) {
	res := conf.RenderAsString("main", nil)
	t.Log("\n", res)
}
