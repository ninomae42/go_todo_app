package store

import (
	"errors"

	"github.com/ninomae42/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskId]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskId
	Tasks  map[entity.TaskId]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskId, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}
