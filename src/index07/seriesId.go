package index07

type SeriesId struct {
	Id   int32
	Time int64
}

func NewSeriesId(id int32, t int64) SeriesId {
	return SeriesId{
		Id:   id,
		Time: t,
	}
}
