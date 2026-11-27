# Repository Guidelines

## Project Structure

Use this repository to learn Go and review it for technical interviews. Each lesson lives on its own branch, named
`lesson-NNN-topic`, such as `lesson-012-functions-and-parameters`. Use a three-digit number and a lowercase,
hyphen-separated English topic.

## Lesson Content

Stay within the current lesson topic. When the user provides material, use its questions and examples as the basis.
Correct inaccuracies and add only edge cases that change the result or clarify a rule of this topic. Mention adjacent
topics briefly or link to them.

Explain what happens and why in enough detail to understand the rule. Keep commands and routine steps brief. Use one
short example per rule by default and add another when it demonstrates different behavior. Cut repetition and prose
that adds no understanding. Let the explanation determine the length.

- `README.md` explains the lesson in Russian and stands on its own. Introduce the topic up front and include the
  context, definitions, and short examples readers need to understand it without opening `hello.go`.
- Use descriptive section headings for major concepts and subheadings for distinct rules, behaviors, and edge cases.
  Place short Go code blocks next to the points they illustrate. Use several focused blocks when a section covers
  multiple ideas, so readers can understand each rule before moving to the next one.
- In README code blocks, use concise Russian comments next to the relevant lines. Explain what the code does and why
  it behaves that way, such as which field changes, what gets copied or shared, or why a line fails to compile or
  panics. Show the expected output beside the call. Use surrounding prose for the rule and its context.
- Link to `hello.go` as a runnable companion for practice. Give the explanations in `README.md` itself.
- Keep the README focused on the current topic. Link to related lesson branches where helpful.
- `hello.go` contains runnable, idiomatic examples in `package main`. Use descriptive Go identifiers and Russian
  comments to explain each example, including edge cases. Show exact output beside the relevant calls. Document examples
  that fail to compile or panic in comments or in the README.

## Markdown Formatting

Limit Markdown source lines to 120 characters, including `README.md` and `AGENTS.md`.

Keep each list item or paragraph focused on one point. Use separate paragraphs or headings for distinct concepts.
Use numbered lists for ordered steps and bullet lists for independent points.

## Official Sources

Verify every technical claim against official Go sources, including claims from material the user provides. Check
command behavior, example output, and Go version differences. Correct inaccuracies and fill gaps that matter to the
current topic. In `Источники`, link to the exact documentation sections or API symbols that support the lesson.

- Use the language specification for syntax and semantic guarantees.
  [Go Language Specification](https://go.dev/ref/spec)
- Use Go's guidance for established idioms and conventions. [Effective Go](https://go.dev/doc/effective_go)
- Use the relevant standard library package documentation and release notes for API behavior and version changes.
  [Standard library](https://pkg.go.dev/std) · [Go release notes](https://go.dev/doc/devel/release)

## Build and Verify

Run these commands from the repository root on the lesson branch:

| Command | Purpose |
| --- | --- |
| `go run hello.go` | Compile and run the lesson. |
| `go build hello.go` | Build `hello.exe` on Windows or `hello` on Linux and macOS. |
| `gofmt -w hello.go` | Apply standard Go formatting, including tab indentation. |
| `go vet hello.go` | Check the example for common mistakes. |

Compare the program's output with the comments near the example calls in `hello.go`. The generated executables are
ignored by Git.

For a multi-file example, also verify the documented command that includes all its files. If the lesson has
a `go.mod`, also verify `go run .`, `go build .`, and `go vet ./...`.

There is no coverage target. If a single-file lesson without a module needs regression tests, use `hello_test.go`
and run `go test hello.go hello_test.go`. For a lesson with a module, run `go test ./...`.
Name the tests `TestSomething` and use Go's `testing` package.

## Commits and Pull Requests

Write short commit subjects in Russian, for example `Добавить урок 012 о функциях и параметрах` or
`Уточнить именованные результаты в уроке 012`. In pull requests, name the lesson, summarize the changes, list the
checks, and link a related issue if one exists.
