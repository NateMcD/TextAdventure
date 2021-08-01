package storybuilder

type StoryNode struct {
	text	string
	Choices *Choices
}

type Choices struct {
	cmd				string
	description		string
	nextNode		*StoryNode
	nextChoice		*Choices
}

func (node *StoryNode) addChoice(cmd string, description string, nextNode *StoryNode) {
	choice := &Choices{cmd, description, nextNode, nil}
	if node.Choices == nil {
		node.Choices = choice
	} else {
		currentChoice := node.Choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}