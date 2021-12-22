const redis = require('redis');
const client = redis.createClient();

client.on("error", function (error) {
    console.error("Error en conexion: ", error)
})

client.on('connect', function() {
  console.log('Connected!');
});