package cache

import (
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
	"sync"
	"time"
)

type question struct {
	q       *chgk_api_client.Question
	expired int64
}

type QuestionCache struct {
	m      map[uint64]*question
	wg     sync.WaitGroup
	mu     sync.Mutex
	stop   chan struct{}
	expire time.Duration
}

func New(cleanupInterval time.Duration, expire int) *QuestionCache {
	m := make(map[uint64]*question)
	qc := &QuestionCache{
		m:      m,
		expire: time.Duration(expire),
		stop:   make(chan struct{}),
	}
	qc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer qc.wg.Done()
		qc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return qc
}

func (qc *QuestionCache) Put(userID uint64, q *chgk_api_client.Question) {
	qc.mu.Lock()
	defer qc.mu.Unlock()

	qc.m[userID] = &question{q, time.Now().Add(qc.expire * time.Second).Unix()}
}

func (qc *QuestionCache) Get(userID uint64) *chgk_api_client.Question {
	qc.mu.Lock()
	defer qc.mu.Unlock()

	question, exists := qc.m[userID]
	if exists && qc != nil {
		if question.expired <= time.Now().Unix() {
			delete(qc.m, userID)
			return nil
		}
		return question.q
	}
	return nil
}

func (qc *QuestionCache) Delete(userID uint64) {
	qc.mu.Lock()
	defer qc.mu.Unlock()

	delete(qc.m, userID)
}

func (qc *QuestionCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-qc.stop:
			return
		case <-t.C:
			qc.mu.Lock()
			for userID, m := range qc.m {
				if m.expired <= time.Now().Unix() {
					delete(qc.m, userID)
				}
			}
			qc.mu.Unlock()
		}
	}
}

func (qc *QuestionCache) stopCleanup() {
	close(qc.stop)
	qc.wg.Wait()
}
