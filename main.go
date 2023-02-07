package werewolf

import (
	"fmt"
	"time"
)

type 事件接口 interface {
	A游戏开始() error
	A身份发牌(玩家 string, 角色 string) error
	// 天黑(玩家 string) error
	// 天亮(玩家 string) error
	// 死亡(玩家 string) error
	// 发表遗言(玩家 string, 遗言 func(文本 string) error, 结束 func() error) error
	// 自由讨论(玩家 string, 讨论 func(文本 string) error, 结束 func() error) error
	// 女巫用药(玩家 string, 死亡玩家 []string, 用药 func(使用 bool, 药 bool, 对象 string) error) error
	// 预言家预言(玩家 string, 预言 func(对象 string) error) error
	// 守卫守卫(玩家 string, 守卫 func(对象 string) error) error
	// 狼人讨论(玩家 string, 讨论 func(文本 string) error, 结束 func() error) error
	// 狼人刀人(玩家 string, 刀人 func(对象 string) error) error
}

type 狼人杀服务端 struct {
	客户端 *狼人杀客户端

	玩家列表 map[string]bool
}

func A新建游戏(客户端 *狼人杀客户端) (加入 func(玩家 string) error,
	退出 func(玩家 string) error, 结束 func() error, 错误 error) {
	服务端 := &狼人杀服务端{
		客户端: 客户端,
	}
	fmt.Println(服务端)
	加入通道 := make(chan string, 1)
	退出通道 := make(chan string, 1)
	结束通道 := make(chan bool, 1)
	go func() {
		select {
		case 玩家 := <-加入通道:
			服务端.玩家列表[玩家] = true
			if len(服务端.玩家列表) >= 客户端.人数 {
				客户端.A游戏开始()
			}
			<-time.Tick(time.Second * 20)
			break
		case 玩家 := <-退出通道:
			delete(服务端.玩家列表, 玩家)
		case <-结束通道:
			return
		}
		for 玩家 := range 服务端.玩家列表 {
			客户端.A身份发牌(玩家, "")
		}
	}()
	return func(玩家 string) error {
			加入通道 <- 玩家
			return nil
		}, func(玩家 string) error {
			退出通道 <- 玩家
			return nil
		}, func() error {
			结束通道 <- true
			return nil
		}, nil
}

type 狼人杀客户端 struct {
	房间 string
	人数 int
}

func (客户端 *狼人杀客户端) A游戏开始() error {
	// do thx
	return nil
}

func (客户端 *狼人杀客户端) A身份发牌(玩家 string, 角色 string) error {
	// do thx
	return nil
}

func main() {
	加入, _, _, 错误 := A新建游戏(&狼人杀客户端{
		房间: "888888",
		人数: 7,
	})
	if 错误 != nil {
		panic(错误)
	}
	加入("123456789")
	加入("123456789")
}
