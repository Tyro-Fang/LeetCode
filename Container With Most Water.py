def maxArea(height):
    maxA=0
    
    for i in range(len(height)//2):
        begin=height[i]
        me=0
        for j in reversed(range(i+1,len(height))):
            end=height[j]
            if end >me:
                me=end
                he=min(begin,end)
                area=he*(j-i)
                if area>maxA:
                    maxA=area
    return maxA

def maxArea1(height):
    begin=0
    end=len(height)-1
    maxA=0
    while begin<end:
        area=min(height[begin],height[end])*(end-begin)
        maxA=max(area,maxA)
        if  height[begin]<height[end]:
            
            begin=begin+1
        elif height[begin]>height[end]:
            end=end-1
        else :
            end-end-1
            begin=begin+1
    return maxA
            

a=[2,3,4,5,18,17,6]
print(maxArea1(a))
        

    