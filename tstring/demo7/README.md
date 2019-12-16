##### 人生就是一场游戏，我们不停的打怪通关，有的时候这个怪物是自己。————陌无崖

### 题目要求
判断一个字符串是否是一个回文串，如madam、我爱我。

### 分析
#### 解法一 两头往中间扫
对于一个回文字符串，我们可以很轻易的找到其中的规律就是首尾逐个字符进行遍历是相同的字符因此我们可以定义两个指针指向首和尾，进行遍历比较，如果遇见不相同的则返回false
#### 代码
```go
func IsPalindrome_one(str string) bool {
	// 非法输入
	if len(str) == 0 {
		return false
	}
	// 定义两个指针
	right := len(str) - 1
	left := 0
	for i := left; i < len(str)/2; i++ {
		if str[i] != str[right] {
			return false
		}
		right--
	}
	return true
}
```
#### 解法二 中间往两头扫
同理，根据规律也可从中间开始往两边进行逐渐扫描
```go
func IsPalindrome_two(str string) bool {
	if len(str) == 0 {
		return false
	}
	right := 0
	left := 0
	// 判断字符串的长度尾偶数or奇数
	if len(str)%2 == 0 {
		right = (len(str) / 2) - 1
		left = (len(str) / 2)
	} else {
		right = (len(str) / 2) - 1
		left = (len(str) / 2) + 1
	}
	// 遍历字符串
	for {
		if right == -1 {
			break
		}
		if str[right] != str[left] {
			return false
		}
		right--
		left++
	}
	return true
}
```