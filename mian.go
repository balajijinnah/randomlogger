package main

import (
	"fmt"
	"time"

	"github.com/bxcodec/faker"
)

func main() {
	ingested := 0
	// Stop the service after 10 gig.
	ingestionLimit := 10 * 1024 * 1024 * 1024
	for {
		name := faker.FirstName()
		accountNo := faker.CCNumber()
		amount := faker.AmountWithCurrency()
		destName := faker.FirstName()
		line := fmt.Sprintf("%s with %s sent to %s for %s", name, accountNo, destName, amount)
		ingested += len(line)
		fmt.Println(line)
		if ingested > ingestionLimit {
			// break the limit.
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
