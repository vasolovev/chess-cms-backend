package entity

import (
	"time"
)

type Tournament struct {
	NbPlayers  int           `json:"nbPlayers" bson:"nbPlayers"`
	Duels      []interface{} `json:"duels" bson:"duels"`
	IsFinished bool          `json:"isFinished" bson:"isFinished"`
	Podium     []struct {
		Name   string `json:"name" bson:"name"`
		Rank   int    `json:"rank" bson:"rank"`
		Rating int    `json:"rating" bson:"rating"`
		Score  int    `json:"score" bson:"score"`
		Sheet  struct {
		} `json:"sheet" bson:"sheet"`
		Nb struct {
			Game    int `json:"game" bson:"game"`
			Berserk int `json:"berserk" bson:"berserk"`
			Win     int `json:"win" bson:"win"`
		} `json:"nb" bson:"nb"`
		Performance int `json:"performance" bson:"performance"`
	} `json:"podium" bson:"podium"`
	PairingsClosed bool `json:"pairingsClosed" bson:"pairingsClosed"`
	Stats          struct {
		Games         int `json:"games" bson:"games"`
		Moves         int `json:"moves" bson:"moves"`
		WhiteWins     int `json:"whiteWins" bson:"whiteWins"`
		BlackWins     int `json:"blackWins" bson:"blackWins"`
		Draws         int `json:"draws" bson:"draws"`
		Berserks      int `json:"berserks" bson:"berserks"`
		AverageRating int `json:"averageRating" bson:"averageRating"`
	} `json:"stats" bson:"stats"`
	Standing struct {
		Page    int `json:"page" bson:"page"`
		Players []struct {
			Name   string `json:"name" bson:"name"`
			Rank   int    `json:"rank" bson:"rank"`
			Rating int    `json:"rating" bson:"rating"`
			Score  int    `json:"score" bson:"score"`
			Sheet  struct {
				Scores string `json:"scores" bson:"scores"`
			} `json:"sheet,omitempty" bson:"sheet"`
			// Sheet0 struct {
			// 	Scores string `json:"scores"`
			// 	Fire   bool   `json:"fire"`
			// } `json:"sheet,omitempty"`
			// Sheet1 struct {
			// 	Scores string `json:"scores"`
			// 	Fire   bool   `json:"fire"`
			// } `json:"sheet,omitempty"`
			// Sheet2 struct {
			// 	Scores string `json:"scores"`
			// 	Fire   bool   `json:"fire"`
			// } `json:"sheet,omitempty"`
		} `json:"players" bson:"players"`
	} `json:"standing" bson:"standing"`
	ID        string    `json:"id" bson:"_id,omitempty"`
	CreatedBy string    `json:"createdBy" bson:"createdBy"`
	StartsAt  time.Time `json:"startsAt" bson:"startsAt"`
	System    string    `json:"system" bson:"system"`
	FullName  string    `json:"fullName" bson:"fullName"`
	Minutes   int       `json:"minutes" bson:"minutes"`
	Perf      struct {
		Key  string `json:"key" bson:"key"`
		Name string `json:"name" bson:"name"`
		Icon string `json:"icon" bson:"icon"`
	} `json:"perf" bson:"perf"`
	Clock struct {
		Limit     int `json:"limit" bson:"limit"`
		Increment int `json:"increment" bson:"increment"`
	} `json:"clock" bson:"clock"`
	Variant     string `json:"variant" bson:"variant"`
	Rated       bool   `json:"rated" bson:"rated"`
	Berserkable bool   `json:"berserkable" bson:"berserkable"`
	Verdicts    struct {
		List     []interface{} `json:"list" bson:"list"`
		Accepted bool          `json:"accepted" bson:"accepted"`
	} `json:"verdicts" bson:"verdicts"`
	Description string `json:"description" bson:"description"`
}
