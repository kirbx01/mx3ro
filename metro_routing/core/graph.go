package main

import "strings"

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
	Stations     []Station           `json:"stations"`
	Edges        []Edge              `json:"edges"`
	Interchanges map[string][]string `json:"interchanges"`
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

func buildNameLookup(stations []Station) map[string]string {
	lookup := make(map[string]string)
	for _, s := range stations {
		lookup[strings.ToLower(s.Name)] = s.Name
	}
	return lookup
}

func buildDistLookup(stations []Station) map[string]float64 {
	lookup := make(map[string]float64)
	for _, s := range stations {
		lookup[s.Name] = s.Dist
	}
	return lookup
}