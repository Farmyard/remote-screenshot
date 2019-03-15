package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/go-vgo/robotgo"
)

type ScreenshotService struct{}

func (p *ScreenshotService) SendBitmap(bitmapstr string, reply *string) error {
	bitmap := robotgo.BitmapStr(bitmapstr)
	robotgo.CopyBitPB(bitmap)
	return nil
}

func main() {
	rpc.RegisterName("ScreenshotService", new(ScreenshotService))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Args[1]))
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}
}
