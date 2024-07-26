package routers

import (
	"fmt"
	jwtCustom "sap-crm/middleware/jwt"
	"sap-crm/pkg/upload"
	"sap-crm/routers/api"
	v1 "sap-crm/routers/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// InitRouter initialize routing information
func InitRouter(r *fiber.App) {
	apiv1 := r.Group("/api/v1")
	// dang ky
	apiv1.Post("/signup", v1.SignUp)
	// dang nhap
	apiv1.Post("/signin", v1.SignIn)

	// // Receipt
	// receiptRoute := apiv1.Group("/receipt")
	// receiptRoute.Get("/", v1.GetReceipt_Router)
	// receiptRoute.Get("/:id", v1.GetByIdReceipt_Router)
	// receiptRoute.Post("/", v1.PostReceipt)
	// receiptRoute.Put("/:id", v1.Put)

	// apiv1.Use(csrf.New(csrf.ConfigDefault))         // CSRF
	apiv1.Use(cors.New(cors.ConfigDefault)) // CORS
	// apiv1.Use(cache.New(cache.ConfigDefault))       // Cache
	apiv1.Use(compress.New(compress.ConfigDefault)) // Compress
	apiv1.Use(logger.New(logger.ConfigDefault))     // Logger
	apiv1.Use(requestid.New())
	fmt.Println(upload.GetImageFullPath())
	apiv1.Static("/upload/images", upload.GetImageFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	apiv1.Static("/upload/images", upload.GetImageFullPath(), fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	// JWT check
	apiv1.Use(jwtCustom.Protected())
	apiv1.Use(jwtCustom.Authorization())
	// Route
	apiv1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👌!")

	})
	// Member
	memberRoute := apiv1.Group("/members")
	memberRoute.Get("/", v1.GetAllUser_Router)
	memberRoute.Get("/:id", v1.GetByIdUser_Router)
	memberRoute.Put("/:id", v1.PutMembers)

	// Role
	roleRoute := apiv1.Group("/roles")
	roleRoute.Post("/", v1.AddRole)

	// Menu
	menuRoute := apiv1.Group("/menus")
	menuRoute.Get("/", v1.GetMenu_Router)
	menuRoute.Get("/:id", v1.GetMenu_Router)
	menuRoute.Post("/", v1.PostMenu)

	// ở là code tạm để gọi API
	// Items
	itemsRoute := apiv1.Group("/items")
	itemsRoute.Get("/", v1.GetItems_Router)
	itemsRoute.Get("/:id", v1.GetById_Router)
	itemsRoute.Post("/", v1.PostItems)
	itemsRoute.Put("/:id", v1.PutItems)

	// BP
	businessRoute := apiv1.Group("/business")
	businessRoute.Get("/", v1.GetBPs_Router)
	businessRoute.Get("/:id", v1.GetByIdPBs_Router)
	businessRoute.Post("/", v1.PostBPs)
	businessRoute.Put("/:id", v1.PutBPs)

	// Staff
	staffRoute := apiv1.Group("/staff")
	staffRoute.Get("/", v1.GetStaff_Router)
	staffRoute.Get("/:id", v1.GetByIdStaff_Router)
	staffRoute.Post("/", v1.PostStaff)
	staffRoute.Put("/:id", v1.PutStaff)

	// Warehouse
	warehouseRoute := apiv1.Group("/warehouse")
	warehouseRoute.Get("/", v1.GetWhs_Router)
	warehouseRoute.Get("/:id", v1.GetByIdWhs_Router)
	warehouseRoute.Post("/", v1.PostWhs)
	warehouseRoute.Put("/:id", v1.PutWhs)

	// Order
	orderRoute := apiv1.Group("/order")
	orderRoute.Get("/", v1.GetOrder_Router)
	orderRoute.Get("/staff", v1.Staff_Oder)
	orderRoute.Get("/:id", v1.GetByIdOrder_Router)
	orderRoute.Post("/allthanhtoan", v1.GetTablenumberNotThanhToan_Router)
	orderRoute.Post("/", v1.PostOrder)
	orderRoute.Post("/mapTableOrderDetails", v1.PostGetOrder_Router)
	orderRoute.Post("/tablenumber", v1.GetTablenumber_Router)
	orderRoute.Put("/:id", v1.PutOrder)
	orderRoute.Put("/updateorderdetail/:id", v1.Post_AppendOrderDetails)

	// Order_details
	orderDetailRoute := apiv1.Group("/orderDetail")
	// orderDetailRoute.Get("/", v1.GetOrder_Router)
	orderDetailRoute.Get("/:id", v1.GetByIdOrderDetail_Router)
	orderDetailRoute.Put("/updateandinsert/:id", v1.PostOrderDetails_Router)
	orderDetailRoute.Put("/:id", v1.PutOrderDetailItem)

	// Order
	categoryRoute := apiv1.Group("/category")
	categoryRoute.Get("/", v1.GetCategory_Router)
	categoryRoute.Get("/:id", v1.GetByIdCategory_Router)
	categoryRoute.Post("/", v1.PostCategory)
	categoryRoute.Put("/:id", v1.PutCategory)

	// Table
	tableRoute := apiv1.Group("/table")
	tableRoute.Get("/", v1.GetTable_Router)
	tableRoute.Get("/:id", v1.GetByIdTable_Router)
	tableRoute.Post("/", v1.PostTable)
	tableRoute.Put("/:id", v1.PutTable)

	posRoute := apiv1.Group("/pos")
	posRoute.Get("/", v1.GetPOS_Router)
	posRoute.Get("/:id", v1.GetByIdPOS_Router)
	posRoute.Post("/", v1.PostPOS)
	posRoute.Put("/:id", v1.PutPOS)

	// Receipt
	receiptRoute := apiv1.Group("/receipt")
	receiptRoute.Get("/", v1.GetAllReceipt_Router)
	receiptRoute.Get("/:id", v1.GetByIdReceipt_Router)
	receiptRoute.Post("/", v1.PostReceipt)
	// receiptRoute.Put("/:id", v1.Put)

	// issuse
	issueRoute := apiv1.Group("/issue")
	issueRoute.Get("/", v1.GetAllIssue_Router)
	issueRoute.Get("/:id", v1.GetByIdIssue_Router)
	issueRoute.Post("/", v1.PostIssue)
	// issueRoute.Put("/:id", v1.Put)

	// ViewNXT
	viewNXTRoute := apiv1.Group("/viewnxt")
	viewNXTRoute.Get("/", v1.GetAllNXT_Router)
	viewNXTRoute.Post("/whscode", v1.GetWhsCodeNXT_Router)
	viewNXTRoute.Post("/whsitem", v1.GetWhsItemCodeNXT_Router)

	apiv1.Post("/upload", api.UploadImageSingle)
	apiv1.Post("/upload/multiple", api.UploadImageMultiple)

	reportsRoute := apiv1.Group("/reports")
	reportsRoute.Get("/years", v1.GetYears_Router)
	// reportsRoute.Post("/", v1.GetCharTopDoanhThu_Router)
	reportsRoute.Post("/itemcode", v1.GetCharTopItemCode_Router)
	reportsRoute.Post("/chartday", v1.GetCharDay_Router)
	reportsRoute.Post("/chartdaycustomer", v1.GetCharDayCus_Router)
	reportsRoute.Post("/chartmonth", v1.GetCharMonth_Router)
	reportsRoute.Post("/chartyear", v1.GetCharYear_Router)
	reportsRoute.Post("/chartdaycircle", v1.GetCharDayCircle_Router)
	reportsRoute.Post("/chartmonthcircle", v1.GetCharMonthCircle_Router)
	reportsRoute.Post("/chartyearcircle", v1.GetCharYearCircle_Router)
	reportsRoute.Post("/chartnxt", v1.GetCharNXT_Router)
	reportsRoute.Post("/chartreceipt", v1.GetChartReceipt_Router)
	reportsRoute.Post("/chartissue", v1.GetChartIssue_Router)

	configSetting := apiv1.Group("/configsetting")
	configSetting.Get("/", v1.GetConfigSettings_Router)
	configSetting.Get("/:id", v1.GetByIdSettings_Router)
	configSetting.Post("/", v1.PostConfigSettings_v1)
	configSetting.Put("/:id", v1.PutConfigSettings)

}
