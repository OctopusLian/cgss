package cg

import (
	"fmt"
)

type Player struct {
	Name  string        "name"  //玩家昵称
	Level int           "level" //等级
	Exp   int           "exp"   //经验值
	Room  int           "room"  //房间号
	mq    chan *Message // 等待收取的消息
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}
	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message:", msg.Content)
		}
	}(player)
	return player
}
