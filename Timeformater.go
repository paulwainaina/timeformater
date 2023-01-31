package timeformater

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TimeFormater struct {
}

func NewTimeFormater() *TimeFormater {
	return &TimeFormater{}
}

func (tm *TimeFormater) ConvertDateToTime(tms string, del string) (time.Time, error) {
	if del == "" {
		del = "-"
	}
	sections := strings.Split(tms, del)
	if len(sections) !=3 {
		return time.Time{}, fmt.Errorf("invalid date string format %v", sections)
	}
	year, err := strconv.ParseInt(sections[0], 10, 32)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid year format %v", sections[0])
	}
	month, err := strconv.ParseInt(sections[1], 10, 32)
	if err != nil  {
		return time.Time{}, fmt.Errorf("invalid month format %v", sections[1])
	}
	day, err := strconv.ParseInt(sections[2], 10, 32)
	if err != nil  {
		return time.Time{}, fmt.Errorf("invalid day format %v", sections[2])
	}
	switch month{
	case 1:{//jan
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month january of year %v cannot have %v days ",year,day)
		}
	}
	case 2:{//feb
		if year%4==0{
			if day>29|| day<1{
				return time.Time{}, fmt.Errorf("month february of year %v cannot have %v days ",year,day)
			}
		}else{
			if day>28|| day<1{
				return time.Time{}, fmt.Errorf("month february of year %v cannot have %v days ",year,day)
			}
		}
	}
	case 3:{//mar
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month march of year %v cannot have %v days ",year,day)
		}
	}
	case 4:{//apr
		if day>30|| day<1{
			return time.Time{}, fmt.Errorf("month april cannot have 31 days ")
		}
	}
	case 5:{//may
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month may of year %v cannot have %v days ",year,day)
		}
	}
	case 6:{//jun
		if day>30|| day<1{
			return time.Time{}, fmt.Errorf("month june cannot have 31 days ")
		}
	}
	case 7:{//jul
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month july of year %v cannot have %v days ",year,day)
		}
	}
	case 8:{//aug
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month august of year %v cannot have %v days ",year,day)
		}
	}
	case 9:{//sep
		if day>30|| day<1{
			return time.Time{}, fmt.Errorf("month september cannot have 31 days ")
		}
	}
	case 10:{//oct
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month october of year %v cannot have %v days ",year,day)
		}
	}
	case 11:{//nov
		if day>30|| day<1{
			return time.Time{}, fmt.Errorf("month november cannot have 31 days ")
		}
		
	}
	case 12:{//dec
		if day>31 || day<1{
			return time.Time{}, fmt.Errorf("month december of year %v cannot have %v days ",year,day)
		}
		}
	default:{
		return time.Time{}, fmt.Errorf("Invalid month %v ",month)
	}
	}

	return time.Date(int(year),time.Month(month),int(day),0,0,0,0,time.UTC),nil
}
