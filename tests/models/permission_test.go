package test

import (
	"testing"

	"github.com/UserStack/ustackweb/models"
	. "github.com/smartystreets/goconvey/convey"
)

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Permissions()\n", t, func() {
		Convey("GroupName()\n", func() {
			So(models.Permissions().Name("perm.users.list"), ShouldEqual, "list_users")
		})

		Convey("Name()\n", func() {
			So(models.Permissions().GroupName("list_users"), ShouldEqual, "perm.users.list")
		})
	})
}
