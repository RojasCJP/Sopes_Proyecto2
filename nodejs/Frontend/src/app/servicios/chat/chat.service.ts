import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { io, Socket } from 'socket.io-client';
import { MensajeInterface } from 'src/app/estructuras/mensaje_interface';


@Injectable({
  providedIn: 'root'
})
export class ChatService {

  private socket: Socket;
  private url = 'http://backend.so1p2g1.tk'; 

  constructor() {
    this.socket = io(this.url)      // Realizar la conexion con el servidor
  }

  listen(eventName: string): Observable<any> {
    return new Observable((subscriber) => {
      this.socket.on(eventName, (data) => {
        subscriber.next(data)
      })
    })
  }

  emit(eventName: string, data: any): void {
    this.socket.emit(eventName, data);
  }

}

