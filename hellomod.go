package hellomod

import (
	"fmt"
)

func SayHello(name, str string) string {
	return fmt.Sprintf("hello,%s,%s,我是依赖包hellomod V2.0.3", name, str)
}
