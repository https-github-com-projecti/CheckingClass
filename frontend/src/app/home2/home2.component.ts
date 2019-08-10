import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home2',
  templateUrl: './home2.component.html',
  styleUrls: ['./home2.component.css']
})

export class Home2Component implements OnInit {

  mypic: string = 'https://via.placeholder.com/150/FFFFFF/000000/?text=Check%20name%20';

  constructor() { }
  readURL(event:any) {
    if (event.target.files && event.target.files[0]) {
      var reader = new FileReader();
      reader.onload = (event:any) => {
       this.mypic = event.target.result;
      }
      reader.readAsDataURL(event.target.files[0]);
    }
  }

  show_state : boolean = !false;
  ngOnInit() {
    var show_ui = document.getElementById("show");
    if (this.show_state == true){
      show_ui.style.display = "block";
    }
  }
}
