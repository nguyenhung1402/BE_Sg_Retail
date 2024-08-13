package v1

import (
	"sap-crm/service/order_service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FormOrder struct {
	TableName    string             `json:"TableName" validate:"Required"`
	TableNumber  string             `json:"TableNumber" validate:"Required"`
	DocDate      time.Time          `json:"DocDate"`
	Remarks      string             `json:"Remarks"`
	Status       bool               `json:"Status"`
	Type         string             `json:"Type"`
	PostingDate  string             `json:"PostingDate"`
	Total        float64            `json:"Total"`
	Discount     int                `json:"Discount"`
	DocTotal     float64            `json:"DocTotal"`
	Customerpay  string             `json:"Customerpay"`
	Refund       float64            `json:"Refund"`
	CardCode     string             `json:"CardCode"`
	CardName     string             `json:"CardName"`
	StaffCode    string             `json:"StaffCode"`
	StaffName    string             `json:"StaffName"`
	VAT          int                `json:"VAT"`
	Dv           string             `json:"Dv"`
	Creator      string             `json:"Creator"`
	CodeAuto     string             `json:"CodeAuto"`
	ViewPayment  string             `json:"ViewPayment"`
	OrderDetails []FormOrderDetails `json:"OrderDetails"`
	// IDVen        int             `json:"IDVen" validate:"Required"`
	// PONum        string          `json:"PONum" validate:"Required"`
	// UoMUnit      string          `json:"UoMUnit"`
	// DeliveryDate string          `json:"DeliveryDate"`
}

type FormOrderDetails struct {
	ItemCode string  `json:"ItemCode" validate:"Required"`
	ItemName string  `json:"ItemName" validate:"Required"`
	Quantity float64 `json:"Quantity" validate:"Required"`
	Price    float64 `json:"Price" validate:"Required"`
	Remarks  string  `json:"Remarks"`
	Tax      float64 `json:"Tax"`
	WhsCode  string  `json:"WhsCode"`
	// IDItem           int     `json:"IDItem" validate:"Required"`
	// UoMCode          string  `json:"UoMCode"`
	// ExpectedDelivery string  `json:"ExpectedDelivery"`
	// TotalFirst       float64 `json:"TotalFirst"`
	// Discount         float64 `json:"Discount"`
	// TotalAfter       float64 `json:"TotalAfter"`
}

type FormSearchOrder struct {
	TableName   string `form:"TableName" valid:"required" json:"tablename"`
	TableNumber string `form:"TableNumber" valid:"required" json:"tablenumber"`
	// PONum    string `form:"PONum" valid:"required" json:"poNum"`
}

type FormSearchDayMonthYear struct {
	Day   string `form:"Day" valid:"required json:"day"`
	Month string `form:"Month" valid:"required json:"month"`
	Year  string `form:"Year" valid:"required json:"year"`
}

type FormTableNumber struct {
	TableNumber string `form:"TableNumber" valid:"required json:"tablenumber"`
	Date        string `form:"Date" valid:"required json:"date"`
}

// ----------ADD Order
func PostOrder(c *fiber.Ctx) error {
	form := &FormOrder{}

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

	var orderDetails []order_service.OrderDetails = make([]order_service.OrderDetails, len(form.OrderDetails))
	for i, v := range form.OrderDetails {
		item := order_service.OrderDetails{
			ItemName: v.ItemName,
			ItemCode: v.ItemCode,
			Quantity: v.Quantity,
			Price:    v.Price,
			Remarks:  v.Remarks,
			Tax:      v.Tax,
			WhsCode:  v.WhsCode,
			// IDItem:           v.IDItem,
			// UoMCode:          v.UoMCode,
			// ExpectedDelivery: v.ExpectedDelivery,
			// TotalFirst:       v.TotalFirst,
			// Discount:         v.Discount,
			// TotalAfter:       v.TotalAfter,
		}
		orderDetails[i] = item
	}
	orderService := order_service.Order{
		TableName:   form.TableName,
		TableNumber: form.TableNumber,
		DocDate:     form.DocDate,
		Remarks:     form.Remarks,
		Status:      form.Status,
		Type:        form.Type,
		PostingDate: form.PostingDate,
		Total:       form.Total,
		Discount:    form.Discount,
		DocTotal:    form.DocTotal,
		Customerpay: form.Customerpay,
		Refund:      form.Refund,
		CardCode:    form.CardCode,
		CardName:    form.CardName,
		StaffName:   form.StaffName,
		StaffCode:   form.StaffCode,
		VAT:         form.VAT,
		Creator:     form.Creator,
		Dv:          form.Dv,
		ViewPayment: form.ViewPayment,
		// IDVen:        form.IDVen,
		// PONum:        form.PONum,
		// UoMUnit:      form.UoMUnit,
		// DeliveryDate: form.DeliveryDate,
		// IDItem:    form.IDItem,
		// ItemCode:  form.ItemCode,
		// ItemName:  form.ItemName,
		// Quantity:  form.Quantity,
		// Price:     form.Price,
		// VAT:       form.VAT,
		// Total:     form.Total,
		OrderDetails: orderDetails,
	}

	if err := orderService.AddOrder(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Post Add Order false",
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
	data["TableName"] = form.TableName

	return c.Status(fiber.StatusOK).JSON(data)
}

// Update Order

func PutOrder(c *fiber.Ctx) error {
	form := &FormOrder{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	orderService := order_service.Order{
		TableName:   form.TableName,
		TableNumber: form.TableNumber,
		DocDate:     form.DocDate,
		Remarks:     form.Remarks,
		Status:      form.Status,
		Type:        form.Type,
		PostingDate: form.PostingDate,
		Total:       form.Total,
		Discount:    form.Discount,
		DocTotal:    form.DocTotal,
		Customerpay: form.Customerpay,
		Refund:      form.Refund,
		CardCode:    form.CardCode,
		CardName:    form.CardName,
		VAT:         form.VAT,
		Creator:     form.Creator,
		ViewPayment: form.ViewPayment,
		// UoMUnit:      form.UoMUnit,
		// DeliveryDate: form.DeliveryDate,
	}

	err := orderService.UpdateOrder(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Update Order false",
		})
	}

	data := make(map[string]string)
	data["TableName"] = form.TableName

	return c.Status(fiber.StatusOK).JSON(data)
}

