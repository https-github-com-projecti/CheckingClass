import { ClassService } from './../service/class.service';
import { classOrder} from './../home2/home2.component';
import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { DatePipe } from '@angular/common';
import {HomeService} from '../service/home.service';
import { stringify } from '@angular/compiler/src/util';
import { Router } from '@angular/router';

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
  myClass: Object = null ;
  guillaumeNery = 'Hello my code';
  pipe = new DatePipe('en-US');
  private nows = Date.now();
  myFormattedDates = this.pipe.transform(this.nows, 'medium');
  Qrcode : Object = null;
  private userdata = null;
  public getPic = null; 
  public check : any = null;


  newQr : createQrcode = {
    time : null,
    user : null,
    passOfCouse : null,
  }

  select_Class: any = {
    id : null,
    t_class_name: null,
    t_class_description: null,
    t_class_id: null,
    user: null,
  }

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private homeService : HomeService,
    private classService : ClassService,
    private router : Router,
  ) {
    this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")
    );
   }

  ngOnInit() {
    this.loadData();
    this.classService.getmyClass().subscribe(data =>{
      this.myClass = data;
      console.log(data);
    });
  }

  createAuthenicate(){
    const now = Date.now();
    const myFormattedDate : string = this.pipe.transform(now, 'medium');
    const user = localStorage.getItem('isLogin');
    this.newQr.time = myFormattedDate;
    this.newQr.user = user;
    this.newQr.passOfCouse = this.myClass[0]["t_class_pass"];

    this.classService.createQr(this.newQr).subscribe(
        data => {
          this.check = data;
          if (this.check  == "Success"){
            this.loadData();
          }
        },
        error  => {
          alert('Error กรุณาลองใหม่');
          console.log('Error', error);
        }     
    );
  }

  Qr(x){
    console.log(x);
  }

  loadData() {
    this.homeService.getUserdata().subscribe(data =>{
      this.userdata = data;
      this.homeService.setID(this.userdata);
      this.homeService.getGetPic().subscribe(data =>{
        this.getPic = data;
        console.log(this.isEmptyOrSpaces(this.getPic));
        if (this.getPic.trim() === ''){}
        else {  this.mypic = this.getPic }
      });
    });
    this.classService.getmyQr().subscribe(
      data =>{
        console.log(data);
        this.Qrcode = data;
      }
    );
  }
  isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
  }

  logout(){
    localStorage.clear();
    this.router.navigate(['/Home']);
  }
}
