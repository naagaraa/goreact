package blog

import (
    "github.com/gofiber/fiber/v2"
    "goreact/api/utils"
)

/**
* blog create
*/
type BlogDataCreate struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func Create(c *fiber.Ctx) error {
    data := []BlogDataCreate{
        {
            Title:        "api blog v1 create",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request create blog successful", data)
        return c.JSON(response.ToMap())
}


/**
* blog list
*/
type BlogDataList struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}
func List(c *fiber.Ctx) error {
    data := []BlogDataList{
        {
            Title:        "api blog v1 list",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request list blog successful", data)
        return c.JSON(response.ToMap())
}
