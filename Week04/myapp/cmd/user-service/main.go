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
	"fmt"
	"log"

	"example.com/myapp/internal/biz"
	"example.com/myapp/internal/data"
	"github.com/google/wire"
)

func main() {
	mgr, err := initUserManager()
	if err != nil {
		log.Fatalln(err.Error())
	}
	user, err := mgr.GetUser("1")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("user.Name = ", user.Name)
}

// provider
func NewUserManager(dao *data.UserDao) (*biz.UserManager, error) {
	mgr := biz.UserManager{Dao: dao}
	return &mgr, nil
}

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
user.Name =  yeahe83

*/
