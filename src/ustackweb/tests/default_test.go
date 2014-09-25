package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	_ "ustackweb/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.TestBeegoInit(apppath)
}

type Session struct {
	username string
}

func recordRequest(r *http.Request, session *Session) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	if session != nil {
		s := beego.GlobalSessions.SessionStart(w, r)
		s.Set("username", session.username)
	}
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
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
	var nilSession *Session
	adminSession := &Session{username: "admin"}

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
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/profile")
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
		data.Add("Password", "bar")
		response := postRequest("POST", "/sign_in", &data, nilSession)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/profile")
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
}
