package uc_implementation

import (
	"RSOI/src/recommendations/models"
	"math"
	"sort"
)

type RecommendationsUsecase struct {
}

type attributes []float64

func manhattanMetric(a attributes, b attributes) float64 {
	dist := 0.0
	for i := 0; i < len(a); i++ {
		dist += math.Abs(a[i] - b[i])
	}
	return dist
}

func difference(id int, lib []attributes) []float64 {
	n := len(lib)
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = manhattanMetric(lib[id], lib[i])
	}
	return res
}

func get_attributes(lib []models.Book) []attributes {
	tags := make([]string, 0)
	maxyear := lib[0].Year
	minyear := lib[0].Year
	maxrating := lib[0].Year
	minrating := lib[0].Year

	for _, book := range lib {
		tags = append(tags, book.Tags...)
		if book.Year > maxyear {
			maxyear = book.Year
		}
		if book.Year < minyear {
			minyear = book.Year
		}
		if book.Rating > maxrating {
			maxrating = book.Rating
		}
		if book.Rating < minrating {
			minrating = book.Rating
		}
	}

	res := make([]attributes, 0)
	for _, book := range lib {
		attr := make([]float64, 0)
		attr = append(attr, float64(book.Year-minyear)/float64(maxyear))
		attr = append(attr, float64(book.Rating-minrating)/float64(maxrating))
		if book.Status {
			attr = append(attr, 1)
		} else {
			attr = append(attr, 0)
		}
		for _, tag := range tags {
			var i int
			for i = 0; i < len(book.Tags) && book.Tags[i] != tag; i++ {
			}
			if book.Tags[i] == tag {
				attr = append(attr, 1)
			} else {
				attr = append(attr, 0)
			}
		}
		res = append(res, attr)
	}
	return res
}

func (rc *RecommendationsUsecase) GetRecommendations(lib []models.Book, prefs *models.PreferencesList) []models.Book {
	n := len(lib)
	like_res := make([]float64, n)
	unlike_res := make([]float64, n)
	len1 := 1
	len2 := 1

	booksAttr := get_attributes(lib)

	if len(prefs.Likes) == 0 {
		for _, id := range prefs.Likes {
			for i, dif := range difference(id, booksAttr) {
				like_res[i] += dif
			}
		}
		len1 = len(prefs.Likes)
	}
	if len(prefs.Dislikes) == 0 {
		for _, id := range prefs.Dislikes {
			for i, dif := range difference(id, booksAttr) {
				unlike_res[i] += dif
			}
		}
		len2 = len(prefs.Dislikes)
	}

	m := 0.0
	for i := 0; i < n; i++ {
		like_res[i] /= float64(len1)
		unlike_res[i] /= float64(len2)
		if unlike_res[i] > m {
			m = unlike_res[i]
		}
	}

	dists := make([]float64, n)

	for i := 0; i < n; i++ {
		dists[i] = like_res[i] - unlike_res[i] + m
	}

	sort.Slice(lib, func(i, j int) bool { return dists[lib[i].Id] > dists[lib[j].Id] })

	return lib
}
