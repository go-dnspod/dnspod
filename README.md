# dnspod
Golang bindings for DNSPOD API

## Install
    go get go-dnspod/dnspod
## Usage

```
package main

import (
	"fmt"
	"go-dnspod/dnspod"
)

func main() {
	argToken := "YOU_DNSPOD_API_TOKEN"
	myDomain:="yourdomain.com"
	
	dns := dnspod.NewDnspod(argToken)
	//List all your domain
	myDomainList, err := dns.Domain.List()
	if err != nil {
		panic(err)
	}
	fmt.Println("MyDomains:", myDomainList)
	//Create a new A record test.yourdomain.com -> 1.2.3.4
	recordID, err := dns.Record.Create(
		myDomain,
		dnspod.RTypeA,
		"1.2.3.4",
		dnspod.RecordOpt{SubDomain: "test"},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create success! Record ID:%d\r\n", recordID)
}
```
##More Examples
You can find a complete DDNS client source code in **example/ddns** directory.
##License
This library is under the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)


