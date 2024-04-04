
#1 -------------------- data types


e = 10 + 15
# e = 10 + 15 + "sat"  #This is allowed in compile stage, But got error only in run time.
print(e)

e1 = 10.4 + 10.6
print(e1,type(e1))

f = "sathish"
f1 = 'sathish111'  #single quotes also same

f2 = """ ksfhlhf
LFHLDHLDAHG
ldfhdlhg """    #multi line string

f21 = ''' ksfhlhf
LFHLDHLDAHG
ldfhdlhg '''

f3 = r"rawstring/0sj#$%"  #raw string

print(f,f3,type(f1),type(f3))


r = True
t = False
if r:
    print("Its true",type(r))



List = [1,2,3,"sat","dat",[10,20,30],(35,65),True,{2:3,4:5}]  #can store different data types  #List(ordered, modify, allow duplicates)
print(List,type(List[0]),type(List[3]))


Tup = (1,2,3,"sat","dat1",[10,20,30],(35,65))  #Tuple(Ordered, can't modify, allow duplicates)
print(Tup)

Set = {1,3,5,"sda","sda"}  # Sets (unordered, Non-duplicate, can modify)
print(Set, type(4)) 


#Dictionary - Unordered, Non-duplicate key, can modify
Dict = {"1":1,"2":2, 3:"sat",4:[2,3,4,5,"aa"],(2,3):"tuple key",3:"sat11"}   #allow different data type as both key and value. 
print(Dict)


#None datatype, Similar as "null" value
g = None
print(type(g))

if None:  #similar to zero or False value
    print("Not executed this if condition")


#2------------> Data type conversions

#converting data type using list(), set(), tuple() between them
et = {2,4,5,6,8,2}
rt = list(et)
rt.append(34)
print(type(et),type(rt), et,id(et),rt,id(rt))

ty = [3,4,5,6,7,6]
fg = set(ty)
up = tuple(ty)
print(type(ty),type(fg),type(up))


#converting data type using str(), int(), float() between them
ty = 24
po = str(ty)
print(type(ty),type(po))

u = int("23")
print(type(u),u)


#string vs list conversion
yu = ["aa","bb","cc","dd","ee"]
# ty = "".join(yu)
ty = "-".join(yu)
print(ty,type(ty))

tp = "satho hello"
p = tp.split() #default blank space
p1 = list(tp)
print(p,type(p))
print(p1)  #['s', 'a', 't', 'h', 'o', ' ', 'h', 'e', 'l', 'l', 'o']


# converts int to int list
ty = 1234
ty = [int(i) for i in str(ty)]
print(ty)

#coonverts int list to int
uo = [6,7,8,9]
uo = [str(i) for i in uo]
ty = int("".join(uo))
print(ty)



#python keywords
#https://www.w3schools.com/python/ref_keyword_del.asp





#3 -----------> all data types and its inbuilt functions

#string

et = "sathish"
print(et.count("sa"))

print(et.find('h'))

print(et.index('t')) #Same as find(), but returns valueerror exception ,its value not found

et.islower()

et.isnumeric()

print(et.removeprefix('sat'))

print(et.replace("s","####",1))

et.strip()  #blank space removed in both sides

print(et.swapcase())

print(et[::-1])  #reverse a string



#list

ui = [4,5,6,7,8]

print(ui.index(4))



ui.insert(1,99)
print(ui)

ui.reverse()

ui.clear()  #empty list

# ui.remove(99)  #removes value

# ui.pop(3)  #remove value in 3rd index
# print(ui)

ui.extend([1,2,3])



ui.copy()   #only shallow copy of list

fg = ui[:]  #deep copy of list


#tuple

rt = (2,4,6,7)
rt.count(4)
rt.index(6)



#set
yf = {2,4,6,8}
yf.add(10)  #adds single value
yf.update([11,12,13]) #adds multiple values
yf.remove(4)
print(yf.difference({5,6,9}))  #prints uncommon values(if any yf set value in {5,6,9}, it elimates that) in yf set.
print(yf.intersection({5,6,9}))  #prints common values on both sets


