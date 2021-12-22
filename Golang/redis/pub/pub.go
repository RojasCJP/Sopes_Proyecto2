package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	N_dose       int    `json:"n_dose"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			fmt.Println("parser")
			panic(err)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			fmt.Println("marshal")
			panic(err)
		}

		// Posting through the chanel
		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			fmt.Println("publish")
			panic(err)
		}

		// Verifying age range and writing to redis - database
		if user.Age >= 0 && user.Age <= 10 {
			redisClient.Incr(ctx, "range0_10")
		} else if user.Age >= 11 && user.Age <= 20 {
			redisClient.Incr(ctx, "range11_20")
		} else if user.Age >= 21 && user.Age <= 30 {
			redisClient.Incr(ctx, "range21_30")
		} else if user.Age >= 31 && user.Age <= 40 {
			redisClient.Incr(ctx, "range31_40")
		} else if user.Age >= 41 && user.Age <= 50 {
			redisClient.Incr(ctx, "range41_50")
		} else if user.Age >= 51 && user.Age <= 60 {
			redisClient.Incr(ctx, "range51_60")
		} else if user.Age >= 61 && user.Age <= 70 {
			redisClient.Incr(ctx, "range61_70")
		} else if user.Age >= 71 && user.Age <= 80 {
			redisClient.Incr(ctx, "range71_80")
		} else {
			redisClient.Incr(ctx, "range81_end")
		}

		// Save name to report the last five vaccinated
		redisClient.LPush(ctx, "users", user.Name)
		return c.SendStatus(200)

	})

	app.Listen(":3000")
}
