package course

import (
	"context"
	"encoding/json"
	"errors"
	"sort"

	"pmc_server/init/es"
	"pmc_server/init/postgres"
	"pmc_server/model"
	esModel "pmc_server/model/es"
	"pmc_server/utils"

	"github.com/olivere/elastic/v7"
)

func GetCourses(pn, pSize int) ([]model.Course, int64) {
	var courses []model.Course
	total := postgres.DB.Find(&courses).RowsAffected
	postgres.DB.Scopes(utils.Paginate(pn, pSize)).Find(&courses)

	return courses, total
}

func GetCourseByID(id int) (*model.Course, error) {
	var course model.Course
	result := postgres.DB.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no course info found")
	}
	return &course, nil
}

func GetClassListByCourseID(id int) (*[]model.Class, int64) {
	var classes []model.Class
	res := postgres.DB.Where("course_id = ?", id).Find(&classes)
	return &classes, res.RowsAffected
}

func GetCoursesBySearch(param model.CourseFilterParams) ([]model.Course, int64, error) {
	query := elastic.NewBoolQuery()
	if param.MinCredit != 0 {
		query = query.Filter(elastic.NewRangeQuery("min_credit").Gte(param.MinCredit))
	}

	if param.MaxCredit != 0 {
		query = query.Filter(elastic.NewRangeQuery("max_credit").Lte(param.MaxCredit))
	}

	if param.StartTime != 0 {
		query = query.Filter(elastic.NewRangeQuery("start_time").Gte(param.StartTime))
	}

	if param.EndTime != 0 {
		query = query.Filter(elastic.NewRangeQuery("end_time").Lte(param.EndTime))
	}

	if param.IsHonor {
		query = query.Must(elastic.NewTermsQuery("is_honor", true))
	}

	if param.MinRating != 0 {
		query = query.Filter(elastic.NewRangeQuery("min_rating").Gte(param.MinRating))
	}

	if len(param.Weekday) != 0 {
		query = query.Must(elastic.NewTermQuery("weekdays", param.Weekday))
	}

	if param.Keyword != "" {
		query = query.Must(elastic.NewMultiMatchQuery(param.Keyword, "title", "description", "designation_catalog", "catalog_course_name"))
	}

	if param.PageSize > 20 {
		param.PageSize = 20
	}
	if param.PageSize < 10 {
		param.PageSize = 10
	}

	searchRes, err := es.Client.Search().
		Index("course").Query(query).
		From(param.PageNumber).
		Size(param.PageSize).
		Do(context.Background())

	if err != nil {
		return nil, -1, err
	}

	total := searchRes.TotalHits()

	var esCourseIdList []int64
	for _, hit := range searchRes.Hits.Hits {
		var course esModel.Course
		err := json.Unmarshal(*&hit.Source, &course)
		if err != nil {
			return nil, -1, err
		}
		esCourseIdList = append(esCourseIdList, course.ID)
	}

	var courseList []model.Course
	for _, id := range esCourseIdList {
		var course model.Course
		res := postgres.DB.Where("id = ?", id).First(&course)
		if res.Error != nil {
			return nil, -1, err
		}
		courseList = append(courseList, course)
	}

	var filteredCourseList []model.Course
	if param.OfferedOnline || param.OfferedOffline {
		for _, course := range courseList {
			var classes []model.Class
			_ = postgres.DB.Where("course_id = ?", course.ID).Find(&classes)
			for _, class := range classes {
				if param.OfferedOnline && (class.OfferDate == "" || class.StartTime == "" || class.EndTime == "") {
					filteredCourseList = append(filteredCourseList, course)
					continue
				}
				if param.OfferedOffline && (class.OfferDate != "") {
					filteredCourseList = append(filteredCourseList, course)
					continue
				}
			}
		}
	}

	if param.RankByRatingLowToHigh {
		sort.SliceStable(filteredCourseList, func(i, j int) bool {
			return filteredCourseList[i].Rating < filteredCourseList[j].Rating
		})
	}

	if param.RankByRatingHighToLow {
		sort.SliceStable(filteredCourseList, func(i, j int) bool {
			return filteredCourseList[i].Rating > filteredCourseList[j].Rating
		})
	}

	return filteredCourseList, total, nil
}
