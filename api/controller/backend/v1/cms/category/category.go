package category

import (
    "github.com/gofiber/fiber/v2"
    "goreact/api/utils"
)

/**
* blog create
*/
type CategoryDataCreate struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func Create(c *fiber.Ctx) error {
    data := []CategoryDataCreate{
        {
            Title:        "api category v1 create",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request create category successful", data)
        return c.JSON(response.ToMap())
}


/**
* category create
*/
type CategoryDataList struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}
func List(c *fiber.Ctx) error {
    data := []CategoryDataList{
        {
            Title:        "api category v1 list",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request list category successful", data)
        return c.JSON(response.ToMap())
}
