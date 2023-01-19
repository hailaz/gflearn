package main

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

// Upload uploads files to /tmp .
func Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("upload-file")
	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			glog.Error(r.GetCtx(), err)
			continue
		}
		data := make([]byte, file.Size)
		n, err := f.Read(data)
		if err != nil {
			glog.Error(r.GetCtx(), err)
			continue
		}
		glog.Debug(r.GetCtx(), n)
		md5Str, _ := gmd5.EncryptBytes(data)
		g.Log().Debug(r.GetCtx(), md5Str)
	}

	names, err := files.Save("/tmp/")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}

// UploadShow shows uploading simgle file page.
func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload File Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}

// UploadShowBatch shows uploading multiple files page.
func UploadShowBatch(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload Files Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}

func main() {
	s := g.Server()
	s.Group("/upload", func(group *ghttp.RouterGroup) {
		group.POST("/", Upload)
		group.ALL("/show", UploadShow)
		group.ALL("/batch", UploadShowBatch)
	})
	s.Run()
}
