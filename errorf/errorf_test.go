package errorf_test

import (
	"testing"

	"github.com/Drelf2018/intime/errorf"
)

var err = errorf.New("code %d", 404)

func TestErrorf(t *testing.T) {
	t.Log(err)
	t.Log(err.Format(400))
}
