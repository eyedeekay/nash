package builtin_test

import (
	"bytes"
	"testing"

	"github.com/NeowayLabs/nash/internal/sh/internal/fixture"
)

func TestLen(t *testing.T) {
	sh, cleanup := fixture.SetupShell(t)
	defer cleanup()

	var out bytes.Buffer

	sh.SetStdout(&out)

	err := sh.Exec(
		"test len",
		`var a = (1 2 3 4 5 6 7 8 9 0)
		 var len_a <= len($a)
		 echo -n $len_a`,
	)

	if err != nil {
		t.Error(err)
		return
	}

	if "10" != string(out.Bytes()) {
		t.Errorf("String differs: '%s' != '%s'", "10", string(out.Bytes()))
		return
	}

	// Having to reset is a bad example of why testing N stuff together
	// is asking for trouble :-).
	out.Reset()

	err = sh.Exec(
		"test len fail",
		`a = "test"
		 var l <= len($a)
		 echo -n $l
		`,
	)

	if err != nil {
		t.Error(err)
		return
	}

	if "4" != string(out.Bytes()) {
		t.Errorf("String differs: '%s' != '%s'", "4", string(out.Bytes()))
		return
	}
}
