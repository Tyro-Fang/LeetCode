class Person(object):
    def __init__(self,name,age):
        self._name=name
        self._age=age
    @property
    def name(self):
        return self._name
    
    @property
    def age(self):
        return self._age
    
    @age.setter
    def age(self,age):
        self._age=age
    
    def play(self):
        print('%s正在玩耍'% self._name)
    def watch(self):
        print('%s正在看电视' % self._name)
    
class Student(Person):
    def __init__(self,name,age,grade):
        super().__init__(name,age)
        self._grade=grade
    
    @property
    def grade(self):
        return self._grade
    
    @grade.setter
    def grade(self,grade):
        self._grade=grade
    
    def study(self,course):
        print ('%s 正在学习 %s' %(self._name,course))

class Teacher(Person):
    def __init__(self,name,age,tittle):
        
    