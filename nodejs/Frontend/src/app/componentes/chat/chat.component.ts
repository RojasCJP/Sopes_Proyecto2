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

  rangos_valores: any
  array_rangos: number[]
  array_usuarios: string[]

  constructor(private chatService: ChatService) {
    this.array_rangos = [0,0,0,0,0,0,0,0,0]
    this.array_usuarios = ["","","","",""]
  }

  ngOnInit(): void {

    this.chatService.listen('chat:report_range').subscribe((data) => {
      this.rangos_valores = data
      if (data) { 
        if(data.valor==null){data.valor=0}
        if (data.id == 'range0_10') {
          this.array_rangos[0]=data.valor
        } else if (data.id == 'range11_20') {
          this.array_rangos[1]=data.valor        
        } else if (data.id == 'range21_30') {
          this.array_rangos[2]=data.valor        
        } else if (data.id == 'range31_40') {
          this.array_rangos[3]=data.valor        
        } else if (data.id == 'range41_50') {
          this.array_rangos[4]=data.valor        
        } else if (data.id == 'range51_60') {
          this.array_rangos[5]=data.valor        
        } else if (data.id == 'range61_70') {
          this.array_rangos[6]=data.valor        
        } else if (data.id == 'range71_80') {
          this.array_rangos[7]=data.valor        
        } else if (data.id == 'range81_end') {
          this.array_rangos[8]=data.valor        
        }
        console.log(this.array_rangos)
      }
      this.graf_barr.update()
      setTimeout(()=>this.chatService.emit("chat:report_range", "0"),1000);
    })

    this.chatService.listen('chat:report_users').subscribe((data) => {
      this.array_usuarios = data
      console.log("redisdb users:", this.array_usuarios)
      setTimeout(()=>this.chatService.emit("chat:report_users", "0"),1000);
    })

    // Activar los canales hacia el servidor
    this.chatService.emit("chat:report_range", "0");
    this.chatService.emit("chat:report_users", "0");

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
        labels: ['0-10','11-20','21-30','31-40','41-50','51-60', '61-70', '71-80', '81-end'],
        datasets: [{
          label: "",
          data: this.array_rangos,
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

}
