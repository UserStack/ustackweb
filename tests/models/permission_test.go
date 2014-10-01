package test

import (
	"testing"

	"github.com/UserStack/ustackweb/models"
	. "github.com/smartystreets/goconvey/convey"
)

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Permission\n", t, func() {
		Convey("Name()\n", func() {
			permission := models.Permission{GroupName: "perm.users.list"}
			So(permission.Name(), ShouldEqual, "list_users")
		})

	})

	Convey("Name()\n", t, func() {
		groupName := models.GroupNameFromPermissionName("list_users")
		So(groupName, ShouldEqual, "perm.users.list")
	})
}
