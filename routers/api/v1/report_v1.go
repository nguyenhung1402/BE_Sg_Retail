package v1

import (
	report_service "sap-crm/service/report_service"

	"github.com/gofiber/fiber/v2"
)

type ChartTopDoanhThu struct {
	CardCode string  `form:"CardCode"`
	Total    float64 `form:"Total"`
}

type ChartTopItemCode struct {
	ItemCode string  `form:"ItemCode"`
	Total    float64 `form:"Total"`
}

// type FormSearchMonthYear struct {
// 	Month string `form:"Month" valid:"required json:"month"`
// 	Year  string `form:"Year" valid:"required json:"year"`
// }

type NameYears struct {
	Name1 string `form:"Name1" json:"name1"`
	Name2 int    `form:"Name2"  json:"name2"`
}
type FormSearchMonthYear struct {
	Day   string `form:"Day" valid:"required json:"day"`
	Month string `form:"Month" valid:"required json:"month"`
	Year  string `form:"Year" valid:"required json:"year"`
}

type FormNXT struct {
	FromDate string `form:"FromDate" json:"fromdate"`
	ToDate   string `form:"ToDate"  json:"todate"`
}

type NXT struct {
	WhsCode  string  `form:"WhsCode" json:"whscode"`
	ItemCode string  `form:"ItemCode" json:"itemcode"`
	BeginQty float64 `form:"BeginQty" json:"beginqty"`
	InQty    float64 `form:"InQty" json:"inqty"`
	OutQty   float64 `form:"OutQty" json:"outqty"`
	EndQty   float64 `form:"EndQty" json:"endqty"`
}

type NhapXuat struct {
	WhsCode  string  `form:"WhsCode" json:"whscode"`
	ItemCode string  `form:"ItemCode" json:"itemcode"`
	Quantity float64 `form:"Quantity" json:"quantity"`
}

// func GetCharTopDoanhThu_Router(c *fiber.Ctx) error {
// 	formSearch := new(FormSearchMonthYear)
// 	item := report_service.ChartTopDoanhThu{}

// 	if err := c.BodyParser(formSearch); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	data, err := item.GetChartTopDoanhThu_Service(formSearch.Month, formSearch.Year)

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

func GetCharTopItemCode_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.ChartTopItemCode{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartTopItemCode_Service(formSearch.Month, formSearch.Year)

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

func GetYears_Router(c *fiber.Ctx) error {
	item := report_service.NameYears{}
	data, err := item.GetYear_Service()

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

func GetCharDay_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.ChartMonth{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartDay_Service(formSearch.Day, formSearch.Month, formSearch.Year)

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
func GetCharDayCircle_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.Chart{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartDayCircle_Service(formSearch.Day, formSearch.Month, formSearch.Year)

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

func GetCharMonth_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.ChartMonth{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartMonth_Service(formSearch.Month, formSearch.Year)

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

func GetCharMonthCircle_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.Chart{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartMonthCircle_Service(formSearch.Month, formSearch.Year)

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

func GetCharYear_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.ChartMonth{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartYear_Service(formSearch.Year)

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

func GetCharYearCircle_Router(c *fiber.Ctx) error {
	formSearch := new(FormSearchMonthYear)
	item := report_service.Chart{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartYearCircle_Service(formSearch.Year)

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

func GetCharNXT_Router(c *fiber.Ctx) error {
	formSearch := new(FormNXT)
	item := report_service.NXT{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartNXT_Service(formSearch.FromDate, formSearch.ToDate)

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

func GetChartReceipt_Router(c *fiber.Ctx) error {
	formSearch := new(FormNXT)
	item := report_service.NhapXuat{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartReceipt_Service(formSearch.FromDate, formSearch.ToDate)

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

func GetChartIssue_Router(c *fiber.Ctx) error {
	formSearch := new(FormNXT)
	item := report_service.NhapXuat{}

	if err := c.BodyParser(formSearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	data, err := item.GetChartIssue_Service(formSearch.FromDate, formSearch.ToDate)

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
