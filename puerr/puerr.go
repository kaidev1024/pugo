package puerr

import "fmt"

var InterestsDoNotMatch = fmt.Errorf("interests do not match")
var PlayerTypeDoNotMatch = fmt.Errorf("player type does not match")
var ScoringTypeDoNotMatch = fmt.Errorf("scoring type does not match")
var GameTypeDoNotMatch = fmt.Errorf("game type does not match")
var MustBeUoi = fmt.Errorf("ID must start with 'uoi'")
var MustBeTeam = fmt.Errorf("ID must start with 'tms'")
var MustBeLeaderboard = fmt.Errorf("ID must start with 'ldb'")
