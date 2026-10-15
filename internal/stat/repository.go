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

func (repo *StatRepository) GetStats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string

	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD' as perios, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM' as perios, sum(clicks)"
	}
	repo.DB.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
