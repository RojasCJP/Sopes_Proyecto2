db.prueba.save({ usuario: "Sopes" })
db.prueba.find({ usuario: "Sopes" })
db.dropDatabase()

db.datos.aggregate([
    { $group: { _id: "$location", total: { $sum: 1 } } },
    { $sort: { total: -1 } },
    { $limit: 3 }
])


