# 算法题【Golang】

### 链表

1. 判断链表是否有环

<details>
<summary>展开查看</summary>
<pre>
<code>
解法一：链表拆分思路
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
解法二：快慢指针思路
func hasCycle( head *ListNode ) bool {
    // 快慢指针
    if head == nil {
        return false
    }fast := head
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

2. 链表逆序

<details>
<summary>展开查看</summary>
<pre>
<code>
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var ret = &ListNode{}
    p := ret
    for l1 != nil && l2!=nil {
        if l1.Val < l2.Val {
            p.Next = l1
            l1 = l1.Next
        } else {
            p.Next = l2
            l2 = l2.Next
        }
        p = p.Next
    }
    if l1 == nil {
        p.Next = l2
    } else {
        p.Next = l1
    }
    return ret.Next
}
// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil && l2 != nil {
        return l2
    } else if l1 != nil && l2 == nil {
        return l1
    }
    if l2.Val < l1.Val {
        l1, l2 = l2, l1
    }
    l1.Next = mergeTwoLists(l1.Next, l2)
    return l1
}
</code>
</pre>
</details>





