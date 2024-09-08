package dashboard

import (
    "github.com/gofiber/fiber/v2"
    "goreact/api/utils"
)

/**
* dashboard list
*/
type DashboardData struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func Dashboard(c *fiber.Ctx) error {
    data := []DashboardData{
        {
            Title:        "api dashboard v1",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request dashboard v2 successful", data)
        return c.JSON(response.ToMap())
}

