package puerr

import "fmt"

var InterestsMismatch = fmt.Errorf("interests do not match")
var PlayerTypeMismatch = fmt.Errorf("player type does not match")
var LeaderboardTypeMismatch = fmt.Errorf("scoring type does not match")
var GameTypeMismatch = fmt.Errorf("game type does not match")
var MustBeUoi = fmt.Errorf("ID must start with 'uoi'")
var MustBeTeam = fmt.Errorf("ID must start with 'tms'")
var MustBeLeaderboard = fmt.Errorf("ID must start with 'ldb'")
