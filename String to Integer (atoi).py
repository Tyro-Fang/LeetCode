def myAtoi(str):
    
    str=str.strip(' ')
    if str=='':
        return 0
    gs=str[0]
    if (gs.isdigit()==False)and(gs!='-')and(gs!='+'):
        return 0
    
    c=''
    result=''
    flag=0
    for i in range(len(str)):
        
        gs=str[i]
        if i==0:
            if gs.isdigit()==False:
                continue
        if (gs.isdigit()==True) &(flag==0):
            flag=1
        if (gs.isdigit()==False):
            break
        if flag==1:
            result+=gs
    if result=='':
        return 0
    revalue=int(result)
    if str[0]=='-':
        revalue=-revalue
    if revalue>2147483647:
        return 2147483647
    if revalue<-2147483648:
        return -2147483648    
    return revalue
    
    
print(myAtoi("   -42"))   