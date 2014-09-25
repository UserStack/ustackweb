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
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.SessionOn = true
	beego.TestBeegoInit(apppath)
}

func getRequest(method string, urlStr string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, urlStr, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
	return w
}

func postRequest(method string, resourcePath string, data *url.Values) *httptest.ResponseRecorder {
	u, _ := url.ParseRequestURI("/")
	u.Path = resourcePath
	urlStr := fmt.Sprintf("%v", u)

	r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())
	return w
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	Convey("Redirect to Sign In\n", t, func() {
		w := getRequest("GET", "/")
		Convey("Redirect", func() {
			So(w.Code, ShouldEqual, 302)
			So(w.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})

	Convey("Successful Sign In\n", t, func() {
		data := url.Values{}
		data.Add("Username", "admin")
		data.Add("Password", "bar")
		response := postRequest("POST", "/sign_in", &data)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/profile")
		})
	})

	Convey("Failed Sign In\n", t, func() {
		data := url.Values{}
		data.Add("Username", "adminx")
		data.Add("Password", "barx")
		response := postRequest("POST", "/sign_in", &data)
		Convey("Redirect", func() {
			So(response.Code, ShouldEqual, 302)
			So(response.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})
}
