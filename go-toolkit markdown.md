A Beginner’s Guide Using AI:  Getting Started with Go API Calls – 

1. Title & Objective
Tech chosen: Go (Golang)
Why: I had never used Go before and wanted to explore a compiled, statically typed language with built-in support for concurrency and networking.
End Goal: Create a small Go program that fetches and prints a joke from a public API — guided by AI prompts and iteration.

2. Quick Summary of the Technology
What is Go?
Go is a statically typed, compiled programming language designed at Google. It’s known for its simplicity, fast execution, and native concurrency support.

Where is it used?
Used in scalable web apps, APIs, infrastructure tools (Docker, Kubernetes), and command-line tools.

Real-world example:
The backend of Kubernetes and Docker is built using Go.

3. System Requirements
OS: Windows 10+
Editor: Visual Studio Code
Go Version: go1.24.6
Other: Git (for version control)

4. Installation & Setup Instructions
A. Install Go
1. Visit https://go.dev/dl
2. Download installer for your OS
3. During install, check “Add Go to PATH”

B. Verify Installation
go version

C. Set Up Project Folder
1. Create a folder (e.g. go-api-toolkit)
2. Open it in VS Code
3. Run:
    git init
    go mod init go-api-toolkit
    touch .gitignore

D. Add .gitignore
/bin/
/pkg/
*.exe

5. Minimal Working Example
What It Does:
Fetches a random joke from a public API and prints its setup and punchline.

Code:
[Code section will be provided in README]

Expected Output:
Here's your joke:
Why don’t eggs tell jokes?
They’d crack each other up.

6. AI Prompt Journal
Prompt | Purpose | Reflection
-------|---------|------------
Help me understand my capstone | Clarify expectations | Full breakdown of deliverables
What language and project should I choose? | Pick tech & scope | Chose Go + API
Set up Go project in VS Code and Git | Setup guide | Installation, Git, Go mod, safe folders
Why is go not recognized in terminal? | Fix PATH issue | Fixed Go environment
Why does git say dubious ownership? | Git error | Moved project to safe folder
Give me a working Go API example | Build runnable code | Fetched a joke
Parse the joke from JSON | Improve output | Used struct and unmarshalling

7. Common Issues & Fixes
Error | Cause | Fix
------|-------|-----
go: command not recognized | Go not in PATH | Added C:\Go\bin to system variables
git: dubious ownership | Project in root directory | Moved to safe folder like D:\Projects\
undefined: fmt.PrintIn | Typo | Changed to fmt.Println
VS Code Go crash | Webview error | Ignored Go extension tab

8. References
Go Docs: https://go.dev/doc/
Joke API: https://github.com/15Dkatz/official_joke_api
StackOverflow: Git safe directory fix
AI Tool: ChatGPT + ai.moringaschool.com