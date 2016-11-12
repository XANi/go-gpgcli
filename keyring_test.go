package gpgcli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFingerprint(t *testing.T) {
	g, _ := New()
	out, err := g.GetFingerprint("some test print")
	Convey("N", t, func() {
		So(err, ShouldEqual, nil)
		So(out, ShouldEqual, "some test fprint")
	})
}
