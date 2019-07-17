def letterCombinations(digits):#动态规划算法
    dic={'2':'abc','3':'def','4':'ghi','5':'jkl','6':'mno','7':'pqrs','8':'tuv','9':'wxyz'}
    dlen=len(digits)
    if dlen==0:
        return []
    loc=0
    lastarr,nowarr=[],[]
    while loc<dlen:
        if not lastarr:
            digit_str=dic[digits[loc]]
            for i in range(len(digit_str)):
                nowarr.append(digit_str[i])
            lastarr=nowarr.copy()
        else:
            nowarr.clear()
            for i in range(len(lastarr)):
                digit_str=dic[digits[loc]]
                have_str=lastarr[i]
                for j in range(len(digit_str)):
                    nowarr.append(have_str+digit_str[j])
            lastarr=nowarr.copy()
        loc+=1
    return nowarr      

def letterCombinations_1(self,digits):#dfs 字典可以定义在类中，这样每个函数调用即可，不必作为参数传入
    dic={'2':'abc','3':'def','4':'ghi','5':'jkl','6':'mno','7':'pqrs','8':'tuv','9':'wxyz'}
    if not digits:
        return[]
    res=[]
    part=''
    self.dfs(digits,dic,part,res)
    return res

def dfs(self,nums,part,res):
    if not nums:
        res.append(part)
    else:
        for i in dic[nums[0]]:
            self.dfs(nums[i:],part+i,res)
    

    
print(letterCombinations_1(self,'23'))