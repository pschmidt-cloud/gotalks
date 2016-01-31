package main

import "fmt"

func typeCheckThis(a int) {
	if (a > 0) {
		fmt.Println("Hi")
	} else {
		fmt.Println("3" + 5)  // will throw an error at compile time
	}
}

func main() {
	typeCheckThis(2)
}
/* whereas in Python for example
def typeCheckThis(a):
    if a > 0:
        print 'Hi'
    else:
        print 5 + '3'
    You would only see the error at runtime with typeCheckThis(-1)  not typeCheckThis(2)
 */