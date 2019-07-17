class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        fake_head=ListNode(-1)
        fake_head.next=head
        point1=point2=fake_head
        while n:
            point2=point2.next
            n-=1
        while point2.next:
            point2=point2.next
            point1=point1.next
        point1.next=(point1.next).next
        return fake_head.next

