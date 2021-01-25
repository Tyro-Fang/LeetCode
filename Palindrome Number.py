def judge(x):
    num=list(str(x))
    numr=num[::-1]
    if num==numr:
        return True
    else:
        return False

print(judge(-121))




