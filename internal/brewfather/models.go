package brewfather

// ListOptions contains pagination parameters common to all list endpoints.
type ListOptions struct {
	Limit      int
	StartAfter string
}

// ListBatchesOptions contains parameters for listing batches.
type ListBatchesOptions struct {
	ListOptions
	Status   string
	Complete bool
}

// APIError represents an error response from the Brewfather API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return "brewfather API error " + itoa(e.StatusCode) + ": " + e.Body
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := [10]byte{}
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[pos:])
}

// Batch represents a brewing batch in Brewfather.
type Batch struct {
	ID      string `json:"_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	BatchNo int    `json:"batchNo"`

	BrewDate              *int64 `json:"brewDate,omitempty"`
	FermentationStartDate *int64 `json:"fermentationStartDate,omitempty"`
	BottlingDate          *int64 `json:"bottlingDate,omitempty"`

	MeasuredOG  float64 `json:"measuredOg,omitempty"`
	MeasuredFG  float64 `json:"measuredFg,omitempty"`
	MeasuredABV float64 `json:"measuredAbv,omitempty"`
	MeasuredBatchSize float64 `json:"measuredBatchSize,omitempty"`
	MeasuredEfficiency float64 `json:"measuredEfficiency,omitempty"`

	Notes  string       `json:"notes,omitempty"`
	Recipe *BatchRecipe `json:"recipe,omitempty"`

	Created int64 `json:"_created"`
	Updated int64 `json:"_updated"`
}

// BatchRecipe is the recipe embedded in a batch.
type BatchRecipe struct {
	Name      string  `json:"name"`
	StyleName string  `json:"styleName,omitempty"`
	BatchSize float64 `json:"batchSize,omitempty"`
	BoilTime  int     `json:"boilTime,omitempty"`
	OG        float64 `json:"og,omitempty"`
	FG        float64 `json:"fg,omitempty"`
	ABV       float64 `json:"abv,omitempty"`
	IBU       float64 `json:"ibu,omitempty"`
	Color     float64 `json:"color,omitempty"`
}

// BatchReading represents a fermentation sensor reading.
type BatchReading struct {
	ID          string  `json:"_id"`
	Comment     string  `json:"comment,omitempty"`
	Gravity     float64 `json:"gravity,omitempty"`
	Temperature float64 `json:"temp,omitempty"`
	Pressure    float64 `json:"pressure,omitempty"`
	Angle       float64 `json:"angle,omitempty"`
	Battery     float64 `json:"battery,omitempty"`
	Created     int64   `json:"_created"`
}

// Recipe represents a brewing recipe.
type Recipe struct {
	ID         string  `json:"_id"`
	Name       string  `json:"name"`
	Author     string  `json:"author,omitempty"`
	StyleName  string  `json:"styleName,omitempty"`
	BatchSize  float64 `json:"batchSize,omitempty"`
	BoilTime   int     `json:"boilTime,omitempty"`
	OG         float64 `json:"og,omitempty"`
	FG         float64 `json:"fg,omitempty"`
	ABV        float64 `json:"abv,omitempty"`
	IBU        float64 `json:"ibu,omitempty"`
	Color      float64 `json:"color,omitempty"`
	Efficiency float64 `json:"efficiency,omitempty"`
	Type       string  `json:"type,omitempty"`
	Created    int64   `json:"_created"`
	Updated    int64   `json:"_updated"`
}

// Fermentable represents a fermentable ingredient in inventory.
type Fermentable struct {
	ID       string  `json:"_id"`
	Name     string  `json:"name"`
	Type     string  `json:"type,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	PPG      float64 `json:"ppg,omitempty"`
	EBC      float64 `json:"ebc,omitempty"`
	Lovibond float64 `json:"lovibond,omitempty"`
	Origin   string  `json:"origin,omitempty"`
	Supplier string  `json:"supplier,omitempty"`
	Notes    string  `json:"notes,omitempty"`
}

// Hop represents a hop ingredient in inventory.
type Hop struct {
	ID      string  `json:"_id"`
	Name    string  `json:"name"`
	Alpha   float64 `json:"alpha,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Unit    string  `json:"unit,omitempty"`
	Year    string  `json:"year,omitempty"`
	Origin  string  `json:"origin,omitempty"`
	Use     string  `json:"use,omitempty"`
	Notes   string  `json:"notes,omitempty"`
}

// Yeast represents a yeast ingredient in inventory.
type Yeast struct {
	ID          string  `json:"_id"`
	Name        string  `json:"name"`
	Laboratory  string  `json:"laboratory,omitempty"`
	ProductID   string  `json:"productId,omitempty"`
	Type        string  `json:"type,omitempty"`
	Form        string  `json:"form,omitempty"`
	Attenuation float64 `json:"attenuation,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Unit        string  `json:"unit,omitempty"`
	Notes       string  `json:"notes,omitempty"`
}

// Misc represents a miscellaneous ingredient in inventory.
type Misc struct {
	ID     string  `json:"_id"`
	Name   string  `json:"name"`
	Type   string  `json:"type,omitempty"`
	Use    string  `json:"use,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Unit   string  `json:"unit,omitempty"`
	Notes  string  `json:"notes,omitempty"`
}
