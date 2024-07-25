package upload

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"sap-crm/pkg/file"
	"sap-crm/pkg/setting"
	"sap-crm/pkg/utils"
	"strings"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

// GetImageName get image name
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	strArrayImg := [3]string{".jpg", ".jpeg", ".png"}
	ext := file.GetExt(fileName)
	for _, allowExt := range strArrayImg {
		fmt.Println(ext)
		fmt.Println(allowExt)
		if strings.Contains(strings.ToUpper(ext), strings.ToUpper(allowExt)) {
			return true
		}
	}

	return false
}

// CheckImageSize check image size
func CheckImageSize(f *multipart.FileHeader) bool {
	fmt.Println("size")
	fmt.Println(f.Size)
	fmt.Println("sizImageMaxSizee")
	fmt.Println(setting.AppSetting.ImageMaxSize)
	return int(f.Size) <= setting.AppSetting.ImageMaxSize
}

// CheckImage check if the file exists
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
