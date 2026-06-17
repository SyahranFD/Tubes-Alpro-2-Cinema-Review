package main

func SeedMovies() (Movies, int) {
	var m Movies
	data := []Movie{
		{"The_Shawshank_Redemption", 1994, "Drama", "Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.", 9.3},
		{"The_Godfather", 1972, "Crime", "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.", 9.2},
		{"Inception", 2010, "Sci-Fi", "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.", 8.8},
		{"Spirited_Away", 2001, "Horror", "During her family's move to the suburbs, a sullen 10-year-old girl wanders into a world ruled by gods, witches, and spirits, and where humans are changed into beasts.", 8.6},
		{"Parasite", 2019, "Horror", "Greed and class discrimination threaten the newly formed symbiotic relationship between the wealthy Park family and the destitute Kim clan.", 8.6},
	}
	for i, v := range data {
		m[i] = v
	}
	return m, len(data)
}

// code di cineReview.go
// tabMovies, n = SeedMovies()
