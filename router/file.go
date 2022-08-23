package router

import (
	"encoding/base64"
	"fmt"
	"mehm8128_study_server/model"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// func postFile(c echo.Context) error {
// 	//ユーザー名とファイルを取り出す
// 	userID := c.FormValue("userID")
// 	userID2, err := uuid.Parse(userID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "invalid request ID: "+userID)
// 	}
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "invalid request file")
// 	}
// 	fileModel := strings.Split(file.Filename, ".")
// 	extension := fileModel[1]
// 	ctx := c.Request().Context()

// 	//中身を取り出す
// 	src, err := file.Open()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
// 	}
// 	defer src.Close()

// 	//ディレクトリにファイルを作る
// 	ID := uuid.New()
// 	dst, err := os.Create("./files/" + ID.String() + "." + extension)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create file")
// 	}
// 	defer dst.Close()

// 	//コピーする
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "failed to copy file")
// 	}

// 	//ファイルの情報を保存
// 	res, err := model.CreateFile(ctx, ID, file.Filename, userID2)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err)
// 	}
// 	return echo.NewHTTPError(http.StatusOK, &res)
// }

func postFile(c echo.Context) error {
	//ユーザー名とファイルを取り出す
	userID := c.FormValue("userID")
	userID2, err := uuid.Parse(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request ID: "+userID)
	}
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request file")
	}
	ctx := c.Request().Context()

	//中身を取り出す
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	defer src.Close()

	size := file.Size
	data := make([]byte, size)
	_, err = src.Read(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	base64 := base64.StdEncoding.EncodeToString(data)

	//ファイルの情報を保存
	res, err := model.CreateFile(ctx, file.Filename, base64, userID2)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, &res)
}

// func getFile(c echo.Context) error {
// 	ID, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
// 	}
// 	ctx := c.Request().Context()
// 	file, err := model.GetFile(ctx, ID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err)
// 	}
// 	extension := strings.Split(file.FileName, ".")[1]
// 	filePath := "./files/" + file.ID.String() + "." + extension
// 	src, err := os.Open(filePath)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err)
// 	}
// 	defer src.Close()
// 	return c.Stream(http.StatusOK, "image/"+extension, src)
// }

func getFile(c echo.Context) error {
	fmt.Printf("aaa")
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	file, err := model.GetFile(ctx, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	extension := strings.Split(file.FileName, ".")[1]

	data, err := base64.StdEncoding.DecodeString(file.File)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.Blob(http.StatusOK, "image/"+extension, data)
}

func getFileInfo(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	file, err := model.GetFile(ctx, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, &file)
}
