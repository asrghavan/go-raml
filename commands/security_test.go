package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOauth2Middleware(t *testing.T) {
	Convey("oauth2 middleware", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/security/dropbox.raml")
		So(err, ShouldBeNil)

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("middleware generation test", func() {
			err = generateSecurity(apiDef, targetdir, "main")
			So(err, ShouldBeNil)

			// oauth 2 in header
			s, err := testLoadFile(filepath.Join(targetdir, "oauth2_oauth_2_0_headerMwr.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/oauth2_oauth_2_0_headerMwr.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// oauth 2 in query params
			s, err = testLoadFile(filepath.Join(targetdir, "oauth2_oauth_2_0_queryMwr.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/security/oauth2_oauth_2_0_queryMwr.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// scope matching
			s, err = testLoadFile(filepath.Join(targetdir, "oauth2_oauth_2_0_query_ADMINISTRATORMwr.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/security/oauth2_oauth_2_0_query_ADMINISTRATORMwr.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Convey("Go routes generation", func() {
			_, err = generateServerResources(apiDef, targetdir, "main", langGo)
			So(err, ShouldBeNil)

			// check route
			s, err := testLoadFile(filepath.Join(targetdir, "deliveries_if.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/deliveries_if.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}