package v1

import (
	orderdetails_service "sap-crm/service/order_details_service"

	"github.com/gofiber/fiber/v2"
)

type FormOrderDetail struct {
	ItemCode string  `json:"ItemCode" validate:"Required"`
	ItemName string  `json:"ItemName" validate:"Required"`
	Quantity float64 `json:"Quantity" validate:"Required"`
	Price    float64 `json:"Price" validate:"Required"`
	Remarks  string  `json:"Remarks"`
	Tax      float64 `json:"Tax"`
	WhsCode  string  `json:"WhsCode"`
	Status   string  `json:"Status"`
	// IDItem           int     `json:"IDItem" validate:"Required"`
	// UoMCode          string  `json:"UoMCode"`
	// ExpectedDelivery string  `json:"ExpectedDelivery"`
	// TotalFirst       float64 `json:"TotalFirst"`
	// Discount         float64 `json:"Discount"`
	// TotalAfter       float64 `json:"TotalAfter"`
}
type FormOrderDetailAddNew struct {
	IDOrder  uint    `json:"IDOrder"`
	ItemCode string  `json:"ItemCode"`
	ItemName string  `json:"ItemName"`
	Quantity float64 `json:"Quantity"`
	Price    float64 `json:"Price"`
	Remarks  string  `json:"Remarks"`
	Tax      float64 `json:"Tax"`
	WhsCode  string  `json:"WhsCode"`
	Status   string  `json:"Status"`
}

// --- get ID order ---
func GetByIdOrderDetail_Router(c *fiber.Ctx) error {
	orderdetail := orderdetails_service.OrderDetails{}
	data, err := orderdetail.GetByIdOrderDetail_Service(c.Params("id"))

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

func PostOrderDetails_Router(c *fiber.Ctx) error {
	form := &FormOrderDetailAddNew{}
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	itemService := orderdetails_service.OrderDetailsAddNew{
		IDOrder:  form.IDOrder,
		ItemCode: form.ItemCode,
		ItemName: form.ItemName,
		Quantity: form.Quantity,
		Price:    form.Price,
		Remarks:  form.Remarks,
		Tax:      form.Tax,
		WhsCode:  form.WhsCode,
		Status:   form.Status,
	}
	if err := itemService.AddOrderDetails_Service(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Add false",
		})
	}

	data := make(map[string]string)
	data["ItemCode"] = form.ItemCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func PutOrderDetailItem(c *fiber.Ctx) error {
	form := &FormOrderDetail{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	orderDetailService := orderdetails_service.OrderDetails{
		ItemCode: form.ItemCode,
		ItemName: form.ItemName,
		Quantity: form.Quantity,
		Price:    form.Price,
		Remarks:  form.Remarks,
		Tax:      form.Tax,
		WhsCode:  form.WhsCode,
		Status:   form.Status,
		// IDItem:           form.IDItem,
		// UoMCode:          form.UoMCode,
		// ExpectedDelivery: form.ExpectedDelivery,
		// TotalFirst:       form.TotalFirst,
		// Discount:         form.Discount,
		// TotalAfter:       form.TotalAfter,
	}

	err := orderDetailService.UpdateOrderDetaiItem_Service(c.Params("id"))

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
