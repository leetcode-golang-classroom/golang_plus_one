# golang_plus_one

You are given a **large integer** represented as an integer array `digits`, where each `digits[i]` is the `ith` digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading `0`'s.

Increment the large integer by one and return *the resulting array of digits*.

## Examples

**Example 1:**

```
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].

```

**Example 2:**

```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].

```

**Example 3:**

```
Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].

```

**Constraints:**

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain any leading `0`'s.

## 解析

給定一個非負整數陣列 digits 

代表一個數值的每個位元 由左至右由最高位元到最低位元

要求實作一個演算法算出原本 digits 數值加1 所出現的新 digits

 因為知道 每個 digit 運算要從最低位元運算

所以只要從最低位元開始計算每個 digit 進位所帶來的 carry 與當下位元相加的結果

由最低位元往最高位元加 最後再從最高位元開始放入新的 array 即可
![](https://i.imgur.com/7lngUSS.png)

這樣做的時間複雜度是 O(n)

空間複雜度只有回傳值的 O(n)

## 程式碼
```go
package sol

func plusOne(digits []int) []int {
	carry := 1
	lastPos := len(digits) - 1
	res := []int{}
	for pos := lastPos; pos >= 0; pos-- {
		result := digits[pos] + carry
		if result > 9 {
			carry = 1
		} else {
			carry = 0
		}
		digits[pos] = result % 10
	}
	if carry != 0 {
		res = append(res, carry)
	}
	for pos := 0; pos <= lastPos; pos++ {
		res = append(res, digits[pos])
	}
	return res
}
```

## 困難點

1. 要看出 digit 加法要從最低位元開始加

## Solve Point

- [x]  初始化 carry = 1, lastPos := len(digits) - 1, res = []
- [x]  從 pos = lasPos 到 pos = 0 開始做以下運算
- [x]  result = digits[pos] + carry
- [x]  if result > 9  更新 carry = 1 否則 carry = 0
- [x]  digits[pos] = result % 10
- [x]  到跑完所有 pos 檢查  if carry ≠ 0 把 carry 新增到 res
- [x]  然後把 digits的每個數值新增到 res
- [x]  回傳 res