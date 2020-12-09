package main

import (
	"database/sql"
	"errors"
	"fmt"

	"./service"
)

func main() {

	user, err := service.GetOne(1)
	if errors.Is(err, sql.ErrNoRows) { // for test errors.Is unwrap
		fmt.Printf("row not exists")
		return
	} else if err != nil { // log
		fmt.Printf("\n%+v\n", err)
		return
	}

	fmt.Println(user)
}

/*


## 作业题目

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

以上作业，要求提交到 Github 上面，Week02 作业提交地址：https://github.com/Go-000/Go-000/issues/8

请务必按照示例格式进行提交，不要复制其他同学的格式，以免格式错误无法抓取作业。

作业提交截止时间 12 月 2 日（下周三）23:59 前。

Github 使用教程： https://u.geekbang.org/lesson/51?article=294701

学号查询方式：PC 端登录 time.geekbang.org, 点击右上角头像进入我的教室，左侧头像下方 G 开头的为学号


## Reference:

- https://www.bilibili.com/video/BV1dZ4y1577v

*/
