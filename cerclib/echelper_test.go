package cerclib

import (
	_ "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpace(t *testing.T) {
	Convey("Test 1 == 1", t, func() {
		x := 1
		Convey("When x incremented", func() {
			x++

			So(x, ShouldEqual, 2)
		})
	})

	ethSignature := "0x61d864a4e35f81c3907a1485a638ff6aab3d4dc9c3ebee77c68ef1e89597c9663a45056f79851453cd1fbcf6c664161170d67ad06fe5f9da3266afcebeec40ce1c"
	ebase := EcRevoerFeilds{
		EthAddress:        "0x627306090abab3a6e1400e9345bc60c78a8bef57",
		Msg:               "idhub",
		EthereumSignature: ethSignature,
	}

	Convey("Test EC Recover", t, func() {
		contain, err := Resolve(ebase)
		So(err, ShouldEqual, nil)
		So(*contain, ShouldEqual, ebase.EthAddress)
	})
}
