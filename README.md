# goadvanta

A Golang client for the [Advanta SMS API](https://www.advantasms.com/bulksms-api).

Functionality includes:
- Sending single/bulk SMS
- Schedule for later delivery
- Sending bulk SMS
- Check SMS delivery status
- Check remaining balance

## Download

```bash
go get github.com/IamFaizanKhalid/goadvanta
```

## Sample Usage

```go
package main

import (
	"log"
	"github.com/IamFaizanKhalid/goadvanta"
)

func main() {
	c := goadvanta.NewClient("apiKey", "partnerID", "senderID")

	resp, err := c.SendSMS("+254774894426", "Hello, World!")
	if err != nil {
		log.Fatalln(err)
	}
	
	if resp.Success {
		log.Println(resp.MessageID)
	} else {
		log.Println(resp.Error)
	}
}
```

## License

[MIT License](https://github.com/IamFaizanKhalid/nishan-go/blob/master/LICENSE)
