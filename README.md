# note-cli-go: Minimal Notes CLI in Go

A fast and structured note-taking CLI built with Go. Supports 3 types of notes:

- **Fleeting Notes**: Quick timestamped thoughts.
- **Concept Notes**: Small unorganized concepts you’re learning.
- **Permanent Notes**: Well-structured, evergreen knowledge entries.

---

## 🚀 Features

```bash
note fleeting "this is a quick thought"
note concept "difference between CPU and GPU"
note permanent "Transformer Attention Mechanism"
```

## ➕ Note Types

Type - Storage - Location - Format
Fleeting - notes/fleeting/YYYYMMDDHHMM.md - Timestamped, fast-capture
Concept - notes/concepts/concepts-YYYY-MM.md - Appended with date + bullets
Permanent - notes/permanent/<slugified-title>.md - Markdown with structured sections

---

## 📦 Installation

1. Clone the repo:

```bash
git clone https://github.com/s19835/note-cli-go.git
cd note-cli-go
```

2. Run it locally:

```bash
go run main.go [command]
```

---

## 🛣 Roadmap

- MVP CLI
- Tag support
- Configurable note folder via viper
- Markdown templates
- Web UI
- AI-powered note search / summary

---

## 🛠 Built With

- Go 🦫
- Cobra (CLI framework)
- Standard Library (for file I/O, time)

---

## 🤝 Contribute

Pull requests are welcome! Start a discussion or issue if you’ve got ideas or want to collaborate.

---

## 📄 License

MIT – do whatever you want, just give credit.

---
