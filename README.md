# SliceOP
This repository hold the slice operation package 

- After each operation sliceOP did ,it creates new slice with current slices capacity and length and then returns it to user.

- Aim of this project to use garbage collectors effectively and use memory effectively

## How to install

Get the sliceOP package by 

``go get github.com/ASaidOguz/sliceOP``


After the installation you can import the package  in your go file and use it as slice operation package.


SliceOP package uses generic function that accepts interface (any) so you can operator inside any kind of slice. For certain slice operation which you need interact more than one slice both slice should be same type unless function will give compiler error.

---

Some of the usage examples as shown below;

```
package main

import (
	"fmt"

	SliceOp "github.com/ASaidOguz/sliceOP"
)

func main() {

	x := make([]int, 0, 0)
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	fmt.Println("initial state of x : "x)

newSlice := SliceOp.DeleteOrder(x, 0)
	fmt.Println("The new slice after the op: ",
		newSlice,
		"Length: ", len(newSlice),
		" Capacity: ", cap(newSlice))

newSlice1 := SliceOp.DeleteOrder(newSlice, 0)
	fmt.Println("The new slice after the op: ",
		newSlice1,
		"Length: ", len(newSlice1),
		" Capacity: ", cap(newSlice1))

}
```

 Any suggestion or any mistake i made in code please feel free to inform me so i can correct them .

 ---






































































