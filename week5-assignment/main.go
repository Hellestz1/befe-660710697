package main

import (
	"github.com/gin-gonic/gin"
)

type Job_Info struct {
	Job_Id    string `json:"job_id"`
	Job_Title string `json:"name"`
	Salary    string `json:"salary"`
}
type Employees struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Tel    string   `json:"tel"`
	Email  string   `json:"email"`
	Job_Id Job_Info `json:"job_id"`
}

var jobs = []Job_Info{
	{Job_Id: "J001", Job_Title: "Chef", Salary: "20000"},
	{Job_Id: "J002", Job_Title: "Sous Chef", Salary: "25000"},
	{Job_Id: "J003", Job_Title: "Head Chef", Salary: "30000"},
}

// 2. สร้างพนักงาน 10 คน สุ่ม job_id จาก job info ด้านบน
var employees = []Employees{
	{ID: "E01", Name: "Neung", Tel: "0891112222", Email: "Neung@example.com", Job_Id: jobs[0]},
	{ID: "E02", Name: "Song", Tel: "0892223333", Email: "Song@example.com", Job_Id: jobs[1]},
	{ID: "E03", Name: "Sam", Tel: "0893334444", Email: "Sam@example.com", Job_Id: jobs[0]},
	{ID: "E04", Name: "Si", Tel: "0894445555", Email: "Si@example.com", Job_Id: jobs[1]},
	{ID: "E05", Name: "Ha", Tel: "0895556666", Email: "Ha@example.com", Job_Id: jobs[0]},
	{ID: "E06", Name: "Hok", Tel: "0896667777", Email: "Hok@example.com", Job_Id: jobs[2]},
}

func getEmp(c *gin.Context) {
	id := c.Query("job_id")
	title := c.Query("job_title")
	if id != "" || title != "" {
		filter := []Employees{}
		for _, employee := range employees {
			if employee.Job_Id.Job_Id == id && id != "" {
				filter = append(filter, employee)
			} else if employee.Job_Id.Job_Title == title {
				filter = append(filter, employee)
			}
		}
		c.JSON(200, filter)
		return
	}
	c.JSON(200, employees)
}

func main() {
	r := gin.Default()
	api := r.Group("/api/v1")
	api.GET("/employee", getEmp)
	r.Run(":8080")
}
