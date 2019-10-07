import { Router, Route } from '@angular/router';
import { Subscription, throwError } from 'rxjs';
import {SingupService} from './../service/singup.service';
import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroupDirective, NgForm, Validators} from '@angular/forms';
import {ErrorStateMatcher} from '@angular/material/core';
import { DataSource } from '@angular/cdk/collections';
import { DomSanitizer } from '@angular/platform-browser';
import { HttpClient } from '@angular/common/http';
import { error } from 'util';

export interface User {
  tFirstName: string;
  tLastName: string;
  userName: string;
  tId: string;
  tEmail: string;
  tWorkPlace: string;
  tPassword: string;
  tPicture: string;
}

export class MyErrorStateMatcher implements ErrorStateMatcher {
  isErrorState(control: FormControl | null, form: FormGroupDirective | NgForm | null): boolean {
    const isSubmitted = form && form.submitted;
    return !!(control && control.invalid && (control.dirty || control.touched || isSubmitted));
  }
}


@Component({
  selector: 'app-signup-component',
  templateUrl: './signup-component.component.html',
  styleUrls: ['./signup-component.component.css']
})

export class SignupComponentComponent implements OnInit {
  tComfirmPassword = null;
  private imagePath;
  imgURL: any = "../../assets/fbb2978e127f2920ab9774076ade2a36.png";
  private message: string;
  private reader = new FileReader();
  private img: any = null;
  private base64 :string = null;
  matcher = new MyErrorStateMatcher();
  hide = false;

  constructor(
    private singupService: SingupService, private httpClient: HttpClient,
    private router : Router,) 
    {}

  newData: User = {
    tFirstName: null,
    tLastName: null,
    userName: null,
    tId: null,
    tEmail: null,
    tWorkPlace: null,
    tPassword: null,
    tPicture: null,
  };

  onFileSelected(files){
    if (files.length === 0)
      return;
    
    var mimeType = files.target.files[0].type;
    if (mimeType.match(/image\/*/) == null) {
      this.message = "Only file image Ex file .png";
      alert(this.message);
      return;
    }
    this.imagePath = files;
    if (files.target.files.length == 1) {
      this.reader.readAsDataURL(files.target.files[0]); 
      this.reader.onload = (_event) => { 
        this.img = this.reader.result;
        this.newData.tPicture = this.img;
        console.log("this.newData.tPicture = " + this.newData.tPicture);
        this.imgURL = this.img;
      }
    }
    console.log("img = " + this.img);
    console.log("newData.tPicture = " + this.newData.tPicture);
    console.log(this.imagePath);
  }
  
  ngOnInit() {
    console.log(this.newData.tPicture);
  }

  checkData() {
    if(this.newData.tFirstName == null){ alert("กรุณาป้อน First Name") }
    else if(this.newData.tLastName == null){ alert("กรุณาป้อน Last Name")}
    else if(this.newData.userName == null){ alert("กรุณาป้อน User name") }
    else if(this.newData.tId == null){ alert("กรุณาป้อน ID") }
    else if(this.newData.tEmail == null){ alert("กรุณาป้อน E-mail") }
    else if(this.newData.tWorkPlace == null){ alert("กรุณาป้อนสถาณที่ทำงานของท่าน") }
    else if(this.newData.tPassword == null){ alert("กรุณาป้อน Password") }
    else if(this.tComfirmPassword == null){ alert("กรุณาป้อน `Confirm password") }
    else if(this.newData.tPassword != this.tComfirmPassword){ alert("กรุณาป้อน password และ confirm password ให้ตรงกัน") }
    else { this.save() }
  }

  save() {
    console.log(this.newData);
    this.singupService.registerUsers(this.newData).subscribe(
      data => {
          alert('success');
          this.router.navigate(['/Home']);
      },
      error  => {
        alert('Error กรุณาลองใหม่');
        console.log('Error', error);
      }   
    );
  }
}
