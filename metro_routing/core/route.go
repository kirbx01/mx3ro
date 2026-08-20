package main

import "fmt"

type StepInfo struct {
	Station string
	Line    string
}

func bfs(graph map[string][]Neighbor, source, destination string) []StepInfo {
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

func findAllRoutes(graph map[string][]Neighbor, current, destination string, visited map[string]bool, path []StepInfo, routes *[][]StepInfo, maxHops int) {
	if len(path) > maxHops {
		return
	}

	visited[current] = true

	if current == destination {
		routeCopy := append([]StepInfo{}, path...)
		*routes = append(*routes, routeCopy)
	} else {
		for _, neighbor := range graph[current] {
			if !visited[neighbor.Station] {
				newPath := append(path, StepInfo{Station: neighbor.Station, Line: neighbor.Line})
				findAllRoutes(graph, neighbor.Station, destination, visited, newPath, routes, maxHops)
			}
		}
	}

	visited[current] = false
}

func estimateMinutes(route []StepInfo, distLookup map[string]float64) float64 {
	const avgSpeedKmph = 33.0
	const interchangePenaltyMin = 3.0

	totalKm := 0.0
	interchanges := 0
	var lastLine string

	for i := 1; i < len(route); i++ {
		prevDist := distLookup[route[i-1].Station]
		currDist := distLookup[route[i].Station]

		diff := currDist - prevDist
		if diff < 0 {
			diff = -diff
		}
		totalKm += diff

		if route[i].Line != lastLine && i > 1 {
			interchanges++
		}
		lastLine = route[i].Line
	}

	travelMin := (totalKm / avgSpeedKmph) * 60
	return travelMin + float64(interchanges)*interchangePenaltyMin
}

func formatDuration(totalMinutes float64) string {
	mins := int(totalMinutes + 0.5)
	hours := mins / 60
	remainder := mins % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, remainder)
	}
	return fmt.Sprintf("%dm", remainder)
}