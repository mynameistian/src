package main

import (
	"fmt"
)

//import "fmt"

// 寻找包含元音字符为偶数的最大字符串
func findTheLongestSubstring(s string) int {
	ans, status := 0, 0 // ans 最大值 status
	pos := make([]int, 1<<5)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] >= 0 {
			ans = Max(ans, i+1-pos[status])
		} else {
			pos[status] = i + 1
		}
	}
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//计算数组内俩元素之和等于 标致数的 坐标
func twoSum(nums []int, target int) (returnInt []int) {

	returnInt = make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if (nums[i] + nums[j]) == target {
				returnInt[0] = i
				returnInt[1] = j
				return returnInt
			}
		}
	}
	return returnInt
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type ListNode struct {
	Val  int
	Next *ListNode
}

func listReverse(l1 *ListNode) {
	top := &ListNode{}
	tmp := &ListNode{}
	p := &ListNode{}
	tmp = l1
	for {
		top = tmp
		tmp = tmp.Next
		p = tmp.Next
		tmp.Next = top
		tmp = p
		if tmp == nil {
			break
		}
	}
}
func myPower(bitVal int, powerVal int) (PowerNum int) {
	PowerNum = 1
	if powerVal == 0 {
		return
	} else {
		for i := 1; i < powerVal+1; i++ {
			PowerNum *= bitVal
		}
	}
	//fmt.Println("myPower", PowerNum)
	return
}
func sumList(p *ListNode) (listNum int) {
	i := 0
	for {
		//fmt.Println(p.Val)
		listNum += (p.Val * myPower(10, i))
		i++
		p = p.Next
		if p == nil {
			break
		}
	}
	return listNum
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var listNums *ListNode
	ptmp := listNums
	p1 := l1
	p2 := l2

	int10 := 0
	int1 := 0
	l1Val := 0
	l2Val := 0
	i := 0
	for {
		if p1 == nil && p2 == nil {
			if int10 != 0 {
				listNum := ListNode{
					Val:  int10,
					Next: nil,
				}
				ptmp.Next = &listNum
				ptmp = ptmp.Next
			}
			break
		}
		if p1 == nil {
			l1Val = 0
		} else {
			l1Val = p1.Val
			p1 = p1.Next
		}
		if p2 == nil {
			l2Val = 0
		} else {
			l2Val = p2.Val
			p2 = p2.Next
		}

		int1 = (l1Val + l2Val + int10) % 10
		int10 = (l1Val + l2Val + int10) / 10

		fmt.Println(l1Val, " ", l2Val, " ", int10)

		if i == 0 {
			listNums = &ListNode{
				Val:  int1,
				Next: nil,
			}
			ptmp = listNums
		} else {
			listNum := ListNode{
				Val:  int1,
				Next: nil,
			}
			ptmp.Next = &listNum
			ptmp = ptmp.Next
		}
		i++
	}

	tmp := listNums
	for {
		fmt.Printf("%d\n", tmp.Val)
		tmp = tmp.Next
		if tmp == nil {
			break
		}
	}

	return listNums
}

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

// 你可以假设 nums1 和 nums2 不会同时为空。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//获取中位数的角标
	nums1Len := len(nums1)
	nums2Len := len(nums2)

	num := nums1Len + nums2Len
	medianNum1 := 0
	medianNum2 := 0
	medianIndex := num / 2

	fmt.Println(nums1Len, "+", nums2Len, "+", medianIndex)
	i := 0
	j := 0
	for z := 0; z <= medianIndex; z++ {
		medianNum1 = medianNum2
		if i == nums1Len {
			medianNum2 = nums2[j]
			j++

		} else if j == nums2Len {
			medianNum2 = nums1[i]
			i++

		} else if nums1[i] < nums2[j] {
			medianNum2 = nums1[i]
			i++
		} else {
			medianNum2 = nums2[j]
			j++
		}
		fmt.Println(medianNum1, "    ", medianNum2)
	}

	if num%2 == 0 {
		return (float64(medianNum1) + float64(medianNum2)) / 2
	} else {
		return float64(medianNum2)
	}
}

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen == 0 {
		return ""
	}
	var medianIndex2 int
	medianIndex1 := strLen / 2
	if strLen%2 == 1 {
		medianIndex2 = medianIndex1
	} else {
		medianIndex2 = medianIndex1 + 1
	}

	var index [2]int
	var palindromeNum int
	var palindromeLen int
	var i, j int
	for i = 0; ; i++ {

		for j = 0; medianIndex1-i-j > 0; j++ {
			fmt.Println(medianIndex1 - i)
			if !palindrome(s[medianIndex1-i-j], s[medianIndex2-i+j]) {
				break
			}
		}
		j--
		palindromeNum = (medianIndex2 - i + j) - (medianIndex1 - i - j)
		if palindromeNum > palindromeLen {
			palindromeLen = palindromeNum
			index[0] = medianIndex1 - i - j
			index[1] = medianIndex2 - i + j
			fmt.Println(index)
		}
		if medianIndex1-i < palindromeNum/2+1 {
			break
		}
	}

	for i = 0; ; i++ {
		for j = 0; medianIndex2+i+j < strLen; j++ {
			fmt.Println(medianIndex2 + i)
			if !palindrome(s[medianIndex1+i-j], s[medianIndex2+i+j]) {
				break
			}
		}
		j--
		palindromeNum = (medianIndex2 + i + j) - (medianIndex1 + i - j)
		if palindromeNum > palindromeLen {
			palindromeLen = palindromeNum
			index[0] = medianIndex1 + i - j
			index[1] = medianIndex2 + i + j
			fmt.Println(index)
		}
		if strLen-i-j < palindromeNum/2+1 {
			break
		}
	}
	fmt.Println(index)
	return string(s[index[0] : index[1]+1])

}

func palindrome(s1 byte, s2 byte) bool {
	fmt.Println(string(s1), " ", string(s2))
	if s1 == s2 {
		return true
	} else {
		return false
	}
}

func main() {

	s := "ac"

	fmt.Println(longestPalindrome(s))

	//fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	//fmt.Println(1 << 5)
}
