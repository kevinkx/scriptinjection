# Script Injection
This is script for building goal injection request body automatically from .xlsx (excel)

## Installation

Install golang => https://golang.org/doc/install

if you done installing golang, clone this repo and go to the directory. then type this
```bash
go mod init
go mod tidy
```

## Usage

simply put the xlsx file in the same directory (like b2c.xlsx and example.xlsx), then just type this

```python
go run script_injection.go
```

#### Info

change this line with your excel file name
```python
 xlsx, err := excelize.OpenFile("<your file name>.xlsx")
```

change this line with your sheet name
```python
rows := xlsx.GetRows("Sheet1")
```

