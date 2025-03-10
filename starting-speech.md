### প্রোগ্রামের বেসিক স্ট্রাকচার
আগে আর্টিকেলে আমরা জেনেছি প্রোগ্রাম কাকে বলে? এই আর্টিকেলে আমরা জানব একটি প্রোগামের বেসিক স্ট্রাকচার।

```go
package main 

func main() {
    // Your code starts from here…..
}
```

আমরা এখন মেইন ফাংশনের ভেতরে কোড লিখব। উপরের কোড snippet-এর ভেতর কমেন্ট করে দেখানো হয়েছে। অন্যান্য জিনিস যেমন, `package main`, `func main()` এগুলো কি সে সম্পর্কে আমরা পরবর্তীতে জানব। 

```go
package main 

import “fmt”

func main() {
    fmt.Print(“Hello World!!\n”)    
}
```
#### বিস্তারিত ব্যাখ্যা 🔍
- **fmt :**
    এখানে `fmt` হচ্ছে প্যাকেজের নাম। `fmt` প্যাকেজের অন্তর্ভুক্ত একটি ফাংশন হল `Print`. এই ফাংশনের সাথে () ব্র্যাকেটে ডাবল কোটেশনের ভেতর আমরা যাই লিখব, সেটাই আউটপুট আকারে কমান্ড প্রম্পটে প্রিন্ট হবে। 🖨️

    `fmt` প্যাকেজের ফাংশন ব্যবহার করা হয়েছে বলে `package main` এর পর `import fmt` লেখা হয়েছে। 
- **\n:**
    প্রোগ্রামে `Hello World!!` এর পরে \n লেখা হয়েছে। এখানে, \n হচ্ছে একটি সিঙ্গেল ক্যারেক্টার যা newline হিসেবে কাজ করে। এটি কমান্ড প্রম্পটে আউটপুটে একটি নতুন লাইন তৈরি করে। ✨

#### 🛠️ প্রোগ্রাম রান করার কমান্ড:
```bash
go run main
```

#### আউটপুটঃ
```code 
Hello World!!
```

এটাই আপনার প্রথম Go প্রোগ্রাম। 🎉🙌 কোডিংয়ের এই যাত্রায় আপনাকে স্বাগতম! 🚀💻
