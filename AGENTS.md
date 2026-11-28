# Repository Guidelines

## Project Purpose

This repository is used to learn Go systematically and prepare for technical interviews.

The reader is an experienced JavaScript and TypeScript developer learning Go. Assume familiarity with general
programming concepts, but do not assume familiarity with Go-specific semantics, idioms, or implementation details.

The goal of each lesson is to build a correct and practical understanding of the topic, not merely to document
syntax or provide a collection of examples.

## Project Structure

Each lesson lives on its own branch, named `lesson-NNN-topic`, such as `lesson-012-functions-and-parameters`.

Use a three-digit lesson number and a lowercase, hyphen-separated English topic.

- `README.md` contains the complete explanation of the lesson in Russian.
- `hello.go` contains runnable examples that reinforce the lesson.
- Additional files may be introduced when required by the topic.

## Teaching Principles

### Build a mental model

Start by explaining what the concept is, how it works, and why it exists in Go.

Introduce the fundamental concepts and terminology before relying on them in subsequent explanations.

Explain the underlying mechanism when it helps predict program behavior. Distinguish between language guarantees
and implementation details.

Do not reduce the lesson to definitions, syntax descriptions, or isolated facts.

After reading the lesson, the reader should be able to reason about unfamiliar examples independently.

### Explain progressively

Organize the material from fundamental concepts to more advanced behavior.

Prefer the following progression when appropriate:

1. Definition and fundamental concepts.
2. Basic syntax and behavior.
3. How and why the mechanism works.
4. Related operations and practical usage.
5. Important differences and limitations.
6. Non-obvious behavior, common mistakes, and edge cases.

Adapt this progression to the topic. Do not force every lesson into the same structure or introduce unnecessary
sections merely to satisfy a template.

Each section should build on concepts introduced earlier.

### Explain causes, not just results

For every important rule, explain both what happens and why it happens.

Do not merely describe what the code does or repeat its output in prose.

When relevant, explain:

- What values are created, copied, shared, modified, or returned.
- What a length, index, capacity, or offset actually measures.
- How Go evaluates an expression or performs an operation.
- Why seemingly similar operations produce different results.
- Why code fails to compile, panics, or produces unexpected behavior.
- Which guarantees come from the language specification.

Include implementation details only when they materially improve understanding of the current topic.

### Use meaningful comparisons

Use contrasting examples when they reveal an important difference in behavior.

When comparing related operations, reuse the same small input so that differences are easy to identify.

Explicitly explain the reason for the difference rather than leaving the reader to infer it from the output.

Compare Go with JavaScript or TypeScript when the comparison helps transfer existing knowledge or prevents an
incorrect assumption.

Do not add comparisons merely because the reader knows another language.

### Explain practical usage

When relevant, explain when and why a feature should be used.

Distinguish between what is technically possible and what is idiomatic Go.

When multiple approaches solve the same problem, explain their differences, limitations, and relevant trade-offs.

Discuss performance, allocations, complexity, or memory behavior only when they affect the choice of approach.

Avoid unsupported performance claims and premature optimization advice.

### Cover important edge cases

Include edge cases that reveal a language rule, change program behavior, or prevent a realistic mistake.

Prioritize cases that are relevant to practical development and technical interviews.

Explain why the edge case behaves differently from the ordinary case.

Do not turn the lesson into an exhaustive catalog of obscure language behavior.

### Balance depth and conciseness

Be thorough without being verbose.

Let the complexity of the topic determine the length of the explanation. Do not impose arbitrary limits on the
number of paragraphs, sections, or examples.

Use as many words as necessary to explain the concept correctly, but no more.

Avoid:

- Generic introductions and motivational statements.
- Repeating the same rule in different words.
- Obvious statements that add no technical understanding.
- Unnecessary historical background.
- Artificially simplified explanations that sacrifice correctness.
- Excessive abstraction, terminology, or academic language.
- Conclusions that merely repeat the entire lesson.

Do not omit important explanations simply to make the lesson shorter.

## Lesson Scope

Stay within the current lesson topic.

When the user provides material, use its questions, examples, and structure as the foundation of the lesson.

Preserve the supplied topic order when it is logically coherent. Reorganize it only when necessary to explain
prerequisites, improve progression, or correct conceptual inconsistencies.

Treat supplied material as a starting point, not as the limit of the explanation.

Correct inaccuracies, clarify ambiguous statements, and fill gaps that prevent a complete understanding of the topic.

