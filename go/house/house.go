package house

import "fmt"

type TerseComponent struct {
	subject, verb string
}

var TerseComponents = []TerseComponent{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func Verse(v int) string {
	// test cases imply v numbering starts at 1 instead of 0
	vv := v - 1
	outStr := ""
	for i := vv; i >= 0; i-- {
		if i == vv {
			outStr += fmt.Sprintf("This is the %s", TerseComponents[i].subject)
		}
		if i != 0 {
			outStr += fmt.Sprintf("\nthat %s the %s", TerseComponents[i].verb, TerseComponents[i-1].subject)
		}
	}
	return outStr
}

func Song() string {
	outStr := ""
	for i := range TerseComponents {
		if i != 0 {
			outStr += "\n\n"
		}
		outStr += Verse(i + 1) // exercise expects indexing to start at 1
	}
	return outStr
}
