# Go API Toolkit: A Beginner’s AI-Guided Learning Journey 🚀

Welcome to a guided project that uses **Go (Golang)** and **public APIs** to build your first real-world tool using **AI prompts**. This README is written for **beginners** — especially those with some exposure to Python — and shows how to learn a new language like Go using structured prompting techniques.

---

## 🧭 Why This Project?

Choosing a new language is a big step. In this project, we document how AI helped guide the decision to learn **Go**, what it took to set it up, how to build something functional quickly (like calling an API), and how to deal with errors — all while documenting the entire journey.

---

## 🔍 Choosing a Language: Go vs Python

| Feature                 | Python                             | Go (Golang)                        |
|------------------------|-------------------------------------|------------------------------------|
| Typing                 | Dynamically typed                   | Statically typed                   |
| Compilation            | Interpreted                         | Compiled (faster execution)        |
| Use Case               | Data, scripting, quick automation   | Scalable APIs, system tools        |
| Syntax Simplicity      | Beginner-friendly, flexible         | Simple but strict, structured      |
| Learning Curve         | Gentle                              | Moderate (for statically typed)    |

**Why Go?**
- It's a compiled language, so it's fast.
- Great for APIs, networking, and web backends.
- Has a clean syntax — just different from Python.

---

## 💡 What is an API, Really?

**API** stands for **Application Programming Interface**. It's like a menu in a restaurant:  
- You (the client) ask the kitchen (the server) for food (data).
- The API is the waiter — it takes your order and brings back the response.

In this case, we used a **public joke API**. Your Go program:
- Sends a request to the API.
- Gets back a joke in JSON format.
- Prints the joke in your terminal.

---

## 🧠 Learning Plan: Mastering Go with AI Support

### **End Goal**: Make an API-calling tool in Go, with timestamp and error handling.

### **Phase 1: Get Set Up**
- Install Go from [https://go.dev/dl](https://go.dev/dl)
- Add to PATH (AI prompt helped debug PATH issue)
- Initialize Go module and Git project

### **Phase 2: Build a Minimal Project**
- Prompt: “Give me a basic Go script to call an API”
- Learn how Go structs handle JSON
- Add JSON unmarshalling, print the joke
- Add timestamp to output

### **Phase 3: Handle Real Errors**
- Prompt: “How to fix `undefined: fmt.PrintIn`”
- Prompt: “Why go run doesn’t work from VS terminal?”
- Handle missing fields in API
- Log raw response when parsing fails

### **Phase 4: Add Features with AI Help**
- Prompt: “How to let user pick between two APIs?”
- Add option for Chuck Norris jokes
- Handle two different JSON response structures

---

## 🧑‍💻 Prompts Used and Why

| Prompt | Purpose |
|--------|---------|
| “How to install Go and set PATH on Windows?” | Fix Go not recognized |
| “How to safely init Git in Windows folder?” | Fix dubious ownership |
| “Build a joke-fetching API call in Go” | Core functionality |
| “Parse JSON into struct with error fallback” | Error-proofing |
| “Format current time in Go” | Logging enhancement |
| “Switch between 2 APIs using input in Go” | Bonus feature |

---

## 🧪 Example Output

```
Choose joke type: 1 for Dad Joke, 2 for Chuck Norris Joke
2
Time fetched: Wed, 07 Aug 2025 10:11:32 EAT
Chuck Norris says:
Chuck Norris doesn’t read books. He stares them down until he gets the information he wants.
```

---

## 🧩 Errors & Fixes (With Explanation)

| Error | What It Means | How AI Helped |
|-------|---------------|---------------|
| `'go' not recognized` | Go isn’t in your system PATH | Prompt helped add `C:\Go\bin` to PATH |
| `git: dubious ownership` | Git can’t trust this folder | AI suggested safe directory config or changing directory |
| `undefined: fmt.PrintIn` | Typo (should be `fmt.Println`) | AI helped spot typo |
| VS Code: webview error | Go extension crash | Ignored and used CLI instead |

---

## 🛠 How to Run This Project

1. Clone the repo
2. Run:
   ```bash
   go run main.go
   ```

3. Choose `1` for Dad Joke or `2` for Chuck Norris Joke

---

## 🧃 Bonus Features

- ✅ Timestamp shown at every fetch
- ✅ Better JSON error logs
- ✅ User chooses joke type
- ✅ Shows use of two real-world APIs

---

## 📚 References

- [Go Docs](https://go.dev/doc/)
- [Official Joke API](https://github.com/15Dkatz/official_joke_api)
- [Chuck Norris API](https://api.chucknorris.io/)
- [VS Code Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

---

## 🎓 Learning Strategy Recap

This project follows a **structured prompting strategy**:
1. **Conceptual prompts**: Understand what Go is and how it differs from Python
2. **Step-by-step prompts**: Get help building each part of the project
3. **Guided implementation**: Ask AI to walk through examples
4. **Verification prompts**: Ask AI to check if your code follows best practice

> This strategy keeps *you* in the driver's seat. AI doesn’t replace thinking — it enhances learning.

---

## 🏁 Conclusion

This isn’t just a Go project — it’s a learning journey.  
By documenting prompts, errors, fixes, and reflections, you’ve built not just a toolkit — but a roadmap for other beginners too.

## Reflection: How I Used AI in This Project
Throughout this project, I used generative AI tools like ChatGPT and ai.moringaschool.com to support my learning. I relied on AI to:

         Break down complex Go concepts into beginner-friendly steps

         Troubleshoot unfamiliar errors

         Refine my prompts through iteration

         Structure documentation clearly and professionally

         I did not copy solutions blindly — I used AI to accelerate my understanding and verify my thinking. Every part of this project reflects my personal learning journey, powered by thoughtful use of generative AI.

Reach out to me via wanjirungari2@gmail.com for any questions or feedback. 

HAPPY CODING!!!!




