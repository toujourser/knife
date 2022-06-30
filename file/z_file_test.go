package file

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFile(t *testing.T) {
	Convey("测试文件操作", t, func() {
		Convey("文件或者文件夹是否存在", func() {
			So(CheckExist("./upload.go"), ShouldBeTrue)
			So(CheckExist("./upload3.go"), ShouldBeFalse)
		})
		Convey("文件后缀", func() {
			So(GetExt("upload.go"), ShouldEqual, ".go")
			So(GetExt("upload.txt"), ShouldEqual, ".txt")
		})
		Convey("创建目录", func() {
			So(PathCreate("./upload_dir"), ShouldBeNil)
		})
		Convey("文件夹不存在则创建", func() {
			So(IsNotExistMkDir("./upload_dir"), ShouldBeNil)
		})
		Convey("创建文件", func() {
			b := bytes.NewBuffer([]byte("hello world"))
			So(FileCreate(b, "read.txt"), ShouldBeNil)
		})
		Convey("获取文件大小", func() {
			So(GetFileSize("./read.txt"), ShouldEqual, 11)
		})
		Convey("获取当前路径", func() {
			So(GetCurrentPath(), ShouldNotBeNil)
		})
		Convey("文件类型", func() {
			ty, err := GetType("/Users/my/Downloads/Video/cangjingkong.mp4")
			So(err, ShouldBeNil)
			t.Log(ty)
		})
	})

}
