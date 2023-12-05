from flask import Flask, render_template, request, redirect,send_file,jsonify,send_from_directory
#from GetReport import get_report 
import datetime
import json
import random
import string
import socket
import time
#from GetDetail import get_detail




class JSONCreator:
    def __init__(self, filename):
        self.filename = filename
        self.data = self.load_data()

    def load_data(self):
        try:
            with open(self.filename, 'r') as file:
                return json.load(file)
        except (FileNotFoundError, json.JSONDecodeError) as e:
            print(f"Error loading data from {self.filename}: {e}")
            return []

    def add_data(self, url, ip, time):
        if isinstance(time, datetime.datetime):
            time = time.strftime('%Y-%m-%d %H:%M:%S')
        data_entry = {"URL": url, "IP": ip, "Time": time}
        self.data.append(data_entry)

    def create_json(self):
        with open(self.filename, 'w') as file:
            json.dump(self.data, file)



# def getHostIP():
#     return socket.gethostbyname(socket.gethostname())



def send_tcp_request(action, short_link, original_url):
    # Connect to the server
    server_address = ('localhost', 6379)
    try:
        conn = socket.create_connection(server_address)
    except ConnectionRefusedError as e:
        print(f"Error connecting to the server: {e}")
        return

    #conn = socket.create_connection(server_address)

    # Initialize the request_message variable
    request_message = ''

    # Send the request
    if action == 'HGET':
        request_message = f"HGET {short_link}"
    elif action == 'HPUSH':
        request_message = f"HPUSH {short_link} {original_url}"
    elif action=='QPUSH':
        request_message=f"QPUSH {short_link}"   
    else:
        print("Invalid action")

    # Encode the string to bytes before sending
    conn.send(request_message.encode())
    answer=conn.recv(1024)
    conn.close()
    return answer




app = Flask(__name__)
server_address = ('localhost', 3333)


@app.route('/', methods=['GET', 'POST'])
def home():
    generated_link = None
    if request.method == 'POST':
        original_link = request.form['user_input']

        short_link = generate_short_link()

        # Save the short link in the hash table
        send_tcp_request( 'HPUSH',short_link, original_link)

        generated_link = f"http://localhost:3333/{short_link}"

    return render_template('index.html', output_link=generated_link)

@app.route('/<short_link>')
def redirect_to_original(short_link):
    # Retrieve the original link from the hash table
    original_link = send_tcp_request('HGET',short_link,"")#short_link_table.search(short_link)
    original_link=original_link.decode()
    symbols_remove="'"
    for i in symbols_remove:
        original_link = original_link.replace(i,'')
    #print(original_link)
    if original_link:
        # Create a list with the required data
        data = [original_link, request.environ['REMOTE_ADDR'], str(datetime.datetime.now().replace(microsecond=0))]
        # Send a QPUSH request to the server with the data
        #send_tcp_request
        #send_tcp_request('QPUSH',json.dumps(data) , '')
        ip=request.environ['REMOTE_ADDR']
        send_tcp_request('QPUSH','{',"")
        #time.sleep(0.2)
        send_tcp_request('QPUSH',f'"URL":"{original_link}", "IP": "{ip}", "Time": "{datetime.datetime.now().replace(microsecond=0)}"' , '')
        #time.sleep(0.2)
        #send_tcp_request('QPUSH',f'"URL" : "{original_link}", "IP" : "{ip}", "Time" : "{datetime.datetime.now().replace(microsecond=0)}"' , '')
        
        send_tcp_request('QPUSH','}'," ")
        send_tcp_request('QPUSH',','," ")
    else:
        # You might want to log an error or handle it appropriately
        return render_template('error.html')

    # Move the create_json outside the if-else block
    #json_creator.create_json()
    
    return redirect(original_link)

# @app.route('/getreport')
# def getreport():
#     return send_file('statistic.json')


# @app.route('/report')
# def Call_report():
#     get_report()
#     return "Report sended"


# @app.route('/detail')
# def Call_detail():
#     get_detail()
#     return "Detail sended"


def generate_short_link():
    charset = string.ascii_letters + string.digits
    key_length = random.randint(1, 6)

    short_link = ''.join(random.choice(charset) for _ in range(key_length))
    return short_link


@app.route('/favicon.ico')
def favicon():
    return send_from_directory('static', 'favicon.ico', mimetype='image/vnd.microsoft.icon')



if __name__ == '__main__':
    app.run(host=f'localhost', port=3333, debug=True)
