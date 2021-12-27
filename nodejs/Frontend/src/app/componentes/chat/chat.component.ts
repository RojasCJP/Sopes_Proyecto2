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
  array_datos_alm: any[]
  array_top_areas: any[]
  array_graf_1: number[]
  array_graf_1_aux: string[]
  array_graf_1_colors: any[]
  array_graf_2: number[]
  array_graf_2_aux: string[]
  array_graf_2_colors: any[]

  constructor(private chatService: ChatService) {
    this.array_rangos = [0,0,0,0,0,0]
    this.array_usuarios = ["","","","",""]
    this.array_datos_alm = []
    this.array_top_areas = []
    this.array_graf_1 = []
    this.array_graf_1_aux = []
    this.array_graf_1_colors= []
    this.array_graf_2 = []
    this.array_graf_2_aux = []
    this.array_graf_2_colors= []
  }

  ngOnInit(): void {

    this.chatService.listen('chat:report_range').subscribe((data) => {
      this.rangos_valores = data
      if (data) { 
        if(data.valor==null){data.valor=0}
        if (data.id == 'range0_11') {
          this.array_rangos[0]=data.valor
        } else if (data.id == 'range12_18') {
          this.array_rangos[1]=data.valor        
        } else if (data.id == 'range19_26') {
          this.array_rangos[2]=data.valor        
        } else if (data.id == 'range27_59') {
          this.array_rangos[3]=data.valor        
        } else if (data.id == 'range60_end') {
          this.array_rangos[4]=data.valor        
        } 
      }
      this.graf_barr.update()
      setTimeout(()=>this.chatService.emit("chat:report_range", "0"),1000);
    })

    this.chatService.listen('chat:report_users').subscribe((data) => {
      this.array_usuarios = data
      //console.log("redisdb users:", this.array_usuarios)
      setTimeout(()=>this.chatService.emit("chat:report_users", "0"),1000);
    })

    this.chatService.listen('chat:report_datos_alm').subscribe((data) => {
      this.array_datos_alm = data
      setTimeout(()=>this.chatService.emit("chat:report_datos_alm", "0"),1000);
    })

    this.chatService.listen('chat:report_top_areas').subscribe((data) => {
      this.array_top_areas = data
      setTimeout(()=>this.chatService.emit("chat:report_top_areas", "0"),1000);
    })

    this.chatService.listen('chat:graf_cir1').subscribe((data) => {
      var auxiliar1 = data.datos
      var auxiliar2 = data.total
      for (let i = 0; i < auxiliar1.length; i++) {
        for (let j = 0; j < auxiliar2.length; j++) {
          if (auxiliar1[i]._id.location == auxiliar2[j]._id.location) { 
            auxiliar1[i].datos = (auxiliar1[i].datos/auxiliar2[j].datos)*100  
          }
        }
      }
      for (let index = 0; index < auxiliar1.length; index++) {
        this.array_graf_1.push((auxiliar1[index].datos))
        this.array_graf_1_aux.push(auxiliar1[index]._id.location)
        this.array_graf_1_colors.push("rgb("+Math.round(Math.random() * 255)+", "+Math.round(Math.random()*255)+", "+Math.round(Math.random()*255)+")")
      }
      //console.log(this.array_graf_1_colors)
      this.graf_cir_dosis.update()
      this.array_graf_1_aux=[]
      this.array_graf_1=[]
      setTimeout(()=>this.chatService.emit("chat:graf_cir1", "0"),1000);
    })

    this.chatService.listen('chat:graf_cir2').subscribe((data) => {
      var auxiliar1 = data.datos
      var auxiliar2 = data.total
      for (let i = 0; i < auxiliar1.length; i++) {
        for (let j = 0; j < auxiliar2.length; j++) {
          if (auxiliar1[i]._id.location == auxiliar2[j]._id.location) { 
            auxiliar1[i].datos = (auxiliar1[i].datos/auxiliar2[j].datos)*100  
          }
        }
      }
      for (let index = 0; index < auxiliar1.length; index++) {
        this.array_graf_2.push((auxiliar1[index].datos))
        this.array_graf_2_aux.push(auxiliar1[index]._id.location)
        this.array_graf_2_colors.push("rgb("+Math.round(Math.random() * 255)+", "+Math.round(Math.random()*255)+", "+Math.round(Math.random()*255)+")")
      }
      //console.log(this.array_graf_2_colors)
      this.graf_cir_esq.update()
      this.array_graf_2_aux=[]
      this.array_graf_2=[]
      
      setTimeout(()=>this.chatService.emit("chat:graf_cir2", "0"),1000);
    })

    // Activar los canales hacia el servidor
    this.chatService.emit("chat:report_range", "0");
    this.chatService.emit("chat:report_users", "0");
    this.chatService.emit("chat:report_datos_alm", "0");
    this.chatService.emit("chat:report_top_areas", "0");
    this.chatService.emit("chat:graf_cir1", "0");
    this.chatService.emit("chat:graf_cir2", "0");
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
        labels: this.array_graf_1_aux,
        datasets: [{
          label: 'My first ds',
          data: this.array_graf_1,
          backgroundColor: this.array_graf_1_colors
        }]
      }
    })

    this.graf_cir_esq = new Chart(this.ctx2, {
      type: 'pie',
      data: {
        labels: this.array_graf_2_aux,
        datasets: [{
          label: 'My first ds',
          data: this.array_graf_2,
          backgroundColor:this.array_graf_2_colors
        }]
      }
    }) 

    this.graf_barr = new Chart(this.ctx3, {
      type: 'bar',
      data: {
        labels: ['Niños', 'Adolescentes', 'Jovenes', 'Adultos', 'Vejez'],
        /* range0_11
        range12_18
        range19_26
        range27_59
        range60_end */
        datasets: [{
          label: "",
          data: this.array_rangos,
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(255, 159, 64, 0.2)',
            'rgba(255, 205, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(54, 162, 235, 0.2)'
          ],
          borderColor: [
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)',
            'rgb(255, 205, 86)',
            'rgb(75, 192, 192)',
            'rgb(54, 162, 235)'
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
