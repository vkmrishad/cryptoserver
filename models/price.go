package models

type Price struct {
	Symbol     string      `json:"symbol"`
	Ask		   string      `json:"ask"`		
	Bid        string      `json:"bid"`
	Open	   string      `json:"open"`
	Last       string      `json:"last"`
	Low        string      `json:"low"`
	High       string      `json:"high"`
}

type Symbol struct {
	Id 				string 		`json:"id"`
	BaseCurrency 	string      `json:"baseCurrency"`
	FeeCurrency		string		`json:"feeCurrency"`
} 

type Currency struct {
	Id 				string 		`json:"id"`
	FullName 	    string      `json:"fullName"`
}

type TotalPrice struct {
	Id				string 		`json:"id"`
	FullName 	    string      `json:"fullName"`
	*Price
	FeeCurrency		string		`json:"feeCurrency"`
}