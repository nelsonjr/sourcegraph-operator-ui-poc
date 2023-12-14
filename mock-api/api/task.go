package api

import (
	"math/rand"
	"time"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Started     bool      `json:"started"`
	Finished    bool      `json:"finished"`
	Weight      int       `json:"weight"`
	Progress    int       `json:"progress"`
	LastUpdate  time.Time `json:"lastUpdate"`
}

func createFakeTasks() []Task {
	return []Task{
		{
			Title:       "Install",
			Description: "Install Sourcegraph",
			Started:     false,
			Finished:    false,
			Weight:      5,
		},
		{
			Title:       "Configure",
			Description: "Configure Sourcegraph",
			Started:     false,
			Finished:    false,
			Weight:      8,
		},
		{
			Title:       "Database",
			Description: "New Database",
			Started:     false,
			Finished:    false,
			Weight:      7,
		},
		{
			Title:       "Migrate",
			Description: "Run Migration Tasks",
			Started:     false,
			Finished:    false,
			Weight:      22,
		},
		{
			Title:       "Start",
			Description: "Start Sourcegraph",
			Started:     false,
			Finished:    false,
			Weight:      11,
		},
	}
}

func createFakeUpgradeTasks() []Task {
	return []Task{
		{
			Title:       "Upgrade",
			Description: "Upgrade Sourcegraph",
			Started:     false,
			Finished:    false,
			Weight:      5,
		},
		{
			Title:       "Migrate",
			Description: "Run migration tasks",
			Started:     false,
			Finished:    false,
			Weight:      13,
		},
	}
}

func progressTasks(tasks []Task) []Task {
	var result []Task

	var previousStarted bool = true
	var previousFinished bool = true

	for _, task := range tasks {
		var beforeStarted bool = task.Started
		task.Started = previousFinished && (task.Started || (previousStarted && rand.Intn(2) == 0))
		previousStarted = task.Started
		task.Finished = beforeStarted && (task.Progress == 100)
		previousFinished = task.Finished
		task.LastUpdate = time.Now()

		result = append(result, task)
	}

	return result
}

func calculateProgress(tasks []Task) ([]Task, int) {
	var result []Task

	var taskWeights int = 0
	for _, t := range installTasks {
		taskWeights += t.Weight
	}

	var progress float32 = 0

	for _, t := range installTasks {
		if t.Finished {
			progress += float32(t.Weight)
		} else if t.Started {
			if installError == "" {
				var dieThrow int = rand.Intn(t.Weight)
				var delta float32 = float32(dieThrow) / float32(t.Weight) * 100 * rand.Float32()
				t.Progress += int(delta)
			}
			if t.Progress > 100 {
				t.Progress = 100
			}
			progress += float32(t.Weight * t.Progress / 100)
		}

		result = append(result, t)
	}

	return result, int(progress / float32(taskWeights) * 100)
}
