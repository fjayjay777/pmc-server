package model

type Class struct {
	Base
	Semester            string  `gorm:"column:semester" json:"semester"`
	Year                string  `gorm:"column:year" json:"year"`
	Session             string  `gorm:"column:session" json:"session"`
	WaitList            string  `gorm:"column:wait_list" json:"waitList"`
	OfferDate           string  `gorm:"column:offer_date" json:"offerDate"`
	StartTime           string  `gorm:"column:start_time" json:"startTime"`
	EndTime             string  `gorm:"column:end_time" json:"endTime"`
	Location            string  `gorm:"column:location" json:"location"`
	RecommendationScore int     `gorm:"column:recommendation_score" json:"recommendationScore"`
	Type                string  `gorm:"column:type" json:"type"`
	Number              string  `gorm:"column:number" json:"number"`
	Component           string  `gorm:"column:component" json:"component"`
	Unit                string  `gorm:"column:unit" json:"unit"`
	SeatAvailable       int32   `gorm:"column:seat_available" json:"seatAvailable"`
	Note                string  `gorm:"column:notes" json:"note"`
	Instructors         string  `gorm:"column:instructors" json:"instructors"`
	CourseTitle         string  `gorm:"column:course_title" json:"courseTitle"`
	CourseCatalogName   string  `gorm:"column:course_catalog_name" json:"courseCatalogName"`
	CourseID            int64   `gorm:"column:course_id" json:"courseID"`
	Course              *Course `gorm:"foreignKey:CourseID" json:"course"`
}
