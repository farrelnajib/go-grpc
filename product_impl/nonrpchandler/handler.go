package nonrpchandler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type HttpHandlerFunc func(c *fiber.Ctx) error

type Handler struct {
	handlerFunc HttpHandlerFunc
}

func NewHttpHandler(handlerFunc HttpHandlerFunc) Handler {
	return Handler{
		handlerFunc: handlerFunc,
	}
}

func (ths Handler) ServeHTTP(c *fiber.Ctx) {
	err := ths.handlerFunc(c)
	if err != nil {
		log.Fatalf("Failed to server HTTP handler: %v", err.Error())
	}
}
