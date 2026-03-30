# ASCII Art Color (Go)

A command-line ASCII art generator written in Go with support for **colorized output**.
This project extends a basic ASCII art renderer by allowing you to apply terminal colors to either the entire text or specific substrings.

---

## ✨ Features

* Convert plain text into ASCII art using banner fonts
* Apply ANSI colors to:

  * Entire text
  * Specific substrings within the 
* Support for escaped newlines (`\n`)
* Command-line interface (CLI)
* Table-driven tests for correctness

---

## 📁 Project Structure

```bash
ascii-art-color/
├── banners/
│   └── standard.txt          # ASCII font file
├── MethodsAndTesting/
│   ├── color.go              # Color logic
│   ├── file_handler.go       # File reader
│   ├── format.go             # ASCII rendering logic
│   └── format_test.go        # Unit tests
├── main.go                   # CLI entry point
```

---

## ⚙️ Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/ascii-art-color.git
cd ascii-art-color
```

2. Make sure Go is installed:

```bash
go version
```

---

## 🚀 Usage

### 🔹 Basic ASCII Output

```bash
go run . "hello"
```

---

### 🔹 With Color (Full Text)

```bash
go run . --color=red "hello"
```

---

### 🔹 Color Specific Substrings

```bash
go run . --color=green he "hello"
```

👉 This will color only the substring `"he"` inside `"hello"`.

---

### 🔹 Multiple Substrings

```bash
go run . --color=blue he ll "hello"
```

---

### 🔹 Newlines Support

```bash
go run . "hello\nworld"
```

---

## 🎨 Available Colors

### Standard Colors

* red
* green
* yellow
* blue
* magenta
* cyan
* white
* black

### Bright Colors

* bright_red
* bright_green
* bright_yellow
* bright_blue
* bright_magenta
* bright_cyan
* bright_white

---

## 🧠 How It Works

### ASCII Rendering

* Each character is represented using **8 lines** from a banner file.
* Characters are mapped using ASCII values (32–126).
* The banner file (`standard.txt`) stores all character patterns.

### Color Application

* ANSI escape codes are used for coloring.
* If no substring is provided → entire text is colored.
* If substrings are provided → only matching parts are colored.

---

## 🧪 Running Tests

Run all tests with:

```bash
go test ./...
```

---

## ⚠️ Error Handling

* Invalid color input → program exits with error
* Missing banner file → prints `"Error"` and stops
* Invalid CLI flag → shows usage message:

```bash
Usage: go run . [OPTION] [STRING]
EX: go run . --color=<color> <substring to be colored> "something"
```

---

## ⚡ Example Output

```bash
go run . --color=cyan "Go"
```

Produces colored ASCII art in the terminal.

---

## 🔧 Future Improvements

* Support multiple banner styles (shadow, thinkertoy, etc.)
* Better CLI flag parsing
* Performance optimization for substring matching
* Unicode character support
* Configurable input/output files

---

## 📄 License

This project is open-source and available under the MIT License.

---

## 👨‍💻 Author

Built as a Go project to explore:

* String manipulation
* File handling
* CLI development
* Terminal styling with ANSI colors

---
