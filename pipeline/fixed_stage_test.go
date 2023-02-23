package pipeline

import (
	"fmt"
	"log"
	"testing"
)

// Stage provides Process method and NextStage method,
// like a linked list, it represent not only itself, but also the next stage.
type Stage interface {
	Process(param ...string) error
	NextStage() Stage
}

// Start simulate the first stage of pipeline
type Start struct{}

func (s *Start) Process(param ...string) error {
	fmt.Printf("Here is Start processing task, param: %+v\n", param)
	return nil
}

func (s *Start) NextStage() Stage {
	return &Middle{}
}

// Middle simulate the second stage of pipeline
type Middle struct{}

func (m *Middle) Process(param ...string) error {
	fmt.Printf("Here is Middle processing task, param: %+v\n", param)
	return nil
}

func (m *Middle) NextStage() Stage {
	return &End{}
}

// End simulate the final stage of pipeline
type End struct{}

func (e *End) Process(param ...string) error {
	fmt.Printf("Here is End processing task, param: %+v\n", param)
	return nil
}

func (e *End) NextStage() Stage {
	return nil
}

// TestFixedStagePipeline test for pipeline version 1
//
// About pipeline version 1: every stage of pipeline is fixed, which means you can't
// make your own pipeline flexibly.
func TestFixedStagePipeline(t *testing.T) {
	var nowStage Stage
	nowStage = &Start{}
	for nowStage != nil {
		if err := nowStage.Process("Hello", "World", "!"); err != nil {
			log.Printf("Error: %s", err)
		}
		nowStage = nowStage.NextStage()
	}
}
