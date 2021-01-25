from typing import List
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        if not lists:
            return None
        a=b=len(lists)
        mul=1
        while a>1:
            j=0
            a=a//2 if a%2==0 else(a//2+1)
            for i in range(0,a):
                if j+mul<b:
                    lists[j]=self.merge(lists[j],lists[j+mul])
                j+=(mul*2)
            mul*=2
        return lists[0]
    def merge(self,l1,l2):
        tempBegin=ListNode(0)
        tempNode=tempBegin
        while l1 and l2:
            if l1.val>=l2.val:
                tempNode.next=l2
                l2=l2.next    
            else:
                tempNode.next=l1
                l1=l1.next
            tempNode=tempNode.next
        if not l1:
            tempNode.next=l2
        else:
            tempNode.next=l1
        l1=tempBegin.next
        return l1


    def __init__(self,x):
        self.mergeKLists(x)
