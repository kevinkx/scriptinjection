package main
// route: {{ perfAPI  }}/v1/objectives/batch
import (
    "fmt"
	"strings"
    "github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {
    xlsx, err := excelize.OpenFile("b2c.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // // Get value from cell by given worksheet name and axis.
    // cell := xlsx.GetCellValue("Human Capital - DONE", "B2")
    // fmt.Println(cell)
	
	// Create result variable
	var result string

	// Add header for the result
	result = `{"objectives":[`

	// Initialize column name variable
	var column []string

	// Get all the rows in the Sheet.
	rows := xlsx.GetRows("B2C Product") //change this

	//initialize hashmap for storing data from excel
	dataMap := make([]map[string]string, 0)
	// loop pertama dari atas ke bawah, loop kedua dari kiri ke kanan
	for i, row := range rows {
		data := make(map[string]string)
		for j, colCell := range row {
			//if i = 0, then its `judul kolom`, else it is the data.
			if i == 0 {
				//check and transform column name into known request param
				if strings.Contains(strings.ToLower(colCell),"dummy id"){column = append(column, "fake_id")
				}else if strings.Contains(strings.ToLower(colCell),"dummy parent id"){column = append(column, "fake_parent_id")
				}else if strings.Contains(strings.ToLower(colCell),"owner id"){column = append(column, "assignee")
				}else if strings.Contains(strings.ToLower(colCell),"reviewer id"){column = append(column, "assigner")
				}else if strings.Contains(strings.ToLower(colCell),"followers"){column = append(column, "follower")
				}else if strings.Contains(strings.ToLower(colCell),"goal title"){column = append(column, "name")
				}else if strings.Contains(strings.ToLower(colCell),"metric"){column = append(column, "metric")
				}else if strings.Contains(strings.ToLower(colCell),"detailed expectation"){column = append(column, "desc")
				}else if strings.Contains(strings.ToLower(colCell),"target"){column = append(column, "target")
				}else if strings.Contains(strings.ToLower(colCell),"start date"){column = append(column, "start_date")
				}else if strings.Contains(strings.ToLower(colCell),"due date"){column = append(column, "due_date")
				}else if strings.Contains(strings.ToLower(colCell),"roll up"){column = append(column, "roll_up")
				}else if strings.Contains(strings.ToLower(colCell),"weight"){column = append(column, "weight")
				}else if strings.Contains(strings.ToLower(colCell),"goal type"){column = append(column, "goal_type")
				}else if strings.Contains(strings.ToLower(colCell),"objective type"){column = append(column, "objective_type")
				}else if strings.Contains(strings.ToLower(colCell),"objective repetetition"){column = append(column, "objective_repetetition")
				}else {column = append(column, "no need")}
			}else {
				if column[j]!="no need" {
					data[column[j]] = colCell
				}
			}
		}
		if i>0{
			dataMap = append(dataMap, data)
		}
	}

	//build the request

	for i, row := range dataMap {
		dataString := "{"

		dataString += `"name": "` + row["name"] + `",`

		if row["objective_type"] != "" {
			dataString += `"type": "` + strings.ToLower(row["objective_type"]) + `",`
		}else{
			dataString += `"type": "goal",`
		}


		if row["fake_id"] != "" {	
			floatNumber, _ := strconv.ParseFloat(row["fake_id"], 64)
			dataString += `"fake_id": ` + strconv.Itoa(int(floatNumber)) + ","
		}else{
			dataString += `"fake_id": null,`
		}

		if row["fake_parent_id"] != "" {
			floatNumber, _ := strconv.ParseFloat(row["fake_parent_id"], 64)
			dataString += `"fake_parent_id": ` + strconv.Itoa(int(floatNumber)) + ","
		}else{
			dataString += `"fake_parent_id": null,`
		}

		if row["desc"] != "" {
			floatNumber, _ := strconv.ParseFloat(row["desc"], 64)
			dataString += `"description": "` + strconv.Itoa(int(floatNumber)) + `",`
		}else{
			dataString += `"description": null,`
		}

		date := strings.Split(row["start_date"], "/")
		dataString += `"start_date": "20`+date[2]+"-"+date[0]+"-"+date[1]+`T07:00:00+07:00",`
		//save start_date for recurrence
		start_date:=`"20`+date[2]+"-"+date[0]+"-"+date[1]+`T07:00:00+07:00"`

		date = strings.Split(row["due_date"], "/")
		dataString += `"due_date": "20`+date[2]+"-"+date[0]+"-"+date[1]+`T23:59:59+07:00",`

		if row["weight"] != "" {
			floatNumber, _ := strconv.ParseFloat(row["weight"], 64)
			dataString += `"weight": ` + strconv.Itoa(int(floatNumber)) + ","
		}else{
			dataString += `"weight": 0,`
		}

		//***
		dataString += `"comment_template_id": null,`

		dataString +=`"involvements": [`
		if row["assignee"] != ""{
			floatNumber, _ := strconv.ParseFloat(row["assignee"], 64)
			dataString += `{"role": "assignee","user_id": `+strconv.Itoa(int(floatNumber))+`}`
		}
		if row["assigner"] != ""{
			floatNumber, _ := strconv.ParseFloat(row["assigner"], 64)
			dataString += `,{"role": "assigner","user_id": `+strconv.Itoa(int(floatNumber))+`}`
		}
		
		//***
		if row["follower"] != ""{
			dataString += `,{"role": "follower","user_id": `+row["follower"]+`}`
		}
		dataString +=`],`

		if row["metric"]!="" && row["target"]!=""{
			dataString +=`"measurement": {`
			unitId:="0"
			if row["metric"] == "$"{unitId = "1"}
			if row["metric"] == "%"{unitId = "2"}
			if row["metric"] == "#"{unitId = "3"}
			if row["metric"] == "IDR" || row["metric"] == "Rp"{unitId = "4"}
			dataString += `"unit_id": `+unitId+","
			dataString += `"starting_value": 0,`
			floatNumber, _ := strconv.ParseFloat(row["target"], 64)
			dataString += `"target_value": `+strconv.Itoa(int(floatNumber))+","
			//***
			if row["roll_up"] != ""{
				rollUp:=""
				if row["roll_up"] == "auto average"{rollUp = "average"}
				if row["roll_up"] == "auto sum"{rollUp = "auto"}
				if row["roll_up"] == "manual"{rollUp = "manual"}
				dataString += `"roll_up": "`+rollUp+`"`
			}
			dataString +=`},`
		}

		if row["goal_type"]!=""{
			goalType:= ""
			if strings.Contains(row["goal_type"],"High"){
				goalType="1"
			}
			if strings.Contains(row["goal_type"],"Medium"){
				goalType="2"
			}
			if strings.Contains(row["goal_type"],"Low"){
				goalType="3"
			}
			dataString += `"objective_category_id": ` + goalType
		}
		

		//***
		if row["objective_repetetition"] != ""{
			every:= "1"
			title:=strings.Title(row["objective_repetetition"])
			recurrenceType:=""
			dataString +=`,"recurrence": {`
			dataString += `"date": `+start_date+","
			if strings.Contains(strings.ToLower(row["objective_repetetition"]),"quarter"){
				every = "3"
				title = "Monthly"
			}
			dataString += `"display": "`+title+`",`
			if strings.Contains(strings.ToLower(row["objective_repetetition"]),"bi"){every = "2"}
			dataString += `"every": `+every+","
			dataString += `"state": "active",`
			if strings.ToLower(row["objective_repetetition"]) == "daily"{recurrenceType = "day"}
			if strings.ToLower(row["objective_repetetition"]) == "weekly"{recurrenceType = "week"}
			if strings.ToLower(row["objective_repetetition"]) == "biweekly"{recurrenceType = "week"}
			if strings.ToLower(row["objective_repetetition"]) == "monthly"{recurrenceType = "month"}
			if strings.ToLower(row["objective_repetetition"]) == "yearly"{recurrenceType = "year"}
			if strings.ToLower(row["objective_repetetition"]) == "quarterly"{recurrenceType = "month"}
			dataString += `"type": "`+recurrenceType+`"`
			dataString +=`}`
		}


		dataString += "}"
		if i != len(dataMap)-1{
			dataString += ","
		}
		result += dataString
	}

	//Add footer for result
	result += "]}"

	fmt.Println(result)
}