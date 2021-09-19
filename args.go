package gonion

import ( 
    "net/url" 
    "strings"
    "fmt"
) 

// see https://metrics.torproject.org/onionoo.html#parameters for further info

type Params struct {

	Type string

	Running bool 

	Search string

	Lookup string

	Country string

	As string

    As_name string

	Flag string

	First_seen_days string

	Last_seen_days string

	Contact string

	Family string

	Version string 

	Os string

	Host_name string

	Recommended_version bool

	Fields []string

	Order []string

	Offset int

	Limit int
}

func valid_Order(s []string) bool {
    for _, v := range s{
        if v != "consensus_weight" || v != "first_seen" {
            return false
        }
    }
    return true
}

func (args Params) QueryParams() (url.Values, error) {
    q := make(url.Values)

    if args.Type != "" {
        if args.Type == "relay" || args.Type == "bridge"{
            q.Add("type", args.Type)
        }
    }

    q.Add("running", fmt.Sprintf("%t", args.Running))

    if args.Search != ""{
        q.Add("search", args.Search)
    }

    if args.Lookup != ""{
        if len([]rune(args.Lookup)) == 40{
            q.Add("lookup", args.Lookup)
        }else{
            return nil, fmt.Errorf("Error : invalid fingerprint on lookup parameter")
        }
    }

    if args.Country != ""{
        if len([]rune(args.Country)) == 2{
            q.Add("country", args.Country)
        }else{
            return nil, fmt.Errorf("Error : Country code cannot be more than 2 characters on country parameter")
        }
    }
    
    if args.As != ""{
        q.Add("as", args.As)
    }

    if args.As_name != ""{
        q.Add("as_name", args.As_name)
    }

    if args.Flag != ""{
        q.Add("flag", args.Flag)
    }
    
    if args.First_seen_days != ""{
        q.Add("first_seen_days", args.First_seen_days)
    }

    if args.Last_seen_days != ""{
        q.Add("last_seen_days", args.Last_seen_days)
    }

    if args.Contact != ""{
        q.Add("contact", args.Contact)
    }

    if args.Family != ""{
        if len([]rune(args.Family)) == 40{
            q.Add("family", args.Family)
        }else{
            return nil, fmt.Errorf("Error : invalid fingerprint on family parameter")
        }
    }

    if args.Version != ""{
        q.Add("version", args.Version)
    }

    if args.Os  != ""{
        q.Add("os", args.Os)
    }

    if args.Host_name != ""{
        q.Add("host_name", args.Host_name)
    }

    q.Add("recommended_version", fmt.Sprintf("%t", args.Recommended_version))

    if len(args.Fields) > 0 {
        q.Add("fields", strings.Join(args.Fields, ","))
    }

    if len(args.Order) > 0 {
        if valid_Order(args.Order){
            q.Add("order", strings.Join(args.Order, ","))
        }
    }

    if args.Offset != 0{
        q.Add("offset", fmt.Sprintf("%d", args.Offset))
    }

    if args.Limit != 0{
        q.Add("limit", fmt.Sprintf("%d", args.Limit))
    }
    
    return q, nil
}
