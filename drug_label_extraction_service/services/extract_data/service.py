import re
from services.user.service import get_user_time_settings

def clean_text(text):
    res = ''
    for line in text.splitlines():
        if line == '':
            continue
        res = res + line + ' '
    return res.strip()


def check_period(text):
    keywords = ['เช้า', 'กลางวัน', 'เย็น', 'ก่อนนอน']
    res = {'เช้า': False, 'กลางวัน': False, 'เย็น': False, 'ก่อนนอน': False}
    period_type = False

    for keyword in keywords:
        found = re.search(keyword, text)
        if found:
            period_type = True
            res[keyword] = True

    if period_type:
        return res
    return None


def check_hour(text):
    pattern = r'.+ ทุกๆ (\d+) (.+)'
    match = re.search(pattern, text)

    if match is None:
        return None

    num = match.group(1)
    unit = match.group(2)
    return {'num': num, 'unit': unit}


def check_frequency(text):
    keywords = ['วันเว้นวัน']
    dict = {'วันเว้นวัน': 'every other day'}

    for keyword in keywords:
        found = re.search(keyword, text)
        if found:
            return dict[keyword]
    return None


def extract_data(drug_name_text, usage_text, user_id):
    drug_name = clean_text(drug_name_text)
    usage = clean_text(usage_text)
    res = {'drug_name': drug_name, 'drug_usage': usage, 'reminder_type': '', 'frequency': 'every day'}

    freq = check_frequency(usage)
    if freq:
        res['frequency'] = freq

    ok = check_period(usage)
    if ok:
        user_time_settings = get_user_time_settings(user_id)
        res['reminder_type'] = 'period'
        if ok['เช้า']:
            res['morning'] = user_time_settings['morning']
        else:
            res['morning'] = None
        if ok['กลางวัน']:
            res['noon'] = user_time_settings['noon']
        else:
            res['noon'] = None
        if ok['เย็น']:
            res['evening'] = user_time_settings['evening']
        else:
            res['evening'] = None
        if ok['ก่อนนอน']:
            res['before_bed'] = user_time_settings['before_bed']
        else:
            res['before_bed'] = None
        return res

    ok = check_hour(usage)
    if ok:
        res['reminder_type'] = 'hour'
        res['every'] = ok['num']
        return res
    return None
