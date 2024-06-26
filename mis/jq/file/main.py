from flask import *
import subprocess

app = Flask(__name__)


@app.route("/")
def get():
    return render_template("index.tmpl")


@app.route("/", methods=["POST"])
def post():
    filter = request.form["filter"]
    print("[i] filter :", filter)
    if len(filter) >= 9:
        return render_template("index.tmpl", error="Filter is too long")
    if ";" in filter or "|" in filter or "&" in filter:
        return render_template("index.tmpl", error="Filter contains invalid character")
    command = "jq '{}' test.json".format(filter)
    ret = subprocess.run(
        command,
        shell=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        encoding="utf-8",
    )
    return render_template("index.tmpl", contents=ret.stdout, error=ret.stderr)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000, debug=True)
