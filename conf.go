package conf

import (
	"bufio"
	"bytes"
	"text/template"
)

type CONF struct {
	path string
	tpl  *template.Template
}

func New() *CONF {
	return &CONF{}
}

func (c *CONF) Init(path string) error {
	var e error
	c.path = path
	if c.tpl, e = template.ParseGlob(c.path); e != nil {
		return e
	} else {
		return nil
	}
}

func (c *CONF) RenderAsString(name string, data interface{}) string {
	var b bytes.Buffer
	var wr = bufio.NewWriter(&b)
	c.tpl.ExecuteTemplate(wr, name, data)
	wr.Flush()
	return b.String()
}
