import json
from collections import defaultdict
from datetime import datetime
import requests
import socket


def get_detail():
    def getHostIP():
        return socket.gethostbyname(socket.gethostname())


    url = f"http://{getHostIP()}:3333/getreport"

    response = requests.get(url)
    if response.status_code == 200:
        data = json.loads(response.text)
        records = []
        for entry in data:
            record = {
                "URL": entry["URL"],
                "IP": getHostIP(),
                "Time": entry["Time"]
            }
            records.append(record)
        with open('outputForDetail.json', 'w') as f:
            json.dump(records, f, indent=4)
    else:
        print("error")
    url = f"http://{getHostIP()}:3333/getreport"
        

    with open('outputForDetail.json', 'r') as f:
        data = json.load(f)

    # Create report
    report = defaultdict(lambda: defaultdict(lambda: defaultdict(int)))
    for entry in data:
        url = entry['URL']
        source_ip = entry['IP']
        time = datetime.strptime(entry["Time"], "%Y-%m-%d %H:%M:%S")
        minute_interval = time.minute
        time_interval = f"{time.hour:02d}:{minute_interval:02d}-{time.hour:02d}:{(minute_interval+1):02d}"
        report[source_ip][time_interval][url] += 1

    # Format report
    formatted_report = {}
    for source_ip, time_data in report.items():
        source_ip_data = {}
        for time_interval, url_data in time_data.items():
            interval_data = {}
            interval_data["Total"] = sum(url_data.values())
            interval_data["URLS"] = {url: f"({count})" for url, count in url_data.items()}
            source_ip_data[time_interval] = interval_data
        formatted_report[source_ip] = source_ip_data

    # Write report to a JSON file
    with open('Detail.json', 'w') as f:
        json.dump(formatted_report, f, indent=4)



