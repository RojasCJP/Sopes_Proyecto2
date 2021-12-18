import { formatDate } from '@angular/common';
import { Component, OnInit, ViewChild } from '@angular/core';
import { MensajeInterface } from 'src/app/estructuras/mensaje_interface';
import { ChatService } from '../../servicios/chat/chat.service'
import { Chart } from 'chart.js'

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit {
  
  // Variables del input
  mensajes: MensajeInterface[] = []
  input_msj: string = ""

  // Variables con los valores de la grafica
  /* private segundos: number[]
  public cantidad: number[] */
  
  // Grafica circular , una dosis
  canvas: any;
  ctx: any;
  @ViewChild('graf1') graf1: any;
  graf_cir_dosis: any;

  // Grafica circular, esquema completo
  canvas2: any;
  ctx2: any;
  @ViewChild('graf2') graf2: any;
  graf_cir_esq: any;

  // Grafica de barras
  canvas3: any;
  ctx3: any;
  @ViewChild('graf3') graf3: any;
  graf_barr: any;

  constructor(private chatService: ChatService) {
    /* this.segundos = [1,2,3,4,5,6,7]
    this.cantidad = [10,20,30,40,50,60,70] */
  }

  ngOnInit(): void {
    this.chatService.listen('chat:message').subscribe((data) => {
      this.recibirMensaje(data);
    })
  }

  ngAfterViewInit() {
    this.canvas = this.graf1.nativeElement;
    this.canvas2 = this.graf2.nativeElement;
    this.canvas3 = this.graf3.nativeElement;

    this.ctx = this.canvas.getContext('2d');
    this.ctx2 = this.canvas2.getContext('2d');
    this.ctx3 = this.canvas3.getContext('2d');

    this.graf_cir_dosis = new Chart(this.ctx, {
      type: 'pie',
      data: {
        labels: [
          'Red',
          'Blue',
          'Yellow'
        ],
        datasets: [{
          label: 'My first ds',
          data: [300, 50, 100],
          backgroundColor: [
            'rgb(255, 99, 132)',
            'rgb(54, 162, 235)',
            'rgb(255, 205, 86)'
          ]
        }]
      }
    })

    this.graf_cir_esq = new Chart(this.ctx2, {
      type: 'pie',
      data: {
        labels: [
          'Morado',
          'Celeste',
          'Verde'
        ],
        datasets: [{
          label: 'My first ds',
          data: [300, 50, 100],
          backgroundColor: [
            'rgb(153, 51, 255)',
            'rgb(102, 178, 255)',
            'rgb(100, 255, 178)'
          ]
        }]
      }
    }) 

    this.graf_barr = new Chart(this.ctx3, {
      type: 'bar',
      data: {
        labels: ['1-15','16-30','31-45','46-60','61-75','76-90'],
        datasets: [{
          label: "",
          data: [65, 60, 80, 81, 60, 90, 0],
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(255, 159, 64, 0.2)',
            'rgba(255, 205, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(153, 102, 255, 0.2)'
          ],
          borderColor: [
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)',
            'rgb(255, 205, 86)',
            'rgb(75, 192, 192)',
            'rgb(54, 162, 235)',
            'rgb(153, 102, 255)'
          ],
          borderWidth: 1
        }]
      },
      options: {
        scales: {
            xAxes: [{
                gridLines: {
                    offsetGridLines: true
                }
            }]
        }
    }

    })

  }

  enviarMensaje() {
    let mensajeNuevo: MensajeInterface = {
      texto: this.input_msj,
      usuario: "Usuario quemado",
      idUsuario: 0,
      fecha: formatDate(new Date(), 'yyyy/MM/dd H:mm:ss', 'en-US', '+0600').toString()
    }
    
    this.chatService.emit("chat:message", mensajeNuevo);
    this.input_msj = ""
  }

  recibirMensaje(data: MensajeInterface) {
    this.mensajes.push(data)
  }

}
