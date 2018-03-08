package main
import(
	"easy_server"
	"fmt"
//	"os"
	"time"
	"runtime"
)

func JustPrint(op easy_server.TcpConnectionOps,bytes []byte){
     fmt.Println(bytes)
}

func CloseTheConnection(op easy_server.TcpConnectionOps,bytes []byte){
     time.Sleep(10*time.Second)
     op.Close()
}

func UdpDataHandler(ops easy_server.UdpPacketOps,bytes []byte){
     fmt.Println(bytes)
     ops.SendData(bytes)
}

func main(){
	runtime.GOMAXPROCS(20)
//	file,_ := os.Create("dxp.log")
//	easy_server.SetEasyLogger(os.Stdout,file,file,file)
	server := easy_server.NewServer(5000)
	server.CreateWorkers()
	handlers := easy_server.NewTcpDataHandlers(nil,JustPrint,CloseTheConnection)
	server.AddTcpListener(":4003",handlers)

	server.AddUdpListener(":4003",UdpDataHandler)
	server.PrintServerInfo()

	server.Stop()
}
