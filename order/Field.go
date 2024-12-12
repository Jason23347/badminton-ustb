package order

import (
	"fmt"
)

const (
	FIELD_BADMINTON = "Y"
	FIELD_PINGPANG  = "M"
)

type Field struct {
	Type      string // FIELD_* constants
	Number    int
	StartTime string
	EndTime   string
	Price     int
}

type IntervalForm struct {
	StartTime string `json:"timemc"`
	EndTime   string `json:"endtimemc"`

	Price1  int
	Price2  int
	Price3  int
	Price4  int
	Price5  int
	Price6  int
	Price7  int
	Price8  int
	Price9  int
	Price10 int
	Price11 int
	Price12 int
	Price13 int
	Price14 int
	Price15 int
	Price16 int
	Price17 int
	Price18 int
	Price19 int
	Price20 int
	LXBH1   string
	LXBH2   string
	LXBH3   string
	LXBH4   string
	LXBH5   string
	LXBH6   string
	LXBH7   string
	LXBH8   string
	LXBH9   string
	LXBH10  string
	LXBH11  string
	LXBH12  string
	LXBH13  string
	LXBH14  string
	LXBH15  string
	LXBH16  string
	LXBH17  string
	LXBH18  string
	LXBH19  string
	LXBH20  string
	CDBH1   int
	CDBH2   int
	CDBH3   int
	CDBH4   int
	CDBH5   int
	CDBH6   int
	CDBH7   int
	CDBH8   int
	CDBH9   int
	CDBH10  int
	CDBH11  int
	CDBH12  int
	CDBH13  int
	CDBH14  int
	CDBH15  int
	CDBH16  int
	CDBH17  int
	CDBH18  int
	CDBH19  int
	CDBH20  int
}

func (f *Field) String() string {
	fieldStr := fmt.Sprintf("%s:%d,%s-%s;", f.Type, f.Number, f.StartTime, f.EndTime)

	return fieldStr
}

// Convert to []Field
func (rf *IntervalForm) FieldList() []Field {
	fieldList := []Field{
		{Type: rf.LXBH1, Number: rf.CDBH1, Price: rf.Price1, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH2, Number: rf.CDBH2, Price: rf.Price2, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH3, Number: rf.CDBH3, Price: rf.Price3, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH4, Number: rf.CDBH4, Price: rf.Price4, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH5, Number: rf.CDBH5, Price: rf.Price5, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH6, Number: rf.CDBH6, Price: rf.Price6, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH7, Number: rf.CDBH7, Price: rf.Price7, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH8, Number: rf.CDBH8, Price: rf.Price8, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH9, Number: rf.CDBH9, Price: rf.Price9, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH10, Number: rf.CDBH10, Price: rf.Price10, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH11, Number: rf.CDBH11, Price: rf.Price11, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH12, Number: rf.CDBH12, Price: rf.Price12, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH13, Number: rf.CDBH13, Price: rf.Price13, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH14, Number: rf.CDBH14, Price: rf.Price14, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH15, Number: rf.CDBH15, Price: rf.Price15, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH16, Number: rf.CDBH16, Price: rf.Price16, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH17, Number: rf.CDBH17, Price: rf.Price17, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH18, Number: rf.CDBH18, Price: rf.Price18, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH19, Number: rf.CDBH19, Price: rf.Price19, StartTime: rf.StartTime, EndTime: rf.EndTime},
		{Type: rf.LXBH20, Number: rf.CDBH20, Price: rf.Price20, StartTime: rf.StartTime, EndTime: rf.EndTime},
	}

	return fieldList
}

func GetDefaultFieldList() []Field {
	return []Field{
		{Type: FIELD_BADMINTON, Number: 1, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 2, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 3, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 4, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 5, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 6, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 7, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 8, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 9, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 10, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 11, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 12, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 13, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 14, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 15, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 16, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 17, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 18, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 19, Price: 10, StartTime: "12:00", EndTime: "13:00"},
		{Type: FIELD_BADMINTON, Number: 20, Price: 10, StartTime: "12:00", EndTime: "13:00"},
	}
}
