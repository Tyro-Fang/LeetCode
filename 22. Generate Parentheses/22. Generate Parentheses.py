def generate(n):
    result=[]
    res='('
    left=n-1
    right=n
    rec(res,left,right,result)
    return result

def rec(res,left,right,result):
    if left==0 and right==0:
        result.append(res)
    if left>0:
        rec(res+'(',left-1,right,result)
    if  right>left:
        rec(res+')',left,right-1,result)     
print(generate(3))     