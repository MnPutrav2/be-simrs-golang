# API DOCS

## User Logout /user/logout

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

## Error Empty Token : 400
```bash
{
	"status":  "failed",
	"errors":  "unauthorization"
}
```

## Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```

## User Status /user/status

## Request Header
```bash
"Authorization: Bearer xxx"
```

## Success : 200
```bash
{
    "employee_id": "",
    "name": "ADMIN",
    "gender": "laki-laki"
}
```

## Error Empty Token : 400
```bash
{
	"status":  "failed",
	"errors":  "unauthorization"
}
```

## Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```

## User Pages /user/pages

## Request Header
```bash
"Authorization: Bearer xxx"
```

## Success : 200
```bash
[
    {
        "name": "Registrasi",
        "path": "/register"
    },
    {
        "name": "Labolatorium",
        "path": "/lab"
    }
]
```

## Error Empty Token : 400
```bash
{
	"status":  "failed",
	"errors":  "unauthorization"
}
```

## Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```