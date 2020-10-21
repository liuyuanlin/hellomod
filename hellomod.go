package hellomod

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("hello,%s,我是依赖包hellomod V1.0.1", name)
}
