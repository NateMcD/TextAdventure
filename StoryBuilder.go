package storybuilder

import "math/rand"

func GetRoom() string {
	rooms := [4]string{
		"ESCAPED! You have found your way out! Congratulations!",
		"DEATH! The room contains an ogre, unfortunately, it'll kill you regardless of what you do.",
		"DEATH! You encounter a trap, instant death!",
		"You find yourself in a dark room, not sure of your surroundings.",
	}
	randomIndex := rand.Intn(len(rooms))
	chosenRoom := rooms[randomIndex]
	return chosenRoom
}