package entity

import "time"

type Tournament struct {
	NbPlayers  int           `json:"nbPlayers"`
	Duels      []interface{} `json:"duels"`
	IsFinished bool          `json:"isFinished"`
	Podium     []struct {
		Name   string `json:"name"`
		Rank   int    `json:"rank"`
		Rating int    `json:"rating"`
		Score  int    `json:"score"`
		Sheet  struct {
		} `json:"sheet"`
		Nb struct {
			Game    int `json:"game"`
			Berserk int `json:"berserk"`
			Win     int `json:"win"`
		} `json:"nb"`
		Performance int `json:"performance"`
	} `json:"podium"`
	PairingsClosed bool `json:"pairingsClosed"`
	Stats          struct {
		Games         int `json:"games"`
		Moves         int `json:"moves"`
		WhiteWins     int `json:"whiteWins"`
		BlackWins     int `json:"blackWins"`
		Draws         int `json:"draws"`
		Berserks      int `json:"berserks"`
		AverageRating int `json:"averageRating"`
	} `json:"stats"`
	Standing struct {
		Page    int `json:"page"`
		Players []struct {
			Name   string `json:"name"`
			Rank   int    `json:"rank"`
			Rating int    `json:"rating"`
			Score  int    `json:"score"`
			Sheet  struct {
				Scores string `json:"scores"`
			} `json:"sheet,omitempty"`
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
		} `json:"players"`
	} `json:"standing"`
	ID        string    `json:"id"`
	CreatedBy string    `json:"createdBy"`
	StartsAt  time.Time `json:"startsAt"`
	System    string    `json:"system"`
	FullName  string    `json:"fullName"`
	Minutes   int       `json:"minutes"`
	Perf      struct {
		Key  string `json:"key"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"perf"`
	Clock struct {
		Limit     int `json:"limit"`
		Increment int `json:"increment"`
	} `json:"clock"`
	Variant     string `json:"variant"`
	Rated       bool   `json:"rated"`
	Berserkable bool   `json:"berserkable"`
	Verdicts    struct {
		List     []interface{} `json:"list"`
		Accepted bool          `json:"accepted"`
	} `json:"verdicts"`
	Description string `json:"description"`
}
