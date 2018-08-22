def add(l1,l2):
    
    list1=[]
    list2=[]
    while l1:
        list1.append(l1.val)
        l1=l1.next
    while l2:
        list2.append(l2.val)
        l2=l2.next
    target=list(map(int,list(str(int(''.join(map(str,list1[::-1])))+int(''.join(map(str,list2[::-1])))))[::-1]))
    return target
    
l1=[2,4]
l2=[1,3]
print(add(l1,l2))
    



