package users

import (
    "github.com/gofiber/fiber/v2"
    "goreact/api/utils"
)


/**
* users create
*/
type UserDataCreate struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func Create(c *fiber.Ctx) error {
    data := []UserDataCreate{
        {
            Title:        "api users v1 create",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request create user successful", data)
        return c.JSON(response.ToMap())
}


/**
* roles list
*/
type UserDataList struct {
    Title        string `json:"title"`
    Descriptions string `json:"descriptions"`
}

func List(c *fiber.Ctx) error {
    data := []UserDataList{
        {
            Title:        "api users v1 list",
            Descriptions: "api",
        },
    }
    
    response, _ := utils.ResponseStandart(true, 200, "golang Request user list successful", data)
        return c.JSON(response.ToMap())
}
