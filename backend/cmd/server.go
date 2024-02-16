package main

import (
  "github.com/gofiber/fiber/v2"
  "kahoot-api/configs"
)

func main() {
  app := fiber.New()
  
  database, setupError := configs.SetUpDatabase()
  defer configs.CloseConnection(database)
  
  if setupError != nil {
    panic(setupError.Error())
  }

  app.Get("/", func(context *fiber.Ctx) error {
    return context.SendString("Hello World")
  })

  app.Listen(":3000")
}
