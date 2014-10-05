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
		Convey("IsPermissionGroupName()\n", func() {
			So(models.Permissions().IsPermissionGroupName("perm.users.list"), ShouldEqual, true)
			So(models.Permissions().IsPermissionGroupName("perm.groups.read"), ShouldEqual, true)

			So(models.Permissions().IsPermissionGroupName("perms.users"), ShouldEqual, false)
			So(models.Permissions().IsPermissionGroupName("perms.users.list"), ShouldEqual, false)
			So(models.Permissions().IsPermissionGroupName("users.list"), ShouldEqual, false)
			So(models.Permissions().IsPermissionGroupName("foo"), ShouldEqual, false)
		})

		Convey("GroupName()\n", func() {
			So(models.Permissions().Name("perm.users.list"), ShouldEqual, "list_users")
		})

		Convey("Name()\n", func() {
			So(models.Permissions().GroupName("list_users"), ShouldEqual, "perm.users.list")
		})

		Convey("Abilities()\n", func() {
			So(models.UserPermissions().Abilities("foo")["list_users"], ShouldEqual, false)
			So(models.UserPermissions().Abilities("admin")["list_users"], ShouldEqual, false)

			backend.Type = backend.Memory
			models.Permissions().CreateAllInternal()
			models.Users().Create("admin", "admin")
			models.UserPermissions().Allow("admin", "list_users")
			So(models.UserPermissions().Abilities("admin")["list_users"], ShouldEqual, true)
			So(models.UserPermissions().Abilities("foo")["list_users"], ShouldEqual, false)

			models.UserPermissions().Deny("admin", "list_users")
			So(models.UserPermissions().Abilities("admin")["list_users"], ShouldEqual, false)
			So(models.UserPermissions().Abilities("foo")["list_users"], ShouldEqual, false)
		})
	})
}
