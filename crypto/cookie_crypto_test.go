package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncrypt(t *testing.T) {
	Convey("Test Encrypt", t, func() {
		c := CookieCrypto{}
		c.UpdateKeys("0F10F6CB2F5369C14D14FA07BAD302267901240CC8C845DD2C645FBD149A11C9", "C985085862F161091EEEFE30F7DC9D62")

		s := "12345678"
		So(c.Decrypt(c.Encrypt(s)), ShouldEqual, s)

		s = "dadj9813bq78yd879ayhduyiahd78y278hui1"
		So(c.Decrypt(c.Encrypt(s)), ShouldEqual, s)

		s = "1"
		So(c.Decrypt(c.Encrypt(s)), ShouldEqual, s)

		s = "this is a text book, test adadfnjkn djkah jkdfjlkab jkasdhjk bjkadnkab jkbajhfa jkbfjhkahjkdf uk"
		So(c.Decrypt(c.Encrypt(s)), ShouldEqual, s)

		s = "123454"
		So(c.Decrypt(c.Encrypt(s)), ShouldEqual, s)

		s = "e941c654c870cb9e32e7f6259972d79452f41a97a3b5095030963050ce0cdfe3d6ade89e919ec3d5"
		b := "1313"
		So(c.Decrypt(s), ShouldEqual, b)

		s = "e941c654c870cb9e32e7f6259972d79452f41a97a3b5095030963050ce0cdfe3d6ade89e919ec3d"
		So(c.Decrypt(s), ShouldNotEqual, b)

		s = "a38a722f0aa5c8a68ae3340265b6d692c55f2dc37643f72b4269534cfaba54b7b19d8b352e6a663a"
		userID := c.Decrypt(s)
		So(userID, ShouldEqual, "3328863")
		t.Log(userID)
	})
}
