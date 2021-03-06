package ui

import (
	"log"
	"reflect"
	"sort"
)

const (
	ordLowerStart = 97
	maxOrd        = 122
)

// MenuList is a list menu, with choices presented with a key for each. Options is the list of menu options, Inputs
// specifies which keys map to which options, keys is a list of valid keypresses, Paginated defines if the list should
// be displayed all at once, or in pages, and highestOrd defines teh maximum ASCII keypress available for this menu
type MenuList struct {
	Options    map[int]string
	Inputs     map[rune]int
	keys       []int
	Paginated  bool
	highestOrd int
}

// NewMenuList creates a new MenuList
func NewMenuList(options map[int]string) *MenuList {
	menuList := MenuList{}
	menuList.Options = make(map[int]string)
	menuList.Inputs = make(map[rune]int)
	menuList.highestOrd = ordLowerStart
	menuList.Create(options)

	return &menuList

}

// Create builds a new MenuList, given a mapping of options
func (ml *MenuList) Create(options map[int]string) {
	ml.Options = options

	ordLower := ordLowerStart

	for identifier := range options {
		if ordLower <= maxOrd {
			ml.Inputs[rune(ordLower)] = identifier
			ml.keys = append(ml.keys, ordLower)
			ordLower++
			ml.highestOrd = ordLower
		}
	}
}

// Update takes a list of options, compares it to the existing list of options, and updates the menu if the new list
// is different. This allows for updating a menu without creating a new one (which can mess with item ordering).
// Returns true if the options were updated, false otherwise
func (ml *MenuList) Update(options map[int]string) bool {
	// First things first, see if the updated options is the same as the original. If it is, do nothing
	eq := reflect.DeepEqual(options, ml.Options)

	if eq {
		return false
	}

	// The two are not equal. We need to rectify the items in the updated list with the original. This is a two step
	// process. First, update the inputs. For each input, if the identifier still exists in the new list, do nothing
	// If it does not exist, we'll clear out the identifier value. Next, we'll iterate over the new list, and
	// for each value that is not mapped to a key, we'll map it to one. This can be either an existing key, or
	// a new key that will be added

	// Before we do anything else, make sure we have maps to work with
	if ml.Inputs == nil {
		ml.Inputs = make(map[rune]int)
	}

	if ml.Options == nil {
		ml.Options = make(map[int]string)
	}

	for key, identifier := range ml.Inputs {
		// Check if the keys identifier is still in the updated list
		if _, ok := options[identifier]; !ok {
			// The identifier is no longer present in the updated list, so remove it from the key mapping
			ml.Inputs[key] = -1
		}
	}

	// Now, walk through the updated list, and assign new items to any empty keys. This will fill in any gaps in the
	// menu.
	for identifier := range options {
		// Loop through the inputs, looking for nulled slots (-1 for the identifier), also checking that the current
		// item is not already in the inputs
		placed := false
		if _, ok := ml.Options[identifier]; !ok {
			// This item is not currently in the existing list, and needs a spot in the input map
			for key, keyIdentifier := range ml.Inputs {
				if keyIdentifier == -1 {
					ml.Inputs[key] = identifier
					placed = true
				}
			}

			if !placed {
				// There was no free spot in an existing key, so we'll add a new one, based on the highest ord rune
				// used previously
				if ml.highestOrd <= maxOrd {
					ml.Inputs[rune(ml.highestOrd)] = identifier
					ml.keys = append(ml.keys, ml.highestOrd)
					ml.highestOrd++
					placed = true
				}
			}

			// At this point, the item should have been placed. If it has not been, something has gone wrong
			if !placed {
				log.Print("Failed to place an item in the menu. Max Length likely exceeded.")
			}
		}
	}

	// Ensure we break any pointers here, in case options was passed as one. This can cause an odd bug where the
	// menu knows about new options before they have been passed for update.
	newOptions := make(map[int]string)

	for entity, name := range options {
		newOptions[entity] = name
	}

	// Finally, now that all the new items have been placed, and items that need to be removed have been removed,
	// set the menu options to the updated options list
	ml.Options = newOptions

	return true
}

// Print displays the options for the MenuList, sorted by the rune chosen to represent it. yOffset is the number of rows
// to skip before printing, and xOffset, similarly, is the number of columns to skip before printing
func (ml *MenuList) Print(height, width, xOffset, yOffset int) {
	lineStart := yOffset

	// Sort the index slice, this will allow for guaranteed printing order of the two data maps
	sort.Ints(ml.keys)

	for _, keyRune := range ml.keys {
		input := ml.Inputs[rune(keyRune)]
		PrintText(xOffset, lineStart, "("+string(keyRune)+")"+ml.Options[input], "", "", 0, 0)
		lineStart++
	}
}
