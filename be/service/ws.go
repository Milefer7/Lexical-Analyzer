package service

import (
	"net/http"
)

func Ws(w http.ResponseWriter, r *http.Request) error {
	//var upgrader = websocket.Upgrader{
	//	ReadBufferSize:  1024,
	//	WriteBufferSize: 1024,
	//	CheckOrigin: func(r *http.Request) bool {
	//		host := r.Host
	//		origin := r.Header["Origin"]
	//		if len(origin) == 0 {
	//			return false
	//		}
	//		originHost := strings.Split(origin[0], "://")[1]
	//		return host == originHost
	//	},
	//}
	//// 升级初始的 GET 请求到一个 WebSocket 连接
	//ws, err := upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer ws.Close()
	//
	//for {
	//	// 读取从 WebSocket 发送来的新消息
	//	_, msg, err := ws.ReadMessage()
	//	if err != nil {
	//		log.Println("read:", err)
	//		return err
	//	}
	//	log.Printf("recv: %s", msg)
	//}
	//return nil
	return nil
}
