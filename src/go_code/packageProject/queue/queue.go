package queue

/**
 * @author  CoreyChen Zhang
 * @version  2021/4/9 17:01
 * @modified
 */

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pull() interface{} {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q Queue) String() string {
	return "The queue"
}

