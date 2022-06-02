package cache

import (
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
	"sync"
	"time"
)

type Question struct {
	q       *chgk_api_client.Question
	expired int64
}

type QuestionCache struct {
	m    map[uint64]*Question
	wg   sync.WaitGroup
	mu   sync.Mutex
	stop chan struct{}
}

func New(cleanupInterval time.Duration) *QuestionCache {
	m := make(map[uint64]*Question)
	qc := &QuestionCache{
		m: m,
	}
	qc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer qc.wg.Done()
		qc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return qc
}

func (q *QuestionCache) Put(userID uint64, question *chgk_api_client.Question, expired int64) {
	q.mu.Lock()
	q.m[userID] = &Question{question, expired}
	q.mu.Unlock()
}

func (q *QuestionCache) Get(userID uint64) *chgk_api_client.Question {
	q.mu.Lock()
	question, exists := q.m[userID]
	if exists && q != nil {
		return question.q
	}
	return nil
}

func (q *QuestionCache) Delete(userID uint64) {
	q.mu.Lock()
	delete(q.m, userID)
	q.mu.Unlock()
}

func (q *QuestionCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-q.stop:
			return
		case <-t.C:
			q.mu.Lock()
			for uid, m := range q.m {
				if m.expired <= time.Now().Unix() {
					delete(q.m, uid)
				}
			}
			q.mu.Unlock()
		}
	}
}

func (q *QuestionCache) stopCleanup() {
	close(q.stop)
	q.wg.Wait()
}
