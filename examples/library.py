import ctypes, json
library = ctypes.cdll.LoadLibrary('./library.so')

hello_world = library.helloWorld
hello_world()

jsonDict = library.jsonDict
test = jsonDict()   

library.jsonArray.restype = ctypes.c_char_p
test = library.jsonArray().decode('utf8')
json_string = json.dumps(test, ensure_ascii=False)
print(type(json_string))
json_dict = json.loads(test)
print(json_dict)