Add missing concepts when they are necessary to understand the topic, predict its behavior, or avoid common mistakes.

Mention adjacent topics briefly when necessary. Link to other lessons for detailed explanations instead of
expanding the scope of the current lesson.

## README Structure

`README.md` is the primary learning material and must stand on its own.

The reader should be able to understand the complete topic without opening `hello.go`, reading another lesson,
or following external documentation links.

### Opening

Start with the lesson title followed by a concise explanation of the topic.

Define its essential concepts and establish the mental model needed for the rest of the lesson.

Do not start with generic introductions, learning objectives, or descriptions of what the lesson will cover.

### Sections

Divide the lesson into logical sections with descriptive headings.

Each major concept should have its own section. Use subsections for distinct rules, behaviors, and edge cases
within that concept.

Prefer headings that identify the actual concept or operation rather than generic headings such as "More Details",
"Additional Information", or "Things to Know".

Arrange sections so that explanations progress naturally and do not depend on concepts introduced later.

### Explanations

Use the following pattern when explaining an important rule:

1. State the rule or concept clearly.
2. Explain the mechanism behind it.
3. Demonstrate the behavior with a focused example.
4. Explain any non-obvious result or practical implication.

This is a teaching pattern, not a mandatory visible template. Do not label every explanation with these steps.

Keep paragraphs focused on one idea.

Place explanations and examples together so that the reader does not need to navigate between sections to
understand a single rule.

Use **bold text** to emphasize important distinctions, limitations, and language guarantees. Avoid excessive
emphasis.

### Tables and quick references

Use compact tables when they make comparisons, operation results, or differences between related concepts
easier to understand.

Prefer tables for structured comparisons and prose for causal explanations.

Do not use a table as a replacement for explaining an unfamiliar concept.

When useful, finish the lesson with a compact quick reference that summarizes important operations, rules,
or distinctions for later review and interview preparation.

Do not add a quick reference if it would simply duplicate an already concise lesson.

### Navigation

Link to `hello.go` as a runnable companion for practice.

Keep the explanations in `README.md` itself.

Link to other lessons at the point where a related concept appears, rather than collecting unrelated links
in a separate section.

Finish the lesson with an `Источники` section containing links to the relevant official documentation.

## Code Examples

### Examples in README

Place short Go code blocks immediately after or near the explanation they illustrate.

Each code block should demonstrate one rule or a small group of closely related rules.

Prefer several focused examples over one large example covering multiple unrelated concepts.

Use the smallest example that clearly demonstrates the behavior without introducing irrelevant complexity.

Avoid repeating `package main`, imports, and surrounding boilerplate when they are not necessary to understand
the example.

Include required context when omitting it would make the example ambiguous.

Use descriptive identifiers that reflect the purpose of the example.

Show exact expected output beside the relevant calls using concise comments.

Write comments in Russian when an explanation is needed.

Comments should clarify non-obvious behavior, such as:

- What is copied or shared.
- Which value is modified.
- What an index or length represents.
- Why a conversion produces a particular result.
- Why an operation fails to compile or panics.

Do not use comments to mechanically restate self-explanatory code.

Use surrounding prose to explain the general rule, its mechanism, and its practical significance.

Examples demonstrating compilation errors or panics must explicitly identify the failure and explain its cause.

### Runnable examples

`hello.go` contains runnable, idiomatic Go examples in `package main`.

Organize examples in the same logical order as the README.

Use separate functions when they improve readability and isolate independent concepts.

Use descriptive Go identifiers and concise Russian comments.

Show exact expected output beside the relevant calls.

Cover the important behaviors and edge cases discussed in the README.

Keep examples deterministic whenever possible.

For non-deterministic behavior, explain which parts of the output are not guaranteed.

Do not include actively executed examples that intentionally fail to compile or panic. Document them in comments
or in the README instead.

Avoid introducing unnecessary abstractions, dependencies, or architectural patterns into educational examples.

## Writing Style

Write all lesson explanations in Russian.

Use natural, direct, and technically precise language.

Use established English technical terms when they are more recognizable or accurate than their Russian
translations. Introduce their meaning when necessary.

Prefer concrete explanations over abstract formulations.

Assume the reader understands programming but needs to understand the specific rules and behavior of Go.

Do not explain fundamental programming concepts unless their behavior in Go differs in a meaningful way.

Avoid filler phrases, conversational padding, rhetorical questions, and unnecessary transitions.

Do not sacrifice technical accuracy for simplicity.

## Markdown Formatting

