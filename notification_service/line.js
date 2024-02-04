const ACCESS_TOKEN = "gFETqcP4VRi5p/ti4eHBQFL1gnBj3UqWjoyU0u52g3UX20vnmRMjehc3gYyVajPjT8OZ/4sdUxk9kf68bVImo1rMrhj4EwERt8UxWh5+m12c6laBlDkLaNs6y4giHJGn+3VQhf4OFnaTlq48My3VDwdB04t89/1O/w1cDnyilFU='"
const URL = "https://api.line.me/v2/bot/message/push"

async function sendNotification(notificationList) {
    try {
        const successSendingList = []
        for (let notification of notificationList) {
            const payload = {
                to: notification.user_id,
                messages: [
                    {
                        type: 'flex',
                        altText: "This is a Flex Message",
                        contents: {
                            "type": "bubble",
                            "body": {
                                "type": "box",
                                "layout": "vertical",
                                "contents": [
                                    {
                                        "type": "text",
                                        "text": "แจ้งเตือนการใช้ยา",
                                        "weight": "bold",
                                        "size": "xl"
                                    },
                                    {
                                        "type": "box",
                                        "layout": "vertical",
                                        "margin": "lg",
                                        "spacing": "sm",
                                        "contents": [
                                            {
                                                "type": "box",
                                                "layout": "baseline",
                                                "contents": [
                                                    {
                                                        "type": "text",
                                                        "text": "ชื่อยา",
                                                        "color": "#aaaaaa",
                                                        "flex": 1,
                                                        "size": "sm"
                                                    },
                                                    {
                                                        "type": "text",
                                                        "text": `${notification.drug_name}`,
                                                        "color": "#666666",
                                                        "flex": 5,
                                                        "size": "sm"
                                                    }
                                                ]
                                            },
                                            {
                                                "type": "box",
                                                "layout": "baseline",
                                                "spacing": "sm",
                                                "contents": [
                                                    {
                                                        "type": "text",
                                                        "text": "เวลา",
                                                        "color": "#aaaaaa",
                                                        "size": "sm",
                                                        "flex": 1
                                                    },
                                                    {
                                                        "type": "text",
                                                        "text": `${notification.time.toLocaleString()}`,
                                                        "wrap": true,
                                                        "color": "#666666",
                                                        "size": "sm",
                                                        "flex": 5
                                                    }
                                                ]
                                            },
                                            {
                                                "type": "box",
                                                "layout": "baseline",
                                                "spacing": "sm",
                                                "contents": [
                                                    {
                                                        "type": "text",
                                                        "text": "สัตว์เลี้ยง",
                                                        "color": "#aaaaaa",
                                                        "size": "sm",
                                                        "flex": 1
                                                    },
                                                    {
                                                        "type": "text",
                                                        "text": `${notification.pet_id}`,
                                                        "wrap": true,
                                                        "color": "#666666",
                                                        "size": "sm",
                                                        "flex": 5
                                                    }
                                                ]
                                            },
                                            {
                                                "type": "box",
                                                "layout": "baseline",
                                                "contents": [
                                                    {
                                                        "type": "text",
                                                        "text": "วิธีใช้",
                                                        "flex": 1,
                                                        "size": "sm",
                                                        "color": "#aaaaaa"
                                                    },
                                                    {
                                                        "type": "text",
                                                        "flex": 5,
                                                        "size": "sm",
                                                        "color": "#666666",
                                                        "text": `${notification.drug_usage}`
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            "footer": {
                                "type": "box",
                                "layout": "vertical",
                                "spacing": "sm",
                                "contents": [
                                    {
                                        "type": "box",
                                        "layout": "vertical",
                                        "contents": [],
                                        "margin": "sm"
                                    }
                                ],
                                "flex": 0
                            }
                        }
                    }
                ]
            }

            const reqOption = {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${ACCESS_TOKEN}`
                },
                body: JSON.stringify(payload)
            }

            res = await fetch(URL, reqOption)
            data = await res.json()
            console.log(data)
            if (res.status === 400) {
                continue
            }

            successSendingList.push(notification.id)
        }

        return successSendingList
    } catch (err) {
        throw err;
    }
}

module.exports = sendNotification