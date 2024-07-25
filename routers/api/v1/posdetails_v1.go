package v1

import (
	posdetails_service "sap-crm/service/pos_details_service"

	"github.com/gofiber/fiber/v2"
)

type FormPOSDetail struct {
	// IDItem           int     `json:"IDItem" validate:"Required"`
	ItemCode string  `json:"ItemCode" validate:"Required"`
	ItemName string  `json:"ItemName" validate:"Required"`
	Quantity float64 `json:"Quantity" validate:"Required"`
	Price    float64 `json:"Price" validate:"Required"`
	Category string  `json:"Remarks"`
	Tax      float64 `json:"Tax"`
	WhsCode  string  `json:"WhsCode"`
	// UoMCode          string  `json:"UoMCode"`
	// ExpectedDelivery string  `json:"ExpectedDelivery"`
	// TotalFirst       float64 `json:"TotalFirst"`
	// Discount         float64 `json:"Discount"`
	// TotalAfter       float64 `json:"TotalAfter"`
}

// --- get ID pos ---
func GetByIdPOSDetail_Router(c *fiber.Ctx) error {
	posdetail := posdetails_service.POSDetails{}
	data, err := posdetail.GetByIdPOSDetail_Service(c.Params("id"))

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

func PutPOSDetailItem(c *fiber.Ctx) error {
	form := &FormPOSDetail{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	posDetailService := posdetails_service.POSDetails{
		ItemCode: form.ItemCode,
		ItemName: form.ItemName,
		Quantity: form.Quantity,
		Price:    form.Price,
		Category: form.Category,
		Tax:      form.Tax,
		WhsCode:  form.WhsCode,
		// IDItem:           form.IDItem,
		// UoMCode:          form.UoMCode,
		// ExpectedDelivery: form.ExpectedDelivery,
		// TotalFirst:       form.TotalFirst,
		// Discount:         form.Discount,
		// TotalAfter:       form.TotalAfter,
	}

	err := posDetailService.UpdatePOSDetaiItem_Service(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update Detail false",
		})
	}
	data := make(map[string]string)
	data["ItemCode"] = form.ItemCode

	return c.Status(fiber.StatusOK).JSON(data)
}
