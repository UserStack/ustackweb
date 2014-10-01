package test

import (
	"testing"

	"github.com/UserStack/ustackweb/backend"
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

		Convey("Abilities()\n", func() {
			So(models.Permissions().Abilities("foo")["list_users"], ShouldEqual, false)
			So(models.Permissions().Abilities("admin")["list_users"], ShouldEqual, false)

			backend.Type = backend.Memory
			models.Permissions().Create()
			models.Users().Create("admin", "admin")
			models.Permissions().Allow("admin", "list_users")
			So(models.Permissions().Abilities("admin")["list_users"], ShouldEqual, true)
			So(models.Permissions().Abilities("foo")["list_users"], ShouldEqual, false)

			models.Permissions().Deny("admin", "list_users")
			So(models.Permissions().Abilities("admin")["list_users"], ShouldEqual, false)
			So(models.Permissions().Abilities("foo")["list_users"], ShouldEqual, false)
		})
	})
}
