package router

import (
	"io"
	"mehm8128_study_server/model"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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
	fileModel := strings.Split(file.Filename, ".")
	extension := fileModel[1]
	ctx := c.Request().Context()

	//中身を取り出す
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	defer src.Close()

	//ディレクトリにファイルを作る
	ID := uuid.New()
	dst, err := os.Create("./files/" + ID.String() + "." + extension)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create file")
	}
	defer dst.Close()

	//コピーする
	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to copy file")
	}

	//ファイルの情報を保存
	res, err := model.CreateFile(ctx, ID, file.Filename, userID2)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, &res)
}

func getFile(c echo.Context) error {
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
	filePath := "./files/" + file.ID.String() + "." + extension
	src, err := os.Open(filePath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer src.Close()
	return c.Stream(http.StatusOK, "image/"+extension, src)
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
