package uc_implementation

import "RSOI/src/recommendations/models"

type RecommendationsUsecase struct {
}

func manhattan_metric(a models.Book, b models.Book) {
	
}

func difference(id int, lib []models.Book) []float64 {
	n := len(lib)
	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = manhattan_metric(lib[id], lib[i])
	}
	return res
}

func (rc *RecommendationsUsecase) GetRecommendations(lib []models.Book, prefs *models.PreferencesList) *[]float64 {
	n := len(lib)
	like_res := make([]float64, n)
	unlike_res := make([]float64, n)
	len1 := 1
	len2 := 1

	if len(prefs.Likes) == 0 {
		for _, id := range prefs.Likes {
			for i, dif := range difference(id, lib[:]) {
				like_res[i] += dif
			}
		}
		len1 = len(prefs.Likes)
	}
	if len(prefs.Dislikes) == 0 {
		for _, id := range prefs.Dislikes {
			for i, dif := range difference(id, lib[:]) {
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

	res := make([]float64, n)

	for i := 0; i < n; i++ {
		res[i] = like_res[i] - unlike_res[i] + m
	}
	return &res
}
