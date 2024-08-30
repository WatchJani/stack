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
    git clone https://github.com/WatchJani/stack.git
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

## API Reference

- **Push(item int)**: Adds an item to the top of the stack.
- **Pop() (int, error)**: Removes and returns the item from the top of the stack. Returns an error if the stack is empty.
- **Peek() (int, error)**: Returns the item from the top of the stack without removing it. Returns an error if the stack is empty.
- **IsEmpty() bool**: Checks if the stack is empty.

## Contributing

Contributions are welcome! If you have any ideas, feel free to fork the repo and submit a pull request.

1. **Fork the Project**
2. **Create your Feature Branch** (`git checkout -b feature/AmazingFeature`)
3. **Commit your Changes** (`git commit -m 'Add some AmazingFeature'`)
4. **Push to the Branch** (`git push origin feature/AmazingFeature`)
5. **Open a Pull Request**

## License

Distributed under the MIT License. See `LICENSE` for more information.
