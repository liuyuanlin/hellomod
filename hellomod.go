/*的风俗的风俗第三方地方三*/
package hellomod

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("hello,%s,我是依赖包hellomod", name)
}
