package util

import (
    "github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context, filename string) (bool, string) {
    file, err := ctx.FormFile("file")
    if err != nil {
        return false, err.Error()
    }

    //filename := filepath.Base(file.Filename)
    if err := ctx.SaveUploadedFile(file, filename); err != nil {
        return false, err.Error()
    }

    return true, ""
}

func UploadFiles(ctx *gin.Context, filenames []string) (bool, string) {
    form, err := ctx.MultipartForm()
    if err != nil {
        return false, err.Error()
    }

    files := form.File["files"]

    for _, file := range files {
        //filename := filepath.Base(file.Filename)
        //filename := filenames[i]
        filename := ""
        if err := ctx.SaveUploadedFile(file, filename); err != nil {
            return false, err.Error()
        }
    }

    return true, ""
}
