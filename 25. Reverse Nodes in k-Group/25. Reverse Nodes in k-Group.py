from typing import List
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if k<=1 or not head:
            return head
        firstNode=ListNode(-1)
        firstNode.next=head
        tNode=firstNode
        i=0
        while (tNode):
            if i%k==0:
                tNode.next=self.reverseK(tNode.next,k)
            tNode=tNode.next
            i+=1
        return firstNode.next
    def reverseK(self,head,k):
        l=0
        x=head
        while x:
            l+=1
            x=x.next
            if l>=k:
                break
        if l<k:
            return head
        pre,cur,nxt=None,head,head.next
        while k:
            k-=1
            cur.next=pre
            pre,cur,nxt=cur,nxt,nxt.next if nxt else None
        head.next=cur
        return pre 
    def __init__(self,x,k):
        self.reverseKGroup(x,k)

a=ListNode(1)
a.next=ListNode(2)
a.next.next=ListNode(3)
a.next.next.next=ListNode(4)
a.next.next.next.next=ListNode(5)
e=Solution(a,3)
