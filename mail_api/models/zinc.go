package models

type SearchUserResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source struct {
				Username string `json:"username"`
				Password string `json:"password"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type ExistsIndexResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}
