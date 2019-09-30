import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import {MatTableDataSource} from '@angular/material/table';
import { AuthenStudentService } from '../service/authen-student.service'
import { Observable } from 'rxjs/observable';
import { DataSource } from '@angular/cdk/collections';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { ClassService } from './../service/class.service';
import {HomeService} from '../service/home.service';
import { Router } from '@angular/router';

export interface PeriodicElement {
  id : number;
  idStudent : string;
  nameStudent : string;
  authenStudent: {     
    date : string;
    stateAuthen : boolean;
  }
}
export class RoomdataSource extends DataSource<any>{
  constructor(private authenStudentService: AuthenStudentService){
    super();
  }
  connect(): Observable<PeriodicElement[]> {
    console.log('Hey.Guy');
    return this.authenStudentService.getAuthenData();
  }
  disconnect(){}
}

@Component({
  selector: 'app-authenicat-student',
  templateUrl: './authenicat-student.component.html',
  styleUrls: ['./authenicat-student.component.css']
})
export class AuthenicatStudentComponent implements OnInit {
  mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  displayedColumns: string[] = ['id','idStudent', 'nameStudent','date'];
  dataSource = new RoomdataSource(this.authenService);
  dataAuthen : any = null;
  tdId : any = null;
  tdIdStudent : any = null;
  tdauthenStudent : any = null;
  myClass: Object = null ;
  public getPic = null; 
  private userdata = null;
  private Username : string = null;

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private authenService : AuthenStudentService,
    private _formBuilder: FormBuilder,
    private classService : ClassService,
    private homeService : HomeService,
    private router : Router,
  ) {   }
   
   ngOnInit() {
    this.loadData();
    this.Username = localStorage.getItem('isLogin');
    this.classService.getmyClass().subscribe(data =>{
      this.myClass = data;
      this.classService.passofClass(this.myClass[0]["t_class_pass"]);
      this.loadData();
    });
     this.authenService.getAuthenData().subscribe(
       data => {
        this.dataAuthen = data;
        this.tdId = this.dataAuthen[0].id;
        this.tdIdStudent = this.dataAuthen[0].idStudent;
        this.tdauthenStudent = this.dataAuthen[0].authenStudent;
         console.log(this.dataAuthen);
         console.log(this.tdId);
         console.log(this.tdIdStudent);
         console.log(this.tdauthenStudent);
       }
     )

     console.log(this.dataSource);
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
