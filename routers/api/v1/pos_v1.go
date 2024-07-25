package v1

import (
	"time"

	"sap-crm/service/pos_service.go"

	"github.com/gofiber/fiber/v2"
)

type FormPOS struct {
	POSCode     string           `json:"POSCode"`
	DocDate     time.Time        `json:"DocDate"`
	Remarks     string           `json:"Remarks"`
	Status      bool             `json:"Status"`
	Type        string           `json:"Type"`
	PostingDate string           `json:"PostingDate"`
	Total       float64          `json:"Total"`
	Discount    float64          `json:"Discount"`
	DocTotal    float64          `json:"DocTotal"`
	VAT         float64          `json:"VAT"`
	CustomerPay string           `json:"CustomerPay"`
	Refund      float64          `json:"Refund"`
	CardCode    string           `json:"CardCode"`
	CardName    string           `json:"CardName"`
	Creator     string           `json:"Creator"`
	POSDetails  []FormPOSDetails `json:"POSDetails"`
	// IDVen        int             `json:"IDVen" validate:"Required"`
	// PONum        string          `json:"PONum" validate:"Required"`
	// TableName   string    `json:"TableName" validate:"Required"`
	// TableNumber string    `json:"TableNumber" validate:"Required"`
	// UoMUnit      string          `json:"UoMUnit"`
	// DeliveryDate string          `json:"DeliveryDate"`
}

type FormPOSDetails struct {
	ItemCode string  `json:"ItemCode"`
	ItemName string  `json:"ItemName"`
	Quantity float64 `json:"Quantity"`
	Price    float64 `json:"Price"`
	Category string  `json:"Category"`
	Tax      float64 `json:"Tax"`
	WhsCode  string  `json:"WhsCode"`
	// IDItem           int     `json:"IDItem" validate:"Required"`
}

type FormSearchPOS struct {
	POSCode string `form:"POSCode" valid:"required" json:"poscode"`
	// TableNumber string `form:"TableNumber" valid:"required" json:"tablenumber"`
	// PONum    string `form:"PONum" valid:"required" json:"poNum"`
}

func PostPOS(c *fiber.Ctx) error {
	form := &FormPOS{}

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

	var posDetails []pos_service.POSDetails = make([]pos_service.POSDetails, len(form.POSDetails))
	for i, v := range form.POSDetails {
		item := pos_service.POSDetails{
			ItemName: v.ItemName,
			ItemCode: v.ItemCode,
			Quantity: v.Quantity,
			Price:    v.Price,
			Category: v.Category,
			Tax:      v.Tax,
			WhsCode:  v.WhsCode,
			// IDItem:           v.IDItem,
			// UoMCode:          v.UoMCode,
			// ExpectedDelivery: v.ExpectedDelivery,
			// TotalFirst:       v.TotalFirst,
			// Discount:         v.Discount,
			// TotalAfter:       v.TotalAfter,
		}
		posDetails[i] = item
	}
	posService := pos_service.POS{
		// POSCode:     form.POSCode,
		DocDate:     form.DocDate,
		Type:        form.Type,
		PostingDate: form.PostingDate,
		Total:       form.Total,
		Discount:    form.Discount,
		DocTotal:    form.DocTotal,
		VAT:         form.VAT,
		CustomerPay: form.CustomerPay,
		Refund:      form.Refund,
		CardCode:    form.CardCode,
		CardName:    form.CardName,
		Creator:     form.Creator,
		POSDetails:  posDetails,
		// IDVen:        form.IDVen,
		// PONum:        form.PONum,
		// TableName:   form.TableName,
		// TableNumber: form.TableNumber,
		// Remarks:     form.Remarks,
		// Status:      form.Status,
		// UoMUnit:      form.UoMUnit,
		// DeliveryDate: form.DeliveryDate,
		// IDItem:    form.IDItem,
		// ItemCode:  form.ItemCode,
		// ItemName:  form.ItemName,
		// Quantity:  form.Quantity,
		// Price:     form.Price,
		// VAT:       form.VAT,
		// Total:     form.Total,
	}

	if err := posService.AddPOS(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Post Add POS false",
		})
	}

	// poDetailService = po_service.PODetails{
	// 	IDItem: poDetailsForm.IDItem,
	// 	ItemCode: poDetailsForm.ItemCode,
	// 	ItemName: poDetailsForm.ItemName,
	// 	Quantity: poDetailsForm.Quantity,
	// 	Price: poDetailsForm.Price,
	// 	VAT: poDetailsForm.VAT,
	// 	Total: poDetailsForm.Total,
	// }

	// if err := poDetailService.AddPO(); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Post Add PO Detail false",
	// 	})
	// }
	data := make(map[string]string)
	data["POSCode"] = form.POSCode

	return c.Status(fiber.StatusOK).JSON(data)
}

// Update POS

func PutPOS(c *fiber.Ctx) error {
	form := &FormPOS{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	posService := pos_service.POS{
		POSCode:     form.POSCode,
		DocDate:     form.DocDate,
		Type:        form.Type,
		PostingDate: form.PostingDate,
		Total:       form.Total,
		Discount:    form.Discount,
		DocTotal:    form.DocTotal,
		VAT:         form.VAT,
		CustomerPay: form.CustomerPay,
		Refund:      form.Refund,
		CardCode:    form.CardCode,
		CardName:    form.CardName,
		Creator:     form.Creator,
		// TableName:   form.TableName,
		// TableNumber: form.TableNumber,
		// Remarks:     form.Remarks,
		// Status:      form.Status,
		// UoMUnit:      form.UoMUnit,
		// DeliveryDate: form.DeliveryDate,
	}

	err := posService.UpdatePOS(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update POS false",
		})
	}

	data := make(map[string]string)
	data["POSCode"] = form.POSCode

	return c.Status(fiber.StatusOK).JSON(data)
}

// get all pos
func GetPOS_Router(c *fiber.Ctx) error {
	pos := pos_service.POS{}
	data, err := pos.GetPOS_Service()

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

func GetByIdPOS_Router(c *fiber.Ctx) error {
	pos := pos_service.POS{}
	data, err := pos.GetByIdPOS_Service(c.Params("id"))

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

func SearcPOS_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchPOS)
	pos := pos_service.POS{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := pos.SearchPOS_Service(formSearch.POSCode)
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
