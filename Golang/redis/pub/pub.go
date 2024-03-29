package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

const urlRedis = "34.135.96.5:6379"

type User struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	N_dose       int    `json:"n_dose"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     urlRedis,
	Password: "rojas", // no password set
	DB:       0,       // use default DB
})

func main() {
	app := fiber.New()

	app.Post("/insert", func(c *fiber.Ctx) error {
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
		if user.Age >= 0 && user.Age <= 11 {
			redisClient.Incr(ctx, "range0_11")
		} else if user.Age >= 12 && user.Age <= 18 {
			redisClient.Incr(ctx, "range12_18")
		} else if user.Age >= 19 && user.Age <= 26 {
			redisClient.Incr(ctx, "range19_26")
		} else if user.Age >= 27 && user.Age <= 59 {
			redisClient.Incr(ctx, "range27_59")
		} else {
			redisClient.Incr(ctx, "range60_end")
		}

		// Save name to report the last five vaccinated
		redisClient.LPush(ctx, "users", user.Name)
		return c.SendStatus(200)

	})

	app.Listen(":3000")
}
