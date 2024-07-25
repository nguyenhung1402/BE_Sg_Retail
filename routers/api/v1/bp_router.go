package v1

import (
	bp_service "sap-crm/service/bp_service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BPs struct {
	CardCode   string    `form:"CardCode" validate:"required"`
	CardName   string    `form:"CardName" validate:"required"`
	Group      string    `form:"Group"`
	Address    string    `form:"Address"`
	Phone      string    `form:"Phone"`
	Remarks    string    `form:"Remarks"`
	DocDate    time.Time `form:"DocDate"`
	UserAdd    string    `form:"UserAdd"`
	UserUpdate string    `form:"UserUpdate"`
	Status     bool      `form:"Status"`
	Type       string    `form:"Type"`
}

type FormSearchBPs struct {
	CardCode string `form:"CardCode" valid:"required json:"CardCode"`
	CardName string `form:"CardName" valid:"required" json:"CardName"`
	Type     string `form:"Type" valid:"required" json:"Type"`
}

func PostBPs(c *fiber.Ctx) error {
	form := &BPs{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	bpService := bp_service.BPs{
		CardCode:   form.CardCode,
		CardName:   form.CardName,
		Group:      form.Group,
		Address:    form.Address,
		Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	if err := bpService.AddBPs(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["CardCode"] = form.CardCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutBPs(c *fiber.Ctx) error {
	form := &BPs{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	bpService := bp_service.BPs{
		CardCode:   form.CardCode,
		CardName:   form.CardName,
		Group:      form.Group,
		Address:    form.Address,
		Phone:      form.Phone,
		Remarks:    form.Remarks,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
	}

	err := bpService.UpdatBPs(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Register false",
		})
	}

	data := make(map[string]string)
	data["CardCode"] = form.CardCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetBPs_Router(c *fiber.Ctx) error {
	item := bp_service.BPs{}
	data, err := item.GetBPs_Service()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func GetByIdPBs_Router(c *fiber.Ctx) error {
	item := bp_service.BPs{}
	data, err := item.GetByIdBPs_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func SearchBP_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchBPs)
	item := bp_service.BPs{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchBPs_Service(formSearch.CardCode, formSearch.CardName, formSearch.Type)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
