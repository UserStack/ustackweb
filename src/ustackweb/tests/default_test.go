package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
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

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Root Endpoint without sign in\n", t, func() {
		Convey("Redirect", func() {
			So(w.Code, ShouldEqual, 302)
			So(w.HeaderMap.Get("Location"), ShouldEqual, "/sign_in")
		})
	})
}
