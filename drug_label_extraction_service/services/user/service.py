import os

import requests


def get_user_time_settings(user_id):
    host = os.environ.get("USER_SERVICE_HOST", "localhost")
    url = "http://" + host + ":8080/user/" + user_id
    resp = requests.get(url)
    return resp.json()