// get all order
func GetOrder_Router(c *fiber.Ctx) error {
	order := order_service.Order{}
	data, err := order.GetOrder_Service()

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

// staff
func Staff_Oder(c *fiber.Ctx) error {
	order := order_service.StaffOrder{}
	formSearch := new(FormNXT)
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := order.StaffOder_Service(formSearch.FromDate, formSearch.ToDate)

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
func PostGetOrder_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := order_service.Order{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.PostGetOrder_Service(formSearch.Day, formSearch.Month, formSearch.Year)

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

func GetByIdOrder_Router(c *fiber.Ctx) error {
	order := order_service.Order{}
	data, err := order.GetByIdOrder_Service(c.Params("id"))

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

func SearcOrder_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchOrder)
	order := order_service.Order{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := order.SearchOrder_Service(formSearch.TableName, formSearch.TableNumber)
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

func Post_AppendOrderDetails(c *fiber.Ctx) error {
	form := &FormOrder{}
	// Check, if received JSON data is valid.
	if err := c.BodyParser(form); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var orderDetails []order_service.OrderDetails = make([]order_service.OrderDetails, len(form.OrderDetails))
	for i, v := range form.OrderDetails {
		item := order_service.OrderDetails{
			ItemName: v.ItemName,
			ItemCode: v.ItemCode,
			Quantity: v.Quantity,
			Price:    v.Price,
			Remarks:  v.Remarks,
			Tax:      v.Tax,
			WhsCode:  v.WhsCode,
		}
		orderDetails[i] = item
	}
	orderService := order_service.Order{
		TableName:    form.TableName,
		TableNumber:  form.TableNumber,
		DocDate:      form.DocDate,
		Remarks:      form.Remarks,
		Status:       form.Status,
		Type:         form.Type,
		PostingDate:  form.PostingDate,
		Total:        form.Total,
		Discount:     form.Discount,
		DocTotal:     form.DocTotal,
		Customerpay:  form.Customerpay,
		Refund:       form.Refund,
		CardCode:     form.CardCode,
		CardName:     form.CardName,
		VAT:          form.VAT,
		Creator:      form.Creator,
		ViewPayment:  form.ViewPayment,
		OrderDetails: orderDetails,
	}

	if err := orderService.AppendOrderDetails_Service(c.Params("id")); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Put Add Order false",
		})
	}
	data := make(map[string]string)
	data["TableName"] = form.TableName

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetTablenumber_Router(c *fiber.Ctx) error {
	formSearch := new(FormTableNumber)
	order := order_service.Order{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := order.GetTablenumber_Service(formSearch.TableNumber, formSearch.Date)
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

// func GetTablenumberNotThanhToan_Router(c *fiber.Ctx) error {
// 	order := order_service.Order{}
// 	data, err := order.GetTablenumberNotThanhToan_Service()

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"success": true,
// 		"data":    data,
// 	})
// }

func GetTablenumberNotThanhToan_Router(c *fiber.Ctx) error {
	formSearch := new(FormTableNumber)
	order := order_service.Order{}
	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	data, err := order.GetTablenumberNotThanhToan_Service(formSearch.TableNumber, formSearch.Date)
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
