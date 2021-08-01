package storybuilder

func GetRoom(index int) string {
	rooms := [10]string{
		"ESCAPED! You have found your way out! Congratulations!",
		"DEATH! The room contains an ogre, unfortunately, it'll kill you regardless of what you do.",
		"DEATH! You encounter a trap, instant death!",
		"You find yourself in a pitch black room, not sure of your surroundings.",
		"A tornado of wind pulls you in every direction, but is otherwise harmless.",
		"A room of ghosts, them dancing around you.",
		"Skeletons are all over the ground, lifeless.",
		"An empty room.",
		"A room full of goblins! You're able to fend them off but barely!",
		"You enter the room, and hear a whisper saying 'Turn back'.",
	}
	return rooms[index]
}

func GetMaxRooms() int {
	return 10
}