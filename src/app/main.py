from flask import Flask, render_template, url_for, request

app = Flask(__name__)


@app.route('/', methods=["GET"])
def home_page():
    return render_template('index.html')


@app.route('/peering.html', methods=["GET"])
def peering():
    return render_template('peering.html')


@app.errorhandler(404)
def not_found_error(error):
    return render_template('404.html'), 404


@app.errorhandler(Exception)
def unhandled_exception(e):
    return render_template('500.html'), 500


@app.errorhandler(500)
def internal_error(error):
    return render_template('500.html'), 500
