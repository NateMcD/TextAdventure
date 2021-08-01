package storybuilder

import "math/rand"

func GetRoom() string {
	rooms := [4]string{"Escape", "Ogre", "Trap", "DarkRoom"}
	randomIndex := rand.Intn(len(rooms))
	chosenRoom := rooms[randomIndex]
	return chosenRoom
}