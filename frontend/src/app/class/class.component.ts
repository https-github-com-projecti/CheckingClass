import { ClassService } from './../service/class.service';
import { classOrder} from './../home2/home2.component';
import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { DatePipe } from '@angular/common';
import {HomeService} from '../service/home.service';
import { stringify } from '@angular/compiler/src/util';

export interface createQrcode {
  time : string;
  user : string;
  passOfCouse : string;
}

@Component({
  selector: 'app-class',
  templateUrl: './class.component.html',
  styleUrls: ['./class.component.css']
})
export class ClassComponent implements OnInit {
  mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  name_Class: string = "Project"
  guillaumeNery = 'Hello my code';
  pipe = new DatePipe('en-US');
  private nows = Date.now();
  myFormattedDates = this.pipe.transform(this.nows, 'medium');
  Qrcode : any;

  // class : classOrder = {
  //   t_class_name: null,
  //   t_class_description: null,
  //   t_class_id: null,
  //   user: null,
  // }

  newQr : createQrcode = {
    time : null,
    user : null,
    passOfCouse : null,
  }

  select_Class: any = {
    t_class_name : null,
    t_class_description : null,
    t_class_id : null,
    user : null,
  }
  class : classOrder [] = [{
    t_class_name: "Test1",
    t_class_description: "Test1",
    t_class_id: "Test1",
    user: "Test1",}
  ];

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private homeService : HomeService,
    private classService : ClassService,
  ) {
    this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")
    );
   }

  ngOnInit() {
  }

  createAuthenicate(){
    const now = Date.now();
    const myFormattedDate : string = this.pipe.transform(now, 'medium');
    const user = localStorage.getItem('isLogin');
    // console.log(myFormattedDate,user);
    this.newQr.time = myFormattedDate;
    this.newQr.user = "NS";
    this.newQr.passOfCouse = "12345"
    console.log(this.newQr);
    this.classService.createQr(this.newQr).subscribe(
      data => {
        console.log(data);
        this.Qrcode = data;
      }     
    );
  }

  Qr(x){
    console.log(x);
  }
}
