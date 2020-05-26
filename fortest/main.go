package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main(){
	// 门店坐标[121.449188,31.224045]
	// 左上角坐标
	x0,y0:=121.434188,31.209045  //
	dt := 0.0009375
	result := make([]Point,0)
	for i:=0;i<=32; i++{
		for j:=0;j<=32; j++{
			p := Point{x0+dt*float64(i),y0+dt*float64(j)}
			result = append(result, p)

		}
	}
	fmt.Println(len(result))
	for _,p:=range result {
		fmt.Printf("[%.6f,%.6f],",p.x,p.y)
	}

}