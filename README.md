# Email

A simple library to parse raw email bodies and save attachments.

### Downloading
```sh
go get github.com/waqqas-abdulkareem/email
```
### Importing
```go
import(
    "github.com/waqqas-abdulkareem/email"
)
```
### Example Usage
```go
package main

import(
    "fmt"
	"time"
	"crypto/tls"
	"github.com/simia-tech/go-pop3"
	"github.com/waqqas-abdulkareem/email"
)

func main(){
    //Dial up your mail server
    client, err := pop3.Dial(
		"outlook.office365.com:995",
		pop3.UseTLS(&tls.Config{}),
		pop3.UseTimeout(timeout),
	)
	defer client.Quit()
    if err != nil{ panic(err) }
    
    //Sign in to your email
	err = client.Auth("MyUsername@outlook.com","MyPassword")
    if err != nil{ panic(err) }
    
    //get body of the earliest email in your inbox
	body, err := client.Retr(uint32(0)
	if err != nil{ panic(err) }
	
	//parse the email
	email,err := email.Parse(body)
	if err != nil{ panic(err) }
	
	if from,ok := email.Headers["From"];ok{
	    fmt.Printf("%s was the first person to email me\n",from)
	    fmt.Printf("His/Her email said:\n%s\n",email.Body)    
    }else{
        fmt.Printf("This library doesnt work")
    }
}
```
### Status

This library is work-in-progress. You may contribute if you wish to get things moving faster.

