package process

type JDRequest struct {
	Company        string `json:"company"`
	Role           string `json:"role"`
	Location       string `json:"location"`
	JobDescription string `json:"jobdescription"`
}
