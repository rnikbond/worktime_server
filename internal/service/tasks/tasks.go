package tasks

type ITickerTask interface {
	Start()
	Stop()
}

type Tasks struct {
	ITickerTask
}
