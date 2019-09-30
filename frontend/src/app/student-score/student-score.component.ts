import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { ClassService } from './../service/class.service';
import {HomeService} from '../service/home.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-student-score',
  templateUrl: './student-score.component.html',
  styleUrls: ['./student-score.component.css']
})
export class StudentScoreComponent implements OnInit {
  mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  panelOpenState = false;
  step = 0;
  myClass: Object = null ;
  public getPic = null; 
  private userdata = null;
  private Username : string = null;

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private classService : ClassService,
    private homeService : HomeService,
    private router : Router,
  ) {}


  ngOnInit() {
    this.loadData();
    this.Username = localStorage.getItem('isLogin');
    this.classService.getmyClass().subscribe(data =>{
      this.myClass = data;
      this.classService.passofClass(this.myClass[0]["t_class_pass"]);
      this.loadData();
    });
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
  }
  isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
  }

  logout(){
    localStorage.clear();
    this.router.navigate(['/Home']);
  }

}
