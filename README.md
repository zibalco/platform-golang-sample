# Zibal Payment Platform Golang Code Sample

You can use the function "postToZibal" from "functions.go" to send requests to Zibal`s platform endpoints.
Function "platformResult" can be called for printing readable messages based on the result code.

# Wallet List Sample Code
```go
// Wallet List Sample Code
var apiKey = "Enter your Access-Token"

var result = postToZibal("v1/wallet/list", "", apiKey)
// Map result to a struct to easily access parameters
var structResult map[string]interface{}
br := bytes.NewReader([]byte(result))
decodedJson := json.NewDecoder(br)
decodedJson.UseNumber()
err := decodedJson.Decode(&structResult)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(structResult)

// Access response parameters
var resultNumber = structResult["result"]
// Change response parameters types to string
resultStringValue := fmt.Sprint(resultNumber)

// Print readable messages based on response result code
fmt.Println(platformResult(resultStringValue))
```


# Wallet Balance Sample Code
```go
// Wallet Balance Sample Code
var apiKey = "Enter your Access-Token"

data := `{
    "id" : "Your wallet id"
}`

var result = postToZibal("v1/wallet/balance", data, apiKey)
// Map result to a struct to easily access parameters
var structResult map[string]interface{}
br := bytes.NewReader([]byte(result))
decodedJson := json.NewDecoder(br)
decodedJson.UseNumber()
err := decodedJson.Decode(&structResult)
if err != nil {
	fmt.Println(err)
	return
}

// Access response parameters
var resultNumber = structResult["result"]
// Change response parameters types to string
resultStringValue := fmt.Sprint(resultNumber)

// Print readable messages based on response result code
fmt.Println(platformResult(resultStringValue))
```

