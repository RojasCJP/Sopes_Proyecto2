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

// ************* conexion a MongoDB
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
// ************* conexion a MongoDB

// ************* conexion a Redis
const redis = require('redis')
const client = redis.createClient({
    host: '0.0.0.0',
    port: '6379',
})

client.set('foo', 'bar', function (err, reply) {
    console.log(reply)
})

client.get('foo', function (err, reply) {
    console.log(reply)
})
/* 
client.set('foo', 'bar', (err, reply) => {
    if (err) throw err
    console.log("antes de repuesta")
    console.log(reply)

    client.get('foo', (err, reply) => {
        if (err) throw err
        console.log(reply)
    })
}) */

// ************* conexion a Redis


ws.on('connection', function (socket) {
    console.log('Nueva conexion: ', socket.id)

    // reportes

    socket.on('chat:message', (data) => {
        ws.sockets.emit('chat:message', data)
    })
    // ...
})