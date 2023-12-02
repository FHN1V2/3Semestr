import json
from datetime import datetime, timedelta
import requests
import socket

def getHostIP():
    return socket.gethostbyname(socket.gethostname())


def find_pid_for_url(url, reports):
    for record in reversed(list(reports.values())):
        if record["URL"] == url:
            return record["Id"]
    return "null"

def generate_json_report(filename, json_file):
    with open(filename, 'r') as file:
        data = json.load(file)

    time_format = "%Y-%m-%d %H:%M:%S"
    interval = timedelta(minutes=1)

    reports = {}
    id = 1
    
    for entry in data:
        url = entry["URL"]
        time = datetime.strptime(entry["Time"], time_format)
        output_data = entry.get("IP", "null")
        if url in reports:
            if (time - reports[url]["Time"]) < interval:  # Check if time difference is greater than 1 minute
                reports[url]["Count"] += 1
            else:
                record = {
                    "Id": id,
                    "Pid": find_pid_for_url(url, reports),
                    "URL": url,
                    "SourceIP": output_data,
                    "Time": time,  # Add the "Time" key to the record
                    "Count": 1
                }
                id += 1
                reports[url + str(id)] = record  # Use a unique key for the new record
        else:
            record = {
                "Id": id,
                "Pid": find_pid_for_url(url, reports),
                "URL": url,
                "SourceIP": output_data,
                "Time": time,  # Add the "Time" key to the record
                "Count": 1
            }
            id += 1
            reports[url] = record

    with open(json_file, 'w') as f:
        json.dump(list(reports.values()), f, default=str, indent=4)

def read_json_from_url(url):
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
        with open('outputForReport.json', 'w') as f:
            json.dump(records, f, indent=4)
    else:
        print("error")


def get_report():
    url = f"http://{getHostIP()}:3333/getreport"
    read_json_from_url(url)
    generate_json_report("outputForReport.json", "report.json")