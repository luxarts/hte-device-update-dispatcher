package main

import (
	"fmt"
	"hte-device-update-dispatcher/internal/controller"
	"hte-device-update-dispatcher/internal/defines"
	"hte-device-update-dispatcher/internal/repository"
	"hte-device-update-dispatcher/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := InitRouter()
	router.Run("localhost:8080")
}
func InitRouter() *gin.Engine {
	r := gin.Default()
	MapRoutes(r)
	return r
}

func MapRoutes(router *gin.Engine) {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
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
