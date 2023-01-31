package structs

type Request struct {
	Id int `json:"id"`
}

type Task struct {
	WorkerId int `json:"workerId"`
}
