package simple

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outfile()
}

type CookMenu struct {
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

func (CookMenu) outfile() {
	fmt.Println("关火")
}

func (CookMenu) cooke() {

}

func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfile()
}

type Xhs struct {
	CookMenu
}

func (*Xhs) cooke() {
	fmt.Println("炒西红柿")
}

type Egg struct {
	CookMenu
}

func (*Egg) cooke() {
	fmt.Println("炒鸡蛋")
}