#dictionary

to = {1:11,2:22,3:33,4:44}

print(to.keys())
print(to.values())
print(to.items())
to.pop(1)  #removes a key
print(to)
print(to.get(2))
to.update({9:99})  #appends another dict
if 2 in to:
    print("key is present in dict")




#other inbuilt functions like sorted(),map(), filter() etc

# https://www.w3schools.com/python/python_ref_functions.asp



import builtins
print(dir(builtins))

print(abs(-12))



fg = bytearray(101)  #bytearrray datatype
print(type(fg))

bytes(100)  #bytes type

print(ord("A"))


dg = [2,4,5,1,6]
dg = "sathish"
dg1 = list(reversed(dg))  #returns Reversed iterator object, So converts into list
dg1 = tuple(reversed(dg))
print(dg1)


gh = [23,36,46,12,3,2]
s = sorted(gh,key=lambda x:x%2==0,reverse=True)
print(s,type(s))



er = [11,22,33]
#adding one for every value
print(list(map(lambda x:x+1,er)))

#converts int list into integer
print(int("".join(map(str,er))))


sd = ["1","2","3","4"]
print(list(filter(lambda x:int(x)%2==0,sd)))

sd = [1,111,23,45]
print("maximum",max(sd))

num = [[1,2,3], [4,5,6,3,2], [10,11,12], [7,8,9]]
print(max(num, key=lambda x:x[1]))   #[10, 11, 12]  11 is biggest number in 1st index

# min()

a = ("John", "Charles", "Mike")
b = ("Jenny", "Christy", "Monica")
x = zip(a, b)
#use the tuple() function to display a readable version of the result:
print(tuple(x))   #(('John', 'Jenny'), ('Charles', 'Christy'), ('Mike', 'Monica'))

list1 = [1, 2, 3]
list2 = ['a', 'b', 'c']
list3 = ['x', 'y', 'z']
zipped = zip(list1, list2, list3)
result = list(zipped)
print(result)   #[(1, 'a', 'x'), (2, 'b', 'y'), (3, 'c', 'z')]


# all() //returns True, If all the values has True or non-empty string or >0 value

# any()  //returns True, If atleast one value has True or non-empty string or >0 value




#shallow/deep copy: 
# Mutuable objects ---- Deep copy only works in nested mutable objects. Shallow copy - Mutable works default like this. 
# Immutable objects ---- Deep copy defaulty works here.  Shallow copy is not possible, whenever Immutable objects value change address changed automatically.


ui.copy()   #only shallow copy of list
fg = ui[:]  #deep copy of list

# import copy
# copy.copy()  //shallow copy
# copy.deepcopy() //deep copy





#if statements with logical operators

if 1>=1 and 1<3 or 2<1:
    print("dd")

if 1!=1:
    print("gg")
elif 2!=2:
    print("lkl")
else:
    if True:
        print("hello")

if not False:
    print("not operator")


#loops

i = 0
while i<5:
    print(i)
    i = i+1

i = 0
while i<10:
    i = i + 1
    if i == 3 or i == 5:
        continue
    print("====",i)
    

test_list = [25, 6, 2288292, 432, 72,4,8,90,2,3,4,4,45]
while 4 in test_list:
    test_list.remove(4)
print(test_list)


while len(test_list) > 0:
    test_list.pop()
print("++++++++++++",test_list)


#for loops

for i in "sath":
    print(i)

for i in range(10):
    print("for loop",i)

for i in range(2,10,2):  #every 2nd value in loop
    print("for loop!!!",i)


#reverse for loop
for i in range(10,-1,-1):
    print("for loop11111",i)


for i in range(10,-1,-2):  #every 2nd value in loop at reverse
    print("for loop2222",i)

for i in range(1,5):
    for j in range(i+1,5):
        print(i,j)

e = "Sathooo"
for i,j in enumerate(e):
    print(i,j)

