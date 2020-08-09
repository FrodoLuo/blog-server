package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

/*
IsPathExist check whether the path is exist
*/
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ParsePageAndSize(ctx *gin.Context) (page uint64, pageSize uint64, err error) {
	page, err = strconv.ParseUint(ctx.DefaultQuery("page", "0"), 10, 0)
	if err != nil {
		return 0, 0, err
	}
	pageSize, err = strconv.ParseUint(ctx.DefaultQuery("pageSize", "10"), 10, 0)
	if err != nil {
		return 0, 0, err
	}
	return page, pageSize, nil
}

/*
Decipher return the origin data of the ciphered password
*/
func Decipher(ciphered string, key string) (string, error) {
	cipheredBytes, _ := base64.StdEncoding.DecodeString(ciphered)
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	iv := keyBytes[:blockSize]

	des := make([]byte, len(cipheredBytes))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(des, cipheredBytes)
	//解填充
	return string(des), nil
}

func UploadToCOS(fileHeader *multipart.FileHeader) string {
	accessKey := os.Getenv("QINIU_ACCESS_KEY")
	secretKey := os.Getenv("QINIU_SECRET_KEY")
	bucket := os.Getenv("QINIU_BUCKET")
	cosDomain := os.Getenv("BLOG_COS_DOMAIN")

	file, _ := fileHeader.Open()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := auth.New(accessKey, secretKey)
	putPolicy.Expires = 7200
	upToken := putPolicy.UploadToken(mac)

	config := storage.Config{}

	config.Zone = &storage.ZoneHuadong
	config.UseHTTPS = false
	config.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&config)
	ret := storage.PutRet{}

	err := formUploader.Put(context.Background(), &ret, upToken, fileHeader.Filename, file, fileHeader.Size, nil)

	if err != nil {
		panic(err)
	}

	u, _ := url.Parse(cosDomain)
	u.Path = path.Join(u.Path, ret.Key)
	return u.String()
}
