package backends

import (
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	// works first time
	uid, _ := backend.CreateUser("test@example.com", "secret")
	if uid <= 0 {
		t.Fatal("expect uid to be greater than 0")
	}

	// but fails second time with EEXIST
	_, berr2 := backend.CreateUser("test@example.com", "secret")
	if berr2.Code != "EEXIST" {
		t.Fatal("should return EEXIST instead of", berr2.Code)
	}

	// and fails with invalid data password ...
	_, berr3 := backend.CreateUser("test@example.com", "")
	if berr3.Code != "EINVAL" {
		t.Fatal("should return EINVAL instead of", berr3.Code)
	}

	// ... or name
	_, berr4 := backend.CreateUser("", "secret")
	if berr4.Code != "EINVAL" {
		t.Fatal("should return EINVAL instead of", berr4.Code)
	}

	// but should work with regular user
	_, berr5 := backend.CreateUser("test1", "secret")
	if berr5 != nil {
		t.Fatal("should not return error", berr5.Code, berr5.Message)
	}
}

func TestEnableDisableUser(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	// it rejects invalid calls
	inval := backend.DisableUser("")
	if inval.Code != "EINVAL" {
		t.Fatal("should return EINVAL instead of", inval.Code)
	}

	// fails to disable unknown user
	enoent := backend.DisableUser("test@example.com")
	if enoent.Code != "ENOENT" {
		t.Fatal("should fail to disable unkown user with code ENOENT but was",
			enoent.Code)
	}

	// create a user and check he can't login after he was disabled
	backend.CreateUser("test@example.com", "secret")
	_ = backend.DisableUser("test@example.com")
	_, err := backend.LoginUser("test@example.com", "secret")
	if err == nil {
		t.Fatal("should fail to login a user")
	}

	backend.EnableUser("test@example.com")
	uid, err := backend.LoginUser("test@example.com", "secret")
	if err != nil {
		t.Fatal("should fail to login a user")
	}
	if uid != 1 {
		t.Fatal("should have logged in the user after activation")
	}
}

func TestSetGetUserData(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	err := backend.SetUserData("test@example.com", "firstname", "Tester")
	if err.Code != "ENOENT" {
		t.Fatal("should fail to set value on non existing user", err.Code)
	}

	err1 := backend.SetUserData("1", "firstname", "Tester")
	if err1.Code != "ENOENT" {
		t.Fatal("should fail to set value on non existing user", err1.Code)
	}

	_, err2 := backend.GetUserData("test@example.com", "firstname")
	if err2.Code != "ENOENT" {
		t.Fatal("should fail to set value on non existing user", err2.Code)
	}

	_, err3 := backend.GetUserData("1", "firstname")
	if err3.Code != "ENOENT" {
		t.Fatal("should fail to set value on non existing user", err3.Code)
	}

	intval := backend.SetUserData("", "firstname", "Tester")
	if intval.Code != "EINVAL" {
		t.Fatal("should fail to set value on non invalid name", intval.Code)
	}

	intval1 := backend.SetUserData("test@example.com", "", "Tester")
	if intval1.Code != "EINVAL" {
		t.Fatal("should fail to set value on non invalid key", intval1.Code)
	}

	intval2 := backend.SetUserData("test@example.com", "firstname", "")
	if intval2.Code != "EINVAL" {
		t.Fatal("should fail to set value on non invalid value", intval2.Code)
	}

	_, intval3 := backend.GetUserData("", "firstname")
	if intval3.Code != "EINVAL" {
		t.Fatal("should fail to set value on non invalid value", intval3.Code)
	}

	_, intval4 := backend.GetUserData("test@example.com", "")
	if intval4.Code != "EINVAL" {
		t.Fatal("should fail to set value on non invalid value", intval4.Code)
	}

	backend.CreateUser("test@example.com", "secret")
	backend.SetUserData("test@example.com", "firstname", "Tester")
	val, _ := backend.GetUserData("test@example.com", "firstname")
	if val != "Tester" {
		t.Fatal("the value should have been 'Tester' but was", val)
	}
}

