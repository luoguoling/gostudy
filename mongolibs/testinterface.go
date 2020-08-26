package mongolibs

import "fmt"

type Animal interface {
	Run()
	Eat()
}
type Cat struct {
	name string
	sex	bool
}

func (c Cat) Run()  {
	fmt.Println(c.name,"跑")


}
func (c Cat) Eat(){
	fmt.Println(c.name ,"吃!!!")
}
func main()  {
	var a Animal
	a = Cat{
		name: "blue",
		sex:  false,
	}
	a.Run()
	a.Eat()

}
