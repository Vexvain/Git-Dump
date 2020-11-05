![Git-Dump](https://i.ibb.co/SPDs1XS/image.png)

-----------------------------------------------------------

# Usage
```bash
usage: git-dump [flags] url [dir]

Flags:
   -f, --force   overrides DIR if it already exists
   -h, --help    help for git-dump
```

# Example
```bash
git-dump example.com
```

# How to install
```bash
go get -u github/Vexvain/Git-Dump
```

# How it works
The tool will check if directory listing is available. If so, it will simply download the .git directory (like wget)

If directory listing is not currently available, the tool will use several methods to find as many files as possible

**Step by step:**
* Fetch all common files
* Find as many refs as possible
* Finding as many objects as possible
* Fetch all objects continually
* Run `.git checkout .` to recover the current tree
