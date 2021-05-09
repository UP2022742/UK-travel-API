package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

func GetGreenList() ([]string, error) {
	resp, err := http.Get("http://localhost:8080/green")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func GetAmberList() ([]string, error) {
	resp, err := http.Get("http://localhost:8080/amber")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func GetRedList() ([]string, error) {
	resp, err := http.Get("http://localhost:8080/red")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println(err.Error())
	}
	return post, nil
}

// HomePage contains the index.
func (route *Router) HomePage(w http.ResponseWriter, r *http.Request) {
	type MetaHome struct{}
	templates, err := template.ParseFiles(
		"templates/index.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "index.html", MetaHome{})
	route.logger.Debug("An Home Request was made.")
}

// NotificationPage contains the template event page.
func (route *Router) AboutPage(w http.ResponseWriter, r *http.Request) {
	type MetaEvent struct {
	}
	templates, err := template.ParseFiles(
		"templates/about.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "about.html", MetaEvent{})
	route.logger.Debug("An About Request was made.")
}

// TablesPage contains the template event page.
func (route *Router) ProjectsPage(w http.ResponseWriter, r *http.Request) {
	type MetaTables struct {
	}
	templates, err := template.ParseFiles(
		"templates/project.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "project.html", MetaTables{})
	route.logger.Debug("An Projects Request was made.")
}

// TypographyPage contains the template event page.
func (route *Router) SocialPage(w http.ResponseWriter, r *http.Request) {
	type MetaTypo struct {
	}
	templates, err := template.ParseFiles(
		"templates/social.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "social.html", nil)
	route.logger.Debug("An Social Request was made.")
}

// TypographyPage contains the template event page.
func (route *Router) TravelPage(w http.ResponseWriter, r *http.Request) {
	type MetaTravel struct {
		GreenListCountries []string
		AmberListCountries []string
		RedListCountries   []string
	}
	templates, err := template.ParseFiles(
		"templates/travel.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}

	greenList, err := GetGreenList()
	if err != nil {
		fmt.Println("Couldn't get greenlist")
	}
	amberList, err := GetAmberList()
	if err != nil {
		fmt.Println("Couldn't get amberlist.")
	}
	redList, err := GetRedList()
	if err != nil {
		fmt.Println("Couldn't get redlist.")
	}

	templates.ExecuteTemplate(
		w, "travel.html", MetaTravel{
			GreenListCountries: greenList,
			AmberListCountries: amberList,
			RedListCountries:   redList,
		},
	)
	route.logger.Debug("An Travel Request was made.")
}
