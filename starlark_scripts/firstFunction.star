"""
This module provides a simple demonstration function for Starlark.
"""

def printHello(input):

    print("In function 'printHello'")
    return "Response from Star-code with input = '"+ str(input) + "'"


def starlark_def_function():
    from_go = my_go_function("from Starlark calling Go!!!")
    print(from_go)
    return "666-" + from_go

#response = printHello(parameter)
print("last row")