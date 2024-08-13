package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TableName    string         `gorm:"column:TableName" json:"tablename"`
	TableNumber  string         `gorm:"column:TableNumber" json:"tablenumber"`
	DocDate      time.Time      `gorm:"column:DocDate" json:"docDate"`
	Remarks      string         `gorm:"column:Remarks" json:"remarks"`
	Status       bool           `gorm:"column:Status" json:"status"`
	Type         string         `gorm:"column:Type" json:"type"`
	PostingDate  string         `gorm:"column:PostingDate" json:"postingdate"`
	Total        float64        `gorm:"column:Total" json:"total"`
	Discount     int            `gorm:"column:Discount" json:"discout"`
	DocTotal     float64        `gorm:"column:DocTotal" json:"doctotal"`
	Customerpay  string         `gorm:"column:Customerpay" json:"customerpay"`
	Refund       float64        `gorm:"column:Refund" json:"refund"`
	CardCode     string         `gorm:"column:CardCode" json:"cardcode"`
	CardName     string         `gorm:"column:CardName" json:"cardname"`
	StaffCode    string         `gorm:"column:StaffCode" json:"staffcode"`
	StaffName    string         `gorm:"column:StaffName" json:"staffname"`
	VAT          int            `gorm:"column:VAT" json:"vat"`
	Creator      string         `gorm:"column:Creator" json:"creator"`
	ViewPayment  string         `gorm:"column:ViewPayment" json:"viewpayment"`
	CodeAuto     string         `gorm:"column:CodeAuto" json:"codeauto"`
	DV           string         `gorm:"column:DV" json:"dv"`
	OrderDetails []OrderDetails `gorm:"foreignKey:IDOrder;" json:"orderDetails"`
	// IDVen        int         `gorm:"column:IDVen" json:"idven"`
	// PONum        string      `gorm:"column:PONum" json:"ponum"`
	// ObjectType   string      `gorm:"column:ObjectType" json:"objecttype"`
	// UoMUnit      string      `gorm:"column:UoMUnit" json:"uomunit"`
	// DeliveryDate string      `gorm:"column:DeliveryDate" json:"deliverydate"`
}

type StaffOrder struct {
	StaffName string  `gorm:"column:StaffName" json:"staffcode"`
	Doanh_Thu string  `gorm:"column:DOANH_THU" json:"doanh_thu"`
	Toc       float64 `gorm:"column:TOC" json:"toc"`
}
type WhsCodeQuantity struct {
	WhsCode  string `gorm:"column:WhsCode" json:"whscode"`
	ItemCode string `gorm:"column:ItemCode" json:"itemcode"`
	Quantity string `gorm:"column:Quantity" json:"quantity"`
}

