# the-rundown-api-go
Go library to use The Rundown Sports Betting API

## Example Use

```
import (
    therundown "github.com/joelhill/the-rundown-api-go"
)

xRapidAPIKey := "Your Rapid API Key asfafasdfasdfasdfasasdfsadfasdfsd"
config := therundown.NewConfig(xRapidAPIKey)
client := therundown.NewService(config)
options := client.NewEventBySportOptions()
rundown, statusCode, err := client.EventBySport(options)
if err != nil {
    log.Println("something went wrong")
}
```