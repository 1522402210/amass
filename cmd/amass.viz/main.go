// Copyright 2017 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/OWASP/Amass/amass/handlers"
	"github.com/OWASP/Amass/amass/utils/viz"
)

var (
	help           = flag.Bool("h", false, "Show the program usage message")
	input          = flag.String("i", "", "The Amass data operations JSON file")
	visjspath      = flag.String("visjs", "", "Path to the Visjs output HTML file")
	graphistrypath = flag.String("graphistry", "", "Path to the Graphistry JSON file")
	gexfpath       = flag.String("gexf", "", "Path to the Gephi Graph Exchange XML Format (GEXF) file")
	d3path         = flag.String("d3", "", "Path to the D3 v4 force simulation HTML file")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Printf("Usage: %s -i infile --visjs of1 --gexf of2 --d3 of3 --graphistry of4\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		return
	}

	if *input == "" {
		fmt.Println("The data operations JSON file must be provided using the '-i' flag")
		return
	}

	f, err := os.Open(*input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	opts, err := handlers.ParseDataOpts(f)
	if err != nil {
		fmt.Println("Failed to parse the provided data operations")
		return
	}

	graph := handlers.NewGraph()
	err = handlers.DataOptsDriver(opts, graph)
	if err != nil {
		fmt.Printf("Failed to build the network graph: %v\n", err)
		return
	}

	nodes, edges := graph.VizData()
	WriteVisjsFile(*visjspath, nodes, edges)
	WriteGraphistryFile(*graphistrypath, nodes, edges)
	WriteGEXFFile(*gexfpath, nodes, edges)
	WriteD3File(*d3path, nodes, edges)
}

func WriteVisjsFile(path string, nodes []viz.Node, edges []viz.Edge) {
	if path == "" {
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	viz.WriteVisjsData(nodes, edges, f)
	f.Sync()
}

func WriteGraphistryFile(path string, nodes []viz.Node, edges []viz.Edge) {
	if path == "" {
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	viz.WriteGraphistryData(nodes, edges, f)
	f.Sync()
}

func WriteGEXFFile(path string, nodes []viz.Node, edges []viz.Edge) {
	if path == "" {
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	viz.WriteGEXFData(nodes, edges, f)
	f.Sync()
}

func WriteD3File(path string, nodes []viz.Node, edges []viz.Edge) {
	if path == "" {
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	viz.WriteD3Data(nodes, edges, f)
	f.Sync()
}