Limit Markdown source lines to 120 characters, including `README.md` and `AGENTS.md`.

Use consistent Markdown formatting:

- Use `#` for the lesson title.
- Use `##` for major concepts.
- Use `###` for distinct rules and subsections.
- Use fenced code blocks with the `go` language identifier for Go examples.
- Use inline code for identifiers, types, expressions, commands, and language keywords.
- Use numbered lists for sequential steps.
- Use bullet lists for independent points.
- Use tables for compact structured comparisons.

Keep each paragraph or list item focused on one point.

Avoid unnecessary nested lists, deeply nested headings, and excessive formatting.

## Official Sources

Verify technical claims against official Go sources, including claims from material provided by the user.

Use:

- [Go Language Specification](https://go.dev/ref/spec) for syntax and semantic guarantees.
- [Effective Go](https://go.dev/doc/effective_go) for established idioms and conventions.
- [Go Blog](https://go.dev/blog/) for official explanations of language concepts.
- [Standard library](https://pkg.go.dev/std) for API behavior.
- [Go release notes](https://go.dev/doc/devel/release) for version-specific changes.
- [Go FAQ](https://go.dev/doc/faq) for language design decisions and common questions.

Prefer primary sources over third-party tutorials.

Use runtime or compiler source code only when implementation-specific behavior is relevant to the lesson.

Distinguish between specification guarantees, documented API behavior, and current implementation details.

Check the installed Go version when version-specific behavior matters.

In the `Источники` section, link to the exact documentation sections, articles, or API symbols that support
the lesson.

Do not include generic documentation links when a more specific reference is available.

Do not claim that a source was verified if it could not be accessed.

## Build and Verify

Run the following commands from the repository root on the lesson branch:

| Command | Purpose |
| --- | --- |
| `gofmt -w hello.go` | Apply standard Go formatting. |
| `go run hello.go` | Compile and run the lesson. |
| `go build hello.go` | Verify that the lesson builds successfully. |
| `go vet hello.go` | Check the example for common mistakes. |

Compare the actual program output with the expected output documented in `hello.go` and `README.md`.

Correct discrepancies before completing the lesson.

The generated executables are ignored by Git.

For a multi-file example, also verify the documented command that includes all required files.

If the lesson has a `go.mod`, also verify:

- `go run .`
- `go build .`
- `go vet ./...`

There is no coverage target.

If a single-file lesson without a module needs regression tests, use `hello_test.go` and run
`go test hello.go hello_test.go`.

For a lesson with a module, run `go test ./...`.

Name tests `TestSomething` and use Go's `testing` package.

## Lesson Quality Checklist

Before completing a lesson, review it from the perspective of a developer encountering the topic for the first time.

Verify that:

- The opening establishes a correct mental model.
- Essential terminology is explained before it is used.
- Sections follow a logical learning progression.
- Every important rule explains both what happens and why.
- Examples demonstrate distinct and meaningful behaviors.
- Non-obvious results are explained rather than merely shown.
- Important practical limitations and common mistakes are covered.
- Comparisons clarify differences instead of repeating information.
- The README is understandable without opening other files.
- Technical claims and expected outputs are correct.
- Adjacent topics do not distract from the main lesson.
- There are no unnecessary explanations, duplicated rules, or filler.

Revise the lesson if any of these conditions are not met.

Prioritize correctness, clarity, logical progression, and practical understanding over brevity or completeness
for its own sake.

### Visual readability and information hierarchy

Make lessons easy to read, scan, and review before technical interviews.

- Use bold text to highlight essential rules, counterintuitive behavior, and important distinctions.
- Use short bullet lists when explaining the properties, components, or results of an operation.
- Prefer compact comparison tables when multiple operations solve similar problems.
- Explain relevant time complexity, memory usage, and allocation behavior when they affect the choice of approach.
- Distinguish guaranteed behavior from compiler optimizations and implementation details.
- Use simple text diagrams when they clarify memory layout, references, execution order, or data representation.
- Place important warnings near the examples they affect.
- Avoid long sequences of uniformly formatted paragraphs.
- Do not use decorative formatting that adds no educational value.

Make the lesson suitable both for initial learning and quick revision without duplicating its content.

## Commits and Pull Requests

Write short commit subjects in Russian, for example:

- `Добавить урок 012 о функциях и параметрах`
- `Уточнить именованные результаты в уроке 012`

In pull requests, name the lesson, summarize the changes, list the performed checks, and link a related issue
if one exists.

Report any checks that could not be performed and explain why.
