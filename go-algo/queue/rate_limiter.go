package queue

type TimestampMessage struct {
	Message   string
	Timestamp int
}

type Limiter struct {
	Queue    []*TimestampMessage
	Messages map[string]bool
	MinTime  int
}

func New(minTime int) *Limiter {
	return &Limiter{
		Queue:   make([]*TimestampMessage, 0),
		MinTime: minTime,
	}
}

func (l *Limiter) ShouldLimit(timestamp int, message string) bool {
	for len(l.Queue) > 0 && timestamp-l.Queue[0].Timestamp > l.MinTime {
		mesg := l.Queue[0].Message
		l.Queue = l.Queue[1:]
		delete(l.Messages, mesg)
	}
	result := true
	if len(l.Queue) == 0 {
		result = true
	}
	if l.Messages[message] {
		result = false
	}
	l.Queue = append(l.Queue, &TimestampMessage{
		Message:   message,
		Timestamp: timestamp,
	})
	l.Messages[message] = true
	return result
}
