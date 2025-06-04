# Learn Go

A simple Go project demonstrating package structure and basic functionality.

## Overview

This project showcases a basic Go application structure with a separate package for greeting functionality. It's designed to help understand how to organize Go code into packages and import them into the main application.

## Project Structure

```
.
├── greeting/
│   └── greeting.go    # Greeting package with greeting functionality
├── main.go           # Main application entry point
├── go.mod           # Go module definition
└── README.md        # This file
```

## Features

- Simple greeting functionality
- Demonstrates package organization in Go
- Basic example of function exports and imports

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd learn-go
   ```

2. Build and run the application:
   ```bash
   go run .
   ```

## Usage

The application will output a greeting message:

```
Hello, Gopher! Welcome to Go!
```