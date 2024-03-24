package logic

import (
	"errors"
)

type player struct {
	RoomNow        string
	ActionLast     string
	InBackpack     []string
	BackPack       bool
	DoorState      bool
	AllowedActions []string
}

func InitPlayer(PName string) (*player, error) {
	if PName == "" {
		return nil, errors.New("player name cannot be empty")
	}
	p := player{
		ActionLast:     PName,
		RoomNow:        "кухня",
		AllowedActions: []string{"идти", "осмотреться", "надеть", "взять", "применить"},
		BackPack:       false,
		DoorState:      false,
	}
	return &p, nil
}

type room struct {
	RoomItems       []string
	NextAllowedRoom []string
}

func InitGame(text string) (player *player, rooms map[string]room, err error) {
	rooms = map[string]room{}
	p, err := InitPlayer(text)
	if err != nil {
		return p, rooms, errors.New("player name cannot be empty")
	}
	for _, name := range []string{"кухня", "коридор", "комната", "улица"} {
		rooms[name] = room{}
		switch name {
		case "кухня":
			rooms[name] = room{RoomItems: []string{"чай"}, NextAllowedRoom: []string{"коридор"}}
		case "коридор":
			rooms[name] = room{NextAllowedRoom: []string{"комната", "кухня", "улица"}}
		case "комната":
			rooms[name] = room{RoomItems: []string{"ключи", "конспекты"}, NextAllowedRoom: []string{"коридор"}}
		case "улица":
			rooms[name] = room{RoomItems: []string{"весна"}, NextAllowedRoom: []string{"коридор"}}
		}
	}
	return p, rooms, err
}
