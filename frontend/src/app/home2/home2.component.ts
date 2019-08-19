import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
// import { Observable } from 'rxjs/observable';
import { DataSource } from '@angular/cdk/collections';
// import { from } from 'rxjs';
import {HomeService} from '../service/home.service';
import { timer } from 'rxjs';


@Component({
  selector: 'app-home2',
  templateUrl: './home2.component.html',
  styleUrls: ['./home2.component.css']
})


export class Home2Component implements OnInit {
  username: String = "";
  password: String = "";
  private activeTodos: any = [];

  // heroes = ['Windstorm', 'Bombasto', 'Magneta', 'Tornado', 'Narawich Saphimarn'];
  heroes:Array<any>;

  mypic: string = 'https://via.placeholder.com/150/FFFFFF/000000/?text=Check%20name%20';
  constructor(private homeService : HomeService,) { }

  readURL(event:any) {
    if (event.target.files && event.target.files[0]) {
      var reader = new FileReader();
      reader.onload = (event:any) => {
       this.mypic = event.target.result;
      }
      reader.readAsDataURL(event.target.files[0]);
    }
  }
  ngOnInit() {
    this.loadData();

    var show_ui = document.getElementById("show");
    var stat_login = localStorage.getItem('stateLogin');
    var img_btn = document.getElementById("blah");
    var panel_prop = document.getElementById("show_prop");
    var state_btn = true;
    var btn_creat = document.getElementById("btn_creat");
    var create = document.getElementById("create");
    var panel_create = document.getElementById("p_create");
    var name = localStorage.getItem('isLogin');

    btn_creat.onclick = function(){
      create.style.display = "block";
    }
    panel_create.onclick = function(){
      create.style.display = "none";
    }

    img_btn.onclick = function(){
      state_btn = !state_btn;
      if(state_btn == false){
        panel_prop.style.display = "block"; 
      }
      else{
        panel_prop.style.display = "none";
      }      
    }

    if (stat_login == 'true'){
      show_ui.style.display = "none";
    }
  }
  login(){
    var show_ui = document.getElementById("show");
    if(this.username == "admin" && this.password == "admin1234"){
      show_ui.style.display = "none";
      this.homeService.login(this.username, 'true');
      this.loadData();
    }else if(this.username == "NS" && this.password == "NS290821"){
      show_ui.style.display = "none";
      this.homeService.login(this.username, 'true');
      this.loadData();
    }else{
      alert('ชื่อผู้ใช้กับระหัสของคุณไม่ถูกต้อง กรุณากรอกใหม่อีกครั้งครับ')
    }
    console.log(this.username);
  }

  logout(){
    localStorage.clear();
    window.location.reload();
  }

  loadData(){
    timer:0.5;
    if(localStorage.getItem('stateLogin') == 'true'){
      if(localStorage.getItem('isLogin') == "admin"){
        this.heroes = ['Windstorm', 'Bombasto', 'Magneta', 'Tornado', 'Narawich Saphimarn','test','Angular','Golang','MongoDB','Base64'];
      }else if(localStorage.getItem('isLogin') == "NS"){
        this.heroes = ['NS', 'Project'];
      }
    }
  }
}
