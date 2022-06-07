package iterator_pattern

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(T *testing.T) {
	dinerMenu := NewDinerMenu()
	dinerMenu.addItem("name1", "description1", true, 10)
	dinerMenu.addItem("name2", "description2", true, 11)
	dinerMenu.addItem("name3", "description3", false, 13)
	iterator := dinerMenu.createIterator()
	for iterator.hasNext() {
		item := iterator.next().(*MenuItem)
		str := fmt.Sprintf("name:%s,decription:%s,isVegetarian:%v,price:%f",
			item.getName(), item.getDescription(), item.isVegetarian(), item.getPrice())
		fmt.Println(str)
	}

}
