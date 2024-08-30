# Go Stack Implementation

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Introduction

Welcome to the Go Stack Implementation! 🎉

This repository contains a custom implementation of the stack data structure using the Go programming language. The stack operates on the Last-In, First-Out (LIFO) principle, making it a fundamental tool in many algorithms and software design patterns.

## Features

- **Dynamic Resizing**: Automatically adjusts the size of the underlying storage to accommodate more elements.
- **Error Handling**: Provides clear error messages for operations on an empty stack.
- **Easy Integration**: Lightweight and simple to integrate into your projects.

## Getting Started

### Prerequisites

- Go 1.18 or later installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

1. Clone the repository:

    ```bash
    git clone [https://github.com/yourusername/go-stack-implementation.git](https://github.com/WatchJani/stack.git)
    cd go-stack-implementation
    ```

2. Run tests to ensure everything is working:

    ```bash
    go test ./...
    ```

### Usage

Here is a simple example of how to use the stack:

```go
package main

import (
    "fmt"
    "log"
    "github.com/WatchJani/stack"
)

func main() {
    s := stack.NewStack()

    s.Push(10)
    s.Push(20)
    s.Push(30)

    top, err := s.Peek()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Top element:", top)

    popped, err := s.Pop()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Popped element:", popped)

    fmt.Println("Is stack empty?", s.IsEmpty())
}
```
