# Homework: Create a folder path, lists folders and files, print file list in the terminal
import os
import sys
#Files, Folders, Paths, Environment stuff i guess
#print(help(os))

from pathlib import Path


# list of all arguments that you supply to python when running
print(sys.argv)
# sys.argv = ["dir.py", "C:\\Users"]

#If the user didn't provide a directory, use the current directory (specified in '.')
if len(sys.argv) < 2:
    directory = os.path.join(str(Path.home()), "Documents")
#    directory = '.'
else:
    directory = sys.argv[1]

print("Looking in:", directory)
print()

#Specify your items variable (Need to put directory in parantheses to specify destination)
items = os.listdir(directory)

#Sort them alphabetically
items.sort()

for item in items:
    print(item)


#Arguments: Extra info you give to your program when running it. Meaning sys.argv is giving the program extra information (like a file directory when you run it)

#Path.home() - Automatically finds the home folder on any OS:

#Windows: C:\Users\username
#Mac: /Users/username
#Linux: /home/username
#os.path.join() - Joins paths correctly for any OS:
#Windows uses backslashes: C:\Users\username\Documents
#Mac/Linux use forward slashes: /Users/username/Documents
#Path.home() returns a special Path object, not a regular string.
#pythonPath.home()           # Returns: WindowsPath('C:/Users/username')
#str(Path.home())      # Returns: 'C:\\Users\\username'
#os.path.join() expects strings, not Path objects. So we convert it:
# age = 25          # This is a number
# age_str = str(25) # This is text: "25"