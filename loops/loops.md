মনে করুন, আপনাকে একটা প্রোগ্রাম লিখতে বোলা হল যেটা রান করলে `Hello World!` প্রিন্ট করবেন। আপনি নিশ্চয় এরকম কোড করে ফেলবেনঃ
```go
package main
import "fmt"
func main() {
   fmt.Print("Hello World!!\n")
}
```
আচ্ছা, এবার যদি বোলা হয় ১০ বার `Hello World!` প্রিন্ট করতে তাহলে কি করবেন?
```go
package main
import "fmt"
func main() {
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
   fmt.Print("Hello World!!\n")
}
```
এবার যদি ১০০০০ বার প্রিন্ট করতে বলা হয়, তাহলে?
১০০০০ বার `fmt.Print(“Hello World!\n”)` লিখবেন?
<div style="display: flex; justify-content: center; margin-bottom:20px;">
<img src="./We_Dont_Do_That_Here.jpg" alt="Flow Chart" width="300" height="200">
</div>

```go
package main

import "fmt"

func main() {
	i := 1
	for i <= 10000 {
		fmt.Print("Hello World!\n")
		i++
	}
}
```
