package cloudmeta

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spotmaxtech/gokit"
	"testing"
)

func TestAWSRegion(t *testing.T) {
	Convey("test use case", t, func() {
		consul := gokit.NewConsul(TestConsulAddress)
		Convey("test consul fetch", func() {
			region := NewAWSRegion(TestConsulRegionKey)
			err := region.Fetch(consul)
			So(err, ShouldBeNil)

			So(region.data["us-east-1"].Name, ShouldEqual, "us-east-1")
		})
	})
}
