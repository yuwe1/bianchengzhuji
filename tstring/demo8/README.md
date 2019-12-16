##### 人生就是一场游戏，我们不停的打怪通关，有的时候这个怪物是自己。————陌无崖

### 题目要求
给定一个字符串，求出它的最长回文子串的长度

### 题目分析
#### 常规思路
分析题目可知，该题可以抓换成两个问题如下
- 求出字符串的所有的子串
- 判断所有的字串中是回文字符串，且长W度最长
对于上面的解法属于常规思路，求子串可以利用切片的性质比较简单
```go
// 常规解法
func LookPanlind_one(str string) string {
	// 找到该字符串的所有字串存放到字符串数组
	re := getAllSubstring(str)
	max := 0
	sign := 0
	// 遍历 数组
	for i := 0; i < len(re); i++ {
		if IsPalindrome_one(re[i]) {
			if max < len(re[i]) {
				max = len(re[i])
				sign = i
			}
		}
	}
	return re[sign]
}

// 获取所有的字串
func getAllSubstring(str string) (re []string) {
	for i := 0; i < len(str); i++ {
		j := len(str)
		for j = len(str); j > i; j-- {
			re = append(re, str[i:j])
		}
	}
	return re
}

// 判断是否是回文字符串
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
虽然可以求出最终结果，但是这样的方法过于复杂，我们是将所有的字串进行查找出后进行循环判断，增加了代码的时间复杂度，下面介绍中心扩展法
#### 中心扩展法
对于一个回文字符串，找到一个中心位置，不断向外扩展，所对应的字符都是相同的，举例“aba”，以b为中心，判断出a和a相同,因此对于一个字符串，只需要找到一个中心轴，判断其两边的字符的相等情况即可。
```go
// 中心扩展法
func LookPanlind_two(str string) (s string) {

	i, j, max, c := 0, 0, 0, 0
	for i = 0; i < len(str); i++ {
		// 回文长度为奇数
		for j = 0; (i+j) < len(str) && (i-j) >= 0; j++ {
			if str[i-j] != str[i+j] {
				break
			}

			c = 2*j + 1
		}
		if max < c {
			s = str[(i - j + 1):(i + j)]
			println("奇数", s)
			max = c
		}
		// 回文长度为偶数
		for j = 0; (i+j+1) < len(str) && (i-j) >= 0; j++ {
			if str[i-j] != str[i+j+1] {
				break
			}

			c = 2*j + 2
		}
		if max < c {
			s = str[(i - j + 1):(i + j + 1)]
			println("偶数", s)
			max = c
		}
	}
	return s
}
```
#### 代码分析
对于寻找奇数的回文，扩展根据i为中心轴进行左右指针的移动，因此我们可以定义一个j初始化为0，i+j则为当前或者下一个字符，i-j则为当前或者前一个字符，然后对字符进行比较。c用于表示临时的回文长度，求出后将之与max进行比对.对于找到最长的字符串，仍可利用切片的性质，这里需要注意的是对于切片为前者包含，后者不包含，但是由于进入if循环后j多+1了一次，因此i-j便是多减去了一个1，所以要+1，后者的由于j多加1刚好满足切片的条件。
```go
for j = 0; (i+j) < len(str) && (i-j) >= 0; j++ {
    if str[i-j] != str[i+j] {
        break
    }

    c = 2*j + 1
}
if max < c {
    s = str[(i - j + 1):(i + j)]
    println("奇数", s)
    max = c
}
```
