import { Component, OnInit } from '@angular/core';
import {HomeService} from '../service/home.service';

@Component({
  selector: 'app-setting-profile',
  templateUrl: './setting-profile.component.html',
  styleUrls: ['./setting-profile.component.css']
})
export class SettingProfileComponent implements OnInit {
  private mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  private getPic = null; 
  private userdata = null;

  constructor(
    private homeService : HomeService,
  ) { }

  ngOnInit() {
    this.loadData()
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

}
