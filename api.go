package cp

// #cgo CFLAGS: -std=c99 -Ichipmunk
// #cgo LDFLAGS: -lm
import "C"
import "fmt"

func (ptr SwigcptrStruct_SS_cpBody) SetPosition(x, y float64) {
	BodySetPosition(ptr, V(x, y))
	fmt.Println("set pos", x, y)
}
