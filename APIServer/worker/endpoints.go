package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Countries struct {
	Red   []string `json:"Red"`
	Amber []string `json:"Amber"`
	Green []string `json:"Green"`
}

func ReturnCountries() (Countries, error) {
	c := colly.NewCollector()
	var countries Countries

	c.OnHTML("div.govspeak", func(body *colly.HTMLElement) {
		body.DOM.Find("table").Each(func(i int, t *goquery.Selection) {
			var iterCounter []string

			t.Find("tr>th").Each(func(_ int, s *goquery.Selection) {
				re := regexp.MustCompile(`([\(\[]).*?([\)\]])`)
				thing := strings.TrimSuffix(s.Text(), "\n")

				switch thing {
				case "Congo (Democratic Republic)": // Our issue, ish?
					iterCounter = append(iterCounter, "CD")
					break
				case "Congo": // Our issue
					iterCounter = append(iterCounter, "CG")
					break
				case "South Sudan": // Definietly google issue...
					iterCounter = append(iterCounter, "SS")
					break
				case "Côte d’Ivoire":
					iterCounter = append(iterCounter, "CI")
					break
				case "The Bahamas":
					iterCounter = append(iterCounter, "BS")
					break
				case "  Israel and Jerusalem":
					iterCounter = append(iterCounter, "IL")
					break
				case "Kosovo":
					iterCounter = append(iterCounter, "XK")
					break
				case "North Macedonia":
					iterCounter = append(iterCounter, "MK")
					break
				case "Timor-Leste":
					iterCounter = append(iterCounter, "TL")
					break
				case "Eswatini":
					iterCounter = append(iterCounter, "SZ")
					break
				case "Réunion":
					iterCounter = append(iterCounter, "RE")
					break
				default:
					new := re.ReplaceAllString(thing, "")
					iterCounter = append(iterCounter, new)
					break
				}
			})
			if strings.Contains(t.Find("tr>th").First().Text(), "Red") {
				countries.Red = iterCounter
			} else if strings.Contains(t.Find("tr>th").First().Text(), "Amber") {
				countries.Amber = iterCounter
			} else if strings.Contains(t.Find("tr>th").First().Text(), "Green") {
				countries.Green = iterCounter
			}
		})
	})

	c.Visit("https://www.gov.uk/guidance/red-amber-and-green-list-rules-for-entering-england")
	if len(countries.Red) == 0 {
		fmt.Println("Unable to fetch results for Red table.")
	}
	if len(countries.Amber) == 0 {
		fmt.Println("Unable to fetch results for Amber table.")
	}
	if len(countries.Green) == 0 {
		fmt.Println("Unable to fetch results for Green table.")
	}

	return countries, nil
}

func (route *Worker) SendJson(w http.ResponseWriter, r *http.Request, class interface{}, pageRequest string) {
	payload, err := json.Marshal(class)
	if err != nil {
		route.logger.Debug(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	route.logger.Debug("A '" + pageRequest + "' request was made.")

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (route *Worker) AllLists(w http.ResponseWriter, r *http.Request) {
	countries, err := ReturnCountries()
	if err != nil {
		route.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	route.SendJson(w, r, countries, "All Countries")
}

func (route *Worker) GreenList(w http.ResponseWriter, r *http.Request) {
	countries, err := ReturnCountries()
	if err != nil {
		route.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	route.SendJson(w, r, countries.Green, "Green List")
}

func (route *Worker) AmberList(w http.ResponseWriter, r *http.Request) {
	countries, err := ReturnCountries()
	if err != nil {
		route.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	route.SendJson(w, r, countries.Amber, "Amber List")
}

func (route *Worker) RedList(w http.ResponseWriter, r *http.Request) {
	countries, err := ReturnCountries()
	if err != nil {
		route.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	route.SendJson(w, r, countries.Red, "Red List")
}
