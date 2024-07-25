package v1

import (
	"sap-crm/service/issue_service"

	"github.com/gofiber/fiber/v2"
)

type FormIssue struct {
	WhsCode      string             `json:"whscode" validate:"Required"`
	CardCode     string             `json:"cardcode"`
	CardName     string             `json:"cardname"`
	Address      string             `json:"address"`
	Phone        string             `json:"phone"`
	Creator      string             `json:"creator"`
	IssueDetails []FormIssueDetails `json:"issuedetails"`
}

type FormIssueDetails struct {
	ItemCode string  `json:"itemCode" validate:"Required"`
	ItemName string  `json:"itemName"`
	Quantity float64 `json:"quantity" validate:"Required"`
}

type FormSearchIssue struct {
	WhsCode  string `form:"WnsCode" valid:"required" json:"whscode"`
	CardCode string `form:"CardCode" valid:"required" json:"cardcode"`
	CardName string `form:"CardName" valid:"required" json:"cardname"`
}

// ----------ADD Issue
func PostIssue(c *fiber.Ctx) error {
	form := &FormIssue{}

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

	var issueDetails []issue_service.IssueDetails = make([]issue_service.IssueDetails, len(form.IssueDetails))
	for i, v := range form.IssueDetails {
		item := issue_service.IssueDetails{
			ItemName: v.ItemName,
			ItemCode: v.ItemCode,
			Quantity: v.Quantity,
		}
		issueDetails[i] = item
	}
	issueService := issue_service.Issue{
		WhsCode:      form.WhsCode,
		CardCode:     form.CardCode,
		CardName:     form.CardName,
		Address:      form.Address,
		Phone:        form.Phone,
		Creator:      form.Creator,
		IssueDetails: issueDetails,
	}

	if err := issueService.AddIssue(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Post Add Issue false",
		})
	}
	data := make(map[string]string)
	data["WhsCode"] = form.WhsCode

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetAllIssue_Router(c *fiber.Ctx) error {
	order := issue_service.Issue{}
	data, err := order.GetIssue_Service()

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

func GetByIdIssue_Router(c *fiber.Ctx) error {
	order := issue_service.Issue{}
	data, err := order.GetByIdIssue_Service(c.Params("id"))

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
