<div align="center">

# down-sorter

![Go](https://img.shields.io/badge/Go-1.22.2-00ADD8?style=for-the-badge&logo=go&logoColor=white)

A simple CLI tool that automatically sorts files in a directory into categorized folders based on their file extension.

> Why am I even building this? I don't know. It's just me trying to code without AI like how it felt when I first started learning to code.

</div>

---

## How It Works

Run the binary inside any directory you want to clean up. The tool will:

1. Create category folders (if they don't exist yet)
2. Scan all files in that directory
3. Move each file into the matching folder based on its extension

---

## Folder Categories

| Folder | Extensions |
|---|---|
| `Document/txt` | `.txt` |
| `Document/word` | `.doc`, `.docx`, `.odt` |
| `Document/ppt` | `.ppt`, `.pptx`, `.odp` |
| `Document/excel` | `.xls`, `.xlsx`, `.ods` |
| `Document/pdf` | `.pdf` |
| `Audio` | `.mp3`, `.wav`, `.flac`, `.aac`, `.ogg`, etc |
| `Video` | `.mp4`, `.mkv`, `.avi`, `.mov`, `.wmv`, etc |
| `Image` | `.jpg`, `.png`, `.gif`, `.svg`, `.webp`, etc |
| `Application` | `.exe`, `.msi`, `.apk`, `.deb`, `.rpm` |
| `Compressed/ZIP` | `.zip` |
| `Compressed/RAR` | `.rar` |
| `Compressed/7z` | `.7z` |
| `Compressed/TAR` | `.tar` |
| `Compressed/GZ` | `.gz` |
| `Font` | `.ttf`, `.otf`, `.woff` |
| `Other` | anything unrecognized |

---

## Project Structure

```
down-sorter/
├── main.go
├── folder/
│   ├── initFolder.go   # initialize all category folders
│   └── makeFolder.go   # helper to create a folder
├── sorter/
│   └── sorter.go       # scan & move logic
└── go.mod
```

---

## Usage

### Build

**Linux / macOS**
```bash
go build -o down-sorter
```

**Windows**
```powershell
go build -o down-sorter.exe
```

### Run

Drop the binary into whatever directory you want sorted, then run:

**Linux / macOS**
```bash
./down-sorter
```

**Windows**
```powershell
.\down-sorter.exe
```

---

## Packages Used

stdlib only — no external dependencies.

1. `os`
2. `log`
3. `path/filepath`
