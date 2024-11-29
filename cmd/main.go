package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mantvmass/go-queue/pkg/queue"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// สร้าง Queue สำหรับเก็บงาน โดยกำหนดขนาด buffer = 10
	// หาก size เท่ากับ 10 ช่อง jobs สามารถเก็บงานได้มากสุด 10 งานก่อนที่ sender จะถูกบล็อกจนกว่า receiver จะรับงานออกจาก channel
	jobQueue := queue.NewQueue(10)

	go jobQueue.ProcessJobs()

	app.Post("/enqueue", func(c *fiber.Ctx) error {

		// var job queue.Job
		// if err := c.BodyParser(&job); err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		// }

		id := uuid.New()

		jobQueue.AddJob(queue.Job{
			Callback: func() {
				time.Sleep(2 * time.Second)
				fmt.Printf("Queue hash: %s\n", id.String())
			},
		})

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"hash":    id.String(),
			"message": "job enqueued",
		})
	})

	app.Post("/hi", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})

	log.Fatal(app.Listen(":3000"))
}
