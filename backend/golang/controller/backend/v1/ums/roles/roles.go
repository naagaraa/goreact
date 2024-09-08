package roles

import (
    "github.com/gofiber/fiber/v2"
    "goreact/api/utils"
)


/**
* roles create
*/
type RolesDataCreate struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func Create(c *fiber.Ctx) error {
    data := []RolesDataCreate{
        {
            Title:        "api roles v1 create",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request create roles successful", data)
        return c.JSON(response.ToMap())
}


/**
* roles list
*/
type RolesDataList struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func List(c *fiber.Ctx) error {
    data := []RolesDataList{
        {
            Title:        "api roles v1 list",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request list user successful", data)
        return c.JSON(response.ToMap())
}
