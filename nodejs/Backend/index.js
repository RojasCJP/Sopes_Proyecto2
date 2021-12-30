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

var router = express.Router()
app.use(router)

router.get("/", function (req, res) {
    res.send("hola buenas desde node")
})
const server = app.listen(app.get('port'), () => {
    var host = server.address().address
    var port = server.address().port
    console.log('server on port', app.get('port'))
    console.log('host ', host, "; port", port)
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
const url = 'mongodb://34.135.96.5:27017/SopesProyecto2'

mongoose.connect(url, {})
    .then(() => console.log('conectado a mongo'))
    .catch((e) => console.log("error de conexion: ", e))

const personaSchema = mongoose.Schema({
    name: String,
    age: Number,
    id: Number
})

const PersonaModel = mongoose.model('datos', personaSchema)

async function get_datos_alm() {
    const personas = await PersonaModel.find()
    ws.sockets.emit('chat:report_datos_alm', personas)
}

async function get_top_areas() {
    const areas = await PersonaModel.aggregate([
        { $group: { _id: "$location", total: { $sum: 1 } } },
        { $sort: { total: -1 } },
        { $limit: 3 }
    ])
    ws.sockets.emit('chat:report_top_areas', areas)
}

async function graf_cir1() {
    const datos = await PersonaModel.aggregate([
        { $match: { "n_dose": { $eq: 1 } } },
        { $group: { _id: { location: "$location", dose: "$n_dose" }, datos: { $sum: 1 } } }
    ])
    const total = await PersonaModel.aggregate([
        { $group: { _id: { location: "$location" }, datos: { $sum: 1 } } }
    ])

    console.log(datos)
    console.log(total)

    ws.sockets.emit('chat:graf_cir1', { datos: datos, total: total })
}

async function graf_cir2() {
    const datos = await PersonaModel.aggregate([
        { $match: { "n_dose": { $eq: 2 } } },
        { $group: { _id: { location: "$location", dose: "$n_dose" }, datos: { $sum: 1 } } }
    ])
    const total = await PersonaModel.aggregate([
        { $group: { _id: { location: "$location" }, datos: { $sum: 1 } } }
    ])

    console.log(datos)
    console.log(total)

    ws.sockets.emit('chat:graf_cir2', { datos: datos, total: total })
}




// mostrar()
// ************* MongoDB *************

// ************* Redis *************
const redis = require('redis')
const bodyParser = require('body-parser')
const client = redis.createClient(
    {
        url: 'redis://juan:rojas@34.135.96.5:6379'
    }
)


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
        get_range("range0_11")
        get_range("range12_18")
        get_range("range19_26")
        get_range("range27_59")
        get_range("range60_end")
    })

    socket.on('chat:report_users', (data) => {
        get_usuarios()
    })

    socket.on('chat:report_datos_alm', (data) => {
        get_datos_alm()
    })

    socket.on('chat:report_top_areas', (data) => {
        get_top_areas()
    })

    socket.on('chat:graf_cir1', (data) => {
        graf_cir1()
    })
    socket.on('chat:graf_cir2', (data) => {
        graf_cir2()
    })

})