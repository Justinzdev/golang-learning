package controllers

import (
	"net/http"

	"myrestapi/4.simple_restapi/database"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	ID                  int    `json:"id"`
	DiscordUsername     string `json:"discordUsername"`
	DiscordAvatar       string `json:"discordAvatar"`
	DiscordID           string `json:"discordID"`
	DiscordToken        string `json:"discordToken"`
	DiscordRefreshToken string `json:"discordRefreshToken"`
	CreateAt            string `json:"createAt"`
}

func GetUser(c *gin.Context) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM discord_accounts")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}
	defer rows.Close()

	var userDataLists []UserData

	for rows.Next() {
		var user UserData
		err := rows.Scan(
			&user.ID,
			&user.DiscordUsername,
			&user.DiscordAvatar,
			&user.DiscordID,
			&user.DiscordToken,
			&user.DiscordRefreshToken,
			&user.CreateAt,
		)

		if err != nil {
			panic(err.Error())
		}

		userDataLists = append(userDataLists, user)
	}

	c.JSON(http.StatusOK, userDataLists)
}
