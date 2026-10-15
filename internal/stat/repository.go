package stat

import (
	"http-server/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.DB
}

func NewStatRepository(db *db.DB) *StatRepository {
	return &StatRepository{
		DB: db,
	}
}

func (repo *StatRepository) AddClick(linkID uint) {
	var stat Stat
	currDate := datatypes.Date(time.Now())
	repo.DB.Find(&stat, "link_id = ? and date = ?", linkID, currDate)
	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkId: linkID,
			Clicks: 1,
			Date:   currDate,
		})
	} else {
		stat.Clicks += 1
		repo.DB.Save(&stat)
	}
}
