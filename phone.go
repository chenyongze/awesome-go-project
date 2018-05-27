package main
import (
	"fmt"
	"github.com/zheng-ji/gophone"
)

func main() {

	phones := [...]string{"15222309717","18500466899","13621086889","18500090688","18583799357","18601047629","13901037039","13916007116","17726777777","13851588860","18611006666","13451789287","13910212905"}


	fmt.Println(phones)

	for _,x:= range phones {
		pr, err := gophone.Find(x)
		if err == nil {
			//fmt.Println(pr)
			fmt.Printf("mobile %s ==>地区 %s \n",x,pr.Province)
		}
		//fmt.Printf("mobile %s ==>地区 %s \n",x,pr.Province)
		//fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}

	//for _,value := rang phones {
	//		pr, err := gophone.Find(value)
	//		if err == nil {
	//			fmt.Println(pr)
	//		}
	//}




	//pr, err = gophone.Find("15920544678")
	//if err == nil {
	//	// 也可以单独获取该号码各个属性
	//	//fmt.Println(pr.PhoneNum)
	//	fmt.Println(pr.Province)
	//	//fmt.Println(pr.AreaZone)
	//	fmt.Println(pr.City)
	//	fmt.Println(pr.ZipCode)
	//}
}
