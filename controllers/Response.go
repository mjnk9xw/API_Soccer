package controllers

import (
	"SoccerBeego/err"
	"SoccerBeego/models"
)

// //NewResponse khoi tao response
// func NewResponse(data interface{}, err error) *Response {
// 	return &Response{
// 		Data:  data,
// 		Error: errSC.NewErr(err),
// 	}
// }

// // Response data + err
// type Response struct {
// 	Data  interface{}
// 	Error *errSC.Err
// }

// ResponseTour Tour + err
type ResponseTour struct {
	Data  []models.TourInfo
	Error *errSC.Err
}

// ResponseRound round + err
type ResponseRound struct {
	Data  []models.RoundInfo
	Error *errSC.Err
}

// ResponseTable table + err
type ResponseTable struct {
	Data  []models.TableInfo
	Error *errSC.Err
}

// ResponsePlayer player + err
type ResponsePlayer struct {
	Data  []models.PlayerInfo
	Error *errSC.Err
}

// ResponseMatch match + err
type ResponseMatch struct {
	Data  []models.MatchInfo
	Error *errSC.Err
}
