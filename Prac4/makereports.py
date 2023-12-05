import socket
import json
from datetime import datetime,timedelta
import time



def handle_socket_error(error):
    print(f"Error: {error}")
    time.sleep(1)  # Wait for a second before attempting to reconnect


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
            if (time - reports[url]["Time"]) < interval:  # Check if time difference is greater than 1 minute
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

# Create a TCP connection to the server on port 6379
# Create a TCP connection to the server on port 6379
server_address = ('localhost', 6379)
conn = socket.create_connection(server_address)
data=''
# Send QPOP requests until a non-error response is received
while True:
    try:
        request_message = 'QPOP'
        conn.send(request_message.encode())
        time.sleep(0.5)
        response = conn.recv(2048)
        
        # Debugging print
        #print("Received response:", response.decode('utf-8'))

        
        if response.decode() == "ERROR":
            break
        data+=response.decode()
        print(data)
    except socket.error as e:
        handle_socket_error(e)

# Close the connection
data=data[:-1]
data0='['
data+=']'
data2=data0+data
print()
print()
print(data2)
conn.close()
json_data = json.loads(data2)
# Write json_data to a file named 'statis.json'
with open('statis.json', 'w') as f:
    json.dump(json_data, f, indent=4)
generate_json_report("statis.json",("report.json"))



















































































# import json
# from datetime import datetime, timedelta
# import requests
# import socket

# def getHostIP():
#     return socket.gethostbyname(socket.gethostname())




# def send_tcp_request(action, data):
#     # Connect to the server
#     server_address = ('localhost', 6379)
#     try:
#         conn = socket.create_connection(server_address)
#     except ConnectionRefusedError as e:
#         print(f"Error connecting to the server: {e}")
#         return

#     # Initialize the request_message variable
#     request_message = ''

#     # Send the request
#     if action == 'QPOP':
#         request_message = "QPOP"
#     else:
#         print("Invalid action")
#         return

#     # Encode the string to bytes before sending
#     conn.send(request_message.encode())

#     # Receive the response
#     response = b''
#     while True:
#         data = conn.recv(1024)
#         if not data:
#             break
#         response += data

#     conn.close()

#     with open('statistic.json', 'w') as f:
#         f.write(response.decode())

#     generate_json_report("statistic.json", "report.json")

#     # Print the report
#     with open("report.json", "r") as f:
#         print(f.read())


# def find_pid_for_url(url, reports):
#     for record in reversed(list(reports.values())):
#         if record["URL"] == url:
#             return record["Id"]
#     return "null"

# def generate_json_report(filename, json_file):
#     with open(filename, 'r') as file:
#         data = json.load(file)

#     time_format = "%Y-%m-%d %H:%M:%S"
#     interval = timedelta(minutes=1)

#     reports = {}
#     id = 1
    
#     for entry in data:
#         url = entry["URL"]
#         time = datetime.strptime(entry["Time"], time_format)
#         output_data = entry.get("IP", "null")
#         if url in reports:
#             if (time - reports[url]["Time"]) < interval:  # Check if time difference is greater than 1 minute
#                 reports[url]["Count"] += 1
#             else:
#                 record = {
#                     "Id": id,
#                     "Pid": find_pid_for_url(url, reports),
#                     "URL": url,
#                     "SourceIP": output_data,
#                     "Time": time, 
#                     "Count": 1
#                 }
#                 id += 1
#                 reports[url + str(id)] = record  
#         else:
#             record = {
#                 "Id": id,
#                 "Pid": find_pid_for_url(url, reports),
#                 "URL": url,
#                 "SourceIP": output_data,
#                 "Time": time,  
#                 "Count": 1
#             }
#             id += 1
#             reports[url] = record

#     with open(json_file, 'w') as f:
#         json.dump(list(reports.values()), f, default=str, indent=4)

# def read_json_from_url(url):
#     response = requests.get(url)
#     if response.status_code == 200:
#         data = json.loads(response.text)
#         records = []
#         for entry in data:
#             record = {
#                 "URL": entry["URL"],
#                 "IP": getHostIP(),
#                 "Time": entry["Time"]
#             }
#             records.append(record)
#         with open('outputForReport.json', 'w') as f:
#             json.dump(records, f, indent=4)
#     else:
#         print("error")


# def get_report():
#    # url = f"http://localhost:3333/getreport"
#     send_tcp_request('QPOP')
#     generate_json_report("statistic.json", "report.json")