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

func estimateMinutes(route []StepInfo) float64 {
	const avgMinutesPerStop = 2.5
	const interchangePenaltyMin = 4.0

	stops := len(route) - 1
	interchanges := 0

	var lastLine string

	for i := 1; i < len(route); i++ {
		currentLine := route[i].Line

		if lastLine != "" && currentLine != lastLine {
			interchanges++
		}

		lastLine = currentLine
	}

	travelMin := float64(stops) * avgMinutesPerStop
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