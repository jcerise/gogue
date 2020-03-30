package ui

import (
	"github.com/jcerise/gogue"
)

type MessageLog struct {
	messages  []string
	MaxLength int
}

func NewMessageLog(maxLength int) *MessageLog {
	messageLog := MessageLog{MaxLength: maxLength}
	messageLog.messages = make([]string, messageLog.MaxLength)
	return &messageLog
}

func (ml *MessageLog) SendMessage(message string) {
	// Prepend the message onto the messageLog slice
	if len(ml.messages) >= ml.MaxLength {
		// Throw away any messages that exceed our total queue size
		ml.messages = ml.messages[:len(ml.messages)-1]
	}
	ml.messages = append([]string{message}, ml.messages...)
}

func (ml *MessageLog) PrintMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, displayNum int) {
	// Print the latest five messages from the messageLog. These will be printed in reverse order (newest at the top),
	// to make it appear they are scrolling down the screen
	clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, 1)

	toShow := 0

	if len(ml.messages) <= displayNum {
		// Just loop through the messageLog, printing them in reverse order
		toShow = len(ml.messages)
	} else {
		// If we have more than {displayNum} messages stored, just show the {displayNum} most recent
		toShow = displayNum
	}

	for i := toShow; i > 0; i-- {
		gogue.PrintText(1, (viewAreaY-1)+i, ml.messages[i-1], "white", "", 1)
	}
}

// ClearMessage clears the defined message area, starting at viewAreaX and Y, and ending at the width and height of
// the message area
func clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer int) {
	gogue.ClearArea(viewAreaX, viewAreaY, windowSizeX, windowSizeY-viewAreaY, 1)
}

// PrintToMessageArea clears the message area, and print a single message at the top
func PrintToMessageArea(message string, viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer int) {
	clearMessages(viewAreaX, viewAreaY, windowSizeX, windowSizeY, layer)
	gogue.PrintText(1, viewAreaY, message, "white", "", 1)
}
