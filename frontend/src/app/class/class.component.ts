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
  passOfCouse : number;
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
  public length : any = null;
  public ShowdDataQr : any = null ;
  private Username : string = null;
  private QrShow : any;

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
    this.matIconRegistry.addSvgIcon(
      "aspectRatio",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/aspect_ratio24px.svg")
    );
    this.matIconRegistry.addSvgIcon(
      "delete",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/delete24px.svg")
    );
   }

  ngOnInit() {
    var showQr = document.getElementById("showQr-contrainer");
    var showQr2 = document.getElementById("showQr-contrainer2");
    this.loadData();
    this.Username = localStorage.getItem('isLogin');
    this.classService.getmyClass().subscribe(data =>{
      this.myClass = data;
      this.classService.passofClass(this.myClass[0]["t_class_pass"]);
      this.loadData();
    });

    window.onclick = function(event) {
      if (event.target == showQr || event.target == showQr2) {
        // showQr.style.display = "none";
        showQr.style.transition = "all ease-out 600ms";
        showQr.style.transform = "rotateX(60deg)";
        showQr.style.top = "-100%";
      }
      if (event.target == showQr2) {
        // showQr.style.display = "none";
        showQr2.style.transition = "all ease-out 600ms";
        showQr2.style.transform = "rotateX(60deg)";
        showQr2.style.top = "-100%";
      }
    }
  }

  createAuthenicate(){
    const now = Date.now();
    const myFormattedDate : string = this.pipe.transform(now, 'medium');
    const user = localStorage.getItem('isLogin');
    this.newQr.time = myFormattedDate;
    this.newQr.user = user;
    this.newQr.passOfCouse = this.myClass[0]["t_class_pass"];
    console.log(this.newQr);

    this.classService.createQr(this.newQr).subscribe(
        data => {
          this.check = data;
          if (this.check  == "Success"){
            this.loadData();
            this.showQrCode();
          }
        },
        error  => {
          alert('Error กรุณาลองใหม่');
          console.log('Error', error);
        }     
    );
  }

  dataAuthen(x){
    var panel = document.getElementById("panelCard");
    var btn_screen = document.getElementById("btnScreen");
    var btn_delete = document.getElementById("btnDelete");
    this.router.navigate(['/DataAuthen']);
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
        // console.log(data);
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

  showQrCode(){
    this.classService.getShowMyQr().subscribe(
      data => {
        this.ShowdDataQr = data;
      }
    )
    var showQr = document.getElementById("showQr-contrainer");
    showQr.style.display = "flex";
    showQr.style.transition = "all ease-out 600ms";
    showQr.style.top = "0";
    showQr.style.transform = "rotateX(0deg)";
  }

  zoomScreen(x){
    console.log(x);
    this.QrShow = x.Qrcode;
    console.log(this.QrShow);
    var showQr = document.getElementById("showQr-contrainer2");
    showQr.style.display = "flex";
    showQr.style.transition = "all ease-out 600ms";
    showQr.style.top = "0";
    showQr.style.transform = "rotateX(0deg)";
  }
}
