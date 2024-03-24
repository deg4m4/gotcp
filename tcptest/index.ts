async function main() {
    const socket = await Bun.connect({
        hostname: "localhost",
        port: 8000,

        socket: {
            data(socket, data) {
                console.log("Data: ", data.toString())
            },
            open(socket) {
                console.log("Connected")
                setInterval(() => {
                    socket.write("gh")
                }, 1000);
            },
            close(socket) {

            },
            drain(socket) {
                console.log("Drain");
                
            },
            error(socket, error) {
                console.log("Error: ", error)
            },

            // client-specific handlers
            connectError(socket, error) {
                console.log("Connection error: ", error)
            }, // connection failed
            end(socket) { }, // connection closed by server
            timeout(socket) { }, // connection timed out
        },
        tls: true
    });

}

main()
