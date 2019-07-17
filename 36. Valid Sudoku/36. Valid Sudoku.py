from typing import List
class Solution:
    def isValidSudoku(self,board: List[List[str]]):
      list1=[]
      colList = [[] * 1 for row in range(9)]
      threeList=[[] * 1 for row in range(9)]
      for i in range(0,9):
        a=[]#行取值
        for j in range(0,9):
          if board[i][j]!='.':
            a.append(board[i][j])
            colList[j].append(board[i][j])
            x=i//3
            y=(j//3)*3
            threeList[x+y].append(board[i][j])
        list1.append(a)
      for i in range(0,9):
        a=self.isValidSet(list1[i])
        b=self.isValidSet(colList[i])
        c=self.isValidSet(threeList[i])
        if a==False or b==False or c==False:
          return False
      return True
            
    def isValidSet(self,s):
      slen=len(s)
      if slen<=1:
        return True
      for i in range(0,slen-1):
        for j in range(i+1,slen):
          if s[i]==s[j] and s[i]!='.':
            return False
      return True             
    
    def __init__(self,x):
        a=self.isValidSudoku(x)
        print(a)
    
a=[["5","3",".",".","7",".",".",".","."],
["6",".",".","1","9","5",".",".","."],
[".","9","8",".",".",".",".","6","."],
["8",".",".",".","6",".",".",".","3"],
["4",".",".","8",".","3",".",".","1"],
["7",".",".",".","2",".",".",".","6"],
[".","6",".",".",".",".","2","8","."],
[".",".",".","4","1","9",".",".","5"],
[".",".",".",".","8",".",".","7","9"]]
b=Solution(a)
