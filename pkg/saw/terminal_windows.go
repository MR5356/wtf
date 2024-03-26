//go:build windows

package saw

import (
	"embed"
	"fmt"
	"github.com/MR5356/wtf/pkg/utils/iputil"
	"github.com/creack/pty"
	"github.com/olahol/melody"
	"github.com/sirupsen/logrus"
	"golang.org/x/term"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

//go:embed index.html index.css index.js
var content embed.FS

var stopFunctions = make([]func(), 0)

func registerStopFunction(f ...func()) {
	stopFunctions = append(stopFunctions, f...)
}

func doStop() {
	for _, f := range stopFunctions {
		f()
	}
}

func Terminal(cmd string, port int) error {
	defer doStop()
	if runtime.GOOS == "windows" {
		return fmt.Errorf("windows is not supported")
	}
	ips := iputil.GetAvailableIP()

	readMsg := "Readonly at:\n"
	ctrlMsg := "Remote access at:\n"

	for _, ip := range ips {
		readMsg += fmt.Sprintf("\thttp://%s:%d\n", ip, port)
		ctrlMsg += fmt.Sprintf("\thttp://%s:%d?rule=ctrl\n", ip, port)
	}

	fmt.Println(readMsg)
	fmt.Println(ctrlMsg)

	fmt.Println("Make sure your browser size is larger than the terminal size!!!")

	if len(cmd) == 0 {
		cmd = "sh"
	}

	// 创建一个命令对象
	c := exec.Command(cmd)

	// 使用pty运行命令
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}

	// 保证在退出时关闭pty
	registerStopFunction(func() {
		_ = ptmx.Close()
	})

	// 创建一个只读管道
	mRead := melody.New()
	teeReaderIn := io.TeeReader(ptmx, os.Stdin)

	// 创建一个读写管道
	mCtrl := melody.New()
	teeReaderOut := io.TeeReader(os.Stdout, ptmx)

	// 设置StdIn为Raw模式
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	// 保证在退出时恢复原始模式
	registerStopFunction(func() {
		_ = term.Restore(int(os.Stdin.Fd()), oldState)
	})

	// 将输入流发送给websocket
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := teeReaderIn.Read(buf)
			if err != nil {
				if err == io.EOF {
					doStop()
					os.Exit(0)
				}
				return
			}
			_ = mRead.Broadcast(buf[:n])
			_ = mCtrl.Broadcast(buf[:n])
		}
	}()

	// 将输出流发送给websocket
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := teeReaderOut.Read(buf)
			if err != nil {
				if err == io.EOF {
					doStop()
					os.Exit(0)
				}
				return
			}
			_ = mCtrl.Broadcast(buf[:n])
		}
	}()

	// 可操作的websocket
	mCtrl.HandleMessage(func(session *melody.Session, msg []byte) {
		_, err := ptmx.Write(msg)
		if err != nil {
			logrus.Errorf("error writing to pty: %s", err)
		}
	})

	// 网页监听
	fs := http.FileServer(http.FS(content))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/see", func(writer http.ResponseWriter, request *http.Request) {
		err := mRead.HandleRequest(writer, request)
		if err != nil {
			logrus.Errorf("error handling request: %s", err)
		}
	})

	http.HandleFunc("/ctrl", func(writer http.ResponseWriter, request *http.Request) {
		err := mCtrl.HandleRequest(writer, request)
		if err != nil {
			logrus.Errorf("error handling request: %s", err)
		}
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
