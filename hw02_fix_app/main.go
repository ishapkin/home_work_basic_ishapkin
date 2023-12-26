package init

import (
	"fmt"

	"github.com/ishapkin/home_work_basic_ishapkin/hw02_fix_app/printer"
	"github.com/ishapkin/home_work_basic_ishapkin/hw02_fix_app/reader"
	"github.com/ishapkin/home_work_basic_ishapkin/hw02_fix_app/types"
)

func init() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
