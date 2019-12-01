package playgroundmain

import (
	"fmt"
	"testing"
)

type wrapTest struct {
	ID     int      `json:"id"`
	Images []string `json:"image"`
}

type wrapTest2 struct {
	ID     int       `json:"id"`
	Images []*string `json:"image"`
}

func TestSth(t *testing.T) {
	originImages := []string{
		"a",
		"b",
		"c",
	}
	originImagesCopy := make([]string, len(originImages)) //make([]string,0,len(originImages)) will copy fail
	copy(originImagesCopy, originImages)

	// todo 使用range循环，无法改变slice
	//for _, value := range originImages {
	//	value = value+"1"
	//}
	for i := 0; i < len(originImages); i++ {
		fmt.Println(originImages[i])
		originImages[i] = originImages[i] + "1"
	}
	fmt.Println("originImages:", originImages)
	fmt.Println("originImagesCopy:", originImagesCopy)

	fmt.Println("===========")

	var originImages2 []*string //:= make([]*string,0)
	//for _, value := range originImages {
	//	fmt.Println(value)
	//	originImages2 = append(originImages2,&value) // 为什么上面打印的value值不同，但指针相同，且都指向最后一个元素c？和range循环有关，使用下面的普通for即可正常遍历指针和赋值
	//}
	for i := 0; i < len(originImages); i++ {
		fmt.Println(originImages[i])
		originImages2 = append(originImages2, &originImages[i])
	}
	//todo 推荐：除了从channel不断拿值使用range 循环外，其他地方都使用普通for，即使只有拿值，然后也没有[]*string的场景

	fmt.Println("originImages:", originImages)
	fmt.Println("originImages2:", originImages2)
	for _, value := range originImages2 {
		fmt.Println("originImages2 v:", *value)
	}

	fmt.Println("change originImages2 =======")

	for i := 0; i < len(originImages2); i++ {
		v := *originImages2[i] + "2"
		originImages2[i] = &v
	}

	fmt.Println("originImages:", originImages)
	for _, value := range originImages2 {
		fmt.Println("originImages2 v:", *value) //对于[]string来说，可以直接用普通for循环来改变slice的内容，不需要使用[]*string，虽然这里也可以达到修改的目的
	}
	//originImages2data,_ := json.Marshal(originImages2)
	//fmt.Printf("originImages2:%+v\n",originImages2data)
	fmt.Println("===========")

	originImages3 := make([]*string, 0)
	s1 := "s1"
	originImages3 = append(originImages3, &s1)
	for _, value := range originImagesCopy {
		originImages3 = append(originImages3, &value)
	}
	fmt.Println("originImages3:", originImages3)

	//todo 将[]*string序列化成json，看看结果

}
