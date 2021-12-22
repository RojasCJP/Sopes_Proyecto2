const express = require('express')
const cors = require('cors')
const app = express()

app.set('port', 10000)
app.use(cors([
    {
        origin: "*",
        credentials: true
    }
]))

const server = app.listen(app.get('port'), '0.0.0.0', () => {
    console.log('server on port', app.get('port'))
})

const ws = require('socket.io')(server, {
    cors: {
        origin: "*",
        "methods": "GET,HEAD,PUT,PATCH,POST,DELETE",
        "preflightContinue": false,
        "optionsSuccessStatus": 204
    }
})

// ************* MongoDB ************* 
const mongoose = require('mongoose')
const url = 'mongodb://localhost/SopesProyecto2'

mongoose.connect(url, {})
    .then(() => console.log('conectado a mongo'))
    .catch((e) => console.log("error de conexion: ", e))

const personaSchema = mongoose.Schema({
    name: String,
    age: Number,
    id: Number
})

const PersonaModel = mongoose.model('datos', personaSchema)

mostrar = async () => {
    const personas = await PersonaModel.find()
    console.log(personas)
}

// mostrar()
// ************* MongoDB *************

// ************* Redis *************
const redis = require('redis');
const client = redis.createClient();

client.on("error", function (error) {
    console.error("Error en conexion: ", error)
})

client.on('connect', function() {
  console.log('Connected!');
});

client.get('range11_20', (err, reply) => {
    if (err) console.log(err);
    console.log(reply);
});

client.lrange('users', 0, 4, (err, reply) => {
    if (err) console.log(err);
    console.log(reply)
})
// ************* Redis *************

ws.on('connection', function (socket) {
    console.log('Nueva conexion: ', socket.id)

    // reportes

    socket.on('chat:message', (data) => {
        ws.sockets.emit('chat:message', data)
    })
    // ...
})