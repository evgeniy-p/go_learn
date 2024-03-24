package logic

import (
	"testing"
)

func TestInitPlayer(t *testing.T) {
	t.Run("Valid player name", func(t *testing.T) {
		player, err := InitPlayer("John")
		if err != nil {
			t.Errorf("InitPlayer() returned an error for valid player name: %v", err)
		}
		if player == nil {
			t.Error("InitPlayer() returned nil player")
		}
	})

	t.Run("Empty player name", func(t *testing.T) {
		_, err := InitPlayer("")
		if err == nil {
			t.Error("InitPlayer() did not return an error for empty player name")
		}
	})
}

func TestInitGame(t *testing.T) {
	t.Run("Valid player name and game initialization", func(t *testing.T) {
		player, rooms, err := InitGame("John")
		if err != nil {
			t.Errorf("InitGame() returned an error for valid player name: %v", err)
		}
		if player == nil {
			t.Error("InitGame() returned nil player")
		}
		if rooms == nil {
			t.Error("InitGame() returned nil rooms")
		}
	})

	t.Run("Empty player name", func(t *testing.T) {
		_, _, err := InitGame("")
		if err == nil {
			t.Error("InitGame() did not return an error for empty player name")
		}
	})
}
