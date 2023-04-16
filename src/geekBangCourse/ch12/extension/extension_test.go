package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	Pet
}

func (d *Dog) speak() {
	fmt.Println("wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")

}
