import csv
import re
import json
from collections import defaultdict


def normalize_line(line):
    line = line.strip()
    if line == "Voilet line":
        line = "Violet line"
    return line


def parse_csv(path):
    """Read the raw CSV and return a clean list of station dicts."""
    stations = []

    with open(path, encoding='utf-8-sig') as f:
        reader = csv.DictReader(f)
        for row in reader:
            name = row['Station Names']
            line = normalize_line(row['Metro Line'])
            dist = row['Dist. From First Station(km)']
            lat = row['Latitude']
            lon = row['Longitude']

            # capture the [Conn: X,Y] text before we strip it
            match = re.search(r'\[Conn:\s*(.*?)\]', name)
            connections = match.group(1) if match else None

            clean_name = re.sub(r'\s*\[Conn:.*?\]', '', name).strip()

            stations.append({
                'name': clean_name,
                'line': line,
                'dist': float(dist),
                'lat': float(lat),
                'lon': float(lon),
                'connections': connections
            })

    return stations


def build_edges(stations):
    """
    Build graph edges two ways:
    1. Consecutive stations on the same line (sorted by distance)
    2. Interchange edges between same-named stations on different lines
    """
    edges = []

    # --- 1. Edges along each line, in physical order ---
    by_line = defaultdict(list)
    for s in stations:
        by_line[s['line']].append(s)

    for line, line_stations in by_line.items():
        line_stations.sort(key=lambda s: s['dist'])
        for i in range(len(line_stations) - 1):
            a = line_stations[i]
            b = line_stations[i + 1]
            edges.append({
                'from': a['name'],
                'to': b['name'],
                'line': line
            })
            edges.append({
                'from': b['name'],
                'to': a['name'],
                'line': line
            })

    # --- 2. Interchange edges: same station name, different lines ---
    by_name = defaultdict(set)
    for s in stations:
        by_name[s['name']].add(s['line'])

    interchange_stations = {name: lines for name, lines in by_name.items() if len(lines) > 1}

    for name, lines in interchange_stations.items():
        lines = list(lines)
        for i in range(len(lines)):
            for j in range(len(lines)):
                if i != j:
                    edges.append({
                        'from': name,
                        'to': name,
                        'line': f'interchange:{lines[i]}->{lines[j]}'
                    })

    return edges, interchange_stations


def main():
    stations = parse_csv('Delhi metro.csv')
    print(f'Parsed {len(stations)} station rows')

    edges, interchanges = build_edges(stations)
    print(f'Built {len(edges)} edges')
    print(f'Found {len(interchanges)} interchange stations')

    output = {
        'stations': stations,
        'edges': edges,
        'interchanges': {name: list(lines) for name, lines in interchanges.items()}
    }

    with open('stations.json', 'w', encoding='utf-8') as f:
        json.dump(output, f, indent=2, ensure_ascii=False)

    print('Wrote stations.json')


if __name__ == '__main__':
    main()