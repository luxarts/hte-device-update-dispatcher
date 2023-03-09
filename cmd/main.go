package main

import (
	"context"
	"fmt"
	"hte-device-update-dispatcher/internal/controller"
	"hte-device-update-dispatcher/internal/defines"
	"hte-device-update-dispatcher/internal/metrics"
	"hte-device-update-dispatcher/internal/repository"
	"hte-device-update-dispatcher/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	metrics.StartServer()

	router := InitRouter()
	router.Run()
}
func InitRouter() *gin.Engine {
	r := gin.Default()
	MapRoutes(r)
	return r
}

func MapRoutes(router *gin.Engine) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(defines.EnvRedisHost),
	})

	ctx := context.Background()
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}

	repo := repository.NewDeviceRepository(redisClient)
	svc := service.NewDeviceService(repo)
	ctrl := controller.NewDispatcherController(svc)
	router.GET(defines.EndpointDeviceUpdate, func(c *gin.Context) {
		WsHandler(c.Writer, c.Request, ctrl.Handle)
	})
}

func WsHandler(w http.ResponseWriter, r *http.Request, handler func([]byte)) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error de conexion de WS")
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error al recibir el mensaje")
			continue
		}
		if t != websocket.TextMessage {
			continue
		}
		handler(msg)
	}
}
