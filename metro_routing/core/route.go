package main

type StepInfo struct {
	Station string
	Line    string
}

func bfs(graph map[string][]Neighbor, source, destination string) []StepInfo {   //since 285 stations aint big we using bfs would scale after learning others 
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
