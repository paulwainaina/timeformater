#TimeFormater

This is a go module that is used to convert string date/datetime to time in go lang. This can be useful when you are dealing with data from a form input of type date/datetime. This module can be used when implementing UnmarshalJSON interface.

    func (member *Member) UnmarshalJSON(data []byte) error {
        var jsonData map[string]interface{}
        err := json.Unmarshal(data, &jsonData)
        if err != nil {
            return err
        }
        for k, v := range jsonData {
            case "dateofbirth":
                {	t,er:=timeformater.ConvertDateToTime(v.(string),"-")
                    if er!=nil{
                        return er
                    }				
                    member.DateofBirth=t
                }
        }
    }

The method ConvertDateToTime takes in two parameters.
1. The date string
2. The delimeter. A blank string will default to  '-'

