const Config = {

    icons: {
        golang: "https://cdn.iconscout.com/icon/free/png-256/go-77-1175166.png",
        nodejs: "https://cdn.iconscout.com/icon/free/png-256/node-js-1174925.png"
    },

    services: [
        {
            type: "backend",
            id: "items-api",
            display_name: "Items API",
            technology: "golang",
            location: "/backend/items-api",
            repository: "",
            docs: ""
        },
        {
            type: "backend",
            id: "users-api",
            display_name: "Users API",
            technology: "golang",
            location: "/backend/users-api",
            repository: "",
            docs: ""
        },
        {
            type: "frontend",
            id: "client",
            display_name: "Frontend Client",
            technology: "nodejs",
            location: "/frontend/client",
            repository: "",
            docs: ""
        },
        {
            type: "frontend",
            id: "server",
            display_name: "Frontend Server",
            technology: "nodejs",
            location: "/frontend/server",
            repository: "",
            docs: ""
        }
    ]
};

module.exports = {Config};