for i in range(2):
    print(i)
else:   #after for loop ends
    print("loop completed finally")

#functions

def hello(a):
    a = a + 2
    return a

print(hello(5))


def sample(a:int,b:str) -> int:   #these are "type hints", uses in arguments, return types. But just for documentation purpose, It not follows the mentioned data type execution only.
    return "ss"

print("sample-function",sample("aa",10))  #still executes fines, Eventhough different datatypes are mentioned in "Type hints"


def values(a=3,b=10):
    global LocalVariable
    LocalVariable = 100
    return a+b

print(values())   #default values will taken from keyword arguments in values()

print(values(2))  #this 2 value will be overrided in "a=3" argument value

print(values(1,2)) #both values are overrided


def arbitraryArgumentsFunction(*arguments, **keywordArguments):  #like golang variadic function, passing dynamic number of arguments
    print(type(arguments),type(keywordArguments))
    print(arguments, keywordArguments)

arbitraryArgumentsFunction()
arbitraryArgumentsFunction(100,110,120)
arbitraryArgumentsFunction(10,11,12,a=20,b=30)


def alwaysPassByReference(a,b):  
    print("inside function",id(a),id(b))  #same memory address from outside
    b.append(100)   #mutuable type - this will reflect outside function
    a = a + 2       #immutable type - Not reflect outside, Unless return it

c=2
d=[3,4,5,6]
print(id(c),id(d))
alwaysPassByReference(c,d)
print(c,id(c),"++++++++",d,id(d))


#every variable outside funtions are global variable
def MakingGlobalVariable():
    global LocalVariable   #making this variable as global variable to access outside function
    LocalVariable = 100

print(LocalVariable)   #after making variable global, Now we can access outside function.



#exception

try:
    e = 2/0
    e = 2
except Exception as err:
    print(err)
else:
    print("exception not happened, That why else case is passed here")
finally: #this part will execute finally when exception is happened or not
    print("code recovered from exception")



# raise("intentionally stops the program here")   #like golang panic


try:
    e = 19
    if e < 18:
        raise ValueError("invalid age")   #intentionally raising exception with message
except ValueError as err:
    print(err)  #above message "invalid age" prints here
else:
    print("Got valid age")


#-----------------------comprehensions

w = [10,20,30,40]
print([i for i in w if i!= 20 if i!=40 if i!=10])   #[30]

er = "sathish"
print([i*3 for i in er])

print([i*j for i in range(1,5) for j in range(i,5)])

#dictionary comprehension
print({i:str(i)*3 for i in range(10)})

#generator comprehension similar with -- () braces
#set comprehension similar with -- {} braces








#regex
import re
e = re.search(r"\s(\d)([@#$])"," 3#dhsjhf")
if e:
    print(e.group(),"-------",e.group(1),"--------",e.group(2))

r = "24435sfhfolh hefho345lndlkn kdshvk445"
t = re.findall(r"\d+",r)   #['24435', '345', '445']
t = re.findall(r"\d+$",r) #['445']
t = re.search(r"\d+",r)  
if t:
    print(t)

#re.findall -- Finds only regex matched values only out of String from above output -- ['445'] from input - "24435sfhfolh hefho345lndlkn kdshvk445"

#re.search -- If regex matches, it prints entire string -- "24435sfhfolh hefho345lndlkn kdshvk445"

#re.findall(r"\d+$",r,re.MULTILINE)  -- re.MULTILINE for multiline




#file handling
# open()


#oops


#inheritance

#https://www.w3schools.com/python/python_inheritance.asp

# https://www.digitalocean.com/community/tutorials/understanding-class-inheritance-in-python-3





class ParentSample:
    def __init__(self) -> None:
        pass
    def parentFunction(self):
        self.p = "plan"
        print("parent function-----",self.p)

    def parentFunctionSecond(self):
        print("parent function Second-----",self.p)  #self.attribute has scope over entire class, even self.p is not declared in __init__() and declared only in above  parentFunction()
    


