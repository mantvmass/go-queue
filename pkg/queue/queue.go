package queue

type Queue struct {
	jobs chan Job
}

func NewQueue(size int) *Queue {
	return &Queue{
		jobs: make(chan Job, size),
	}
}

func (q *Queue) AddJob(job Job) {
	q.jobs <- job
}

func (q *Queue) ProcessJobs() {
	for job := range q.jobs { // ใช้ range เพื่อดึงงานออกจาก channel ทีละงาน
		job.Callback() // เรียก Callback เพื่อประมวลผลงาน
		// งานที่ประมวลผลเสร็จจะถูกลบออกจาก channel โดยอัตโนมัติ
	}
}
