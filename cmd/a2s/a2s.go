// Copyright 2012 - 2018 The txt2svg Contributors
// All rights reserved.

// The t2s tool generates SVG output given an txt diagram input.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pippi101/txt2svg"
)

const logo = ` .-------------------------.
 |                         |
 |         .-----. .-----. |
 |         '--'  | |  '--' |
 |         '  '--' '--'  ' |
 |  t      '-----' '-----' |
 |  txt     2      svg   ' |
 |                         |
 '-------------------------'

https://github.com/pippi101

[1,0]: {"fill":"#88d","t2s:delref":1}
`

func mainImpl() error {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", logo)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	in := flag.String("i", "-", "Path to input text file. If set to \"-\" (hyphen), stdin is used.")
	out := flag.String("o", "-", "Path to output SVG file. If set to \"-\" (hyphen), stdout is used.")
	noBlur := flag.Bool("b", false, "Disable drop-shadow blur.")
	font := flag.String("f", "Consolas,Monaco,Anonymous Pro,Anonymous,Bitstream Sans Mono,monospace", "Font family to use.")
	scaleX := flag.Int("x", 9, "X grid scale in pixels.")
	scaleY := flag.Int("y", 16, "Y grid scale in pixels.")
	tabWidth := flag.Int("t", 8, "Tab width.")
	doLogo := flag.Bool("L", false, "Generate SVG of t2s logo.")
	flag.Parse()

	var input []byte
	var err error
	if *doLogo {
		input = []byte(logo)
	} else {
		if *in == "-" {
			input, err = ioutil.ReadAll(os.Stdin)
		} else {
			input, err = ioutil.ReadFile(*in)
		}
	}
	if err != nil {
		return err
	}

	canvas, err := txt2svg.NewCanvas(input, *tabWidth, *noBlur)
	if err != nil {
		return err
	}
	svg := txt2svg.CanvasToSVG(canvas, *noBlur, *font, *scaleX, *scaleY)
	if *out == "-" {
		_, err := os.Stdout.Write(svg)
		return err
	}
	return ioutil.WriteFile(*out, svg, 0666)
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "t2s: %s\n", err)
		os.Exit(1)
	}
}
