package webscrape

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"

	"github.com/comediq-api/database"
	"github.com/comediq-api/helpers"
	"github.com/comediq-api/types"
)

var client *colly.Collector
var daysOfTheWeek = map[string]struct{}{
	"Monday":    {},
	"Tuesday":   {},
	"Wednesday": {},
	"Thursday":  {},
	"Friday":    {},
	"Saturday":  {},
	"Sunday":    {},
}

func Init() {
	client = colly.NewCollector()

	client.OnHTML("body", func(e *colly.HTMLElement) {
		var currentDay string = ""

		e.ForEachWithBreak("tr", func(i int, el *colly.HTMLElement) bool {

			if el.DOM.Find("b").Text() == "Open Mics in the United StatesOpen Mics Worldwide" {
				return false
			}

			day := strings.Split(el.Text, " ")[0]
			_, exists := daysOfTheWeek[day]
			if exists {
				currentDay = day
				return true
			}

			duration := "5 min"
			city := "Los Angeles"
			uuid := uuid.New()
			openMic := types.HistoricalMic{
				UniqueIdentifier: uuid,
				Day:              &currentDay,
				StageTime:        &duration,
				City:             &city,
			}

			// fmt.Println("Row", i, ":", el.Text)
			el.ForEach("td", func(j int, td *colly.HTMLElement) {
				if j == 0 {
					openMic.StartTime = &td.Text
					endTime, err := helpers.ClockAdd(td.Text, 60)

					if err != nil {
						fmt.Println("Error adding time:", err)
						return
					}
					openMic.LatestEndTime = &endTime
				} else {
					a := td.DOM.Find("a")
					venue := a.Find("b").Text()
					location := a.Text()[len(venue):]
					openMic.Location = &location
					href, exists := a.Attr("href")
					openMic.VenueName = &venue

					if !exists {
						fmt.Println("No details on", venue)
					}
					ScrapeMic(&openMic, href)
					fmt.Println(openMic.UniqueIdentifier.String() + "\n" + *openMic.VenueName + "\n" + *openMic.Location + "\n" +
						href + "\n" + *openMic.Cost + "\n" + *openMic.StartTime + "\n" + *openMic.LatestEndTime + "\n" + *openMic.Day + "\n" + *openMic.OpenMic + "\n---")

					active := true
					lastVerified := "Unverified"

					openMic.Active = &active
					openMic.LastVerified = &lastVerified
					database.UpsertData(openMic, "open_mics_historical")

					// fmt.Println("Col", j, ":", td.Text)
				}
			})
			return true
		})
	})

	client.Visit("https://badslava.com/open-mics.php?city=Los%20Angeles&state=CA")
}

var fieldMap = map[string]func(*types.HistoricalMic) **string{
	"Event Name": func(h *types.HistoricalMic) **string { return &h.OpenMic },
	"Cost":       func(h *types.HistoricalMic) **string { return &h.Cost },
}

func ScrapeMic(openMic *types.HistoricalMic, href string) {
	micClient := colly.NewCollector()

	micClient.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("td", func(i int, el *colly.HTMLElement) {
			// fmt.Println("Mic Row", i, ":", el.Text)
			rowData := strings.Split(el.Text, ": ")
			if len(rowData) < 2 {
				openMic.SignUpInstructions = &el.Text
			}
			if setter, ok := fieldMap[rowData[0]]; ok {
				fieldPtr := setter(openMic)
				*fieldPtr = &rowData[1]
			}
		})
	})

	micClient.Visit(href)
}
