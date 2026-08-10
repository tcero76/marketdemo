package demo

import (
	"strings"

	logger "github.com/tcero76/marketplace/config/log"
	"gorm.io/gorm"
)

func cleanWords(words []string) []string {
	var result []string
	for _, w := range words {
		if w = strings.TrimSpace(w); w != "" {
			result = append(result, w)
		}
	}
	return result
}

type TextSpecProduct struct {
	Words []string
	Log   *logger.LoggerLogstash
}

func (s TextSpecProduct) Apply(db *gorm.DB) *gorm.DB {
	s.Log.Info("Words son: ", s.Words)
	clean := cleanWords(s.Words)
	if len(clean) > 0 {
		text := strings.Join(clean, " ")
		return db.Where("search_vector @@ plainto_tsquery('english', ?)", text)
	}
	return db
}

type CategorySpecProduct struct {
	CategoryIDs []int
	Log         *logger.LoggerLogstash
}

func (s CategorySpecProduct) Apply(db *gorm.DB) *gorm.DB {
    if len(s.CategoryIDs) == 0 {
        return db
    }
    s.Log.Info("Filtering products by categories: ", s.CategoryIDs)
    return db.Where("categoryid IN ?", s.CategoryIDs)
}

type Specification interface {
	Apply(db *gorm.DB) *gorm.DB
}

func ApplySpecifications(db *gorm.DB, specs ...Specification) *gorm.DB {
	for _, spec := range specs {
		db = spec.Apply(db.Debug())
	}
	return db
}