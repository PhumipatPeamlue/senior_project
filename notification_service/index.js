const cron = require("node-cron");

const {connectDB, getNotSentNotificationList, changeNotificationStatus} = require("./database");
const sendNotification = require("./line");

async function main() {
    let connection;

    try {
        connection = await connectDB()
        const list = await getNotSentNotificationList(connection)
        const successList = await sendNotification(list)
        await changeNotificationStatus(connection, successList)
    } catch (err) {
        console.log("error: ", err);
    } finally {
        await connection.end()
    }
}

cron.schedule('* * * * *', () => {
    console.log("run cronjob")
    main()
})

