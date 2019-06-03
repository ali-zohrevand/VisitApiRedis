package Services

import (
	"github.com/ali-zohrevand/VisitApi/models"
	"time"
)

func VisitHappend(Path string, username string) {
	var visit models.Visit
	visit.Path = Path
	visit.Username = username
	visit.TimeVisit = time.Now().Unix()

}
