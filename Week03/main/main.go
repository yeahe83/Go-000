/*
Week03 作业题目：

1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

以上作业，要求提交到 Github 上面，Week03 作业提交地址：
https://github.com/Go-000/Go-000/issues/69

请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。
作业提交截止时间 12 月 9 日（周三）23:59 前。

Github 使用教程： https://u.geekbang.org/lesson/51?article=294701

学号查询方式：PC 端登录 time.geekbang.org, 点击右上角头像进入我的教室，左侧头像下方 G 开头的为学号

*/

package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		srv := &http.Server{Addr: ":8080"}

		// https://golang.google.cn/pkg/net/http/#example_Server_Shutdown
		// https://golang.google.cn/pkg/net/http/#example_ListenAndServe
		helloHandler := func(w http.ResponseWriter, req *http.Request) {

			if req.URL.Query().Get("q") == "y" { // http://localhost:8080/hello?q=y
				srv.Shutdown(ctx)
			}

			io.WriteString(w, "Hello, world!\n")
		}

		http.HandleFunc("/hello", helloHandler)

		go func() {
			<-ctx.Done() // cancel
			srv.Shutdown(ctx)
		}()

		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			cancel() // to cancel signal
			return err
		}

		return nil

	})

	g.Go(func() error {

		// https://golang.google.cn/pkg/os/signal/#example_Notify
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT)

		select {

		case <-ctx.Done(): // cancel
			fmt.Println("signal cancelled")
			return errors.New("signal cancelled")
		case s := <-c:
			signal.Reset(syscall.SIGINT)
			fmt.Println("Got signal:", s)
			cancel() // to cancel http.server
			return errors.New(fmt.Sprint("Got signal:", s))

		}

	})

	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

/*

浏览器：http://localhost:8080/hello?q=y

PS C:\Users\yeahe83\Documents\Dev\demo\Go-000\Week03\main> go run main.go
http: Server closed
signal cancelled
2020/12/10 02:45:51 http: Server closed
exit status 1


键盘：Ctrl+Z

PS C:\Users\yeahe83\Documents\Dev\demo\Go-000\Week03\main> go run main.go
Got signal: interrupt
http: Server closed
2020/12/10 02:46:03 Got signal:interrupt
exit status 1



*/
