package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	authController "goreact/backend/golang/controller/auth"
	DASHBOARD_dashboard_controller_v1 "goreact/backend/golang/controller/backend/v1/dashboard"
	DASHBOARD_dashboard_controller_v2 "goreact/backend/golang/controller/backend/v2/dashboard"
	CMS_blog_controller_v1 "goreact/backend/golang/controller/backend/v1/cms/blog"
	CMS_category_controller_v1 "goreact/backend/golang/controller/backend/v1/cms/category"
	UMS_roles_controller_v1 "goreact/backend/golang/controller/backend/v1/ums/roles"
	UMS_users_controller_v1 "goreact/backend/golang/controller/backend/v1/ums/users"
)

// Exported function Main (needs to be uppercase to be exported)
func Routes() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/help", func(c *fiber.Ctx) error {
		return c.SendString("help pages")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("about pages")
	})

	// api routes
	// Create a group with a common prefix
	v1 := app.Group("/api/v1")
	v2 := app.Group("/api/v2")

	// version 1
	// dashboard
	v1_dashboard := v1.Group("/dashboard")
	v1_dashboard.Get("/",DASHBOARD_dashboard_controller_v1.Dashboard)

	// auth
	v1_auth := v1.Group("/auth")
	v1_auth.Get("/login", authController.Login )
	v1_auth.Get("/register", authController.Register )

	// cms
	v1_cms := v1.Group("/cms")
	v1_cms.Get("blog", CMS_blog_controller_v1.List)
	v1_cms.Get("category", CMS_category_controller_v1.List)
	
	// cms
	v1_ums := v1.Group("/ums")
	v1_ums.Get("roles", UMS_roles_controller_v1.List)
	v1_ums.Get("users", UMS_users_controller_v1.List)

	// end version 1


	// version 2
	v2_dashboard := v2.Group("/dashboard")
	v2_dashboard.Get("/",DASHBOARD_dashboard_controller_v2.Dashboard)


	app.Listen(":3000")
}
