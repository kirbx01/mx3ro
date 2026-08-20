package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Station struct {
	Name        string  `json:"name"`
	Line        string  `json:"line"`
	Dist        float64 `json:"dist"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Connections *string `json:"connections"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Line string `json:"line"`
}

type MetroData struct {
	Stations      []Station              `json:"stations"`
	Edges         []Edge                 `json:"edges"`
	Interchanges  map[string][]string    `json:"interchanges"`
}
type Neighbor struct {
	Station string
	Line    string
}

func buildGraph(edges []Edge) map[string][]Neighbor {
	graph := make(map[string][]Neighbor)

	for _, e := range edges {
		graph[e.From] = append(graph[e.From], Neighbor{Station: e.To, Line: e.Line})
	}

	return graph
}

type StepInfo struct {
	Station string
	Line    string
}

func bfs(graph map[string][]Neighbor, source, destination string) []StepInfo { //since 285 stations aint big we using bfs would scale after learning others 
	visited := map[string]bool{source: true}
	queue := [][]StepInfo{{{Station: source, Line: ""}}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1].Station

		if current == destination {
			return path
		}

		for _, neighbor := range graph[current] {
			if !visited[neighbor.Station] {
				visited[neighbor.Station] = true
				newPath := append([]StepInfo{}, path...)
				newPath = append(newPath, StepInfo{Station: neighbor.Station, Line: neighbor.Line})
				queue = append(queue, newPath)
			}
		}
	}

	return nil 
}

func buildNameLookup(stations []Station) map[string]string {
	lookup := make(map[string]string)
	for _, s := range stations {
		lookup[strings.ToLower(s.Name)] = s.Name
	}
	return lookup
}

func main() {
	data, err := os.ReadFile("/home/pansi/mx3ro/metro_routing/stations.json")
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

	color.New(color.FgHiCyan, color.Bold).Println("🚇 Welcome to Delhi Metro Router")
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
	
	route := bfs(graph, source, destination)

	if route == nil {
		fmt.Println("No route found.")
	} else {
		fmt.Println("Route found:")
		for _, step := range route {
			fmt.Println(" -", step.Station, step.Line)
		}
	}
}