func TestLoginUser(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	_, err := backend.LoginUser("", "secret")
	if err.Code != "EINVAL" {
		t.Fatal("should have detected a blank username", err.Code)
	}

	_, err1 := backend.LoginUser("test0@example.com", "")
	if err1.Code != "EINVAL" {
		t.Fatal("should have detected a blank password", err1.Code)
	}

	_, err2 := backend.LoginUser("test0@example.com", "secret")
	if err2.Code != "ENOENT" {
		t.Fatal("should have reported, that user is unknown but was", err2.Code)
	}

	uid, _ := backend.CreateUser("test@example.com", "secret")
	uid2, _ := backend.LoginUser("test@example.com", "secret")
	if uid != uid2 {
		t.Fatal("should have found the same user with uid", uid,
			"but found", uid2)
	}
}

func TestChangeUserPassword(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	ierr1 := backend.ChangeUserPassword("", "secret2", "secret2")
	if ierr1.Code != "EINVAL" {
		t.Fatal("should fail since user value is empty")
	}

	ierr2 := backend.ChangeUserPassword("test@example.com", "", "secret2")
	if ierr2.Code != "EINVAL" {
		t.Fatal("should fail since passwd value is empty")
	}

	ierr3 := backend.ChangeUserPassword("test@example.com", "secret2", "")
	if ierr3.Code != "EINVAL" {
		t.Fatal("should fail since new passwd value is empty")
	}

	backend.CreateUser("test@example.com", "secret")
	backend.LoginUser("test@example.com", "secret")
	serr := backend.ChangeUserPassword("test@example.com", "secret2", "secret2")
	// fails if password is wrong
	if serr.Code != "ENOENT" {
		t.Fatal("Should have failed to change with wrong password")
	}
	backend.ChangeUserPassword("test@example.com", "secret", "secret2")
	_, err := backend.LoginUser("test@example.com", "secret2")
	if err != nil {
		t.Fatalf("User passwd should have been changed: %s, (%s)", err.Code, err.Message)
	}
}

func TestChangeUserName(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	ierr1 := backend.ChangeUserName("", "secret2", "test2@example.com")
	if ierr1.Code != "EINVAL" {
		t.Fatal("should fail since name value is empty")
	}

	ierr2 := backend.ChangeUserName("test@example.com", "", "test2@example.com")
	if ierr2.Code != "EINVAL" {
		t.Fatal("should fail since passwd value is empty")
	}

	ierr3 := backend.ChangeUserName("test@example.com", "secret2", "")
	if ierr3.Code != "EINVAL" {
		t.Fatal("should fail since new name value is empty")
	}

	backend.CreateUser("test@example.com", "secret")
	backend.LoginUser("test@example.com", "secret")
	serr := backend.ChangeUserName("test@example.com", "secret2", "test2@example.com")
	if serr.Code != "ENOENT" {
		t.Fatal("Should have failed to change with wrong password")
	}
	backend.ChangeUserName("test@example.com", "secret", "test2@example.com")
	_, err := backend.LoginUser("test2@example.com", "secret")
	if err != nil {
		t.Fatalf("User name should have been changed: %s, (%s)", err.Code, err.Message)
	}
}

func TestDeleteUser(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	err := backend.DeleteUser("")
	if err.Code != "EINVAL" {
		t.Fatal("should error on missing parameter")
	}

	// create two users
	backend.CreateUser("test0@example.com", "secret")
	backend.CreateUser("test-1411718662776791574", "secret")

	// now has two users
	users, _ := backend.Users()
	if len(users) != 2 {
		t.Fatal("user count should have been 2 but was", len(users))
	}

	// delete one using uid
	derr1 := backend.DeleteUser("1")
	if derr1 != nil {
		t.Fatal("should not error on delete", derr1.Code, derr1.Message)
	}
	users, _ = backend.Users()
	if len(users) != 1 {
		t.Fatal("user count should have been 0 but was", len(users))
	}

	// delete one using name
	derr := backend.DeleteUser("test-1411718662776791574")
	if derr != nil {
		t.Fatal("should not error on delete", derr.Code, derr.Message)
	}
	users, _ = backend.Users()
	if len(users) != 0 {
		t.Fatal("user count should have been 1 but was", len(users))
	}
}

func TestUsers(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	users, _ := backend.Users()
	if len(users) != 0 {
		t.Fatal("users should be empty and is ", len(users))
	}

	// create a user
	uid, _ := backend.CreateUser("test@example.com", "secret")
	if uid <= 0 {
		t.Fatal("expect uid to be greater than 0")
	}

	// the list should have one enties now
	users, _ = backend.Users()
	if len(users) != 1 {
		t.Fatal("users should be empty and is ", len(users))
	}
	if users[0].Name != "test@example.com" {
		t.Fatal("name should have been 'test@example.com' but was", users[0].Name)
	}
	if users[0].Uid != 1 {
		t.Fatal("name should have been 1 but was", users[0].Uid)
	}
}

