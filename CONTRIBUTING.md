# Contributing to go-github-apps-helpers

First off, thank you for considering contributing to go-github-apps-helpers! We appreciate your time and effort, and value your input in making this project better.

The following is a set of guidelines for contributing to go-github-apps-helpers. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Table of Contents

- [Contributing to go-github-apps-helpers](#contributing-to-go-github-apps-helpers)
  - [Table of Contents](#table-of-contents)
  - [Code of Conduct](#code-of-conduct)
  - [How to Contribute](#how-to-contribute)
    - [Reporting Bugs](#reporting-bugs)
    - [Suggesting Enhancements](#suggesting-enhancements)
    - [Pull Requests](#pull-requests)
  - [Style Guides](#style-guides)
    - [Git Commit Messages](#git-commit-messages)
    - [Golang Style Guide](#golang-style-guide)

## Code of Conduct

This project and everyone participating in it is governed by the [go-github-apps-helpers Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [example-email@example.com](mailto:example-email@example.com).

## How to Contribute

### Reporting Bugs

- Before submitting a bug report, please perform a search to see if the problem has already been reported. If it has, please add any additional information in the comments.
- When you are creating a bug report, please include as many details as possible. Fill out the required template, as the information it asks for helps us resolve issues faster.

### Suggesting Enhancements

- Before creating an enhancement suggestion, please perform a search to see if the idea has already been suggested. If it has, please add any additional information in the comments.
- When you are creating an enhancement suggestion, please include as many details as possible. Fill out the required template, as the information it asks for helps us understand your idea and make decisions.

### Pull Requests

- Fork the repository and create your branch from `main`.
- If you've added code that should be tested, add tests.
- If you've changed APIs, update the documentation.
- Ensure the test suite passes.
- Make sure your code lints.
- Make sure you run pre-commit
- Create a pull request with a clear title and description.

## Style Guides

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line
- Consider starting the commit message with an applicable emoji:
  - :art: `:art:` when improving the format/structure of the code
  - :racehorse: `:racehorse:` when improving performance
  - :non-potable_water: `:non-potable_water:` when plugging memory leaks
  - :memo: `:memo:` when writing docs
  - :bug: `:bug:` when fixing a bug
  - :fire: `:fire:` when removing code or files
  - :green_heart: `:green_heart:` when fixing the CI build
  - :white_check_mark: `:white_check_mark:` when adding tests
  - :lock: `:lock:` when dealing with security
  - :arrow_up: `:arrow_up:` when upgrading dependencies
  - :arrow_down: `:arrow_down:` when downgrading dependencies
  - :shirt: `:shirt:` when removing linter warnings

### Golang Style Guide

- Follow the Effective Go guidelines for writing idiomatic and efficient Go code.
- Format your code using gofmt or goimports.
- Organize your code using packages, and separate your library code from executable code.
- Use clear and concise variable and function names.
- Write comments for exported functions, types, and variables following the Godoc conventions.
- Keep functions small and focused on a single responsibility.
- Write tests for your code to ensure it works as expected and to prevent regressions.

By following these guidelines and best practices, you will contribute to a more maintainable, readable, and efficient codebase. Once again, thank you for your interest in contributing to go-github-apps-helpers!