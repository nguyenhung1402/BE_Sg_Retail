package v1

import (
	// issuedetails_service "sap-crm/service/issue_details_service"
	issuedetails_service "sap-crm/service/issue_details_service"

	"github.com/gofiber/fiber/v2"
)

type FormIssueDetail struct {
	// IDItem           int     `json:"IDItem" validate:"Required"`
	ItemCode string  `json:"ItemCode" validate:"Required"`
	ItemName string  `json:"ItemName" validate:"Required"`
	Quantity float64 `json:"Quantity" validate:"Required"`
	// UoMCode          string  `json:"UoMCode"`
	// ExpectedDelivery string  `json:"ExpectedDelivery"`
	// TotalFirst       float64 `json:"TotalFirst"`
	// Tax              float64 `json:"Tax"`
	// Discount         float64 `json:"Discount"`
	// TotalAfter       float64 `json:"TotalAfter"`
}

// --- get ID issue ---
func GetByIdIssueDetail_Router(c *fiber.Ctx) error {
	issuedetail := issuedetails_service.IssueDetails{}
	data, err := issuedetail.GetByIdIssueDetail_Service(c.Params("id"))

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

func PutIssueDetailItem(c *fiber.Ctx) error {
	form := &FormIssueDetail{}
	// form := &ItemsUpdate{}
	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	issueDetailService := issuedetails_service.IssueDetails{
		// IDItem:           form.IDItem,
		ItemCode: form.ItemCode,
		ItemName: form.ItemName,
		Quantity: form.Quantity,

		// UoMCode:          form.UoMCode,
		// ExpectedDelivery: form.ExpectedDelivery,
		// TotalFirst:       form.TotalFirst,
		// Tax:              form.Tax,
		// Discount:         form.Discount,
		// TotalAfter:       form.TotalAfter,
	}

	err := issueDetailService.UpdateIssueDetaiItem_Service(c.Params("id"))

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
