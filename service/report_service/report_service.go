package reports_service

import (
	"fmt"
	"sap-crm/models"
)

type ChartTopDoanhThu struct {
	CardCode string  `gorm:"column: CardCode" json:"cardcode"`
	Total    float64 `gorm:"column: Total" json:"total"`
}

type ChartTopItemCode struct {
	ItemCode string  `gorm:"column: ItemCode" json:"itemcode"`
	Total    float64 `gorm:"column: Total" json:"total"`
}
type NameYears struct {
	Name1 string `gorm:"column: Name1" json:"name1"`
	Name2 int    `gorm:"column: Name2" json:"name2"`
}
type Chart struct {
	Group    string  `gorm:"column: Group" json:"group"`
	Total    float64 `gorm:"column: Total" json:"total"`
	Discount float64 `gorm:"column: Discount" json:"discount"`
}
type ChartMonth struct {
	Group    string  `gorm:"column: Group" json:"group"`
	ItemCode string  `gorm:"column: ItemCode" json:"itemcode"`
	Total    float64 `gorm:"column: Total" json:"total"`
	Discount float64 `gorm:"column: Discount" json:"discount"`
}

type NXT struct {
	WhsCode  string  `gorm:"column: WhsCode" json:"whscode"`
	ItemCode string  `gorm:"column: ItemCode" json:"itemcode"`
	BeginQty float64 `gorm:"column: BeginQty" json:"beginqty"`
	InQty    float64 `gorm:"column: InQty" json:"inqty"`
	OutQty   float64 `gorm:"column: OutQty" json:"outqty"`
	EndQty   float64 `gorm:"column: EndQty" json:"endqty"`
}
type NhapXuat struct {
	WhsCode  string  `gorm:"column: WhsCode" json:"whscode"`
	ItemCode string  `gorm:"column: ItemCode" json:"itemcode"`
	Quantity float64 `gorm:"column: Quantity" json:"quantity"`
}

// func (a *ChartTopDoanhThu) GetChartTopDoanhThu_Service(month string, year string) (*[]models.ChartTopDoanhThu, error) {
// 	item, err := models.GetChartTopDoanhThu_Model(month, year)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//fmt.Println(item)
// 	return item, nil
// }

func (a *ChartTopItemCode) GetChartTopItemCode_Service(month string, year string) (*[]models.ChartTopItemCode, error) {
	item, err := models.GetChartTopItemCode_Model(month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *NameYears) GetYear_Service() (*[]models.NameYears, error) {
	item, err := models.GetYear_Model()
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *ChartMonth) GetChartDay_Service(day string, month string, year string) (*[]models.ChartMonth, error) {
	item, err := models.GetChartDay_Model(day, month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}
func (a *Chart) GetChartDayCircle_Service(day string, month string, year string) (*[]models.Chart, error) {
	item, err := models.GetChartDayCircle_Model(day, month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *ChartMonth) GetChartMonth_Service(month string, year string) (*[]models.ChartMonth, error) {
	item, err := models.GetChartMonth_Model(month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}
func (a *Chart) GetChartMonthCircle_Service(month string, year string) (*[]models.Chart, error) {
	item, err := models.GetChartMonthCircle_Model(month, year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *ChartMonth) GetChartYear_Service(year string) (*[]models.ChartMonth, error) {
	item, err := models.GetChartYear_Model(year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *Chart) GetChartYearCircle_Service(year string) (*[]models.Chart, error) {
	item, err := models.GetChartYearCircle_Model(year)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *NXT) GetChartNXT_Service(fromdate string, todate string) (*[]models.NXT, error) {
	item, err := models.GetChartNXT_Model(fromdate, todate)
	if err != nil {
		return nil, err
	}
	fmt.Println("dhjkasdhjas")
	fmt.Println(item)

	return item, nil
}

func (a *NhapXuat) GetChartReceipt_Service(fromdate string, todate string) (*[]models.NhapXuat, error) {
	item, err := models.GetChartReceipt_Model(fromdate, todate)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}

func (a *NhapXuat) GetChartIssue_Service(fromdate string, todate string) (*[]models.NhapXuat, error) {
	item, err := models.GetChartIssue_Model(fromdate, todate)
	if err != nil {
		return nil, err
	}
	//fmt.Println(item)
	return item, nil
}
