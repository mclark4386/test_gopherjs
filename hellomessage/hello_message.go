package hellomessage

import (
	"fmt"

	"myitcv.io/react"
)

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
	food  string
}

//Step 4
func HelloMessage(p HelloMessageProps) *HelloMessageElem {
	return buildHelloMessageElem(p)
}

func (r HelloMessageDef) Render() react.Element {
	name := r.Props().Name
	count := r.State().count
	food := r.State().food
	return react.Div(nil,
		react.P(nil,
			react.S(fmt.Sprintf("Hola %s:%d. Says here your favorite food is %s.", name, count, food)),
		),
		react.Input(&react.InputProps{
			OnChange:    foodChange{r},
			Type:        "text",
			Placeholder: "What is your favorite food?",
			Value:       food,
		}),
		react.Button(
			&react.ButtonProps{
				OnClick: countClick{r},
			},
			react.S("Tick Count"),
		),
	)
}

func (r HelloMessageDef) GetInitialState() {
	r.SetState(HelloMessageState{count: 0})
}

type countClick struct{ HelloMessageDef }

func (c countClick) OnClick(e *react.SyntheticMouseEvent) {
	s := c.State()
	s.count++
	c.SetState(s)
}

type foodChange struct{ HelloMessageDef }

func (f foodChange) OnChange(e *react.SyntheticEvent) {
	s := f.State()
	s.food = e.Target().NodeValue()
	f.SetState(s)
}
