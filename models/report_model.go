package models

import (
	"fmt"
	"gorm.io/gorm"
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
type Customer struct {
	Sl       string  `gorm:"column: Sl" json:"Sl"`
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
	WhsName  string  `gorm:"column: WhsName" json:"whsname"`
}

type NhapXuat struct {
	WhsCode  string  `gorm:"column: WhsCode" json:"whscode"`
	ItemCode string  `gorm:"column: ItemCode" json:"itemcode"`
	Quantity float64 `gorm:"column: Quantity" json:"quantity"`
}

// func GetChartTopDoanhThu_Model(month string, year string) (*[]ChartTopDoanhThu, error) {

// 	item := []ChartTopDoanhThu{}

// 	result := db.Raw("select top 6 t0.CardCode, t2.CardName, SUM(isnull(t1.TotalAfter,0)) Total from sos t0 " +
// 		" inner join so_details t1 on t0.id = t1.IDSO " +
// 		" inner join b_ps t2 on t0.CardCode = t2.CardCode " +
// 		" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' " +
// 		" group by t0.CardCode, t2.CardName " +
// 		" order by SUM(isnull(t1.TotalAfter,0)) DESC ").Find(&item)
// 	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
// 		return nil, result.Error
// 	}

// 	return &item, nil
// }

// SELECT        WhsCode, ItemCode, SUM(Quantity) AS Quantity
// FROM            (SELECT        t0.WhsCode, t1.ItemCode, SUM(t1.Quantity) AS Quantity
//                           FROM            dbo.receipts AS t0 INNER JOIN
//                                                     dbo.receipt_details AS t1 ON t0.id = t1.IDReceipt
//                           WHERE        (t0.WhsCode IS NOT NULL) OR
//                                                     (t0.WhsCode <> '')
//                           GROUP BY t0.WhsCode, t1.ItemCode
//                           UNION ALL
//                           SELECT        t0.WhsCode, t1.ItemCode, - SUM(t1.Quantity) AS Quantity
//                           FROM            dbo.receipts AS t0 INNER JOIN
//                                                    dbo.issue_details AS t1 ON t0.id = t1.IDIssue
//                           WHERE        (t0.WhsCode IS NOT NULL) OR
//                                                    (t0.WhsCode <> '')
//                           GROUP BY t0.WhsCode, t1.ItemCode
//                           UNION ALL
//                           SELECT        WhsCode, ItemCode, - SUM(Quantity) AS Quantity
//                           FROM            dbo.order_details AS t0
//                           WHERE        (WhsCode IS NOT NULL) OR
//                                                    (WhsCode <> '')
//                           GROUP BY WhsCode, ItemCode
//                           UNION ALL
//                           SELECT        WhsCode, ItemCode, - SUM(Quantity) AS Quantity
//                           FROM            dbo.pos_details AS t0
//                           WHERE        (WhsCode IS NOT NULL) OR
//                                                    (WhsCode <> '')
//                           GROUP BY WhsCode, ItemCode) AS A
// GROUP BY WhsCode, ItemCode

func GetChartTopItemCode_Model(month string, year string) (*[]ChartTopItemCode, error) {

	item := []ChartTopItemCode{}

	result := db.Raw("select top 6 t1.ItemCode, t2.ItemName, SUM(isnull(t0.DocTotal,0)) Total from pos t0 " +
		" inner join pos_details t1 on t0.id = t1.IDPOS " +
		" inner join items t2 on t1.ItemCode = t2.ItemCode " +
		" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
		" group by t1.ItemCode, t2.ItemName " +
		" order by SUM(isnull(t0.DocTotal,0)) DESC").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return &item, nil
}

func GetYear_Model() (*[]NameYears, error) {
	item := []NameYears{}
	result := db.Raw("select A.Name1, A.Name2 from ( " +
		" select distinct 'Year' Name1, YEAR(t0.DocDate) Name2 from pos t0 " +
		" union " +
		" select distinct 'Year' Name1, YEAR(t0.DocDate) Name2 from orders t0 " +
		" ) A  order by YEAR(A.Name2) DESC").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return &item, nil
}

//	func GetChartDay_Model(day string, month string, year string) (*[]ChartMonth, error) {
//		item := []ChartMonth{}
//		result := db.Raw("select A.[Group], A.ItemCode, SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//			" select distinct t3.CategoryName [Group],null ItemCode, (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
//			" from orders t0 " +
//			" inner join order_details t1 on t0.id = t1.IDOrder " +
//			" inner join items t2 on t1.ItemCode = t2.ItemCode " +
//			" inner join categories t3 on t2.[Group] = t3.CategoryCode " +
//			" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//			" ) A" +
//			" group by A.[Group], A.ItemCode " +
//			" order by A.[Group] ").Find(&item)
//		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//			return nil, result.Error
//		}
//		return &item, nil
//	}
//func GetChartDay_Model(day string, month string, year string) (*[]ChartMonth, error) {
//	item := []ChartMonth{}
//	//result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//	//	" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
//	//	" from orders t0 " +
//	//	" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//	//	" ) A ").Find(&item)
//	result := db.Raw("select A.Total Total, A.Discount Discount from ( " +
//		" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
//		" from orders t0 " +
//		" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//		" ) A ").Find(&item)
//	fmt.Println(&item)
//	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//		return nil, result.Error
//	}
//	return &item, nil
//}

//	func GetChartDayCircle_Model(day string, month string, year string) (*[]Chart, error) {
//		item := []Chart{}
//		result := db.Raw("select A.[Group], SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//			" select distinct t3.CategoryName [Group], SUM(isnull(t0.DocTotal,0)) Total, SUM((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount  " +
//			" from orders t0  " +
//			" inner join order_details t1 on t0.id = t1.IDOrder " +
//			" inner join items t2 on t1.ItemCode = t2.ItemCode " +
//			" inner join categories t3 on t2.[Group] = t3.CategoryCode " +
//			" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//			" group by t3.CategoryName) " +
//			" A " +
//			" group by A.[Group] " +
//			" order by A.[Group]").Find(&item)
//		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//			return nil, result.Error
//		}
//		return &item, nil
//	}
func GetChartDayCircle_Model(day string, month string, year string) (*[]Chart, error) {
	item := []Chart{}
	result := db.Raw("select A.Total Total, A.Discount Discount from ( " +
		" select SUM(isnull(t0.DocTotal,0)) Total, SUM((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount  " +
		" from orders t0  " +
		" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
		" ) A ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

//	func GetChartMonth_Model(month string, year string) (*[]ChartMonth, error) {
//		item := []ChartMonth{}
//		result := db.Raw("select A.[Group], A.ItemCode, SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//			" select distinct t3.CategoryName [Group],null ItemCode, (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
//			" from orders t0 " +
//			" inner join order_details t1 on t0.id = t1.IDOrder " +
//			" inner join items t2 on t1.ItemCode = t2.ItemCode " +
//			" inner join categories t3 on t2.[Group] = t3.CategoryCode " +
//			" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//			" ) A" +
//			" group by A.[Group], A.ItemCode " +
//			" order by A.[Group] ").Find(&item)
//		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//			return nil, result.Error
//		}
//		return &item, nil
//	}
func GetChartMonth_Model(month string, year string) (*[]ChartMonth, error) {
	item := []ChartMonth{}
	result := db.Raw("select A.Total Total, A.Discount Discount from ( " +
		" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
		" from orders t0 " +
		" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
		" ) A ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

//	func GetChartMonthCircle_Model(month string, year string) (*[]Chart, error) {
//		item := []Chart{}
//		result := db.Raw("select A.[Group], SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//			" select distinct t3.CategoryName [Group], SUM(isnull(t0.DocTotal,0)) Total, SUM((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount  " +
//			" from orders t0  " +
//			" inner join order_details t1 on t0.id = t1.IDOrder " +
//			" inner join items t2 on t1.ItemCode = t2.ItemCode " +
//			" inner join categories t3 on t2.[Group] = t3.CategoryCode " +
//			" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
//			" group by t3.CategoryName) " +
//			" A " +
//			" group by A.[Group] " +
//			" order by A.[Group]").Find(&item)
//		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//			return nil, result.Error
//		}
//		return &item, nil
//	}
func GetChartMonthCircle_Model(month string, year string) (*[]Chart, error) {
	item := []Chart{}
	result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
		" select SUM(isnull(t0.DocTotal,0)) Total, SUM((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount  " +
		" from orders t0  " +
		" where MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
		" ) A  ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

//	func GetChartYear_Model(year string) (*[]ChartMonth, error) {
//		item := []ChartMonth{}
//		result := db.Raw("select A.[Group], A.ItemCode, SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
//			" select distinct t3.CategoryName [Group],null ItemCode, (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
//			" from orders t0 " +
//			" inner join order_details t1 on t0.id = t1.IDOrder " +
//			" inner join items t2 on t1.ItemCode = t2.ItemCode " +
//			" inner join categories t3 on t2.[Group] = t3.CategoryCode " +
//			" where cast(YEAR(t0.DocDate) as nvarchar(10)) ='" + year + "' and t0.ViewPayment = 'off' " +
//			" ) A" +
//			" group by A.[Group], A.ItemCode " +
//			" order by A.[Group] ").Find(&item)
//		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
//			return nil, result.Error
//		}
//		return &item, nil
//	}
func GetChartYear_Model(year string) (*[]ChartMonth, error) {
	item := []ChartMonth{}
	result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
		" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
		" from orders t0 " +
		" where cast(YEAR(t0.DocDate) as nvarchar(10)) ='" + year + "' and t0.ViewPayment = 'off' " +
		" ) A ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

func GetChartYearCircle_Model(year string) (*[]Chart, error) {
	item := []Chart{}
	result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
		" select SUM(isnull(t0.DocTotal,0)) Total, SUM((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount  " +
		" from orders t0  " +
		" where cast(YEAR(t0.DocDate) as nvarchar(10)) ='" + year + "' and t0.ViewPayment = 'off' " +
		" ) A  ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

// test week
func GetChartDay_Model(fromdate string, todate string) (*[]ChartMonth, error) {
	item := []ChartMonth{}
	//result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
	//	" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
	//	" from orders t0 " +
	//	" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
	//	" ) A ").Find(&item)
	result := db.Raw("select A.Total Total, A.Discount Discount from ( " +
		" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
		" from orders t0 " +
		" where t0.DocDate >= '" + fromdate + "' AND t0.DocDate < DATEADD(DAY, 1,'" + todate +
		"') and t0.ViewPayment = 'off' " +
		" ) A ").Find(&item)
	fmt.Println(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

// cus
func GetChartDayCus_Model(fromdate string, todate string) (*[]Customer, error) {
	item := []Customer{}
	//result := db.Raw("select SUM(isnull(A.Total,0)) Total, SUM(isnull(A.Discount,0)) Discount from ( " +
	//	" select (isnull(t0.DocTotal,0)) Total, ((isnull(t0.Total,0) * (isnull(t0.Discount,0)))/100) Discount " +
	//	" from orders t0 " +
	//	" where DAY(t0.DocDate) = '" + day + "' and MONTH(t0.DocDate) = '" + month + "' and YEAR(t0.DocDate) ='" + year + "' and t0.ViewPayment = 'off' " +
	//	" ) A ").Find(&item)
	result := db.Raw("select COUNT(DISTINCT A.CardCode) AS Sl, SUM(A.Total) AS Total,  SUM(A.Discount) AS Discount  from ( " +
		" select t0.CardCode,isnull(t0.DocTotal,0)  AS Total, (isnull(t0.Total,0) * isnull(t0.Discount,0))/100 AS Discount " +
		" from orders t0 " +
		" where t0.DocDate >= '" + fromdate + "' AND t0.DocDate < DATEADD(DAY, 1,'" + todate +
		"') and t0.ViewPayment = 'off' " +
		" ) A ").Find(&item)
	fmt.Println(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}
func GetChartNXT_Model(fromdate string, todate string) (*[]NXT, error) {
	item := []NXT{}
	result := db.Raw("select " +
		" C.WhsCode, W.WhsName, C.ItemCode," +
		" SUM(isnull(C.BeginQty,0)) BeginQty, " +
		" SUM(isnull(C.InQty,0)) InQty, " +
		" SUM(isnull(C.OutQty,0)) OutQty, " +
		" SUM(isnull(C.EndQty,0)) EndQty " +
		" from( " +
		" select  " +
		" B.WhsCode, B.ItemCode, " +
		" (case when(cast(B.docdate as date) < CAST('" + fromdate + "' as date)) then (isnull(B.inQty,0)-isnull(B.outQty,0)) else 0 end) BeginQty, " +
		" (case when(cast(B.docdate as date) between CAST('" + fromdate + "' as date) and CAST('" + todate + "' as date)) then (isnull(B.inQty,0)) else 0 end) InQty, " +
		" (case when(cast(B.docdate as date) between CAST('" + fromdate + "' as date) and CAST('" + todate + "' as date)) then (isnull(B.outQty,0)) else 0 end) OutQty, " +
		" (case when(cast(B.docdate as date) <= CAST('" + todate + "' as date)) then (isnull(B.inQty,0)-isnull(B.outQty,0)) else 0 end) EndQty " +
		" from ( " +
		" SELECT " +
		" WhsCode, ItemCode, " +
		" (CASE when A.objtype = 'RP' THEN A.Quantity ELSE 0 END) AS inQty, " +
		" (CASE when A.objtype in ('ISS','ORD','ORD') THEN A.Quantity ELSE 0 END) AS outQty, " +
		" A.docdate " +
		" FROM  " +
		" (SELECT         " +
		" t0.WhsCode, t2.ItemName ItemCode, SUM(t1.Quantity) AS Quantity, 'RP' objtype, cast(t0.created_at as date) docdate " +
		" FROM dbo.receipts AS t0  " +
		" INNER JOIN dbo.receipt_details AS t1 ON t0.id = t1.IDReceipt " +
		" INNER JOIN dbo.items AS t2 ON t1.ItemCode = t2.ItemCode " +
		" WHERE (t0.WhsCode IS NOT NULL) OR (t0.WhsCode <> '') " +
		" GROUP BY t0.WhsCode, t2.ItemName,cast(t0.created_at as date) " +
		" UNION ALL " +
		" SELECT         " +
		" 	t0.WhsCode, t2.ItemName ItemCode, SUM(t1.Quantity) AS Quantity, 'ISS' objtype, cast(t0.created_at as date) docdate " +
		" FROM dbo.issues AS t0  " +
		" INNER JOIN dbo.issue_details AS t1 ON t0.id = t1.IDIssue " +
		" INNER JOIN dbo.items AS t2 ON t1.ItemCode = t2.ItemCode " +
		" WHERE (t0.WhsCode IS NOT NULL) OR (t0.WhsCode <> '') " +
		" GROUP BY t0.WhsCode, t2.ItemName,cast(t0.created_at as date) " +
		" UNION ALL " +
		" SELECT         " +
		" 	WhsCode, t2.ItemName ItemCode, SUM(t0.Quantity) AS Quantity, 'ORD' objtype, cast(t0.created_at as date) docdate " +
		" FROM dbo.order_details AS t0 " +
		" INNER JOIN dbo.items AS t2 ON t0.ItemCode = t2.ItemCode " +
		" WHERE (WhsCode IS NOT NULL) OR (WhsCode <> '') and t0.Status = 'off' " +
		" GROUP BY WhsCode, t2.ItemName,cast(t0.created_at as date)" +
		" UNION ALL " +
		" SELECT         " +
		" WhsCode, t2.ItemName ItemCode, SUM(t0.Quantity) AS Quantity, 'POS' objtype, cast(t0.created_at as date) docdate " +
		" FROM dbo.pos_details AS t0 " +
		" INNER JOIN dbo.items AS t2 ON t0.ItemCode = t2.ItemCode " +
		" WHERE (WhsCode IS NOT NULL) OR (WhsCode <> '') " +
		" GROUP BY WhsCode, t2.ItemName,cast(t0.created_at as date)) " +
		" AS A " +
		" where cast(A.docdate as date) between CAST('" + fromdate + "' as date) and CAST('" + todate + "' as date) " +
		" ) AS B " +
		" where (TRIM(B.WhsCode) is not null) AND (TRIM(B.WhsCode) <> '') " +
		" ) AS C " +
		"INNER JOIN dbo.warehouses W ON C.WhsCode = W.WhsCode" +
		" group by C.WhsCode, W.WhsName, C.ItemCode " +
		" order by C.WhsCode ").Find(&item)
	fmt.Println(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

// func GetChartNXT_Model(fromdate string, todate string) (*[]NXT, error) {
// 	item := []NXT{}
// 	result := db.Raw("select " +
// 		" C.WhsCode, C.ItemCode, " +
// 		" SUM(isnull(C.BeginQty,0)) BeginQty, " +
// 		" SUM(isnull(C.InQty,0)) InQty, " +
// 		" SUM(isnull(C.EndQty,0)) EndQty " +
// 		" from( " +
// 		" select  " +
// 		" B.WhsCode, B.ItemCode, " +
// 		" (case when(cast(B.docdate as date) < CAST('" + fromdate + "' as date)) then (isnull(B.inQty,0)-isnull(B.outQty,0)) else 0 end) BeginQty, " +
// 		" (case when(cast(B.docdate as date) between CAST('" + fromdate + "' as date) and CAST('" + todate + "' as date)) then (isnull(B.inQty,0)-isnull(B.outQty,0)) else 0 end) InQty, " +
// 		" (case when(cast(B.docdate as date) <= CAST('" + todate + "' as date)) then (isnull(B.inQty,0)-isnull(B.outQty,0)) else 0 end) EndQty " +
// 		" from ( " +
// 		" SELECT " +
// 		" WhsCode, ItemCode, " +
// 		" (CASE when A.objtype = 'RP' THEN A.Quantity ELSE 0 END) AS inQty, " +
// 		" (CASE when A.objtype in ('ISS','ORD','ORD') THEN A.Quantity ELSE 0 END) AS outQty, " +
// 		" A.docdate " +
// 		" FROM  " +
// 		" (SELECT         " +
// 		" t0.WhsCode, t1.ItemCode, SUM(t1.Quantity) AS Quantity, 'RP' objtype, cast(t0.created_at as date) docdate " +
// 		" FROM dbo.receipts AS t0  " +
// 		" INNER JOIN dbo.receipt_details AS t1 ON t0.id = t1.IDReceipt " +
// 		" WHERE (t0.WhsCode IS NOT NULL) OR (t0.WhsCode <> '') " +
// 		" GROUP BY t0.WhsCode, t1.ItemCode,cast(t0.created_at as date) " +
// 		" UNION ALL " +
// 		" SELECT         " +
// 		" 	t0.WhsCode, t1.ItemCode, - SUM(t1.Quantity) AS Quantity, 'ISS' objtype, cast(t0.created_at as date) docdate " +
// 		" FROM dbo.issues AS t0  " +
// 		" INNER JOIN dbo.issue_details AS t1 ON t0.id = t1.IDIssue " +
// 		" WHERE (t0.WhsCode IS NOT NULL) OR (t0.WhsCode <> '') " +
// 		" GROUP BY t0.WhsCode, t1.ItemCode,cast(t0.created_at as date) " +
// 		" UNION ALL " +
// 		" SELECT         " +
// 		" 	WhsCode, ItemCode, - SUM(Quantity) AS Quantity, 'ORD' objtype, cast(t0.created_at as date) docdate " +
// 		" FROM dbo.order_details AS t0 " +
// 		" WHERE (WhsCode IS NOT NULL) OR (WhsCode <> '') " +
// 		" GROUP BY WhsCode, ItemCode,cast(t0.created_at as date) " +
// 		" UNION ALL " +
// 		" SELECT         " +
// 		" WhsCode, ItemCode, - SUM(Quantity) AS Quantity, 'POS' objtype, cast(t0.created_at as date) docdate " +
// 		" FROM dbo.pos_details AS t0 " +
// 		" WHERE (WhsCode IS NOT NULL) OR (WhsCode <> '') " +
// 		" GROUP BY WhsCode, ItemCode,cast(t0.created_at as date)) " +
// 		" AS A " +
// 		" where cast(A.docdate as date) between CAST('" + fromdate + "' as date) and CAST('" + todate + "' as date) " +
// 		" ) AS B " +
// 		" where (TRIM(B.WhsCode) is not null) AND (TRIM(B.WhsCode) <> '') " +
// 		" ) AS C " +
// 		" group by C.WhsCode, C.ItemCode " +
// 		" order by C.WhsCode ").Find(&item)
// 	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
// 		return nil, result.Error
// 	}
// 	return &item, nil
// }

func GetChartReceipt_Model(fromdate string, todate string) (*[]NhapXuat, error) {
	item := []NhapXuat{}
	result := db.Raw("SELECT " +
		" t2.WhsName WhsCode, t1.ItemCode, SUM(t1.Quantity) AS Quantity " +
		" FROM dbo.receipts AS t0  " +
		" INNER JOIN dbo.receipt_details AS t1 ON t0.id = t1.IDReceipt " +
		" INNER JOIN dbo.warehouses AS t2 on t0.WhsCode = t2.WhsCode " +
		" WHERE ((t0.WhsCode IS NOT NULL) AND (t0.WhsCode <> '')) AND (cast(t0.created_at as date) between cast('" + fromdate + "' as date) and cast('" + todate + "' as date)) " +
		" GROUP BY t2.WhsName, t1.ItemCode " +
		" order by t2.WhsName ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}

func GetChartIssue_Model(fromdate string, todate string) (*[]NhapXuat, error) {
	item := []NhapXuat{}
	result := db.Raw("SELECT " +
		" t2.WhsName WhsCode, t1.ItemCode, SUM(t1.Quantity) AS Quantity " +
		" FROM dbo.issues AS t0  " +
		" INNER JOIN dbo.issue_details AS t1 ON t0.id = t1.IDIssue " +
		" INNER JOIN dbo.warehouses AS t2 on t0.WhsCode = t2.WhsCode " +
		" WHERE ((t0.WhsCode IS NOT NULL) AND (t0.WhsCode <> '')) AND (cast(t0.created_at as date) between cast('" + fromdate + "' as date) and cast('" + todate + "' as date)) " +
		" GROUP BY t2.WhsName, t1.ItemCode " +
		" order by t2.WhsName ").Find(&item)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &item, nil
}
