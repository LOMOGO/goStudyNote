package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var path = "./files/"

func main() {
	router := gin.Default()

	//上传单个文件
	router.POST("/file", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}

		//判断文件夹是否存在，如果不存在那么就创建一个文件夹。
		_, err = os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(path, os.ModePerm) // 0777
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		err = c.SaveUploadedFile(file, path + file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"文件名称": file.Filename,
			"文件大小": file.Size,
			"文件头": file.Header,
		})
	})

	router.POST("/files", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"]

		var filenames []string

		for _, file := range files {
			filenames = append(filenames, file.Filename)
		}

		c.JSON(http.StatusOK, filenames)
	})

	router.Run(":8080")
}
