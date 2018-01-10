package hellomessage

import "myitcv.io/react"

//The following line is required for go generate to work:
//go:generate reactGen

//Step 1
type HelloMessageDef struct { //this type MUST have the 'Def' suffix, because 'Definition'
	react.ComponentDef //must have (at least) an anon embed of this, but can have other fields
}

//Step 2
type HelloMessageProps struct { //[Optional] contains props, NOTE *Props suffix
	Name string
}

//Step 3
type HelloMessageState struct { //[Optional] contains state, NOTE *State suffix
	count int
}
