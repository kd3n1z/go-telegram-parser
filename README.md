# go-telegram-parser

A lightweight parser for validating and extracting data from Telegram Web App's initialization data. This library ensures data validation using HMAC-SHA256 hashing.

## Usage

```go
package main

import (
	"fmt"
	telegramparser "github.com/kd3n1z/go-telegram-parser"
)

func main() {
	botToken := "your-bot-token-here"

	parser := telegramparser.CreateParser(botToken)

	initData, err := parser.Parse("query_id=123&auth_date=1234567890&hash=abcdef...")

	if err == nil {
		fmt.Println("Validated data:")
		fmt.Println("QueryId:", initData.QueryId)
		fmt.Println("Hash:", initData.Hash)
		fmt.Println("User.FirstName:", initData.User.FirstName)
		// ...
	} else {
		fmt.Println("Validation failed.")
	}
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
