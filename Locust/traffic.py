import json
import time
from locust import HttpUser, task
import json


class testApi(HttpUser):

    contador = 0
    entrada = []

    @task
    def acces_model(self):
        response = self.client.post("/insert",
                                    json=testApi.entrada[testApi.contador])
        testApi.contador += 1
        if testApi.contador == len(testApi.entrada):
            testApi.contador = 0
            exit()
        json_var = response
        print("contador: ", testApi.contador)

    def getJson(self):
        f = open('traffic.json')
        data = json.load(f)
        return data

    def on_start(self):
        testApi.entrada = self.getJson()
        print(testApi.entrada[0])
