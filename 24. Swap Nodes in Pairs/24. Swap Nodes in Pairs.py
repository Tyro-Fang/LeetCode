from typing import List
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        firstNode=ListNode(-1)
        firstNode.next=head
        node=firstNode
        while True:
            if node.next==None or node.next.next==None :
                break
            A=node.next
            B=node.next.next
            A.next=B.next
            B.next=A
            node.next=B
            node=node.next.next
        return firstNode.next
    def __init__(self,x):
        self.swapPairs(x)

a=ListNode(1)
a.next=ListNode(2)
a.next.next=ListNode(3)
a.next.next.next=ListNode(4)
e=Solution(a)
