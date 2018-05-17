package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/podhmo/go-tools/loader"
	xloader "golang.org/x/tools/go/loader"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	// disable gc
	debug.SetGCPercent(-1)

	fmt.Println("##", "load once")
	for i := 0; i < 3; i++ {
		if err := run0(); err != nil {
			return err
		}
	}
	fmt.Println("##", "load twice (without cache)")
	for i := 0; i < 3; i++ {
		if err := run1(); err != nil {
			return err
		}
	}
	fmt.Println("##", "load twice (with cache)")
	for i := 0; i < 3; i++ {
		if err := run2(); err != nil {
			return err
		}
	}
	return nil
}

func run0() error {
	c := loader.Config{}
	c.Import("github.com/podhmo/go-tools/loader")
	st := time.Now()
	p, err := c.Load()
	if err != nil {
		return err
	}
	fmt.Println(len(p.AllPackages), "packages", time.Since(st))
	return nil
}

func run1() error {
	st := time.Now()
	{
		c := xloader.Config{}
		c.Import("github.com/podhmo/go-tools/loader")
		_, err := c.Load()
		if err != nil {
			return err
		}
	}
	{
		c := xloader.Config{}
		c.Import("github.com/podhmo/go-tools/loader")
		p, err := c.Load()
		if err != nil {
			return err
		}
		fmt.Println(len(p.AllPackages), "packages", time.Since(st))
	}
	return nil
}

func run2() error {
	st := time.Now()
	c := loader.Config{}
	c.Import("github.com/podhmo/go-tools/loader")
	p, err := c.Load()
	if err != nil {
		return err
	}

	{
		c := loader.Config{}
		c.Import("github.com/podhmo/go-tools/loader")
		p, err := c.LoadWith(p)
		if err != nil {
			return err
		}
		fmt.Println(len(p.AllPackages), "packages", time.Since(st))
	}
	return nil
}
