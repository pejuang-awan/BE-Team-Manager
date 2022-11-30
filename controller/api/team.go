package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/pejuang-awan/BE-Team-Manager/models"
	"github.com/pejuang-awan/BE-Team-Manager/services"
)

func Join(c *gin.Context) {
	db := services.Database

	var input models.TeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	team := models.Team{
		CaptainID:    input.CaptainID,
		TournamentID: input.TournamentID,
		TeamName:     input.TeamName,
		Members:      strings.Join(input.Members, ","),
	}

	db.Create(&team)

	c.JSON(http.StatusOK, models.Response{Data: input})
}

func Leave(c *gin.Context) {
	db := services.Database

	var input models.TeamLeave
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Error: err.Error()})
		return
	}

	db.Delete(&models.Team{}, "captain_id = ? AND tournament_id = ?", input.CaptainID, input.TournamentID)

	c.JSON(http.StatusOK, models.Response{Data: "Successfully left tournament"})
}

func Participants(c *gin.Context) {
	db := services.Database

	var rawTeams []models.Team
	if result := db.Model(&models.Team{}).Where("tournament_id = ?", c.Param("id")).Find(&rawTeams); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: "An error occured"})
		return
	}

	var teams []models.TeamInput
	for _, rawTeam := range rawTeams {
		teams = append(teams, models.TeamInput{
			CaptainID:    rawTeam.CaptainID,
			TournamentID: rawTeam.TournamentID,
			TeamName:     rawTeam.TeamName,
			Members:      strings.Split(rawTeam.Members, ","),
		})
	}

	c.JSON(http.StatusOK, models.Response{Data: teams})
}

func Tournaments(c *gin.Context) {
	db := services.Database

	var rawTeams []models.Team
	if result := db.Model(&models.Team{}).Where("captain_id = ?", c.Param("id")).Find(&rawTeams); result.Error != nil {
		c.JSON(http.StatusNotFound, models.Response{Error: "An error occured"})
		return
	}

	var teams []models.TeamInput
	for _, rawTeam := range rawTeams {
		teams = append(teams, models.TeamInput{
			CaptainID:    rawTeam.CaptainID,
			TournamentID: rawTeam.TournamentID,
			TeamName:     rawTeam.TeamName,
			Members:      strings.Split(rawTeam.Members, ","),
		})
	}

	c.JSON(http.StatusOK, models.Response{Data: teams})
}
