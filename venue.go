package foursquarego

type Venue struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Contact       Contact    `json:"contact"`
	Location      Location   `json:"location"`
	CanonicalUrl  string     `json:"canonicalUrl"`
	Categories    []Category `json:"categories"`
	Verified      bool       `json:"verified"`
	Stats         Stats      `json:"stats"`
	Url           string     `json:"url"`
	Price         Price      `json:"price"`
	HasMenu       bool       `json:"hasMenu"`
	Likes         Likes      `json:"likes"`
	Like          bool       `json:"like"`
	Dislike       bool       `json:"dislike"`
	Ok            bool       `json:"ok"`
	Rating        float32    `json:"rating"`
	RatingColor   string     `json:"ratingColor"`
	RatingSignals int        `json:"ratingSignals"`
	Menu          Menu       `json:"menu"`
	Specials      Specials   `json:"specials"`
	Photos        Photos     `json:"photos"`
	Reasons       Reasons    `json:"reasons"`
	CreatedAt     int        `json:"createdAt"`
	Tips          Tips       `json:"tips"`
	Tags          []string   `json:"tags"`
	ShortUrl      string     `json:"shortUrl"`
	TimeZone      string     `json:"timeZone"`
	Listed        Listed     `json:"listed"`
	Phrases       []Phrase   `json:"phrases"`
	Hours         Hours      `json:"hours"`
	Popular       Popular    `json:"popular"`
	PageUpdates   Omit       `json:"-"`
	Inbox         Omit       `json:"-"`
	VenueChains   Omit       `json:"-"`
	Attributes    Attributes `json:"attributes"`
	BestPhoto     Photo      `json:"bestPhoto"`
}

type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
}

type Location struct {
	Address          string  `json:"address"`
	CrossStreet      string  `json:"crossStreet"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	PostalCode       string  `json:"postalCode"`
	Cc               string  `json:"cc"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	Country          string  `json:"country"`
	FormattedAddress string  `json:"formattedAddress"`
}

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       Icon   `json:"icon"`
	Primary    bool   `json:"primary"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type Price struct {
	Tier     int    `json:"tier"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
}

type Likes struct {
	Count   int    `json:"count"`
	Groups  Omit   `json:"-"` //TODO: take care fo this later
	Summary string `json:"summary"`
}

type Menu struct {
	Type      string `json:"type"`
	Label     string `json:"lablel"`
	Anchor    string `json:"anchor"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobileUrl"`
}

type Specials struct {
	Count int  `json:"count"`
	Items Omit `json:"-"` //TODO  take care fo this later
}

type Photos struct {
	Count  int          `json:"count"`
	Groups []PhotoGroup `json:"groups"`
}

type PhotosResponse struct {
	Count        int     `json:"count"`
	Items        []Photo `json:"items"`
	DupesRemoved int     `json:"dupesRemoved"`
}

type PhotoGroup struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

type Photo struct {
	ID         string      `json:"id"`
	CreatedAt  int         `json:"createdAt"`
	Source     PhotoSource `json:"source"`
	Prefix     string      `json:"prefix"`
	Suffix     string      `json:"suffix"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	User       User        `json:"user"`
	Visibility bool        `json:"visibility"`
}

type PhotoSource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	Photo     UserPhoto `json:"photo"`
}

type UserPhoto struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type HereNow struct {
	Count   int    `json:"count"`
	Summary string `json:"summary"`
	Groups  Omit   `json:"groups"` //TODO: take care fo this later
}

type Reasons struct {
	Count int      `json:"count"`
	Items []Reason `json:"items"`
}

type Reason struct {
	Summary    string `json:"summary"`
	Type       string `json:"type"`
	ReasonName string `json:"reasonName"`
}

type Tips struct {
	Count  int        `json:"count"`
	Groups []TipGroup `json:"groups"`
}

type TipGroup struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Items []Tip  `json:"items"`
}

type Tip struct {
	ID           string `json:"id"`
	CreatedAt    int    `json:"createdAt"`
	Text         string `json:"text"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	CanonicalUrl string `json:"canonicalurl"`
	Photo        Photo  `json:"photo"`
	PhotoUrl     string `json:"photoUrl"`
	Flags        Omit   `json:"flags"` //TODO:  take care fo this later
	Likes        Likes  `json:"likes"`
	Like         bool   `json:"like"`
	LogView      bool   `json:"logView"`
	EditedAt     int    `json:"editedAt"`
	Todo         Omit   `json:"todo"` //TODO:  take care fo this later
	User         User   `json:"user"`
}

type Listed struct {
	Count  int         `json:"count"`
	Groups []ListGroup `json:"groups"`
}

type ListGroup struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Items []List `json:"items"`
}

type List struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	Type          string      `json:"type"`
	User          User        `json:"user"`
	Editable      bool        `json:"editable"`
	Public        bool        `json:"public"`
	Collaborative bool        `json:"collaborative"`
	Url           string      `json:"url"`
	CanonicalUrl  string      `json:"canonicalUrl"`
	CreatedAt     int         `json:"createdAt"`
	UpdatedAt     int         `json:"updatedAt"`
	Photo         Photo       `json:"photo"`
	Followers     Followers   `json:"followers"`
	ListItems     []ListItems `json:"listItems"`
}

type Followers struct {
	Count int `json:"count"`
}

type ListItems struct {
	Count int        `json:"count"`
	Items []ListItem `json:"items"`
}

type ListItem struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Tip       Tip    `json:"tip"`
	Photo     Photo  `json:"photo"`
}

type Phrase struct {
	Phrase string `json:"phrase"`
	Sample Sample `json:"sample"`
	Count  int    `json:"count"`
}

type Sample struct {
	Entities []Entitie `json:"entities"`
	Text     string    `json:"text"`
	Count    int       `json:"count"`
}

type Entitie struct {
	Indicies []int  `json:"indicies"`
	Type     string `json:"type"`
}

type Hours struct {
	Status     string      `json:"status"`
	IsOpen     bool        `json:"isOpen"`
	TimeFrames []TimeFrame `json:"timeFrames"`
}

type TimeFrame struct {
	Days          string `json:"days"`
	IncludesToday bool   `json:"includesToday"`
	Open          []Open `json:"open"`
	Segments      Omit   `json:"-"` //TODO:  take care fo this later
}

type Open struct {
	RenderedTime string `json:"renderedTime"`
}

type Popular struct {
	Status     string      `json:"status"`
	IsOpen     bool        `json:"isOpen"`
	TimeFrames []TimeFrame `json:"timeFrames"`
}

type Attributes struct {
	Groups []Attribute `json:"groups"`
}

type Attribute struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Count   int    `json:"count"`
	Items   []Omit `json:"-"`
}
