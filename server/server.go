package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"gopkg.in/olahol/melody.v1"
	"net"
	"net/http"
	"os"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

var w2sessions [] *melody.Session

func UDPlisten() {
	/* Lets prepare a address at any address at port 10001*/
	ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Получено ",string(buf[0:n]), " от ",addr, " bytes n ", n)

		if err != nil {
			fmt.Println("Error: ",err)
		}

		msg := WSdata {
			Namespace: "objects",
			Action: "NewData",
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
			Data: string(buf[0:n]),
		}

		if data, err := json.Marshal(msg); err == nil {
			go func () {
				for _, s := range w2sessions {
					s.Write([]byte(data))
				}
			}()
		} else {
			fmt.Printf("Ошибка анмаршалинга сообщения: %s \n %+v \n", err, msg)
		}
	}
}

func RandomValue()  {
	//tick :=
}

type WSdata struct {
	Namespace string `json:"namespace"`
	Action string `json:"action"`
	Timestamp string `json:"timestamp"`
	Data string `json:"data"`
}

func WSlisten()  {
	sever := http.NewServeMux()
	m := melody.New()

	sever.Handle("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		m.HandleRequest(w, r)
	}))

	frontFS := http.FileServer(http.Dir("../views/dist"))
	sever.Handle("/dist", http.StripPrefix("/dist", frontFS))
	sever.Handle("/", frontFS)

	m.HandleConnect(func(s *melody.Session) {
		w2sessions = append(w2sessions, s)
		fmt.Printf("Новое ws-подключение: %d \n" , len(w2sessions))
	})

	//m.HandleMessage(func(s *melody.Session, msg []byte) {
	//	m.Broadcast(msg)
	//})

	http.ListenAndServe(":5000", sever)
}

func main() {

	go UDPlisten()
	WSlisten()
}