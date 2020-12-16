/*

## Week04 作业题目：

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

️ 以上作业，要求提交到 Github 上面，Week04 作业提交地址：
https://github.com/Go-000/Go-000/issues/76

请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。

️Github 使用教程：https://u.geekbang.org/lesson/51?article=294701

学号查询方式：
PC 端登录 time.geekbang.org,点击右上角头像进入【我的教室】，左侧头像下方 G 开头的为学号。

## Reference

https://blog.csdn.net/uisoul/article/details/108776073

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

	"example.com/myapp/internal/user/biz"
	"example.com/myapp/internal/user/data"
	"github.com/google/wire"
	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		srv := &http.Server{Addr: ":8080"}

		userHandler := func(w http.ResponseWriter, req *http.Request) {

			mgr, err := initUserManager()
			if err != nil {
				log.Fatalln(err.Error())
			}

			userID := req.URL.Query().Get("id")
			user, err := mgr.GetUser(userID)
			if err != nil {
				io.WriteString(w, err.Error())
			} else {
				io.WriteString(w, user.Name)
			}

		}

		http.HandleFunc("/user", userHandler)

		go func() { // 监听ctx的关闭
			<-ctx.Done() // cancel
			srv.Shutdown(context.TODO())
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
		c := make(chan os.Signal, 4)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

		select {

		case <-ctx.Done(): // cancel
			fmt.Println("signal cancelled")
			return errors.New("signal cancelled")
		case s := <-c:
			signal.Reset(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
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

// NewUserManager provider to UserManager
func NewUserManager(dao *data.UserDao) (*biz.UserManager, error) {
	mgr := biz.UserManager{Dao: dao}
	return &mgr, nil
}

// NewUserDao provider to UserDao
func NewUserDao() (*data.UserDao, error) {
	dao := data.UserDao{}
	return &dao, nil
}

func initUserManager() (*biz.UserManager, error) {
	wire.Build(NewUserManager, NewUserDao)
	return nil, nil
}

/*

PS C:\Users\yeahe83\Documents\dev\demo\Go-000\Week04\myapp\cmd\user-service> wire
wire: example.com/myapp/cmd/user-service: wrote C:\Users\yeahe83\Documents\dev\demo\Go-000\Week04\myapp\cmd\user-service\wire_gen.go
PS C:\Users\yeahe83\Documents\dev\demo\Go-000\Week04\myapp\cmd\user-service> go run .\wire_gen.go


http://localhost:8080/user?id=1
yeahe83

*/
