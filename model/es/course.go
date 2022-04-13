package es

type Course struct {
	ID                 int64     `json:"id"`
	DesignationCatalog string    `json:"designation_catalog"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	CatalogCourseName  string    `json:"catalog_course_name"`
	CatalogName        string    `json:"catalog_name"`
	Prerequisites      string    `json:"prerequisites"`
	Component          string    `json:"component"`
	MaxCredit          float32   `json:"max_credit"`
	MinCredit          float32   `json:"min_credit"`
	Subject            string    `json:"subject"`
	IsHonor            bool      `json:"is_honor"`
	FixedCredit        bool      `json:"fixed_credit"`
	Rating             float32   `json:"rating"`
	Professors         []string  `json:"professors"`
	OffersOnline       bool      `json:"offersOnline"`
	OffersOffline      bool      `json:"offersOffline"`
	Dates              []int32   `json:"dates"`
	StartTimes         []float32 `json:"startTimes"`
	EndTimes           []float32 `json:"endTimes"`
}

func (Course) GetMapping() string {
	return `
{
   "mappings":{
         "properties":{
            "id":{
               "type":"integer"
            },
            "title":{
               "type":"text"
            },
            "description":{
               "type":"text"
            },
            "designation_catalog":{
               "type":"text"
            },
            "catalog_course_name":{
               "type":"text"
            },
			"catalog_name":{
 				"type":"text"
			}
            "prerequisites":{
               "type":"text"
            },
            "component":{
               "type":"text"
            },
            "max_credit":{
               "type":"number"
            },
            "min_credit":{
               "type":"number"
            },
            "subject":{
               "type":"text"
            },
            "is_honor":{
               "type":"boolean"
            },
            "fixed_credit":{
               "type":"boolean"
            },
			"rating": {
				"rating":"rating"
			},
			"classes": {
				
			}
         }
      }
}
`
}

func (Course) GetIndexName() string {
	return "course"
}
