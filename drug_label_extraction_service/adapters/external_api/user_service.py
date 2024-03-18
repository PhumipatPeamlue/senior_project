import os
import requests


def get_user_time_settings(user_id):
    host = os.environ.get("USER_SERVICE_HOST", "localhost")
    port = os.environ.get("USER_SERVICE_PORT", "8080")
    url = "http://" + host + ":" + port + "/user/" + user_id
    resp = requests.get(url)
    return resp.json()
