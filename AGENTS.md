# Agent Notes

Small Go CLI that scaffolds a ticket folder.

## Run / build

```bash
# Run directly (interactive prompts follow)
go run . init <7-digit-ticket-number>

# Build the binary (matches .gitignore entry)
go build -o makeTicket .
./makeTicket init <7-digit-ticket-number>
```

## Commands

- `init <ticket>` — create the ticket folder and write the PM template.
- `status` — placeholder; currently returns empty.
- `config` — print the project/version stored in the current folder's `.config`.
- `help` — list commands.

## What `init` does

Creates a directory named after the Ticket number in the current working directory and writes:

- `<ticket>/<ticket>.txt~` — filled PM template (Spanish)
- `<ticket>/<ticket>.xml~` — empty placeholder
- `<ticket>/evEntregable.png~` — empty placeholder
- `<ticket>/evSVN.png~` — empty placeholder
- `<ticket>/.config` — CSV metadata: `ticket,project,version`

## Configuration

`config.json` supplies the `projects` and `versions` maps. The CLI looks for it in this order:

1. The user's config directory: `~/.config/makeTicket/config.json` on Linux/macOS, `%APPDATA%\makeTicket\config.json` on Windows.
2. Falls back to `config.json` in the current working directory.

## Repo quirks

- Module name is `github.com/vramirez/makeTicket`; repo path is `.../ticket`.
- No tests, no CI, no Makefile.
- Several version entries have `svnURL: "ACTUALIZAR"` (placeholder).
- Output and error messages are in Spanish.
- The `.gitignore` ignores the compiled `makeTicket` binary.
