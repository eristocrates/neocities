package main

import "os"

// Song struct
type Song struct {
	SongID   string `json:"song_id"`
	SongName string `json:"song_name"`
}

// Team struct
type Team struct {
	TeamID   string `json:"team_id"`
	TeamName string `json:"team_name"`
	Songs    []Song `json:"songs"`
}

// Event struct
type Event struct {
	Id                int    `json:"id"`
	FullName          string `json:"fullName"`
	AbbrevName        string `json:"abbrevName"`
	ShortName         string `json:"shortName"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	BannerType1       string `json:"bannerType1"`
	RegistrationStart string `json:"registrationStart"`
	RegistrationEnd   string `json:"registrationEnd"`
	ImpressionStart   string `json:"impressionStart"`
	ImpressionEnd     string `json:"impressionEnd"`
	PeriodStart       string `json:"periodStart"`
	PeriodEnd         string `json:"periodEnd"`
	EntryCount        int    `json:"entryCount"`
	ImpressionCount   int    `json:"impressionCount"`
	InfoLink          string `json:"infoLink"`
	DetailLink        string `json:"detailLink"`
	ListLink          string `json:"listLink"`
	IsBof             bool   `json:"isBof"`
	LogoType1         string `json:"logoType1"`
	LogoType2         string `json:"logoType2"`
	LogoType3         string `json:"logoType3"`
	LogoType4         string `json:"logoType4"`
	LogoType5         string `json:"logoType5"`
	TitleJpg          string `json:"titleJpg"`
	BannerType2       string `json:"bannerType2"`
	Video             string `json:"video"`
	HeaderJpg         string `json:"headerJpg"`
	HeaderPng         string `json:"headerPng"`
	BackJpg           string `json:"backJpg"`
	BackPng           string `json:"backPng"`
	IsModern          bool   `json:"isModern"`
	Teams             []Team `json:"teams"`
}

type CommaWriter struct {
	file      *os.File
	needComma bool
}
