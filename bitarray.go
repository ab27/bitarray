package bitarray

// usage:
// func main() {
// 	b := New(40)
// 	b.Set(11)
// 	b.Set(34)
// 	b.Set(2)
// 	b.Set(40)
// 	b.Set()
// 	fmt.Println(b.OnBits())
// 	p(b)
// }
import (
	"fmt"
	"strings"
)

const LENGTH int = 8

var p = fmt.Println

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type BitArray struct {
	array []uint8
	Size  int
}

func New(length int) *BitArray {
	byte_count := length / LENGTH
	if length%LENGTH != 0 {
		byte_count += 1
	}
	return &BitArray{make([]uint8, byte_count), length}
}

func (b *BitArray) String() string {
	result := make([]string, len(b.array))
	for i, _ := range b.array {
		result[i] = fmt.Sprintf("%08b", b.array[len(b.array)-1-i])
	}
	return strings.Join(result, "")
}

func (b *BitArray) Set(index int) {
	if index > b.Size || index <= 0 {
		panic("Error: index out of range")
	}
	array_index := index / LENGTH

	if array_index == len(b.array) {
		array_index -= 1
	}

	b.array[array_index] |= 1 << uint((index-1)%LENGTH)
}

func (b *BitArray) ClearBit(index int) {
	if index > b.Size || index <= 0 {
		panic("Error: index out of range")
	}
	array_index := index / LENGTH

	b.array[array_index] &= uint8(0xff ^ (1 << uint((index-1)%LENGTH)))
}

func (b *BitArray) Get(index int) bool {
	if index > len(b.array)*LENGTH || index <= 0 {
		panic("index out of range")
	}
	array_index := index / LENGTH

	if array_index == len(b.array) {
		array_index -= 1
	}

	return (b.array[array_index])&(1<<uint((index-1)%LENGTH)) == (1 << uint((index-1)%LENGTH))
}

func (b *BitArray) OnBits() []int {
	str := b.String()

	var ons []int

	for i := 0; i < b.Size; i++ {
		//fmt.Println(str, len(str), b.Size)
		char := str[len(str)-1-i]
		//fmt.Println(i, char)
		if string(char) == "1" {
			ons = append(ons, i+1)
		}
	}

	return ons
}
