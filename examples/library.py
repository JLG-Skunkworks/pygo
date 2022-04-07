import ctypes, json
library = ctypes.cdll.LoadLibrary('./library.so')

hello_world = library.helloWorld
hello_world()

jsonDict = library.jsonDict
test = jsonDict()   

library.jsonArray.restype = ctypes.c_char_p
print(library.jsonArray())