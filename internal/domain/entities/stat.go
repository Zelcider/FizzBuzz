package entities

type Stat struct {
	Request StatRequest
	Hits    int64
}

type StatRequest struct {
	Limit int64
	Int1  int64
	Str1  string
	Int2  int64
	Str2  string
}
