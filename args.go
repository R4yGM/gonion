package gonion

// see https://metrics.torproject.org/onionoo.html#parameters for further info

type Params struct {

	Type string

	Running bool 

	Search string

	Lookup string

	Country string

	As int

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

	Order string

	Offset int

	Limit int
}


func (args Params) QueryParams() url.Values {
    q := make(url.Values)

    if args.LatLon != nil {
        q.Add("lat", strconv.FormatFloat(args.LatLon.Lat, 'f', -1, 64))
        q.Add("lon", strconv.FormatFloat(args.LatLon.Lon, 'f', -1, 64))
    }

    if args.LocationID != "" {
        q.Add("location_id", args.LocationID)
    }
    if args.UnitSystem != "" {
        q.Add("unit_system", args.UnitSystem)
    }

    if len(args.Fields) > 0 {
        q.Add("fields", strings.Join(args.Fields, ","))
    }

    if !args.StartTime.IsZero() {
        q.Add("start_time", args.StartTime.Format(time.RFC3339))
    }
    if !args.EndTime.IsZero() {
        q.Add("end_time", args.EndTime.Format(time.RFC3339))
    }

    return q
}
