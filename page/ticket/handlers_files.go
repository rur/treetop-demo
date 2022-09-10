package ticket

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"

	"github.com/rur/treetop"
	"github.com/rur/treetop-demo/page/ticket/inputs"
	"github.com/rur/treetop-demo/site"
)

// -------------------------
// ticket Handlers
// -------------------------

// Ref: attachment-file-list
// Block: attachment-list
// Doc: Default attachment file list template handler, parse file info from query string
func attachmentFileListHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	// load file info from query
	query := req.URL.Query()
	data := struct {
		Files []*inputs.FileInfo
	}{}

	for _, enc := range query["attachment"] {
		info := &inputs.FileInfo{}
		if err := info.UnmarshalBase64([]byte(enc)); err != nil {
			// skip it
			env.ErrorLog.Println(err)
		} else {
			data.Files = append(data.Files, info)
		}
	}
	return data
}

// Ref: uploaded-files
// Block: attachment-list
// Method: POST
// Doc: Load a list of uploaded files, save to storage and return metadata to the form
func uploadedFilesHandler(env *site.Env, rsp treetop.Response, req *http.Request) interface{} {
	data := struct {
		Files []*inputs.FileInfo
	}{}

	if err := req.ParseMultipartForm(1024 * 1024 * 16 /*16 MiB*/); err != nil {
		env.ErrorLog.Println(err)
		http.Error(rsp, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil
	}

	for _, fh := range req.MultipartForm.File["file-upload"] {
		info := inputs.FileInfo{
			Name: fh.Filename,
		}
		if fh.Size < 2048 {
			info.Size = fmt.Sprintf("%dB", fh.Size)
		} else {
			info.Size = fmt.Sprintf("%.0fKB", float64(fh.Size)/1024)
		}
		f, err := fh.Open()
		if err != nil {
			env.ErrorLog.Println(err)
			http.Error(rsp, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return nil
		}
		defer f.Close()

		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			env.ErrorLog.Println(err)
			http.Error(rsp, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return nil
		}
		info.SHA1 = fmt.Sprintf("%x", h.Sum(nil))
		if b64, err := info.MarshalBase64(); err != nil {
			env.ErrorLog.Println(err)
			http.Error(rsp, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return nil
		} else {
			info.Encoded = string(b64)
		}
		data.Files = append(data.Files, &info)
	}
	return data
}
