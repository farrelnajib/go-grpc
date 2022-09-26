package nonrpchandler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func SampleHandler() HttpHandlerFunc {
	return func(c *fiber.Ctx) error {
		successResponse := make(map[string]interface{})
		successResponse["response"] = "Success"
		jsonSuccessResponse, _ := json.Marshal(successResponse)
		c.Send(jsonSuccessResponse)
		return nil
	}
}
