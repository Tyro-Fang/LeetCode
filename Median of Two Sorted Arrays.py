def judge(l1,l2):
    l3=[]
    if (len(l1)>0)&(len(l2)>0):
        a=l1.pop()  
        b=l2.pop()
        while True:
            if a>=b:
                l3.append(a) 
                c=1
            else:
                l3.append(b)
                c=2
            
            if c==1:
                if len(l1)>0:
                    a=l1.pop()
                else:
                    l3.append(b)
                    while len(l2)>0:
                        l3.append(l2.pop())
                        if len(l2)==0:
                            break
                    break
            if c==2:
                if len(l2)>0:
                    b=l2.pop()
                else:
                    l3.append(a)
                    while len(l1)>0:
                        l3.append(l1.pop())
                        if len(l1)==0:
                            break
                    break
    elif len(l1)==0:
        l3=l2
    elif len(l2)==0:
        l3=l1
    
    if len(l3)%2==0:
        index=int(len(l3)/2)
        x=l3[index-1]
        y=l3[index]
        return (x+y)/2
    else:
        index=int((len(l3)+1)/2-1)
        return l3[index]

#只需要找到处在中位的一个或两个数字，无需得到排序完的数组
def findm(l1,l2):
    totallen=len(l1)+len(l2)
    if totallen%2==0:
        tm=totallen//2
    else:
        tm=(totallen-1)//2
    index1=0
    index2=0
    tmp=0
    for i in range(tm+1):
        tmp1=tmp
        if (index1<len(l1))&(index2<len(l2)):
            if l1[index1]<=l2[index2]:
                tmp=l1[index1]                
                index1=index1+1                   
            else:
                tmp=l2[index2]               
                index2=index2+1                
        elif index1<len(l1):
            tmp=l1[index1]
            index1=index1+1        
        elif index2<len(l2):
            tmp=l2[index2]
            index2=index2+1         
    if totallen%2==0:
        tmp=(tmp+tmp1)/2
    return tmp
    
    


l1=[1,3]
l2=[2]
#print(judge(l1,l2))
print(findm(l1,l2))
