db.prueba.save({ usuario: "Sopes" })
db.prueba.find({ usuario: "Sopes" })
db.dropDatabase()

db.datos.aggregate([
    { $group: { _id: "$location", total: { $sum: 1 } } },
    { $sort: { total: -1 } },
    { $limit: 3 }
])


db.datos.count()
db.datos.aggregate([
    { $match: { "n_dose": { $eq: 1 } } },
    { $group: { _id: { location: "$location", dose: "$n_dose" }, datos: { $sum: 1 } } }
])
db.datos.aggregate([
    { $group: { _id: { location: "$location" }, datos: { $sum: 1 } } }
])