func AddOrder(data map[string]interface{}, dataDetail []map[string]interface{}) error {
	orderDetails := make([]OrderDetails, len(dataDetail))
	for i, v := range dataDetail {
		item := OrderDetails{
			ItemCode: v["ItemCode"].(string),
			ItemName: v["ItemName"].(string),
			Quantity: v["Quantity"].(float64),
			Price:    v["Price"].(float64),
			Remarks:  v["Remarks"].(string),
			Tax:      v["Tax"].(float64),
			WhsCode:  v["WhsCode"].(string),
			Status:   v["Status"].(string),
		}
		orderDetails[i] = item
	}
	order := Order{
		TableName:    data["tablename"].(string),
		TableNumber:  data["tablenumber"].(string),
		DocDate:      data["docdate"].(time.Time),
		Remarks:      data["remarks"].(string),
		Status:       data["status"].(bool),
		Type:         data["type"].(string),
		PostingDate:  data["postingdate"].(string),
		Total:        data["total"].(float64),
		Discount:     data["discount"].(int),
		DocTotal:     data["doctotal"].(float64),
		Customerpay:  data["customerpay"].(string),
		Refund:       data["refund"].(float64),
		CardCode:     data["cardcode"].(string),
		CardName:     data["cardname"].(string),
		DV:           data["dv"].(string),
		StaffCode:    data["staffcode"].(string),
		VAT:          data["vat"].(int),
		Creator:      data["creator"].(string),
		CodeAuto:     data["codeauto"].(string),
		ViewPayment:  data["viewpayment"].(string),
		OrderDetails: orderDetails,
		// IDVen:        data["idven"].(int),
		// PONum:        data["ponum"].(string),
		// ObjectType:   "22",
		// UoMUnit:      data["uomunit"].(string),
		// DeliveryDate: data["deliverydate"].(string),
	}
	result := db.Create(&order)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

// Edit Member modify a single Member
func UpdateOrder_Model(id string, data map[string]interface{}) error {
	order := Order{
		TableName:   data["tablename"].(string),
		TableNumber: data["tablenumber"].(string),
		DocDate:     data["docdate"].(time.Time),
		Remarks:     data["remarks"].(string),
		Status:      data["status"].(bool),
		Type:        data["type"].(string),
		PostingDate: data["postingdate"].(string),
		Total:       data["total"].(float64),
		Discount:    data["discount"].(int),
		DocTotal:    data["doctotal"].(float64),
		Customerpay: data["customerpay"].(string),
		Refund:      data["refund"].(float64),
		CardCode:    data["cardcode"].(string),
		CardName:    data["cardname"].(string),
		VAT:         data["vat"].(int),
		Creator:     data["creator"].(string),
		ViewPayment: data["viewpayment"].(string),
		// IDVen:        data["idven"].(int),
		// PONum:        data["ponum"].(string),
		// ObjectType:   "22",
		// UoMUnit:      data["uomunit"].(string),
		// DeliveryDate: data["deliverydate"].(string),
	}
	if err := db.Model(&order).Where("id = ?", id).Updates(map[string]interface{}{
		"TableName":   order.TableName,
		"TableNumber": order.TableNumber,
		"Remarks":     order.Remarks,
		"Status":      order.Status,
		"Type":        order.Type,
		"PostingDate": order.PostingDate,
		"Total":       order.Total,
		"Discount":    order.Discount,
		"DocTotal":    order.DocTotal,
		"Customerpay": order.Customerpay,
		"Refund":      order.Refund,
		"CardCode":    order.CardCode,
		"CardName":    order.CardName,
		"VAT":         order.VAT,
		"Creator":     order.Creator,
		"ViewPayment": order.ViewPayment,
		// "IDVen":        po.IDVen,
		// "PONum":        po.PONum,
		// "UoMUnit":      po.UoMUnit,
		// "DeliveryDate": po.DeliveryDate,
	}).Error; err != nil {
		return err
	}
	return nil
}

// get all po
func GetOrder_Model() (*[]Order, error) {

	order := []Order{}
	err := db.Debug().Order("created_at desc").Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &order, nil
}

//Staff Oder

func StaffOrder_Model(fromdate string, todate string) (*[]StaffOrder, error) {
	order := []StaffOrder{}

	err := db.Raw(`WITH EmployeeCounts AS (
    SELECT 
        o.ID,
        COUNT(*) AS EmployeeCount
    FROM 
        orders o
        CROSS APPLY STRING_SPLIT(o.StaffCode, ',') AS staff_split
    WHERE 
        o.DocDate >= ?
        AND o.DocDate < DATEADD(DAY, 1, ?)
    GROUP BY 
        o.ID
),
EmployeeRevenue AS (
    SELECT 
        o.ID,
        s.StaffName,
        CASE 
            WHEN o.DV = 0 THEN o.Total / ec.EmployeeCount
            ELSE 0
        END AS RevenuePerEmployee,
        CASE 
            WHEN o.DV = 1 THEN 1
            ELSE 0
        END AS TOCPerEmployee
    FROM 
        orders o
        CROSS APPLY STRING_SPLIT(o.StaffCode, ',') AS staff_split
        JOIN staffs s ON s.StaffCode = staff_split.value
        JOIN EmployeeCounts ec ON o.ID = ec.ID
    WHERE 
        o.DocDate >= ?
        AND o.DocDate < DATEADD(DAY, 1, ?)
)
SELECT 
    StaffName,
    SUM(RevenuePerEmployee) AS DOANH_THU,
    SUM(TOCPerEmployee) AS TOC
FROM 
    EmployeeRevenue
GROUP BY 
    StaffName`, fromdate, todate, fromdate, todate).Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &order, nil

}

func PostGetOrder_Model(day string, month string, year string) (*[]Order, error) {
	dateString := year + month + day
	order := []Order{}
	fmt.Println(day)
	fmt.Println(month)
	fmt.Println(year)
	fmt.Println(dateString)
	err := db.Debug().Where("cast(created_at as date) = ?", dateString).Order("created_at desc").Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(order)

	return &order, nil
}

// Get id po
func GetByIdOrder_Model(id string) (*Order, error) {
	var order Order

	err := db.Where("id = ?", id).Preload("OrderDetails").First(&order).Error
	// err := db.Raw("select * from sos t0  " +
	// 	" inner join so_details t1 " +
	// 	" on t0.id = t1.IDSO " +
	// 	" where t0.id = " + id + " ",
	// ).Find(&pos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &order, nil
}

func SearchOrder_Model(tablename string, tablenumber string) (*[]Order, error) {
	order := []Order{}

	err := db.Raw(
		"select * from order a " +
			"where (charindex(N'" + tablename + "', a.TableName) > 0) " +
			"union " +
			"select * from order a " +
			"where (charindex(N'" + tablenumber + "', a.TableNumber) > 0) ",
	).Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &order, nil
}

func GetOrderTableName(tablename string) (*Order, error) {
	var order Order
	err := db.Model(&Order{}).Where("username = ?", tablename).First(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &order, nil
}

func AppendOrderDetails_Model(data map[string]interface{}, dataDetail []map[string]interface{}, id string) error {
	// var order Order
	orderDetails := make([]OrderDetails, len(dataDetail))
	for i, v := range dataDetail {
		item := OrderDetails{
			ItemCode: v["ItemCode"].(string),
			ItemName: v["ItemName"].(string),
			Quantity: v["Quantity"].(float64),
			Price:    v["Price"].(float64),
			Remarks:  v["Remarks"].(string),
			Tax:      v["Tax"].(float64),
			WhsCode:  v["WhsCode"].(string),
			Status:   v["Status"].(string),
		}
		orderDetails[i] = item
	}

	order := Order{
		TableName:    data["tablename"].(string),
		TableNumber:  data["tablenumber"].(string),
		DocDate:      data["docdate"].(time.Time),
		Remarks:      data["remarks"].(string),
		Status:       data["status"].(bool),
		Type:         data["type"].(string),
		PostingDate:  data["postingdate"].(string),
		Total:        data["total"].(float64),
		Discount:     data["discount"].(int),
		DocTotal:     data["doctotal"].(float64),
		Customerpay:  data["customerpay"].(string),
		Refund:       data["refund"].(float64),
		CardCode:     data["cardcode"].(string),
		CardName:     data["cardname"].(string),
		VAT:          data["vat"].(int),
		Creator:      data["creator"].(string),
		ViewPayment:  data["viewpayment"].(string),
		OrderDetails: orderDetails,
	}
	fmt.Println(order)
	if err := db.Model(&order).Where("id = ?", id).Association("order_details").Append(orderDetails); err != nil {
		return err
	}

	return nil
}

func GetORDERIDAuto_Model() (string, error) {
	var countID2 string
	// err := db.Raw("select ('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").Scan(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) AS IDPONum").First(&item).Error
	// err := db.Raw("select cast(('SO'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '-' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from ( select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from sos where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(PONum) > 10) AS IDPONum").Scan(&item).Error
	// item := db.Raw("select top 1 PONum countID2 from sos").First(&item)
	rows, err := db.Raw("select cast(('HD'+LEFT(CONVERT(VARCHAR(10),getdate(),12),4) + '' +  RIGHT(REPLICATE('0',4) +  cast((COUNT(countID)+1) as nvarchar(10)),5)) as nvarchar(100)) countID2 from (select CONVERT(VARCHAR(5),ROW_NUMBER() OVER(PARTITION BY LEFT(CONVERT(VARCHAR(10),DocDate,12),4) ORDER BY ID)) countID from orders where (MONTH(DocDate) = MONTH(getdate()) and YEAR(DocDate) = YEAR(GETDATE())) and LEN(CodeAuto) > 10) AS IDPONum").Rows()

	if err != nil {
		return countID2, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&countID2)
	}

	return countID2, nil
}

func GetTablenumber(code string, date string) (*[]Order, error) {
	order := []Order{}
	fmt.Println(code)

	// db.Debug().Order("created_at desc").Find(&order).Error
	// db.Debug().Where("tablenumber = ?", code).Order("created_at desc")
	// err := db.Model(&Order{}).Where("tablenumber = ?", code).First(&order).Error
	err := db.Debug().Where("tablenumber = ? AND cast(created_at as date) = ?", code, date).Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println(err)
	fmt.Println(order)

	return &order, nil
}

func GetTablenumberNotThanhToan(code string, date string) (*[]Order, error) {
	order := []Order{}
	data := "save"
	err := db.Where("viewpayment = ? AND tablenumber = ? AND cast(created_at as date) = ?", data, code, date).Find(&order).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &order, nil
}
