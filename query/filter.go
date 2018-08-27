package query

import (
	"github.com/golang/glog"
	"github.com/tstromberg/campwiz/result"
)

// filter applies post-fetch criteria filtering.
func filter(c Criteria, res result.Results) result.Results {
	glog.V(1).Infof("Filtering %d results ...", len(res))
	var filtered result.Results

	for _, r := range res {
		if c.IncludeGroup && r.Availability[0].Group > 0 {
			filtered = append(filtered, r)
			continue
		}
		if c.IncludeBoatIn && r.Availability[0].Boat > 0 {
			filtered = append(filtered, r)
			continue
		}
		if c.IncludeWalkIn && r.Availability[0].WalkIn > 0 {
			filtered = append(filtered, r)
			continue
		}
		if c.IncludeStandard && r.Availability[0].Standard > 0 {
			filtered = append(filtered, r)
			continue
		}
	}
	return filtered
}
