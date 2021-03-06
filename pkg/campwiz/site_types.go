package campwiz

type SiteKind string

const (
	RV                 SiteKind = "🚚"
	AccessibleRV       SiteKind = "♿🚚"
	Standard           SiteKind = "🚗"
	AccessibleStandard SiteKind = "♿🚗"
	Lodging            SiteKind = "🛏️"
	Tent               SiteKind = "⛺"

	Group      SiteKind = "🧑‍🤝‍🧑"
	Day        SiteKind = "🥪"
	Equestrian SiteKind = "🏇"
	Boat       SiteKind = "⛵"
	Walk       SiteKind = "🥾" // May be a tiny walk

	// Features
	Biking                 = 4001
	Boating                = 4002
	EquipmentRental        = 4003
	Fishing                = 4004
	Golf                   = 4005
	Hiking                 = 4006
	HorsebackRiding        = 4007
	Hunting                = 4008
	RecreationalActivities = 4009
	ScenicTrails           = 4010
	Sports                 = 4011
	Beach                  = 4012
	Winter                 = 4013
)
