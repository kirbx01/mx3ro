package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"mx3ro/ui"
	"github.com/fatih/color"
)

func main() {
	data, err := os.ReadFile("/home/pansi/mx3ro/metro_routing/core/stations.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var metro MetroData
	err = json.Unmarshal(data, &metro)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	graph := buildGraph(metro.Edges)
	nameLookup := buildNameLookup(metro.Stations)
	distLookup := buildDistLookup(metro.Stations)
	color.New(color.FgHiCyan, color.Bold).Println("waddup mx3ro user જ⁀➴")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter source station: ")
	source, _ := reader.ReadString('\n')
	source = strings.TrimSpace(source)

	fmt.Print("Enter destination station: ")
	destination, _ := reader.ReadString('\n')
	destination = strings.TrimSpace(destination)

	source, ok1 := nameLookup[strings.ToLower(source)]
	destination, ok2 := nameLookup[strings.ToLower(destination)]

	if !ok1 || !ok2 {
		fmt.Println("Invalid station name(s) entered.")
		return
	}

	shortest := bfs(graph, source, destination)
	if shortest == nil {
		fmt.Println("No route found.")
		return
	}
	shortestLen := len(shortest)

	visited := map[string]bool{}
	var routes [][]StepInfo
	startPath := []StepInfo{{Station: source, Line: ""}}
	findAllRoutes(graph, source, destination, visited, startPath, &routes, shortestLen+4)

	sort.Slice(routes, func(i, j int) bool {
		return len(routes[i]) < len(routes[j])
	})

	if len(routes) > 5 {
		routes = routes[:5]
	}

	fmt.Println()
	fmt.Printf("Found %d route(s):\n\n", len(routes))

	for i, route := range routes {
		minutes := estimateMinutes(route, distLookup)
		duration := formatDuration(minutes)
		color.New(color.FgHiWhite, color.Bold).Printf("Route %d (%d stops, ~%s):\n", i+1, len(route)-1, duration)
		var lastLine string
		for j, step := range route {
			printer := ui.LineColor(step.Line)
			if step.Line != lastLine && j > 0 {
				color.New(color.FgHiWhite, color.Bold).Printf("  [Interchange -> %s]\n", step.Line)
			}
			printer("  %s\n", step.Station)
			lastLine = step.Line
		}
		fmt.Println()
	}
}