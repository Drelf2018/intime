package intime_test

import (
	"html/template"
	"os"
	"testing"

	"github.com/Drelf2018/intime"
)

func TestNew(t *testing.T) {
	d := intime.New(nil)
	t.Log(d.Time())

	d = intime.New("20240321", "format", "%Y%m%d")
	t.Log(d)
}

type Post struct {
	Time intime.Time
}

func TestPost(t *testing.T) {
	post := Post{Time: intime.Now()}
	post.Time.Format()
	tmpl, err := template.New("example").Parse(`{{.Time.Format "%H:%M:%S"}}`)
	if err != nil {
		t.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, post)
	if err != nil {
		t.Fatal(err)
	}
}