class Sample(ParentSample):
    def __init__(self,a,b,c):  #Gets passing value into class, 
        self.a = a #assigning values with self.Variable allows scope over entire class
        self.b = b
        self.c = c

    def sss(self):
        self.d = 5   #directly creates new variable for the class   
        print("ggg",self.a)

    def sss1(self):
        print("sss1",self.d)   #self.d value assigned in above class method

    def values(self,x,y):   #gets variables from outside and assigns into class variable self.multiply
        self.multiply = x*y
        return self.multiply
    
    # def parentFunction(self):
    #     print("child class same parent function name")

s = Sample("2",10,20)
s.sss()
s.sss1()
print(s.values(10,30))
s.parentFunction()
s.parentFunctionSecond()


class Vehicle:

    def __init__(self, name, max_speed, mileage):
        self.name = name
        self.max_speed = max_speed
        self.mileage = mileage

    def Values(self):
        print(self.name,self.max_speed,self.mileage)

class Bus(Vehicle):
    def __init__(self,name, max_speed, mileage):
        Vehicle.__init__(self,name, max_speed, mileage)   
        
        #why should call the parent class __init__ and pass the same arguments?  When we creates object for child class "e = Bus("School",180,12)", If both parent and child class having __init__ method, Parent class __init__ method is overridden(not called) due to presence of child class __init__ method, So Parent class __init__ method attributes like self.name, self.max_speed,self.mileage are not called and these values are not accessible to child class object. So we should make call seperately Vehicle.__init__(), So child pass object can access all __init__ method attributes like self.name, self.max_speed etc and Values() function will print those values.

        pass

    def printValues(self):
        print(self.name,self.max_speed,self.mileage)
    
e = Bus("School",180,12)
e.printValues()


#composition
class NormalClass:
    def __init__(self,a):
        self.a = a 
    def normalFunction(self):
        return self.a*self.a
    
class CompositionClass:
    def __init__(self,b):
        self.b = b
        self.normalClass = NormalClass(b)  #calls normalClass and its constructor

    def compositionFunction(self):
        print("accessing normalClass function inside composition class ",self.normalClass.normalFunction())
        print("gggpppp")

d = CompositionClass(10)
# d.normalFunction()   #we can't access normal function directly with compositionClass object, (its difference from inheritance)
print("accessing normalClass function using compositionClass object outside----",d.normalClass.normalFunction())   
print(d.compositionFunction())

#static variable and function
class A:
    staticVariable = 10

    def staticFunction(a):
        print(a)

    @staticmethod
    def staticFunction1(a):
        print(a)


#access static methods and variables with both class name and class objects
print(A.staticVariable)
A.staticFunction(1001)   
s =A()
print(s.staticVariable)
s.staticFunction1(200)


#python abstract class and interface
#class with one or more abstract method is called abstract class.

#unlike golang, We can't create object for abstract class/interface
from abc import ABC, abstractmethod

class bank(ABC):
    @abstractmethod
    def payment(self,amount):
        pass

    def test(self):
        print("Non-abstract method")

class AxisBank(bank): 
    def payment(self, amount):   #AxisBank class is inherited from bank class, Bank class is inherited from abstract "ABC" class. 

    #Main use of abstract class in python (forces child class to implement all abstract method)

    #So this AxisBank class is also child class of abstract class "bank". If we didn't implemented this payment method, then its throws compilation error>
        print("axis payment",amount)

class ICICIBank(bank): 
    def payment(self, amount): 
        print("ICICI payment",amount)

s = AxisBank()
s.payment(100)
d = AxisBank()
d.payment(100)




#Decorator doing below three combinations with functions

#Nested function
def s():
    print("start")

    def inner():
        print("inner")
    
    inner()   #calling above inner fn
    print("end")

s()


#Return function from another function
def bbc(a):
    return a+2

def abc(a):
    
    a = a + 1
    #doing some operations and passing to another function and returning from it
    return bbc(a)

print(abc(10))


#passing another function as argument
def add(x,y):
    return x+y

