package api

import (
	"fmt"
	"sap-crm/pkg/upload"

	"github.com/gofiber/fiber/v2"
)

func UploadImageSingle(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})

	}

	if file == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})

	}

	imageName := upload.GetImageName(file.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		fmt.Println(upload.CheckImageExt(imageName))
		fmt.Println(upload.CheckImageSize(file))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FORMAT",
		})

	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FAIL",
		})

	}

	if err := c.SaveFile(file, src); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "ERROR_UPLOAD_SAVE_IMAGE_FAIL",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":          true,
		"msg":            "SUCCESS",
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}

func UploadImageMultiple(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})

	}
	files := form.File["image"]
	if files == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "INVALID_PARAMS",
		})

	}
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	for _, file := range files {
		imageName := upload.GetImageName(file.Filename)
		src := fullPath + imageName

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			fmt.Println(upload.CheckImageExt(imageName))
			fmt.Println(upload.CheckImageSize(file))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FORMAT",
			})

		}
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			fmt.Println(upload.CheckImageExt(imageName))
			fmt.Println(upload.CheckImageSize(file))
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FORMAT",
			})

		}
		err = upload.CheckImage(fullPath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_CHECK_IMAGE_FAIL",
			})

		}
		if err := c.SaveFile(file, src); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   "ERROR_UPLOAD_SAVE_IMAGE_FAIL",
			})

		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":          true,
		"msg":            "SUCCESS",
		"image_url":      upload.GetImagePath(),
		"image_save_url": savePath,
	})
}
