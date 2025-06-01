# API DOCS

## USER LOGIN /login

## Request Body
```bash
{
	"username": "",
	"password" : ""
}
```
- password max 10 character

## Success : 200
```bash
{
	"status": "success",
	"token": ""
}
```

## Error Username or Password : 400
```bash
{
	"status":  "failed",
	"errors":  "Login failed : Check your username or password"
}
```

## Error empty request body : 400
```bash
{
	"status":  "failed",
	"errors":  "No JSON data"
}
```