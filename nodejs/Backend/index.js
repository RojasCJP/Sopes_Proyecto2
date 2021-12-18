const express = require('express');
const cors = require('cors');
const app = express();

app.set('port', 3000);
app.use(cors([
    {
        origin: "*",
        credentials: true
    }
]));

const server = app.listen(app.get('port'), '0.0.0.0', () => {
    console.log('server on port', app.get('port'));
})

const ws = require('socket.io')(server, {
    cors: {
        origin: "*",
        "methods": "GET,HEAD,PUT,PATCH,POST,DELETE",
        "preflightContinue": false,
        "optionsSuccessStatus": 204
    }
});

ws.on('connection', function (socket)  {
    console.log('Nueva conexion: ', socket.id);
    socket.on('chat:message', (data) => {
        ws.sockets.emit('chat:message', data);
    })
    // ...
});