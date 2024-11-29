package queue

type Job struct {
	Callback func()
}
