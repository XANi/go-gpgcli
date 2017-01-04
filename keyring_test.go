package gpgcli

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFingerprint(t *testing.T) {
	g, _ := New()
	// TODO include some test gpg config
	out, err := g.GetFingerprintById("a")
	Convey("N", t, func() {
		So(err, ShouldEqual, nil)
		So(out, ShouldContainSubstring, "0")
	})
}
