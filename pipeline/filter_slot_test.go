package pipeline

import (
	"fmt"
	"log"
	"testing"
)

type Context struct {
	Param1 string
	Param2 string
	Param3 string
}

// Filter provides Process method, Filter means a single function block,
// for example, it you are going to take a swim, first you need to warm up yourself,
// second you need to change your suit, and finally you can dive into water.
// As for swimming action, your Filter list will be Filter(warm up) -> Filter(change your suit) -> Filter(swim)
type Filter interface {
	Process(ctx *Context) error
}

type Warm struct{}

func (w *Warm) Process(ctx *Context) error {
	fmt.Printf("I am warming myself up, now context: param1: %s; param2: %s; param3: %s\n", ctx.Param1, ctx.Param2, ctx.Param3)
	ctx.Param1 = "Warming up!"
	return nil
}

type ChangeSuit struct{}

func (c *ChangeSuit) Process(ctx *Context) error {
	fmt.Printf("I am changing suit, now context: param1: %s; param2: %s; param3: %s\n", ctx.Param1, ctx.Param2, ctx.Param3)
	ctx.Param2 = "Changing suit!"
	return nil
}

type Swim struct{}

func (s *Swim) Process(ctx *Context) error {
	fmt.Printf("I am going to swim, now context: param1: %s; param2: %s; param3: %s\n", ctx.Param1, ctx.Param2, ctx.Param3)
	ctx.Param3 = "Swimming!"
	return nil
}

type GoHome struct{}

func (g *GoHome) Process(ctx *Context) error {
	fmt.Printf("I am going home, now context: param1: %s; param2: %s; param3: %s\n", ctx.Param1, ctx.Param2, ctx.Param3)
	return nil
}

func Pipeline(ctx *Context, filterList ...Filter) {
	for _, filter := range filterList {
		if err := filter.Process(ctx); err != nil {
			log.Printf("Error: %s\n", err)
		}
	}
}

// TestFilterSlotPipeline test for pipeline version 2
//
// About pipeline version 2: now you can define you own pipeline by combining
// different Filter as you wish.
func TestFilterSlotPipeline(t *testing.T) {
	var ctx Context
	Pipeline(&ctx, &Warm{}, &ChangeSuit{}, &Swim{}, &GoHome{})
}
