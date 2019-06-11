import os
from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello():
    hostname = os.environ['HOSTNAME']
    return "I'm Python MS - POD: {0}".format(hostname)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=3000)