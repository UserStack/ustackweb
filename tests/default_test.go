package test

import (
	"bytes"
	"fmt"
	_ "github.com/UserStack/ustackweb/routers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/UserStack/ustackweb/backend"
	"github.com/UserStack/ustackweb/models"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	backend.Type = backend.Memory
	beego.TestBeegoInit(apppath)
}

type Session struct {
	username string
}

func SetupDatabase() {
	models.Permissions().Create()
	models.Users().Create("admin", "admin")
	models.Users().Create("enduser", "enduser")
	models.Permissions().AllowAll("admin")
}

func recordRequest(r *http.Request, session *Session) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	if session != nil {
		s := beego.GlobalSessions.SessionStart(w, r)
		s.Set("username", session.username)
	}
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	// beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
	return w
}

func getRequest(method string, urlStr string, session *Session) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, urlStr, nil)
	return recordRequest(r, session)
}

func postRequest(method string, resourcePath string, data *url.Values, session *Session) *httptest.ResponseRecorder {
	u, _ := url.ParseRequestURI("/")
	u.Path = resourcePath
	urlStr := fmt.Sprintf("%v", u)

	r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	return recordRequest(r, session)
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	SetupDatabase()
	var nilSession *Session
	adminSession := &Session{username: "admin"}
	enduserSession := &Session{username: "enduser"}

	Convey("Redirect to Sign In\n", t, func() {
		response := getRequest("GET", "/", nilSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})

	Convey("Redirect to Profile when already Signed In\n", t, func() {
		response := getRequest("GET", "/", adminSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 200)
			So(response.Body.String(), ShouldContainSubstring, "Home")
		})
	})

	Convey("Shows Sign In\n", t, func() {
		response := getRequest("GET", "/sign_in", nilSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 200)
			So(response.Body.String(), ShouldContainSubstring, "Sign In")
		})
	})

	Convey("Successful Sign In\n", t, func() {
		data := url.Values{}
		data.Add("Username", "admin")
		data.Add("Password", "admin")
		response := postRequest("POST", "/sign_in", &data, nilSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/")
		})
	})

	Convey("Failed Sign In\n", t, func() {
		data := url.Values{}
		data.Add("Username", "adminx")
		data.Add("Password", "barx")
		response := postRequest("POST", "/sign_in", &data, nilSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})

	Convey("Users without Sign In\n", t, func() {
		response := postRequest("GET", "/users", &url.Values{}, nilSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})

	Convey("Users\n", t, func() {
		response := postRequest("GET", "/users", &url.Values{}, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 200)
		})
	})

	Convey("Users Unauthorized\n", t, func() {
		response := postRequest("GET", "/users", &url.Values{}, enduserSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 401)
			So(response.Body.String(), ShouldContainSubstring, "Unauthorized")
		})
	})

	Convey("Create User\n", t, func() {
		data := url.Values{}
		data.Add("Username", "mikes")
		data.Add("Password", "micke")
		response := postRequest("POST", "/users", &data, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldStartWith, "/users/")
		})
	})

	Convey("Create User Error\n", t, func() {
		response := postRequest("POST", "/users", &url.Values{}, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 200)
			So(response.Body.String(), ShouldContainSubstring, "Can not be empty")
		})
	})

	Convey("Create User Unauthorized\n", t, func() {
		response := postRequest("POST", "/users", &url.Values{}, enduserSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 401)
			So(response.Body.String(), ShouldContainSubstring, "Unauthorized")
		})
	})

	Convey("Create Group\n", t, func() {
		data := url.Values{}
		data.Add("Name", "foo")
		response := postRequest("POST", "/groups", &data, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldStartWith, "/groups")
		})
	})

	Convey("Create Group Error\n", t, func() {
		response := postRequest("POST", "/groups", &url.Values{}, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 200)
			So(response.Body.String(), ShouldContainSubstring, "Can not be empty")
		})
	})

	Convey("Create Group Unauthorized\n", t, func() {
		response := postRequest("POST", "/groups", &url.Values{}, enduserSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 401)
			So(response.Body.String(), ShouldContainSubstring, "Unauthorized")
		})
	})

	Convey("Create Permission Group\n", t, func() {
		data := url.Values{}
		data.Add("Name", "perm.foo1.bar")
		response := postRequest("POST", "/groups", &data, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldStartWith, "/groups")
		})
	})

	Convey("Create Permission Group Unauthorized\n", t, func() {
		models.Permissions().Deny("admin", "create_permissions")
		data := url.Values{}
		data.Add("Name", "perm.foo2.bar")
		response := postRequest("POST", "/groups", &data, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 401)
			So(response.Body.String(), ShouldContainSubstring, "Unauthorized")
		})
	})

	Convey("Delete Permission Group\n", t, func() {
		models.Permissions().Allow("admin", "delete_groups")
		models.Permissions().Allow("admin", "delete_permissions")
		models.Groups().Create("perm.foo.bar")
		response := postRequest("GET", "/groups/perm.foo.bar/delete", &url.Values{}, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldStartWith, "/groups")
		})
	})

	Convey("Delete Permission Group Unauthorized\n", t, func() {
		models.Permissions().Deny("admin", "delete_permissions")
		response := postRequest("GET", "/groups/perm.foo.bar/delete", &url.Values{}, adminSession)
		Convey("Render", func() {
			So(response.Code, ShouldEqual, 401)
			So(response.Body.String(), ShouldContainSubstring, "Unauthorized")
		})
	})
}
