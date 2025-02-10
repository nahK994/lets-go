### Golang-এ লুপ 
সহজ কথায়, লুপ মানে বনবন করে ঘোরা। প্রোগামিং-য়ের ভাষায়, এক/একাধিক instruction বারবার execute হতে থাকাকে `লুপ` বলে। 
মনে করুন, আপনাকে একটা প্রোগ্রাম লিখতে বলা হল যেটা রান করলে `Hello World!` প্রিন্ট করবেন। আপনি নিশ্চয় এরকম কোড করে ফেলবেনঃ
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
প্রথমে আমরা একটা ফ্লো-চার্টের মাধ্যমে এই সমস্যার সমাধান বোঝার চেষ্টা করি। 
<div style="display: flex; justify-content: center; margin-bottom:20px;">
<img src="./loop_flow_chart1.jpg" alt="Flow Chart" width="450" height="550">
</div>

উপরের ফ্লো-চার্টে `i` নামে একটা ভ্যারিয়েবল declare করে তাতে ১ assign করা হয়েছে। এরপর, একটা লুপের ভেতর `Hello World!` প্রিন্ট করা হচ্ছে আর সাথে সাথে `i` এর ভ্যালু ১ করে বাড়ানো হচ্ছে।

যতক্ষণ, `i` এর ভ্যালু  ১০০০০ এর সমান বা ছোট থাকে(i<=১০০০০), ততক্ষণ লুপটি চলবে। যখন, `i` এর ভ্যালু ১০০০১ হবে (অর্থ্যাৎ, i<=১০০০০ শর্তটি মিথ্যা হবে), তখন লুপ থেকে বের হয়ে প্রোগ্রামের execution শেষ হয়ে যাবে।  
```go
package main

import "fmt"

func main() {
	var i int
i = 1
	for i <= 10000 {
		fmt.Print("Hello World!\n")
		i++
	}
}
```

উপরের কোডে,

```go
for i <= 10000 {
		// instructions... 
	}
```

এই অংশটি হচ্ছে লুপ। এখানে, `for` হচ্ছে লুপের কিওয়ার্ড। `i<=10000` হচ্ছে লুপ থেকে বের হবার exit condition. এই condition break করলে(false হলে) এই লুপ থেকে বের হয়ে যাবে।

এবার কিছুক্ষণ সময় নিয়ে উপরোক্ত কোডটিকে ফ্লো-চার্ট-এর মিলিয়ে দেখার চেষ্টা করুন। আশা করি, সব ক্লিয়ার হয়ে যাবে।
 
এখন যদি বলি, এমন একটি প্রোগ্রাম লিখুন যা ১ থেকে ১০ পর্যন্ত প্রিন্ট করবে। পারবেন ????

খুবই সহজ, `Hello World!` এর জায়গায় `i` এর ভ্যালু প্রিন্ট করবেন। ব্যস।

```go
package main

import "fmt"

func main() {
	var i int
i = 1
	for i <= 10000 {
		fmt.Print(i,"\n")
		i++
	}
}
```

এবার আপনারা বের করবেন নিম্নোক্ত কোডের আউটপুট কি হবে? আমি বলব না।

```go
package main

import "fmt"

func main() {
	var i int
   i = 1
	for i <= 10000 {
		i++
		fmt.Print(i,"\n")
	}
}
```
