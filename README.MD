# MOCKETER
> Mock made fast

## What is Moketer?
Moketer is a proyect designed to stop you from creating "quick-start" apis to validate external tools.
Use cases:
- Test an external tool that requires any endpoint to understand/test it's functionality (yes kong, I'm speaking about you).
- Validating that I'm hitting the right endpoints
- Validate if I'm sending the right parameters (for those cases where you just aren't sure if your fronted is sending what it's supossed to send and the api doesn't give you much information)

### TODO:
- Define a yaml/json spec to allow the users to define custom codes and respondes for specifc endpoints
- Add Compiled version

## Installation
```
git clone https://github.com/stashconsulting/moketer.git
cd moketer
go build -o moketer
sudo chmod u+x moketer
```

## USAGE

### Example
Start the webserver, appends the header and cookies to the response and also prints it to the console <br/>
```
./moketer --host=0.0.0.0 --port=8080 --uri --headers --cookies --std
curl 0.0.0.0:8080
```

### Params
Each part of the request is constomizable, so you always get what you specify

<pre>
-basicAuth
    Request BasicAuth
-body
    Request body (default true)
-cookies
    Request cookies
-headers
    Request headers
-host string
    Host [Required]
-port string
    Port (default "8080")
-q
    Removes log prefixes
-std
    Prints the request response in the terminal
-uri
    Request URI
</pre>
