package migrate

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"pmc_server/model"
	esModel "pmc_server/model/es"

	"github.com/olivere/elastic/v7"
	pos "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func postgresToES() {
	url := "https://search-pmc-search-h5lyuv34ujl5f6gby7gxbsrnnm.us-east-1.es.amazonaws.com/"
	username := "admin1"
	password := "Admin123!"
	logger := log.New(os.Stdout, "pmc", log.LstdFlags)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth(username, password),
		elastic.SetTraceLog(logger),
	)

	if err != nil {
		fmt.Println("failed to init es")
		return
	}

	dsn := "host=pmc1.ccyv4mlgftmr.us-east-1.rds.amazonaws.com user=admin1 password=admin123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(pos.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("failed to init postgres")
		return
	}

	var courses []model.Course
	res := db.Find(&courses)
	if res.Error != nil || res.RowsAffected == 0 {
		fmt.Println("failed to fetch courses")
		return
	}

	for _, course := range courses {
		esCourse := esModel.Course{
			ID:                 course.ID,
			DesignationCatalog: course.DesignationCatalog,
			Title:              course.Title,
			Description:        course.Description,
			CatalogCourseName:  course.CatalogCourseName,
			Prerequisites:      course.Prerequisites,
			Component:          course.Component,
			MaxCredit:          course.MaxCredit,
			MinCredit:          course.MinCredit,
			IsHonor:            course.IsHonor,
			FixedCredit:        course.FixedCredit,
		}

		_, err := client.Index().Index(esCourse.GetIndexName()).BodyJson(esCourse).Id(strconv.Itoa(int(course.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

}