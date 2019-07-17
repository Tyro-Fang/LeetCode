def divide(dividend: int, divisor: int) -> int:
        flag=0
        if (dividend>0 and divisor<0) or (dividend<0 and divisor>0):
            flag=1
        if divisor==1:
            return dividend
        if divisor==-1:
            t=abs(dividend)
            if t>2147483647 or t<-2147483648:
                return 2147483647
            else:
                return -dividend
        divisor=abs(divisor)
        dividend=abs(dividend)
        if dividend<divisor:
            return 0
        curr=divisor
        ans=1
        while curr+curr<=dividend:
            ans<<=1
            curr<<=1
        ans+=divide(dividend-curr,divisor)
        if flag==1:
            ans=-ans
        if ans>2147483647 or ans<-2147483648:
            return 2147483647
        return ans
        
       
a=divide(10,3)