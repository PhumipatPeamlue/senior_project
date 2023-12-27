import re


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

    for keyword in keywords:
        found = re.search(keyword, text)
        if found:
            return keyword
    return None


def extract_data(drug_name_text, usage_text):
    drug_name = clean_text(drug_name_text)
    usage = clean_text(usage_text)
    res = {'drug_name': drug_name, 'usage': usage, 'notification_type': '', 'notify': [], 'frequency': 'ทุกๆวัน'}

    freq = check_frequency(usage)
    if freq:
        res['frequency'] = freq

    ok = check_period(usage)
    if ok:
        res['notification_type'] = 'period'
        res['notify'] = ok
        return res

    ok = check_hour(usage)
    if ok:
        res['notification_type'] = 'hour'
        res['notify'] = ok
        return res
    return None
