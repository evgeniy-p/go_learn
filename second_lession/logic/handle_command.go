package logic

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func getRemainingItems(value string, items []string) (returnSlice []string) {
	var newSlice []string

	for _, item := range items {
		if item != value {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func lookUp(player *player, rooms map[string]room, value string) (string, error) {
	var textResponse string
	switch player.RoomNow {
	case "кухня":
		if !player.BackPack {
			textResponse = "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		} else {
			textResponse = "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		player.ActionLast = value
		return textResponse, nil
	case "комната":
		if len(rooms[player.RoomNow].RoomItems) < 1 {
			return "пустая комната. можно пройти - коридор", nil
		}
		if !player.BackPack {
			textResponse = fmt.Sprintf("на столе: %s, на стуле: рюкзак. можно пройти - коридор", strings.Join(rooms[player.RoomNow].RoomItems, `, `))
		} else {
			textResponse = fmt.Sprintf("на столе: %s. можно пройти - коридор", strings.Join(rooms[player.RoomNow].RoomItems, `, `))
		}
		return textResponse, nil
	default:
		return "ничего интересного", nil
	}

}

func moveTo(player *player, rooms map[string]room, value string) (string, error) {
	containsMoveTo := slices.Contains(rooms[player.RoomNow].NextAllowedRoom, value)
	if !containsMoveTo {
		myAnswer := fmt.Sprintf("нет пути в %s", value)
		return myAnswer, nil
	}
	if value == "коридор" {
		player.RoomNow = value
		return "ничего интересного. можно пройти - комната, кухня, улица", nil
	}
	if value == "комната" {
		player.RoomNow = value
		return "ты в своей комнате. можно пройти - коридор", nil
	}
	if value == "улица" {
		if !player.DoorState {
			return "дверь закрыта", nil
		}
		containsAction := slices.Contains(player.InBackpack, "ключи")
		if !containsAction {
			return "нет ключей", nil
		} else {
			player.RoomNow = value
			return "на улице весна. можно пройти - коридор", nil
		}
	}
	if value == "кухня" {
		player.RoomNow = value
		return "кухня, ничего интересного. можно пройти - коридор", nil
	} else {
		return "test", errors.New("Not checked")
	}
}

func apply(player *player, value string, lastAction string) (string, error) {

	if value != "дверь" {
		containsItem := slices.Contains(player.InBackpack, value)
		if !containsItem && lastAction != "ключи" {
			myAnswer := fmt.Sprintf("нет предмета в инвентаре - %s", lastAction)
			return myAnswer, nil
		} else {
			return "не к чему применить", nil
		}
		return "не к чему применить", nil
	} else {
		if lastAction == "ключи" && slices.Contains(player.InBackpack, "ключи") {
			player.DoorState = true
			return "дверь открыта", nil
		} else {
			if slices.Contains(player.InBackpack, lastAction) {
				return "не к чему применить", nil
			} else {
				myAnswer := fmt.Sprintf("нет предмета в инвентаре - %s", lastAction)
				return myAnswer, nil
			}
		}

	}
}

func getIn(player *player) (string, error) {
	if !player.BackPack {
		player.BackPack = true
		return "вы надели: рюкзак", nil
	} else {
		return "такого предмета нет!", nil
	}

}

func getItem(player *player, rooms map[string]room, value string) (string, error) {
	if !player.BackPack {
		return "некуда класть", nil
	}
	containsItem := slices.Contains(rooms[player.RoomNow].RoomItems, value)
	if !containsItem {
		return "нет такого", nil
	}
	player.InBackpack = append(player.InBackpack, value)
	newSlice := getRemainingItems(value, rooms[player.RoomNow].RoomItems)
	rooms[player.RoomNow] = room{RoomItems: newSlice, NextAllowedRoom: rooms[player.RoomNow].NextAllowedRoom}
	toBackPack := fmt.Sprintf("предмет добавлен в инвентарь: %s", value)
	return toBackPack, nil

}

func HandleCommand(command string, player *player, rooms map[string]room) (answer string, err error) {
	var actions []string
	var textResponse string
	var lastAction string
	haveLastAction := false
	textResponse = "test"
	actions = strings.Split(command, " ")
	for index, value := range actions {
		if index == 0 { // first element = action
			containsAction := slices.Contains(player.AllowedActions, value)
			if !containsAction {
				myAnswer := fmt.Sprintf("неизвестная команда %s", value)
				return myAnswer, nil
			}
			switch value {
			case "осмотреться":
				if len(actions) > 1 {
					fmt.Println("в случае действия \"осмотреться\" достаточно указать только действие, " +
						"все стальное будет проигнорировано")
					break
				} else {
					return lookUp(player, rooms, value)
				}
			case "применить", "идти":
				if len(actions) < 2 {
					return "требуется больше аргументов", nil
				}
			}
			player.ActionLast = value
		} else {
			switch player.ActionLast {
			case "взять":
				return getItem(player, rooms, value)
			case "надеть":
				return getIn(player)
			case "идти":
				return moveTo(player, rooms, value)
			case "применить":
				if haveLastAction {
					return apply(player, value, lastAction)
				} else {
					haveLastAction = true
					lastAction = value
				}
			}
		}
	}

	return textResponse, nil
}
