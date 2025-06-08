# API DOCS

## Register /register/create

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Request Body
```bash
{
  "treatment_number": "20250806000002",
  "register_number": "001",
  "register_date": "2025-06-08",
  "medical_record": "000002",
  "payment_method": "UMUM",
  "policlinic": "BED",
  "doctor": "DR000001"
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

## Get Regiter data /register/get

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Params
- date : registration date
- limit : limit data
- search : search data from medical_record or patient name

### Success : 200
```bash
[
    {
        "treatment_number": "20250806000001",
        "register_number": "001",
        "register_date": "2025-06-08",
        "medical_record": "000002",
        "name": "Dummy 2",
        "gender": "laki - laki",
        "payment_method": "UMUM",
        "policlinic_id": "BED",
        "policlinic_name": "POLIKLINIK BEDAH",
        "doctor_id": "DR000001",
        "doctor_name": "dr. Albert, Sp.B"
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

<!-- ## Update Patient /patient/update

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
``` -->

## Delete Register /register/delete?treatmentNumber=

### Request Header
```bash
"Authorization: Bearer xxx"
```

### Params
- treatmentNumber : data treatment_number patient

### Success : 200
```bash
{
	"status": "success",
	"response": "deleted"
}
```

### Error Method Not Allowed : 400
```bash
{
    "status": "failed",
    "errors": "method not allowed"
}
```