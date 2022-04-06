import ctypes
library = ctypes.cdll.LoadLibrary('./examples/library.so')
hello_world = library.helloWorld
hello_world()