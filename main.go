package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")

	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not found: ", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/api/hello", responser)

	port := Config.GetInt("port")
	err := r.Run(":" + fmt.Sprint(port))
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}

func responser(c *gin.Context) {
	failureRate := Config.GetFloat64("failure_rate")
	maxDelay := Config.GetInt("max_delay")

	if failureRate > 0 {
		if rand.Float64() < failureRate {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}

	if maxDelay > 0 {
		time.Sleep(time.Duration(rand.Intn(maxDelay)) * time.Millisecond)
	}

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
