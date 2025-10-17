# ISit-25 | Fractional Hex Number to Binary

A tiny terminal app written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea) to convert a **fractional hexadecimal number of the form `0.ABC`** into its **binary fraction** by mapping each hex digit to 4 bits.

---

## How it works

- Type hex digits (0–9, **a–f**).
- Press **Enter** to convert.
- The app prints: `Fractional hex Number 0.<hex> converted to binary is equal to 0.<bits>`
- Controls: **Backspace/Delete** to erase, **Ctrl+C** to quit.

> Example: `0.a3f` -> output `0.101000111111`.

The conversion is done step-by-step: each hex digit is converted to decimal is turned into a 4‑bit binary chunk and concatenated after `0.`.

---

## Requirements

- Go 1.21+ (or newer)
- Module dependency: `github.com/charmbracelet/bubbletea`

---

## Setup & Run

```bash
# 1) Create a module (skip if you already have one)
git clone https://github.com/vjstripeika/hex-to-binary-ISit-25.git

# 2) Pull deps
go mod tidy

# 3) Run
go run .
```

---

## Notes & limitations

- This tool specifically handles **fractional** parts as per given asignment; it does not convert whole-number hex parts.

---

## License

MIT
