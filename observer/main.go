package main

import "github.com/hzhyvinskyi/ood-patterns/observer/cases"

func main() {
	subject := cases.NewCase("Tmp Case", "OPEN")
	observerOne := &cases.CaseOwner{1}
	observerTwo := &cases.CaseOwner{2}
	subject.register(observerOne)
	subject.register(observerTwo)
	subject.notifyAll()
}
