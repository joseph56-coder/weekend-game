import { Manager } from "https://cdn.socket.io/4.4.1/socket.io.esm.min.js";

const manager = new Manager("https://f840-2804-431-c7e5-fd83-dcc-90d0-780b-3795.sa.ngrok.io/socket.io");
const socket = manager.socket("/"); // main namespace
const adminSocket = manager.socket("/user"); // admin namespace
manager.open((err) => {
    if (err) {
        console.log("err")
    } else {
        console.log("no err")
    }
});