func TestGroup(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	_, err := backend.CreateGroup("")
	if err.Code != "EINVAL" {
		t.Fatal("should return EINVAL instead of", err.Code)
	}

	gid, cerr := backend.CreateGroup("developers")
	if gid <= 0 {
		t.Fatal("should have created group", cerr.Code, cerr.Message)
	}

	_, eerr := backend.CreateGroup("developers")
	if eerr.Code != "EEXIST" {
		t.Fatal("should return EEXIST instead of", eerr.Code, eerr.Message)
	}
}

func TestDeleteGroup(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	err := backend.DeleteUser("")
	if err.Code != "EINVAL" {
		t.Fatal("should error on missing parameter")
	}

	// create two users
	backend.CreateGroup("dev")
	backend.CreateGroup("sales")

	// now has two users
	groups, _ := backend.Groups()
	if len(groups) != 2 {
		t.Fatal("user count should have been 2 but was", len(groups))
	}

	// delete one using uid one using name
	derr1 := backend.DeleteGroup("dev")
	if derr1 != nil {
		t.Fatal("should not error on delete", derr1.Code, derr1.Message)
	}
	derr1 = backend.DeleteGroup("2")
	if derr1 != nil {
		t.Fatal("should not error on delete", derr1.Code, derr1.Message)
	}
	groups, _ = backend.Groups()
	if len(groups) != 0 {
		t.Fatal("user count should have been 0 but was", len(groups))
	}
}

func TestGroups(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	gid0, _ := backend.CreateGroup("developers")
	gid1, _ := backend.CreateGroup("sales")
	gid2, _ := backend.CreateGroup("admins")
	groups, _ := backend.Groups()

	expected := []Group{
		Group{Gid: gid0, Name: "developers"},
		Group{Gid: gid1, Name: "sales"},
		Group{Gid: gid2, Name: "admins"},
	}

	if !reflect.DeepEqual(expected, groups) {
		t.Fatalf("expected %v\nto equal %v\n", expected, groups)
	}
}

func TestUsersAndGroupsAssociations(t *testing.T) {
	backend, dberr := NewSqliteBackend(":memory:")
	if dberr != nil {
		t.Fatal(dberr)
	}
	defer backend.Close()

	backend.CreateGroup("developers")
	backend.CreateGroup("admins")

	backend.CreateUser("joe", "secret")
	backend.CreateUser("mike", "secret")

	// User Groups
	groups, _ := backend.UserGroups("joe")
	if len(groups) != 0 {
		t.Fatal("expected joe not to have any groups")
	}
	backend.AddUserToGroup("joe", "1")
	backend.AddUserToGroup("1", "admins")
	expectedGroups := []Group{
		Group{Gid: 1, Name: "developers"},
		Group{Gid: 2, Name: "admins"},
	}
	groups, _ = backend.UserGroups("joe")
	if !reflect.DeepEqual(expectedGroups, groups) {
		t.Fatalf("expected joe have groups %v but has %v", expectedGroups, groups)
	}

	// Group Users
	users, _ := backend.GroupUsers("admins")
	if len(users) != 1 {
		t.Fatalf("expected to have one admin, got %v", users)
	}
	backend.AddUserToGroup("mike", "admins")
	users, _ = backend.GroupUsers("admins")
	expectedUsers := []User{
		User{Uid: 1, Name: "joe"},
		User{Uid: 2, Name: "mike"},
	}
	if reflect.DeepEqual(expectedUsers, users) {
		t.Fatalf("expected admins have users %v but has %v", expectedUsers, users)
	}

	// Remove associations
	backend.RemoveUserFromGroup("mike", "admins")
	backend.RemoveUserFromGroup("joe", "admins")
	admins, _ := backend.GroupUsers("admins")
	if len(admins) != 0 {
		t.Fatalf("expected to have no admin, got %v", admins)
	}
	users, _ = backend.GroupUsers("admins")
	if len(users) != 0 {
		t.Fatalf("expected to have no admin, got %v", users)
	}
}
