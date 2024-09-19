## rq: simple http requests

The idea of ​​the library is to separate the user from the redundant process of working with Json and sending simple requests.

## Examples

With the net/http package you would send the request something like this:
```go
import (
    "io"
    "bytes"
    "net/http"
    "encoding/json"
)
func main() {
    dataMap := map[string]string {
        "email": "msunyc@github.com",
    }
    jsonBytes, err := json.Marshal(dataMap)
    if err != nil {
        panic(err)
    }
    req, err := http.NewRequest("POST", "google.com", bytes.NewBuffer(jsonBytes))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    jsonObject := make(map[string]any)
    err = json.Unmarshal(body, &jsonObject)
    if err != nil {
        panic(err)
    }
}
```
With this wrapper you can reduce the number of unnecessary imports and delegate the marshalling process to the library:
```go
import (
    "github.com/msuny-c/rq"
)

func main() {
    req, err := rq.NewRequest("POST", "google.com")
    if err != nil {
        panic(err)
    }
    req.Http.Header.Set("Content-Type", "application/json")
    jsonData := map[string]string {
        "email": "msunyc@github.com",
    }
    res, err := req.SendJson(jsonData)
    if err != nil {
        panic(err)
    }
    jsonObject := make(map[string]any)
    res.Json(&jsonObject)
}
```
Even simpler usage example:
```go
import (
    "github.com/msuny-c/rq"
)

func main() {
    jsonMap := make(map[string]any)
    rs, err := rq.Get("google.com")
    if err != nil {
        panic(err)
    }
    rs.Json(&jsonMap)
}
```

## Todo:
- Expand the library with advanced http capabilities.