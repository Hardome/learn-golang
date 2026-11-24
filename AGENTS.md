# Repository Guidelines

## Project Structure

Use this repository to learn Go and review it for technical interviews. Each lesson lives on its own branch, named `lesson-NNN-topic`, such as `lesson-012-functions-and-parameters`. Use a three-digit number and a lowercase, hyphen-separated English topic.

## Lesson Content

Cover the topic end to end at senior interview depth. Explain language guarantees, important subtleties, practical trade-offs, and all relevant edge cases. Identify implementation details and version-dependent behavior explicitly when they affect the result.

- `README.md` contains the theory, explanations, and expected console output in Russian. Keep it focused on the current topic and link to related lesson branches where helpful.
- `hello.go` contains runnable, idiomatic examples in `package main`. Use descriptive Go identifiers and Russian comments to explain each example, including edge cases. Show exact output beside the relevant calls. Document examples that fail to compile or panic in comments or in the README.

## Official Sources

Add an `Источники` section to every lesson README. Use official sources only and link directly to the sections or API symbols that support the lesson.

- Use the language specification for syntax and semantic guarantees. [Go Language Specification](https://go.dev/ref/spec)
- Use Go's guidance for established idioms and conventions. [Effective Go](https://go.dev/doc/effective_go)
- Use the relevant standard library package documentation and release notes for API behavior and version changes. [Standard library](https://pkg.go.dev/std) · [Go release notes](https://go.dev/doc/devel/release)

## Build and Verify

Run these commands from the repository root on the lesson branch:

| Command | Purpose |
| --- | --- |
| `go run hello.go` | Compile and run the lesson. |
| `go build hello.go` | Build `hello.exe` on Windows or `hello` on Linux and macOS. |
| `gofmt -w hello.go` | Apply standard Go formatting, including tab indentation. |
| `go vet hello.go` | Check the example for common mistakes. |

Compare the program's output with the README and the comments near the example calls. The generated executables are ignored by Git.

There is no `go.mod` or coverage target. If a lesson needs regression tests, use Go's `testing` package in `hello_test.go`, name functions `TestSomething`, and run `go test hello.go hello_test.go`.

## Commits and Pull Requests

Write short commit subjects in Russian, for example `Добавить урок 012 о функциях и параметрах` or `Уточнить именованные результаты в уроке 012`. In pull requests, name the lesson, summarize the changes, list the checks, and link a related issue if one exists.
