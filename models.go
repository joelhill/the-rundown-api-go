package therundown

type TheRunDownAPILines struct {
	Meta   Meta    `json:"meta"`  
	Events []Event `json:"events"`
}

type Event struct {
	EventID            string                `json:"event_id"`            
	SportID            int64                 `json:"sport_id"`            
	EventDate          string                `json:"event_date"`          
	RotationNumberAway int64                 `json:"rotation_number_away"`
	RotationNumberHome int64                 `json:"rotation_number_home"`
	Score              Score                 `json:"score"`               
	Teams              []Team                `json:"teams"`               
	TeamsNormalized    []TeamsNormalized     `json:"teams_normalized"`    
	LinePeriods        map[string]LinePeriod `json:"line_periods"`        
}

type LinePeriod struct {
	PeriodFullGame     Period `json:"period_full_game"`    
	PeriodFirstHalf    Period `json:"period_first_half"`   
	PeriodSecondHalf   Period `json:"period_second_half"`  
	PeriodFirstPeriod  Period `json:"period_first_period"` 
	PeriodSecondPeriod Period `json:"period_second_period"`
	PeriodThirdPeriod  Period `json:"period_third_period"` 
	PeriodFourthPeriod Period `json:"period_fourth_period"`
}

type Period struct {
	LineID    int64     `json:"line_id"`  
	Teams     []Team    `json:"teams"`    
	Moneyline Moneyline `json:"moneyline"`
	Spread    Spread    `json:"spread"`   
	Total     Total     `json:"total"`    
	Affiliate Affiliate `json:"affiliate"`
}

type Affiliate struct {
	AffiliateID   int64         `json:"affiliate_id"`  
	AffiliateName AffiliateName `json:"affiliate_name"`
	AffiliateURL  string        `json:"affiliate_url"` 
}

type Moneyline struct {
	LineID             int64   `json:"line_id"`             
	MoneylineAway      float64 `json:"moneyline_away"`      
	MoneylineAwayDelta float64 `json:"moneyline_away_delta"`
	MoneylineHome      float64 `json:"moneyline_home"`      
	MoneylineHomeDelta float64 `json:"moneyline_home_delta"`
	DateUpdated        string  `json:"date_updated"`        
	Format             Format  `json:"format"`              
}

type Spread struct {
	LineID                    int64   `json:"line_id"`                      
	PointSpreadAway           float64 `json:"point_spread_away"`            
	PointSpreadAwayDelta      float64 `json:"point_spread_away_delta"`      
	PointSpreadHome           float64 `json:"point_spread_home"`            
	PointSpreadHomeDelta      float64 `json:"point_spread_home_delta"`      
	PointSpreadAwayMoney      float64 `json:"point_spread_away_money"`      
	PointSpreadAwayMoneyDelta float64 `json:"point_spread_away_money_delta"`
	PointSpreadHomeMoney      float64 `json:"point_spread_home_money"`      
	PointSpreadHomeMoneyDelta float64 `json:"point_spread_home_money_delta"`
	DateUpdated               string  `json:"date_updated"`                 
	Format                    Format  `json:"format"`                       
}

type Team struct {
	TeamID           int64  `json:"team_id"`           
	TeamNormalizedID int64  `json:"team_normalized_id"`
	IsAway           bool   `json:"is_away"`           
	IsHome           bool   `json:"is_home"`           
	Name             string `json:"name"`              
}

type Total struct {
	LineID               int64   `json:"line_id"`                
	TotalOver            float64 `json:"total_over"`             
	TotalOverDelta       float64 `json:"total_over_delta"`       
	TotalUnder           float64 `json:"total_under"`            
	TotalUnderDelta      float64 `json:"total_under_delta"`      
	TotalOverMoney       float64 `json:"total_over_money"`       
	TotalOverMoneyDelta  float64 `json:"total_over_money_delta"` 
	TotalUnderMoney      float64 `json:"total_under_money"`      
	TotalUnderMoneyDelta float64 `json:"total_under_money_delta"`
	DateUpdated          string  `json:"date_updated"`           
	Format               Format  `json:"format"`                 
}

type Score struct {
	EventID           string      `json:"event_id"`            
	EventStatus       EventStatus `json:"event_status"`        
	ScoreAway         int64       `json:"score_away"`          
	ScoreHome         int64       `json:"score_home"`          
	WinnerAway        int64       `json:"winner_away"`         
	WinnerHome        int64       `json:"winner_home"`         
	ScoreAwayByPeriod interface{} `json:"score_away_by_period"`
	ScoreHomeByPeriod interface{} `json:"score_home_by_period"`
	VenueName         string      `json:"venue_name"`          
	VenueLocation     string      `json:"venue_location"`      
	GameClock         int64       `json:"game_clock"`          
	DisplayClock      string      `json:"display_clock"`       
	GamePeriod        int64       `json:"game_period"`         
	Broadcast         Broadcast   `json:"broadcast"`           
	EventStatusDetail string      `json:"event_status_detail"` 
}

type TeamsNormalized struct {
	TeamID       int64  `json:"team_id"`     
	Name         string `json:"name"`        
	Mascot       string `json:"mascot"`      
	Abbreviation string `json:"abbreviation"`
	IsAway       bool   `json:"is_away"`     
	IsHome       bool   `json:"is_home"`     
	Ranking      string `json:"ranking"`     
	Record       string `json:"record"`      
}

type Meta struct {
	DeltaLastID string `json:"delta_last_id"`
}

type AffiliateName string
const (
	BetOnline AffiliateName = "BetOnline"
	Betcris AffiliateName = "betcris"
	Bookmaker AffiliateName = "Bookmaker"
	Intertops AffiliateName = "Intertops"
	JustBet AffiliateName = "JustBet"
	LowVig AffiliateName = "LowVig"
	Matchbook AffiliateName = "Matchbook"
	Pinnacle AffiliateName = "Pinnacle"
	RedZone AffiliateName = "RedZone"
	SportsBetting AffiliateName = "SportsBetting"
	The5Dimes AffiliateName = "5Dimes"
	TigerGaming AffiliateName = "TigerGaming"
)

type Format string
const (
	American Format = "American"
)

type Broadcast string
const (
	BroadcastESPN Broadcast = "ESPN"
	Empty Broadcast = ""
	Espn Broadcast = "ESPN+"
)

type EventStatus string
const (
	StatusScheduled EventStatus = "STATUS_SCHEDULED"
)
