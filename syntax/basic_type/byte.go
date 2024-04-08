package main

func Byte() {
	//var a byte = 'a'
	var str string = "This is string"
	var bs []byte = []byte(str)
	bs[0] = 'A'
	println(str)
	println(bs)
}
