# API DOCS

## User Logout

## Request Header
```bash
"Authorization: Bearer xxx"
```

## Success : 200
```bash
{
	"status": "success",
	"response": "logout"
}
```

## Error token : 400
```bash
{
	"status":  "failed",
	"errors":  "unauthorization"
}
```