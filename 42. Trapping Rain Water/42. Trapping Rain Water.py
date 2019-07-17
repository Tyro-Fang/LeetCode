from typing import List
class S:
    def trap(self,height):
        res=0
        hlen=len(height)
        if hlen<=2:
            return 0
        lefth=[0]*hlen
        righth=[0]*hlen
        maxh=0
        for i in range(hlen):
            maxh=max(height[i],maxh)
            lefth[i]=maxh
        maxh=0
        for i in range(hlen-1,-1,-1):
            maxh=max(height[i],maxh)
            righth[i]=maxh
        for i in range(hlen):
            res+=min(lefth[i],righth[i])-height[i]
        return res
            
    def __init__(self,a):
        b=self.trap(a)
        a=n
        print(b)

a=[5,2,1,2,1,5]
b=S(a)
  