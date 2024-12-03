package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamieyoung5/kwikmedical-cas/internal/handler"
	"go.uber.org/zap"
	"os"
)

const port = "4444"

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Error("Error initializing logger", zap.Error(err))
		os.Exit(1)
	}

	apiHandler := handler.NewHandler(logger)

	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/authenticate", apiHandler.Authenticate)

	logger.Info("Starting Api server", zap.String("port", port))
	if err = router.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

// TODO: add proper rules if deploying anywhere not local
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
