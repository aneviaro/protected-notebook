package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/xrustalik/protected-notebook/file"

	"github.com/gin-gonic/gin"
)

func GetFileListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, file.Get())
}

func AddFileHandler(c *gin.Context) {
	fileItem, statusCode, err := convertHTTPBodyToFile(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{
		"id": file.Add(fileItem.Name, fileItem.Content)})
}

func DeleteFileHandler(c *gin.Context) {
	fileID := c.Param("id")
	if err := file.Delete(fileID); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func convertHTTPBodyToFile(httpBody io.ReadCloser) (file.File, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return file.File{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToFile(body)
}

func convertJSONBodyToFile(jsonBody []byte) (file.File, int, error) {
	var fileItem file.File
	err := json.Unmarshal(jsonBody, &fileItem)
	if err != nil {
		return file.File{}, http.StatusBadRequest, err
	}
	return fileItem, http.StatusOK, nil
}