def secondFunction(func,a,b):
    return func(a,b)

print(secondFunction(add,10,5))


#decorator
def decorator(func):
    def inner():
        print("decorator function")
        func()  #calling ordinary function 
    return inner

@decorator
def ordinary():
    print("ordinary function")

ordinary()


#normal function has arguments
def decorator1(func):
    def inner(*args):   # *args holds the ordinary1()'s aa argument
        print("decorator function",args)
        func(*args) #calling ordinary1-- *args is passed from inner(*args) to func(*args)   
    return inner

@decorator1
def ordinary1(aa):
    print("ordinary function",aa)

ordinary1(10)

#both normal and decorator function has arguments
def decorator2(*args):  #decorator arguments -- "hello"
    def inner(func):   #ordinary2() as argument
        def wrap(*args): # args - ordinary2() arguments 2000
            print("decorator function2",args)
            func(*args)  # *args gets passed from -- def wrap(*args)
        return wrap
    return inner

@decorator2("hello")
def ordinary2(dd):
    print("ordinary function2",dd)

ordinary2(2000)


#Decorator in class:
class ErrorCheck:
    def __init__(self, function):  #__init__  getting normal function as argument
        print (function)
        self.function = function

    def __call__(self, *params):    #For oops, Under this method only, should write the decorator function logic.
        print (params)
        if any([isinstance(i, str) for i in params]):
            raise TypeError("parameter cannot be a string !!")
        else:
            return self.function(*params)   #calls the normal function


@ErrorCheck
def add_numbers(*numbers):
    return sum(numbers)
#  returns 6
print(add_numbers(1, 2, 3))


#generator

# https://www.tutorialsteacher.com/python/python-generator

#generator function returns iterator object -- iter(data_type), 

r = iter([10,12,13])
print(r,next(r))

#can also create generator using generator comprehension

#generator function useful for memory efficient, generator function returns iteration memory_Address, memory_Address is lesser space than actual value. So Handling bigger datas like reading logs,files etc, In such cases generator can use.



def generatorFunction():
    yield 1   #Assume each yield shares smaller data(From bigger data from generatorfunction return reference) one by one   
    yield 2

w = generatorFunction()

# print(next(w))  #can also use next()

print("generator object memory address--",w)  

for i in w:
    print(i)   #here we handling the large data on one by one with generator memory object address.













#multi-threading, waiting,lock,unlock,communication in threads

import threading

def add1(a,b):
    print(a+b)

def square(a):
    print(a*a)

# if __name__ =="__main__":
t1 = threading.Thread(target=add1,args=(10,20))
t2 = threading.Thread(target=square,args=(4,)) #should add suffix "," after value 4,

t1.start()
t2.start()

t2.join()   #waiting for thread to complete, Similar to golang wg.wait()
t1.join()

print(t1.name,t1.is_alive,"=============",t1.isDaemon)

print("thread completed")


#We can use Queue to transfer the data between the python threads like golang channels.
from queue import Queue
queue = Queue()
queue.put()  
queue.get()

#Can use the mutex lock and unlock like golang to prevent racing conditions.
queue.mutex
queue.mutex.acquire()
queue.mutex.locked()
#https://superfastpython.com/thread-queue/


#file handling

r = open("dummy.txt","r+")     #can also use "with Open("dummy.txt","r+") as file:"
# r.write("sssssssssssss")
# r.write("\nttttttttttt")

# print(r.read())


for i in r.readlines():
    print("line-----",i)
print(r.mode, r.name,r.closed)
r.close()





#rest API connection, requests package

# import requests

# x = requests.get('https://w3schools.com/python/demopage.htm',PARAMS = {'address_key':"actual_location_Value"},headers={'Authorization': 'access_token actual_Token_value_here'})   #passing value to fetch matched API resources with token passing.

# print(type(x.text))


# x1 = requests.post('https://w3schools.com/python/demopage.htm',data={"Data_key":"insert_Value"})



#pytest/unit test



