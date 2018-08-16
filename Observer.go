package main

import "fmt"

func main() {
	subject:=Subject{}
	subject.state=1
	subject.observers=[]Observer{Observer{"小A"},Observer{"小B"}}
	subject.update()
}

type ObserverLintener interface {
	Notify(state int)
}

type SubjectLintener interface {
	setState(state int)
	update()
}

type Observer struct {
	name string
}

type Subject struct {
	state int
	observers []Observer
}

func (subject *Subject)setState(state int)  {
	subject.state=state;
}
func (subject *Subject)update()  {
	for _,v:=range subject.observers{
		v.Notify(subject.state)
	}
}
func (observer *Observer)Notify(state int)  {
	if state ==1{
		fmt.Println("观察到observer状态是1;")
	}
}


