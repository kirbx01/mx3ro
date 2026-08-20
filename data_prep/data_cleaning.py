import csv
import re

with open('Delhi metro.csv', encoding='utf-8-sig') as f:
    reader = csv.DictReader(f)
    for row in reader:
        print(row)
        for row in reader:
            name = row['Station Names']
            line = row['Metro Line']
            dist = row['Dist. From First Station(km)']
            lat = row['Latitude']
            lon = row['Longitude']
            
            clean_name = re.sub(r'\s*\[Conn:.*?\]', '', name).strip()

            print(clean_name, '|', line , '|', dist, '|', lat, '|', lon)
    