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

#### Date
Dont forget to change format in date column (Format -> Number -> Plain Text)

and you must modify this code 
```python
		date := strings.Split(row["start_date"], "/")
		dataString += `"start_date": "20`+date[2]+"-"+date[0]+"-"+date[1]+`T07:00:00+07:00",`
		//save start_date for recurrence
		start_date:=`"20`+date[2]+"-"+date[0]+"-"+date[1]+`T07:00:00+07:00"`

		date = strings.Split(row["due_date"], "/")
		dataString += `"due_date": "20`+date[2]+"-"+date[0]+"-"+date[1]+`T23:59:59+07:00",`
```

you must adjust the array of date with date format in excel. that code valid with case MM/DD/YY.
