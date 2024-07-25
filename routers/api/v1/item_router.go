package v1

import (
	item_service "sap-crm/service/item_service"

	"github.com/gofiber/fiber/v2"
)

type Items struct {
	ItemCode   string  `form:"ItemCode" valid:"required"`
	ItemName   string  `form:"ItemName" valid:"required"`
	Quantity   int     `form:"Quantity" valid:"required"`
	Group      string  `form:"Group"`
	InStock    float64 `form:"InStock"`
	Price      float64 `form:"Price"`
	UserAdd    string  `form:"UserAdd"`
	UserUpdate string  `form:"UserUpdate"`
	Status     bool    `form:"Status"`
	Type       string  `form:"Type"`
	Barcode    string  `form:"Barcode"`
	Image      string  `form:"Image"`
	WhsCode    string  `form:"WhsCode"`
	CardCode   string  `form:"CardCode"`
}

type ItemsUpdate struct {
	ID         string  `form:"ID"`
	ItemCode   string  `form:"ItemCode" valid:"required"`
	ItemName   string  `form:"ItemName" valid:"required"`
	Quantity   int     `form:"Quantity" valid:"required"`
	Group      string  `form:"Group"`
	InStock    float64 `form:"InStock"`
	Price      float64 `form:"Price"`
	UserAdd    string  `form:"UserAdd"`
	UserUpdate string  `form:"UserUpdate"`
	Status     bool    `form:"Status"`
	Type       string  `form:"Type"`
	Barcode    string  `form:"Barcode"`
	Image      string  `form:"Image"`
	WhsCode    string  `form:"WhsCode"`
	CardCode   string  `form:"CardCode"`
}

type FormSearch struct {
	ItemCode string `form:"ItemCode" valid:"required json:"ItemCode"`
	ItemName string `form:"ItemName" valid:"required" json:"ItemName"`
}

func PostItems(c *fiber.Ctx) error {
	form := &Items{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	itemService := item_service.Items{
		ItemCode:   form.ItemCode,
		ItemName:   form.ItemName,
		Quantity:   form.Quantity,
		Group:      form.Group,
		InStock:    form.InStock,
		Price:      form.Price,
		UserAdd:    form.UserAdd,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
		Barcode:    form.Barcode,
		Image:      form.Image,
		WhsCode:    form.WhsCode,
		CardCode:   form.CardCode,
	}
	if err := itemService.AddItems(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Add false",
		})
	}

	data := make(map[string]string)
	data["ItemCode"] = form.ItemCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutItems(c *fiber.Ctx) error {
	form := &Items{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	itemService := item_service.Items{
		// ItemCode:   form.ItemCode,
		ItemName:   form.ItemName,
		Quantity:   form.Quantity,
		Group:      form.Group,
		InStock:    form.InStock,
		Price:      form.Price,
		UserUpdate: form.UserUpdate,
		Status:     form.Status,
		Type:       form.Type,
		Barcode:    form.Barcode,
		Image:      form.Image,
		WhsCode:    form.WhsCode,
		CardCode:   form.CardCode,
	}

	err := itemService.UpdateItems(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update false",
		})
	}

	data := make(map[string]string)
	data["ItemCode"] = form.ItemCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetItems_Router(c *fiber.Ctx) error {
	item := item_service.Items{}
	data, err := item.GetItems_Service()

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

func GetById_Router(c *fiber.Ctx) error {
	item := item_service.Items{}
	data, err := item.GetById_Service(c.Params("id"))

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

func SearchItem_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearch)
	item := item_service.Items{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := item.SearchItems_Service(formSearch.ItemCode, formSearch.ItemName)
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
