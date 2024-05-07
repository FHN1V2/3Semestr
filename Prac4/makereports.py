import socket
import json
from datetime import datetime, timedelta
from flask import Flask, send_from_directory,send_file
import time


app = Flask(__name__)

def handle_socket_error(error):
    print(f"Error: {error}")
    time.sleep(1)  

def find_pid_for_url(url, reports):
    for record in reversed(list(reports.values())):
        if record["URL"] == url:
            return record["Id"]
    return "null"

def generate_json_report(filename, json_file):
    with open(filename, 'r') as file:
        data = json.load(file)

    time_format = "%Y-%m-%d%H:%M:%S"
    interval = timedelta(minutes=1)

    reports = {}
    id = 1
    
    for entry in data:
        url = entry["URL"]
        time = datetime.strptime(entry["Time"], time_format)
        output_data = entry.get("IP", "null")
        if url in reports:
            if (time - reports[url]["Time"]) < interval:  
                reports[url]["Count"] += 1
            else:
                record = {
                    "Id": id,
                    "Pid": find_pid_for_url(url, reports),
                    "URL": url,
                    "SourceIP": output_data,
                    "Time": time, 
                    "Count": 1
                }
                id += 1
                reports[url + str(id)] = record  
        else:
            record = {
                "Id": id,
                "Pid": find_pid_for_url(url, reports),
                "URL": url,
                "SourceIP": output_data,
                "Time": time,  
                "Count": 1
            }
            id += 1
            reports[url] = record

    with open(json_file, 'w') as f:
        json.dump(list(reports.values()), f, default=str, indent=4)


def make_statistic():
    server_address = ('localhost', 6379)
    conn = socket.create_connection(server_address)
    data=''

    while True:
        try:
            request_message = 'QPOP'
            conn.send(request_message.encode())
            time.sleep(0.5)
            response = conn.recv(2048)



            
            if response.decode() == "ERROR":
                break
            data+=response.decode()
            #print(data)
        except socket.error as e:
            handle_socket_error(e)


    data=data[:-1]
    data0='['
    data+=']'
    data2=data0+data
    #print()
    #print()
    #print(data2)
    conn.close()
    json_data = json.loads(data2)

    with open('statis.json', 'w') as f:
        json.dump(json_data, f, indent=4)


@app.route('/sendreport', methods=['GET'])
def send_report():
    make_statistic()
    generate_json_report("statis.json", "report.json")

    with open('report.json', 'r') as f:
        report_data = f.read()

    return report_data


while True:
    if __name__ == '__main__':
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_address = ('10.241.84.120', 6060)
        server_socket.bind(server_address)
        server_socket.listen(1)

        print("Waiting for connection...")
        client_socket, client_address = server_socket.accept()
        print(f"Connection from {client_address}")
            
        data = client_socket.recv(1024).decode()
            
        if data == "REPORT":
            print("Generating report...")
            response = send_report()
            client_socket.send(response.encode())
            client_socket.close()
        else:
            print("Invalid request")
        #client_socket.close()
