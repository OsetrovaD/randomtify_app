package entities

import (
	"fmt"
	"strings"
)

type Alphabets []string

func (a *Alphabets) String() string {
	return strings.Join(*a, "\n")
}

type Song struct {
	Name   string
	Year   int
	Rating string
	Album  string
}

func (s *Song) String() string {
	return fmt.Sprintf("Name: %s    Year: %d    Rating: %s    Album: %s", s.Name, s.Year, s.Rating, s.Album)
}

type Artist struct {
	Name       string `json:"name"`
	Rating     string `json:"rating"`
	SpotifyUrl string `json:"spotifyUrl"`
}

func (a *Artist) String() string {
	return fmt.Sprintf("Name: %s    Rating: %s    Spotify URL: %s", a.Name, a.Rating, a.SpotifyUrl)
}

type ArtistExtendedInfo struct {
	Artist
	Songs []*Song
}

func (a *ArtistExtendedInfo) String() string {
	artistInfo := a.Artist.String()
	var res []string
	for _, song := range a.Songs {
		res = append(res, song.String())
	}
	songs := strings.Join(res, "\n")
	return fmt.Sprintf("%s\n\nSongs:\n%s", artistInfo, songs)
}

type AllArtistsInfo []Artist

func (a *AllArtistsInfo) String() string {
	var res []string
	for _, artist := range *a {
		res = append(res, artist.String())
	}
	return strings.Join(res, "\n")
}
