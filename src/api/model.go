package api

type SearchRes struct {
	ID       string `json:"id" pterm:"2, right"`
	Name     string `json:"name" pterm:"2, elastic"`
	Leechers string `json:"leechers" pterm:"2, right, LE"`
	Seeders  string `json:"seeders" pterm:"SE, 3, right"`
	Size     string `json:"size" pterm:"right"`
	Username string `json:"username" pterm:"ignore"`
	Added    string `json:"added" pterm:"ignore"`
	Status   string `json:"status" pterm:"ignore"`
	Category string `json:"category" pterm:"ignore"`
}

type PageRes struct {
	ID           int         `json:"id"`
	Category     int         `json:"category"`
	Name         string      `json:"name"`
	NumFiles     int         `json:"num_files"`
	Size         int         `json:"size"`
	Seeders      int         `json:"seeders"`
	Leechers     int         `json:"leechers"`
	Username     string      `json:"username"`
	Added        int64       `json:"added"`
	Descr        string      `json:"descr"`
	InfoHash     string      `json:"info_hash"`
}