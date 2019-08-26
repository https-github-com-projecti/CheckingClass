import { Subscription } from 'rxjs';
import {newUser} from './../Entity/newUser.entity';
import {SingupService} from './../service/singup.service';
import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroupDirective, NgForm, Validators} from '@angular/forms';
import {ErrorStateMatcher} from '@angular/material/core';
import { DataSource } from '@angular/cdk/collections';
import { DomSanitizer } from '@angular/platform-browser';
import { HttpClient } from '@angular/common/http';

export interface User {
  tName: string;
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
  imgURL: any = "../../assets/fbb2978e127f2920ab9774076ade2a36.png";;
  private message: string;
  private reader = new FileReader();
  private img: any = null;
  private base64 :string = null;


  matcher = new MyErrorStateMatcher();
  constructor(private singupService: SingupService, private httpClient: HttpClient) {
  }

  newData: User = {
    tName: null,
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
    this.reader.readAsDataURL(files.target.files[0]); 
    this.reader.onload = (_event) => { 
      this.img = this.reader.result;
      this.base64 = btoa(this.img);
      this.newData.tPicture = this.base64;
      this.imgURL = this.img;
    }
  }
  
  ngOnInit() {
  
  }

  checkData() {
    if(this.newData.tName == null){ alert("กรุณาป้อน Name") }
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
        if (data){
          alert(data);
        }
        else {
          alert('success');
        }
      }      
    );
  }
}
