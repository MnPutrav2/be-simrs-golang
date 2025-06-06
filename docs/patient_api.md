# API DOCS

## Create Patient /patient/create

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Request Body
```bash
{
  "medical_record": "000004",
  "name": "John Doe",
  "gender": "laki - laki",
  "wedding": "Married",
  "religion": "Islam",
  "education": "Bachelor",
  "birth_place": "Jakarta",
  "birth_date": "1990-01-01",
  "work": "Engineer",
  "address": "Jl. Merdeka No. 1",
  "village": 1,
  "district": 1,
  "regencie": 1,
  "province": 1,
  "nik": "1234567890123456",
  "bpjs": "9876543210",
  "phone_number": "081234567890",
  "parent_name": "Jane Doe",
  "parent_gender": "perempuan"
}
```

### Success : 200
```bash
{
    "status": "success",
    "response": "created"
}
```

### Error Duplicate Data : 400
```bash
{
    "status": "failed",
    "errors": "duplicate entry"
}
```

### Error empty request body : 400
```bash
{
	"status":  "failed",
	"errors":  "No JSON data"
}
```

### Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```

## Get Patient /patient/get

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Params
- limit : limit data
- search : search data from medical_record or patient name

### Success : 200
```bash
[
    {
        "medical_record": "000003",
        "name": "John Doe",
        "gender": "laki - laki",
        "wedding": "kawin",
        "religion": "Islam",
        "education": "Magister (S2)",
        "birth_place": "Jakarta",
        "birth_date": "1990-01-01",
        "work": "Engineer",
        "address": "Jl. Merdeka No. 1",
        "village": 1,
        "district": 1,
        "regencie": 1,
        "province": 1,
        "nik": "1234567890123456",
        "bpjs": "9876543210",
        "phone_number": "081234567890",
        "parent_name": "Jane Doe",
        "parent_gender": "perempuan"
    }
]
```

### Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```

## Update Patient /patient/update

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Request Body
```bash
{
    "medical_record" : "000002",
    "update": {
        "medical_record": "000001",
        "name": "Dummy",
        "gender": "laki - laki",
        "wedding": "Menikah",
        "religion": "Islam",
        "education": "SMA/Sederajat",
        "birth_place": "Jakarta",
        "birth_date": "1990-01-01",
        "work": "Engineer",
        "address": "Jl. Merdeka No. 1",
        "village": 1,
        "district": 1,
        "regencie": 1,
        "province": 1,
        "nik": "1234567890123456",
        "bpjs": "9876543210",
        "phone_number": "081234567890",
        "parent_name": "Jane Doe",
        "parent_gender": "Perempuan"
    }
}
```

### Success : 200
```bash
{
	"status":"success",
	"response":"updated"
}
```

### Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```

## Delete Patient /patient/delete?mr=

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Params
- mr : data medical_record patient

### Success : 200
```bash
{
	"status":"success",
	"response":"deleted"
}
```

### Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```