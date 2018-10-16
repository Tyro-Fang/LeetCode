def intToRoman(num):
    dictNum={0:'I',1:'V',2:'X',3:'L',4:'C',5:'D',6:'M',7:'Y',8:'Z'} 
    result=''
    b=num
    for i in reversed(range(4)):
        a=b//pow(10,i)
        b=b%pow(10,i)
        c=a//5
        if c==1:
            a=a-5
            if a!=4:
                result=result+dictNum[i*2+1]
            num1=dictNum[i*2]
            num2=dictNum[i*2+2] 
        else :
            num1=dictNum[i*2]
            num2=dictNum[i*2+1] 
        rNum=''
        if a<4:
            for i in range(a):
                rNum=rNum+num1
        elif a==4:
            rNum=num1+num2
        result=result+rNum
    return result
        
print(intToRoman(19))       
        


        


        



