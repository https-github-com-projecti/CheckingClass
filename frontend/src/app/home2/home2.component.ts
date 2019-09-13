import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DataSource } from '@angular/cdk/collections';
import {HomeService} from '../service/home.service';
import { Config } from 'protractor';
import { error } from 'util';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { timer } from 'rxjs';
import { timeout } from 'q';

export interface userLogin {
  username: string;
  password: string;
}

export interface classOrder {
  t_class_name : string;
  t_class_description : string;
  t_class_id : string;
  user : string;
}

@Component({
  selector: 'app-home2',
  templateUrl: './home2.component.html',
  styleUrls: ['./home2.component.css']
})


export class Home2Component implements OnInit {
  mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  teacherClass: object = null;
  name = localStorage.getItem('isLogin');
  private userData : any;
  

  constructor
  (
    private homeService : HomeService,
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private router : Router,
  ) { 
    this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")
    );
  }

  userlogins: userLogin = {
    username: null,
    password: null,
  };

  select_Class : any = {
    t_class_name : null,
    t_class_description : null,
    t_class_id : null,
    user : null,
  }

  newclass : classOrder = {
    t_class_name: null,
    t_class_description: null,
    t_class_id: null,
    user: null,
  };

  // Testclass : classOrder[] = [{
  //   t_class_name: "Project",
  //   t_class_description: "",
  //   t_class_id: "523495",
  //   user: "Test1",}
  // ];


  ngOnInit() {
    this.loadData();
    this.MaintainStatus()
    var show_ui = document.getElementById("mainId");
    var stat_login = localStorage.getItem('stateLogin');
    var img_btn = document.getElementById("blah");
    var panel_prop = document.getElementById("show_prop");
    var state_btn = true;
    var btn_creat = document.getElementById("btn_creat");
    var create = document.getElementById("create");
    var panel_create = document.getElementById("p_create");

    this.homeService.getUser().subscribe(data => {
      console.log(data);
    });

    this.homeService.getAllClass().subscribe(data =>{
      console.log(data);
    });

    if (stat_login == 'true'){
      show_ui.style.display = "none";
    }

    console.log(this.userData);
    // console.log(name);
    // btn_creat.onclick = function(){
    //   create.style.display = "block";
    // }
    // panel_create.onclick = function(){
    //   create.style.display = "none";
    // }

    // img_btn.onclick = function(){
    //   state_btn = !state_btn;
    //   if(state_btn == false){
    //     panel_prop.style.display = "block"; 
    //   }
    //   else{
    //     panel_prop.style.display = "none";
    //   }      
    // }
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
        this.userData = data;
        this.homeService.login(this.userlogins.username, 'true');
        this.loadData();
        show_ui.style.display = "none";
        console.log(data);
        console.log("check login = ", this.userData);
      },
      error => {
        console.log('Error', error);
        window.location.reload();
    });
  }

  logout(){
    localStorage.clear();
    window.location.reload();
  }

  checkClassOrder(){
    if(this.newclass.t_class_name == null){
      alert('กรุณาป้อนชื่อคลาสเรียนด้วยครับ')
    }else if(this.newclass.t_class_id == null){
      alert('กรุณาระบุคลาสไอดีด้วยครับ')
    }else{
      this.setUser();
      this.classSave();
    }
  }
  setUser(){
    this.newclass.user = localStorage.getItem('isLogin');
    if(this.newclass.user == null){
      console.log('no User');
    }
  }
  classSave(){
    var panel_create = document.getElementById("create");
    this.homeService.CreateClass(this.newclass).subscribe(
      data =>{
      var checkOK;
      checkOK = data;
      console.log(checkOK);
      if(checkOK == "Success"){
        alert('Create success');
        panel_create.style.display = "none";
        this.loadData()
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

  MaintainStatus(){
    timer:0.5;
    var show_ui  = document.getElementById("mainId");
    if(localStorage.getItem('stateLogin') == 'true' && localStorage.getItem('isLogin') == this.userlogins.username){
      show_ui.style.display = "none";
      this.loadData();
    }
  }

  loadData() {
    this.homeService.getClass().subscribe(data =>{
      this.teacherClass = data;
      console.log(typeof(data));
      console.log(this.teacherClass);
    });
  }

  selectClass(x){
    // console.log(x);
    this.select_Class.t_class_name = x.t_class_name;
    this.select_Class.t_class_description = x.t_class_description;
    this.select_Class.t_class_id = x.t_class_id;
    this.select_Class.user = x.user;
    console.log(this.select_Class);
    this.router.navigate(['/Class']);
  }  
}