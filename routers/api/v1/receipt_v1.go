package v1

import (
	"sap-crm/service/receipt_service"

	"github.com/gofiber/fiber/v2"
)

type FormReceipt struct {
	WhsCode        string               `json:"whsCode" validate:"Required"`
	CardCode       string               `json:"cardcode"`
	CardName       string               `json:"Cardname"`
	Address        string               `json:"address"`
	Phone          string               `json:"phone"`
	Creator        string               `json:"creator"`
	ReceiptDetails []FormReceiptDetails `json:"receiptdetails"`
}

type FormReceiptDetails struct {
	ItemCode string  `json:"itemcode" validate:"Required"`
	ItemName string  `json:"itemname"`
	Quantity float64 `json:"quantity" validate:"Required"`
}

type FormSearchReceipt struct {
	WhsCode  string `form:"WnsCode" valid:"required" json:"whscode"`
	CardCode string `form:"CardCode" valid:"required" json:"cardcode"`
	CardName string `form:"CardName" valid:"required" json:"cardname"`
}

// ----------ADD Receipt
func PostReceipt(c *fiber.Ctx) error {
	form := &FormReceipt{}

	// Đông test

	// End Đông test

	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var receiptDetails []receipt_service.ReceiptDetails = make([]receipt_service.ReceiptDetails, len(form.ReceiptDetails))
	for i, v := range form.ReceiptDetails {
		item := receipt_service.ReceiptDetails{
			ItemName: v.ItemName,
			ItemCode: v.ItemCode,
			Quantity: v.Quantity,
		}
		receiptDetails[i] = item
	}
	receiptService := receipt_service.Receipt{
		WhsCode:        form.WhsCode,
		CardCode:       form.CardCode,
		CardName:       form.CardName,
		Address:        form.Address,
		Phone:          form.Phone,
		Creator:        form.Creator,
		ReceiptDetails: receiptDetails,
	}

	if err := receiptService.AddReceipt(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Post Add Receipt false",
		})
	}
	data := make(map[string]string)
	data["WhsCode"] = form.WhsCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetAllReceipt_Router(c *fiber.Ctx) error {
	order := receipt_service.Receipt{}
	data, err := order.GetReceipt_Service()

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

func GetByIdReceipt_Router(c *fiber.Ctx) error {
	order := receipt_service.Receipt{}
	data, err := order.GetByIdReceipt_Service(c.Params("id"))

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
