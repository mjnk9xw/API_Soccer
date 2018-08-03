package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Tour
type TourController struct {
	beego.Controller
}

// @Title CreateTour
// @Description create tour
// @Param	body		body 	models.Tour	true		"body for tour content"
// @Success 200 {ResponseTour} ResponseTour
// @Failure 403 body is empty
// @router / [post]
func (t *TourController) Post() {
	var tour models.Tour
	json.Unmarshal(t.Ctx.Input.RequestBody, &tour)
	tour.SetTourID()
	err := CheckTourName(tour.Name)
	if err == errSC.Success {
		serverstore.Tours.Set(&tour)
		serverstore.TourIDToListRoundID.Set(tour.GenIDKey())
		serverstore.TourIDToListPlayerID.Set(tour.GenIDKey())
	}
	t.Data["json"] = ResponseTour{Data: []models.TourInfo{&tour}, Error: errSC.NewErr(err)}
	t.ServeJSON()
}

// @Title GetAll
// @Description get all tours
// @Success 200 {ResponseTour} ResponseTour
// @Failure 403 :TourID is empty
// @router / [get]
func (t *TourController) GetAll() {
	tours, err := serverstore.Tours.GetAll()
	t.Data["json"] = ResponseTour{Data: tours, Error: errSC.NewErr(err)}
	t.ServeJSON()
}

// @Title Get
// @Description get tour by TourID
// @Param	TourID		path 	string	true		"The key for staticblock"
// @Success 200 {ResponseTour} ResponseTour
// @Failure 403 :TourID is empty
// @router /:TourID [get]
func (t *TourController) Get() {
	TourID := t.Ctx.Input.Param(":TourID")
	TourIDInt, err := CheckTour(TourID)
	if err == errSC.Success {
		tour, err := serverstore.Tours.GetByID(TourIDInt)
		t.Data["json"] = ResponseTour{Data: []models.TourInfo{tour}, Error: errSC.NewErr(err)}
	} else {
		t.Data["json"] = ResponseTour{Error: errSC.NewErr(err)}
	}
	t.ServeJSON()
}

// @Title Update
// @Description update the tour
// @Param	body		body 	models.Tour	true		"body for tour content"
// @Success 200 {ResponseTour} ResponseTour
// @Failure 403 not Tour
// @router / [put]
func (t *TourController) Put() {
	var tour models.Tour
	json.Unmarshal(t.Ctx.Input.RequestBody, &tour)
	_, err := serverstore.Tours.GetByID(tour.GetTourID())
	if err == errSC.Success {
		serverstore.Tours.Set(&tour)
	}
	t.Data["json"] = ResponseTour{Data: []models.TourInfo{&tour}, Error: errSC.NewErr(err)}
	t.ServeJSON()
}

// @Title Delete
// @Description delete the tour
// @Param	TourID		path 	string	true		"The TourID you want to delete"
// @Success 200 {ResponseTour} ResponseTour
// @Failure 403 TourID is empty
// @router /:TourID [delete]
func (t *TourController) Delete() {
	TourID := t.Ctx.Input.Param(":TourID")
	TourIDInt, err := CheckTour(TourID)
	if err == errSC.Success {
		err = serverstore.Tours.Remove(TourIDInt)
		serverstore.TourIDToListRoundID.Remove(TourID)
		serverstore.TourIDToListPlayerID.Remove(TourID)
	}
	t.Data["json"] = ResponseTour{Error: errSC.NewErr(err)}
	t.ServeJSON()
}
