package storybuilder

type storyNode struct {
	text	string
	choices *choices
}

type choices struct {
	cmd				string
	description		string
	nextNode		*storyNode
	nextChoice		*choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}