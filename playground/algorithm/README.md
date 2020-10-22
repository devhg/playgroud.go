
# 算法题【Golang】

### 链表

1. 判断链表是否有环
<details>
<summary>展开查看</summary>
解法一：链表拆分思路
<pre>
<code>
package main
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 
 * @param head ListNode类 
 * @return bool布尔型
*/
func hasCycle( head *ListNode ) bool {
    // 拆分法
    p := head
    for p != nil {
        aft := p.Next
        if aft == head {
            return true
        }
        p.Next = head
        p = aft
    }
    return false
}
</code>
</pre>
解法二：快慢指针思路
<pre>
<code>
func hasCycle( head *ListNode ) bool {
    // 快慢指针
    if head == nil {
        return false
    }
    
    fast := head
    slow := head
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            return true
        }
    }
    return false
}
</code>
</pre>
</details>


