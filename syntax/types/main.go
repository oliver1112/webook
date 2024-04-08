package main

//func main() {
//	//NewUser()
//	//UseFish()
//	//UseListV1()
//	Components()
//}

func UseListV1() {
	l := &LinkedList{}
	err := l.Add(1, 123)
	if err != nil {
		println("error")
	}
	l.Add(1, "123")
	l.Add(1, nil)
}
