package main

import (
	"net/http"
	"strconv"

	"github.com/archit-p/MicroserviceTemplate/pkg/dto"
	"github.com/gorilla/mux"

	"github.com/jinzhu/copier"
)

// createSample godoc
// @Summary Create a Sample in the database
// @Description Create a Sample in the database
// @Tags samples
// @Accept json
// @Produce json
// @Param sample body dto.Sample true "Create sample"
// @Success 200 ID string
// @Router /sample/create [post]
func (app *application) createSample(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	var s dto.Sample
	err := s.FromJSON(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	id, err := app.samples.Insert(s.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Created ", id)
	w.Write([]byte("Created document"))
}

// getSample godoc
// @Summary Retrive a Sample from the database
// @Description Retrive a Sample from the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [get]
func (app *application) getSample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	snip, err := app.samples.Get(vars["id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var s dto.Sample
	copier.Copy(&s, &snip)

	err = s.ToJSON(w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Read ", vars["id"])
}

// updateSample godoc
// @Summary Update a Sample in the database
// @Description Update a Sample in the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Param sample body dto.Sample true "Update sample"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [put]
func (app *application) updateSample(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)

	var s dto.Sample
	err := s.FromJSON(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	res, err := app.samples.Update(vars["id"], s.Content)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infoLog.Printf("Updated : %d : %s", res, vars["id"])
	w.Write([]byte("Updated document"))
}

// deleteSample godoc
// @Summary Delete a Sample in the database
// @Description Delete a Sample in the database
// @Tags samples
// @Accept  json
// @Produce  json
// @Param id path string true "ID to look for"
// @Success 200 {object} dto.Sample
// @Router /sample/{id} [delete]
func (app *application) deleteSample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := app.samples.Delete(vars["id"])
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infoLog.Printf("Deleted : %d : %s", res, vars["id"])
	w.Write([]byte("Deleted document"))
}

// searchSample godoc
// @Summary Search using keywords
// @Description Search the keywords in the content
// @Tags samples
// @Accept  json
// @Produce  json
// @Param q query string true "Search query"
// @Success 200 {object} dto.Samples
// @Router /sample [get]
func (app *application) searchSample(w http.ResponseWriter, r *http.Request) {
	keywords := r.URL.Query().Get("q")

	res, err := app.samples.Search(keywords)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var s dto.Samples
	copier.Copy(&s, res)

	s.ToJSON(w)
	app.infoLog.Printf("Searched : %s", keywords)
}

// searchSample godoc
// @Summary Find top Samples
// @Description Top samples based on parameter and limit
// @Tags samples
// @Accept  json
// @Produce  json
// @Param param query string true "Parameter to sort by"
// @Param limit query int64 true "Number of results to limit to"
// @Success 200 {object} dto.Samples
// @Router /sample/top [get]
func (app *application) topSamples(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	res, err := app.samples.Top(param, limit)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var s dto.Samples
	copier.Copy(&s, res)

	s.ToJSON(w)
	app.infoLog.Printf("Top results : %s : %d", param, limit)
}
