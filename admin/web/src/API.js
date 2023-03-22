const API = {

    build: async () => {
        try {
            const response = await fetch('http://localhost:9999/versions?service=items-api', {
                method: 'GET',
                headers: {
                    'Content-type': 'application/json; charset=UTF-8',
                },
            });
            const data = await response.json();
            console.log(data);
            return data;
        } catch (err) {
            console.log(err.message);
            return err;
        }
    },

    start: () => {
        console.log("Start...")
    },

    stop: () => {
        console.log("Stop...")
    },
};

module.exports = {API};