# PYTHON & GO CHEAT SHEET

## PYTHON BASICS

### Comments
```python
# This is a comment
```

### Variables
```python
name = "John"           # Create a variable
directory = sys.argv[1] # Get from command line
```

### Lists
```python
items = ["apple", "banana", "orange"]
items[0]        # First item (apple)
items[1]        # Second item (banana)
items.sort()    # Sort alphabetically
```

### If/Else Statements
```python
if condition:
    # Do this if true
else:
    # Do this if false

# Examples:
if len(sys.argv) < 2:
    # User didn't provide argument
    
if "-l" in sys.argv:
    # User provided -l flag
```

### For Loops
```python
for item in items:
    print(item)  # Do something with each item
```

### Importing
```python
import sys      # For command line arguments
import os       # For file system operations
from pathlib import Path  # For home directory
```

---

## COMMON PYTHON FUNCTIONS

### Command Line Arguments
```python
sys.argv        # List of all arguments
sys.argv[0]     # Program name
sys.argv[1]     # First argument user typed
len(sys.argv)   # How many arguments
```

### File System Operations
```python
os.listdir(directory)           # List files in directory
os.path.join(folder, file)      # Safely combine paths (cross-platform!)
os.path.isdir(path)             # Check if it's a directory
os.stat(path)                   # Get file information
os.stat(path).st_size           # File size in bytes
os.stat(path).st_mtime          # Modification time
Path.home()                     # Get home directory
```

### Printing
```python
print("Hello")              # Print text
print(variable)             # Print variable
print(item + " <DIR>")      # Combine strings
```

---

## GO BASICS

### Comments
```go
// This is a comment
```

### Variables
```go
var directory string           // Declare a variable
directory := os.Args[1]        // Short declaration (Go figures out type)
```

### If/Else Statements
```go
if condition {
    // Do this if true
} else {
    // Do this if false
}

// Examples:
if len(os.Args) < 2 {
    // User didn't provide argument
}
```

### For Loops
```go
for _, item := range items {
    fmt.Println(item.Name())
}
// _ means "ignore this value"
// range loops through each item
```

### Importing
```go
import (
    "fmt"  // For printing
    "os"   // For file system operations
)
```

---

## COMMON GO FUNCTIONS

### Command Line Arguments
```go
os.Args         // List of all arguments
os.Args[0]      // Program name
os.Args[1]      // First argument
len(os.Args)    // How many arguments
```

### File System Operations
```go
os.ReadDir(directory)       // List files in directory
os.UserHomeDir()            // Get home directory (returns 2 values!)
item.Name()                 // Get name of file/folder
item.IsDir()                // Check if it's a directory
```

### Printing
```go
fmt.Println("Hello")           // Print text
fmt.Println(variable)          // Print variable
fmt.Println(item.Name() + " <DIR>")  // Combine strings
```

### Running Go Programs
```bash
go run filename.go             # Run without compiling
go build filename.go           # Compile to executable
```

---

## KEY DIFFERENCES: PYTHON vs GO

| Feature | Python | Go |
|---------|--------|-----|
| Comments | `#` | `//` |
| Variables | `x = 5` | `x := 5` |
| Lists/Arrays | `sys.argv` | `os.Args` |
| Print | `print()` | `fmt.Println()` |
| Indentation | Required! | Uses `{ }` brackets |
| Unused variables | Allowed | NOT allowed - causes error |
| Unused imports | Allowed | NOT allowed - auto-removed |

---

## GIT COMMANDS

### First Time Setup (once per computer)
```bash
git config --global user.name "Your Name"
git config --global user.email "your.email@gmail.com"
```

### Starting a Repository
```bash
git init                       # Create new repository
git clone URL                  # Download existing repository
```

### Daily Workflow
```bash
git status                     # See what's changed
git add .                      # Stage all changes
git commit -m "message"        # Save snapshot with description
git push                       # Upload to GitHub
git pull                       # Download from GitHub
```

