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
const redis = require('redis')
const client = redis.createClient()


client.on("error", function (error) {
    console.error("Error en conexion: ", error)
})

client.on('connect', function () {
    console.log('conectado a redis')
})

async function get_range(range) {
    var respuesta = ""
    client.get(range, (err, reply) => {
        if (err) { console.log(err) }
        respuesta = reply
        ws.sockets.emit('chat:report_range', { valor: respuesta, id: range })
    })
}

async function get_usuarios() {    
    client.lrange('users', 0, 4, (err, reply) => {
        if (err) console.log(err)
        ws.sockets.emit('chat:report_users', reply)
    })
}
// ************* Redis *************

var resultado = []
ws.on('connection', async function (socket) {
    console.log('Nueva conexion: ', socket.id)

    socket.on('chat:report_range', async (data) => {
        get_range("range0_10")
        get_range("range11_20")
        get_range("range21_30")
        get_range("range31_40")
        get_range("range41_50")
        get_range("range51_60")
        get_range("range61_70")
        get_range("range71_80")
        get_range("range81_end")
    })

    socket.on('chat:report_users', (data) => {
        get_usuarios()
    })

})