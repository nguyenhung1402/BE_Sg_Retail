package main

import (
	"sap-crm/models"
	"sap-crm/pkg/setting"
	"sap-crm/routers"

	// "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := fiber.New(fiber.Config{
		ReadTimeout:  setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		// JSONEncoder:  sonic.Marshal,
		// JSONDecoder:  sonic.Unmarshal,
	})
	r.Use(cors.New())

	routers.InitRouter(r)

	r.Listen(":1604")
}