### Working with Branches
```bash
git branch                     # See all branches
git checkout main              # Switch to main branch
git checkout -b new-branch     # Create and switch to new branch
git branch -m old new          # Rename a branch
git branch -d branch-name      # Delete a branch locally
git push -u origin branch      # Push branch to GitHub
git push origin --delete name  # Delete branch on GitHub
```

### Checking Information
```bash
git remote -v                  # See where origin points
which go                       # See where a program is located
```

---

## LINUX TERMINAL COMMANDS

### Navigation
```bash
cd /path/to/folder             # Change directory
cd ..                          # Go up one level
cd ~                           # Go to home directory
pwd                            # Show current directory
ls                             # List files
ls -a                          # List all files (including hidden)
ls -l                          # List with details
```

### File Operations
```bash
mkdir folder                   # Create directory
touch file.txt                 # Create empty file
rm file.txt                    # Delete file
rm -rf folder                  # Delete folder and contents
mv old new                     # Rename/move file
cp source dest                 # Copy file
cat file.txt                   # View file contents
```

### Finding Things
```bash
find ~ -name "filename"        # Search for file
which program                  # Find where program is installed
```

### Special Directories
```bash
.                              # Current directory
..                             # Parent directory
~                              # Home directory
/                              # Root directory
```

---

## SSH & GITHUB

### Generate SSH Key
```bash
ssh-keygen -t ed25519 -C "your.email@gmail.com"
# Press Enter for default location
# Press Enter twice to skip passphrase
```

### View SSH Key (to add to GitHub)
```bash
cat ~/.ssh/id_ed25519.pub
# Copy entire output and add to GitHub → Settings → SSH Keys
```

### Test SSH Connection
```bash
ssh -T git@github.com
# Should say "Hi username!"
```

### Switch Repository to SSH
```bash
git remote set-url origin git@github.com:username/repo.git
```

---

## KEY CONCEPTS

### What is a Flag?
Optional argument that changes program behavior:
```bash
ls           # Just list files
ls -l        # List with details (-l is a flag)
ls -a        # List all files (-a is a flag)
```

### What is Origin?
Nickname for your GitHub repository URL. Each repository has its own origin stored in its `.git` folder.

### What is a Branch?
A separate version/timeline of your code. `main` is the official version, other branches are for working on changes.

### What is a Pull Request?
A request to merge your branch into main after review. Used in professional development for code review.

### What is a Remote?
Anything not on your current computer - usually GitHub acts as the remote that syncs between multiple computers.

### What is an HTTP Server?
A program that listens for requests (like from a web browser) and responds with data (like webpages). It uses the HTTP protocol which is how websites communicate.

---

## TROUBLESHOOTING

### "command not found"
- Program not installed OR
- Not in your PATH

### "permission denied"
- Need `sudo` (run as administrator) OR
- File permissions issue

### "No such file or directory"
- Check spelling and capitalization
- Make sure you're in the right directory (`pwd`)
- Use `ls` to see what files exist

### Go removes my imports when I save
- Go auto-removes unused imports
- Only import when you're actually using it

### Git asking for password
- Set up SSH key authentication instead

### "index out of range" error
- Trying to access sys.argv[1] or os.Args[1] when user didn't provide an argument
- Check length first with `if len(sys.argv) < 2`

---

## TIPS FOR RETENTION

1. **Use it multiple times** - The more you use something, the more it sticks
2. **Understand WHY, not just WHAT** - Understanding logic helps memory
3. **Keep notes** - Reference this cheat sheet often!
4. **Don't memorize everything** - Even experienced programmers look things up
5. **Build things** - The best way to learn is by doing
6. **Ask questions** - If you don't understand, ask!

---

**Remember:** Programming is not about memorizing everything. It's about:
- Knowing what's possible
- Knowing where to look
- Understanding the logic
- Practicing until it becomes natural
