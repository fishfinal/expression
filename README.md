# expression

<div align="center">

<img src="https://img.shields.io/badge/Go-%3E%3D1.16-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
<img src="https://img.shields.io/badge/license-Apache--2.0-blue?style=for-the-badge" alt="License" />
<img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white" alt="Build Status" />
<img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge" alt="PRs Welcome" />

</div>

A lightweight and out-of-the-box Go regular expression wrapper. This package provides convenient helper functions for validating common formats (such as email, IP addresses, domains, and phone numbers) and performing string match/replace operations.

## Installation

```shell
go get -u github.com/channellabs/expression
```

## Features

- [x] Validate IPv4 and Internal IPv4 addresses
- [x] Validate Chinese cellphone numbers (with or without +86 prefix)
- [x] Validate Chinese landline telephone numbers
- [x] Validate email addresses
- [x] Validate web domains (URLs and IPs)
- [x] Validate time strings (HH:MM format)
- [x] Simple string matching and regular expression replacement

## Quick Start

### Validators

The package provides a set of ready-to-use validators. They return `true` if the input matches the specific pattern, and `false` otherwise.

```go
package main

import (
	"fmt"
	
	"github.com/channellabs/expression"
)

func main() {
	// Email Validation
	fmt.Println(expression.IsEmail("hello@tom.com"))         // true
	fmt.Println(expression.IsEmail("hello.@foxmail.com"))    // false

	// IPv4 Validation
	fmt.Println(expression.IsIPv4("192.168.0.1"))            // true
	fmt.Println(expression.IsInternalIPv4("127.0.0.1"))      // true
	fmt.Println(expression.IsInternalIPv4("172.10.0.1"))     // false

	// Phone Number Validation (China)
	fmt.Println(expression.IsCellphoneOfChina("+86 13717995000")) // true
	fmt.Println(expression.IsTelephoneOfChina("010-92832839"))    // true

	// Web Domain & URL Validation
	fmt.Println(expression.IsWebDomain("https://www.baidu.com"))  // true
	fmt.Println(expression.IsWebDomain("http:192.168.0.1"))       // false

	// Time Validation (HH:MM)
	fmt.Println(expression.IsTimeHourPoint("23:59"))              // true
	fmt.Println(expression.IsTimeHourPoint("24:01"))              // false
}
```

### Regular Expression Utilities

If you need to match or replace arbitrary regular expressions, the package also provides general utility functions.

```go
package main

import (
	"fmt"
	
	"github.com/channellabs/expression"
)

func main() {
	// Match string
	matched, err := expression.RegularMatchString(`no`, `/usr/bin/which: no hello in ...`)
	fmt.Println(matched, err) // true, <nil>

	// Replace string
	replaced, err := expression.RegularReplaceMatch(
		`\{\{\.NAME\}\}`, 
		"The {{.NAME}} Application starting...", 
		"CallMe",
	)
	fmt.Println(replaced, err) // The CallMe Application starting..., <nil>
}
```

## API Overview

| Function | Description |
| --- | --- |
| `IsIPv4(str string) bool` | Checks if the string is a valid IPv4 address. |
| `IsInternalIPv4(str string) bool` | Checks if the string is an internal (private) IPv4 address (e.g., 127.x, 10.x, 192.168.x). |
| `IsCellphoneOfChina(str string) bool` | Checks if the string is a valid Chinese mobile phone number (supports +86 prefix). |
| `IsTelephoneOfChina(str string) bool` | Checks if the string is a valid Chinese landline telephone number. |
| `IsEmail(str string) bool` | Checks if the string is a valid email address. |
| `IsWebDomain(str string) bool` | Checks if the string is a valid web domain or IP-based URL. |
| `IsTimeHourPoint(str string) bool` | Checks if the string represents a valid time of day (HH:MM). |
| `RegularMatchString(pattern, s string) (bool, error)` | Reports whether the string `s` contains any match of the regular expression `pattern`. |
| `RegularReplaceMatch(pattern, original, replace string) (string, error)` | Replaces all matches of `pattern` in `original` with `replace`. |

## Documentation

For more detailed information, please refer to the [GoDoc](https://pkg.go.dev/github.com/channellabs/expression).

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
