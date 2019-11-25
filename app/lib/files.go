package lib

import (
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"../../app"
)

type filesOutput struct {
	Name   string
	Path   string
	Status bool
}

func FilesUpload(e *app.Env, r *http.Request, inputname string, path string, is_rename bool) (filesOutput, error) {
	uploadedFile, handler, err := r.FormFile(inputname)
	var upload_path string

	if path == "" {
		upload_path = e.UPLOAD_PATH
	} else {
		upload_path = path
	}

	result := filesOutput{"", upload_path, false}

	if err != nil {
		return result, err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		return result, err
	}

	filename := handler.Filename

	if is_rename {
		alias := strings.ToLower(RandomString(16))
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, upload_path, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return result, err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return result, err
	}

	result = filesOutput{filename, upload_path, true}
	return result, nil
} // end func

func FilesMultiUpload(e *app.Env, r *http.Request, path string, is_rename bool) ([]filesOutput, error) {
	basePath, _ := os.Getwd()
	var upload_path string

	if path == "" {
		upload_path = e.UPLOAD_PATH
	} else {
		upload_path = path
	}

	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	results := []filesOutput{}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		filename := part.FileName()
		fileLocation := filepath.Join(basePath, upload_path, filename)

		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}

		if err != nil {
			return nil, err
		}

		upload_status := true
		if _, err := io.Copy(dst, part); err != nil {
			upload_status = false
		}

		// log.Printf("%+v", fileLocation)
		result := filesOutput{filename, upload_path, upload_status}
		results = append(results, result)
	}

	return results, nil
}
