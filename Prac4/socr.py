from flask import Flask, render_template, request, redirect,send_file,jsonify
from GetReport import get_report 
import datetime
import json
import random
import string
import socket
from GetDetail import get_detail


class HashTable:
    def __init__(self, size):
        self.size = size
        self.table = [None] * size

    def hash_function(self, key):
        return hash(key) % self.size

    def insert(self, key, value):
        index = self.hash_function(key)
        if self.table[index] is None:
            self.table[index] = [(key, value)]
        else:
            for i, (existing_key, _) in enumerate(self.table[index]):
                if existing_key == key:
                    # Если ключ уже существует, обновляем значение
                    self.table[index][i] = (key, value)
                    break
            else:
                # Если ключ не найден, добавляем новую пару ключ-значение
                self.table[index].append((key, value))

    def search(self, key):
        index = self.hash_function(key)
        if self.table[index] is not None:
            for existing_key, value in self.table[index]:
                if existing_key == key:
                    return value
        # Если ключ не найден
        return None

    def delete(self, key):
        index = self.hash_function(key)
        if self.table[index] is not None:
            for i, (existing_key, _) in enumerate(self.table[index]):
                if existing_key == key:
                    del self.table[index][i]
                    break


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



def getHostIP():
    return socket.gethostbyname(socket.gethostname())


app = Flask(__name__)
server_address = ('{getHostIP()}', 3333)
short_link_table = HashTable(size=512)  # Adjust the size based on your requirements

@app.route('/', methods=['GET', 'POST'])
def home():
    generated_link = None
    if request.method == 'POST':
        original_link = request.form['user_input']

        short_link = generate_short_link()

        # Save the short link in the hash table
        short_link_table.insert(short_link, original_link)

        generated_link = f"http://{getHostIP()}:3333/{short_link}"

    return render_template('index.html', output_link=generated_link)

@app.route('/<short_link>')
def redirect_to_original(short_link):
    # Retrieve the original link from the hash table
    original_link = short_link_table.search(short_link)
    
    if original_link:
        json_creator = JSONCreator('statistic.json')
        json_creator.add_data(original_link + "(" + short_link + ")", request.environ['REMOTE_ADDR'], datetime.datetime.now().replace(microsecond=0))
        # Move the create_json outside the if-else block
    else:
        # You might want to log an error or handle it appropriately
        return render_template('error.html')

    # Move the create_json outside the if-else block
    json_creator.create_json()
    
    return redirect(original_link)

@app.route('/getreport')
def getreport():
    return send_file('statistic.json')


@app.route('/report')
def Call_report():
    get_report()
    return "Report sended"


@app.route('/detail')
def Call_detail():
    get_detail()
    return "Detail sended"


def generate_short_link():
    charset = string.ascii_letters + string.digits
    key_length = random.randint(1, 6)

    short_link = ''.join(random.choice(charset) for _ in range(key_length))
    return short_link



if __name__ == '__main__':
    app.run(host=f'{getHostIP()}', port=3333, debug=False)


