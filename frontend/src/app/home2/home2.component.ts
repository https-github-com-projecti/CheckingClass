import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataSource } from '@angular/cdk/collections';
import {HomeService} from '../service/home.service';
import { error } from 'util';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import {ClassService } from '../service/class.service';

export interface userLogin {
  username: string;
  password: string;
}

export interface classOrder {
  subject_id : number;
  TSName : string;
  TSDescription : string;
  TSID : string;
  TSTeacher : string;
}

@Component({
  selector: 'app-home2',
  templateUrl: './home2.component.html',
  styleUrls: ['./home2.component.css']
})

export class Home2Component implements OnInit {
  mypic: string = "data:image/gif;base64,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";
  teacherClass: object = null;
  private userdata = null;
  loginas : string = null;
  loginReply : any;
  private reader = new FileReader();
  public getPic = null; 
  hide = false;
  joinName : string = null;
  private Username : string = null;

  constructor(
    private homeService : HomeService,private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,private router : Router,private classservice : ClassService,)
    { 
    this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")
    );
    this.matIconRegistry.addSvgIcon(
      "delete",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/delete24px.svg")
    );
  }

  userlogins: userLogin = {
    username: null,
    password: null,
  };

  newclass : classOrder = {
    subject_id : null,
    TSName: null,
    TSDescription: null,
    TSID: null,
    TSTeacher: null,
  };

  select_class : classOrder = {
    subject_id : null,
    TSName: null,
    TSDescription: null,
    TSID: null,
    TSTeacher: null,
  };

  ngOnInit() {
    this.homeService.getStatus().subscribe(
      data => {
        var statusBackend = data;
        if (statusBackend == "online"){
          this.homeService.status(statusBackend);
        }
      } ,
      error => {
        this.logout();
      }
    );
    this.loadData();
    var show_ui = document.getElementById("mainId");
    var stat_login = localStorage.getItem('stateLogin');
    var state_btn = true;

    var btn_creat = document.getElementById("bt_createClass");
    var create = document.getElementById("create");
    var panel_create = document.getElementById("CreateClass");
    var img;

    var btnJoinClass = document.getElementById("btJoinClass");
    var panelJoinClass = document.getElementById("joinContainer");
    this.Username = localStorage.getItem('isLogin');

    if (stat_login == 'true'){
      show_ui.style.display = "none";
    }
    btn_creat.onclick = function(){
      panel_create.style.display = "flex";
    }

    btnJoinClass.onclick = function(){
      panelJoinClass.style.display = "flex";
    }

    window.onclick = function(event) {
      if (event.target == panel_create) {
        panel_create.style.display = "none";
      }
      if (event.target == panelJoinClass) {
        panelJoinClass.style.display = "none";
      }
    }
  }

  checkLogin(){
    if(this.userlogins.username == null){
      alert("กรุณาป้อน User name");
    }else if(this.userlogins.password == null){
      alert("กรุณาป้อน Password");
    }else{
      this.login();
    }
  }
  
  login(){
    var show_ui = document.getElementById("mainId");
    this.homeService.LoginUser(this.userlogins).subscribe(
      data => {
        this.loginReply = data;
        console.log("state Login = " + this.loginReply);
        if(this.loginReply == "Success"){
          // console.log(this.loginReply);
          this.homeService.login(this.userlogins.username, 'true');
          show_ui.style.display = "none";
          this.loadData();
        }
        else {
          alert("Username และ Password ของคุณไม่ถูกต้องกรุณาล็อคอินใหม่")
        }
      },
      error => {
        console.log('Error', error);
        alert("server ERROR")
        // window.location.reload();
    }
    );
  }

  logout(){
    localStorage.clear();
    window.location.reload();

  }

  checkClassOrder(){
    // var inputName = document.getElementById('clear_class1');
    // var inputDescription = document.getElementById('clear_class2');
    // var inputSubjectID = document.getElementById('clear_class3');
    // var btnCreateClass = document.getElementById('bt_saveClass');
    if(this.newclass.TSName == null){
      alert('กรุณาป้อนชื่อคลาสเรียนด้วยครับ')
    }else if(this.newclass.TSID == null){
      alert('กรุณาระบุคลาสไอดีด้วยครับ')
    }else{
      (<HTMLInputElement>document.getElementById('clear_class1')).disabled = true;
      (<HTMLInputElement>document.getElementById('clear_class2')).disabled = true;
      (<HTMLInputElement>document.getElementById('clear_class3')).disabled = true;
      (<HTMLInputElement>document.getElementById('bt_saveClass')).disabled = true;
      this.setUser();
      this.classSave();
    }
  }
  setUser(){
    this.newclass.TSTeacher = localStorage.getItem('id');
    if(this.newclass.TSTeacher == null){
      console.log('no User');
    }
  }
  classSave(){
    var panel_create = document.getElementById("CreateClass");
    this.homeService.CreateClass(this.newclass).subscribe(
      data =>{
      var checkOK;
      checkOK = data;
      console.log(checkOK);
      if(checkOK == "Success"){
        panel_create.style.display = "none";
        (<HTMLInputElement>document.getElementById('clear_class1')).value = ' ';
        (<HTMLInputElement>document.getElementById('clear_class2')).value = ' ';
        (<HTMLInputElement>document.getElementById('clear_class3')).value = ' ';
        (<HTMLInputElement>document.getElementById('clear_class1')).disabled = false;
        (<HTMLInputElement>document.getElementById('clear_class2')).disabled = false;
        (<HTMLInputElement>document.getElementById('clear_class3')).disabled = false;
        (<HTMLInputElement>document.getElementById('bt_saveClass')).disabled = false;
        this.loadData();
      }
      else{
        console.log('Error', error);
      }
    },
    error => {
      console.log('Error', error);
      alert('Error กรุณาป้อนข้อมูลใหม่อีกครั้งครับ');
    }
    );
  }

  loadData() {
    this.homeService.getUserdata().subscribe(data =>{  //get ID
      this.userdata = data; //มาแค่ID
      // console.log(this.userdata);
      this.homeService.setID(this.userdata); //สำหรับserver Test
      // this.homeService.setID(this.userdata[0]['user_id']); //สำหรับserver DB
      this.homeService.getGetPic().subscribe(data =>{ //get Picture
        this.getPic = data;
        console.log(this.getPic);
        if (this.getPic.trim() === ''){}
        else {  this.mypic = this.getPic }
      });
      this.homeService.getClass().subscribe(data =>{ //get My class
        this.teacherClass = data;
        console.log(this.teacherClass);
      });
    });
    this.Username = localStorage.getItem('isLogin');
  }

  selectClass(x){
    this.classservice.classSelect(x.subject_id);
    this.router.navigate(['/Class']);
  }  

  isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
  }

  join(){
    var panel = document.getElementById("joinContainer");
    var x = this.joinName;
    panel.style.display = "none";
    (<HTMLInputElement>document.getElementById('inputJoinClass')).value = ' ';
    console.log(x);
  }
}

