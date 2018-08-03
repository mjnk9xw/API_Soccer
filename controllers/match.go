package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
	"SoccerBeego/serverstore"
	"encoding/json"

	"github.com/astaxie/beego"
)

type MatchController struct {
	beego.Controller
}

// @Title Get
// @Description get match by MatchID
// @Param	MatchID		path 	string	true		"The key for staticblock"
// @Success 200 {ResponseMatch} ResponseMatch
// @Failure 403 :MatchID is empty
// @router /:MatchID [get]
func (m *MatchController) Get() {
	TourID := m.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := m.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := m.Ctx.Input.Param(":TableID")
			_, err := CheckTable(TableID)
			if err == errSC.Success {
				MatchID := m.Ctx.Input.Param(":MatchID")
				MatchIDInt, err := CheckMatch(MatchID)
				if err == errSC.Success {
					match, err := serverstore.Matchs.GetByID(MatchIDInt)
					m.Data["json"] = ResponseMatch{Data: []models.MatchInfo{match}, Error: errSC.NewErr(err)}
				} else {
					m.Data["json"] = ResponseMatch{Error: errSC.NewErr(err)}
				}
			} else {
				m.Data["json"] = ResponseMatch{Error: errSC.NewErr(err)}
			}
		} else {
			m.Data["json"] = ResponseMatch{Error: errSC.NewErr(err)}
		}
	} else {
		m.Data["json"] = ResponseMatch{Error: errSC.NewErr(err)}
	}
	m.ServeJSON()
}

// @Title CreateMatch
// @Description create match
// @Param	body		body 	models.Match	true		"body for match content"
// @Success 200 {ResponseMatch} ResponseMatch
// @Failure 403 body is empty
// @router / [post]
func (m *MatchController) Post() {
	var match models.Match
	TourID := m.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := m.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := m.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				json.Unmarshal(m.Ctx.Input.RequestBody, &match)
				match.SetMatchID()
				serverstore.Matchs.Set(&match)
				serverstore.TableIDToListMatchID.SetMatch(TableIDInt, match.GetMatchID())
			}
		}
	}
	m.Data["json"] = ResponseMatch{Data: []models.MatchInfo{&match}, Error: errSC.NewErr(err)}
	m.ServeJSON()
}

// @Title Update
// @Description update the match
// @Param	body		body 	models.Match	true		"body for match content"
// @Success 200 {ResponseMatch} ResponseMatch
// @Failure 403 not Match
// @router / [put]
func (m *MatchController) Put() {
	var match models.Match
	TourID := m.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := m.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := m.Ctx.Input.Param(":TableID")
			_, err := CheckTable(TableID)
			if err == errSC.Success {
				json.Unmarshal(m.Ctx.Input.RequestBody, &match)
				serverstore.Matchs.Set(&match)
			}
		}
	}
	m.Data["json"] = ResponseMatch{Data: []models.MatchInfo{&match}, Error: errSC.NewErr(err)}
	m.ServeJSON()
}

// @Title Delete
// @Description delete the match
// @Param	MatchID		path 	string	true		"The MatchID you want to delete"
// @Success 200 {ResponseMatch} ResponseMatch
// @Failure 403 MatchID is empty
// @router /:MatchID [delete]
func (m *MatchController) Delete() {
	TourID := m.Ctx.Input.Param(":TourID")
	_, err := CheckTour(TourID)
	if err == errSC.Success {
		RoundID := m.Ctx.Input.Param(":RoundID")
		_, err := CheckRound(RoundID)
		if err == errSC.Success {
			TableID := m.Ctx.Input.Param(":TableID")
			TableIDInt, err := CheckTable(TableID)
			if err == errSC.Success {
				MatchID := m.Ctx.Input.Param(":MatchID")
				MatchIDInt, err := CheckMatch(MatchID)
				if err == errSC.Success {
					serverstore.Matchs.Remove(MatchIDInt)
					serverstore.TableIDToListMatchID.RemoveMatch(TableIDInt, MatchIDInt)
				}
			}
		}
	}
	m.Data["json"] = ResponseMatch{Error: errSC.NewErr(err)}
	m.ServeJSON()
}
