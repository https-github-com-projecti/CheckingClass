import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";

@Component({
  selector: 'app-people',
  templateUrl: './people.component.html',
  styleUrls: ['./people.component.css']
})

export class PeopleComponent implements OnInit {
  mypic: string = '../../assets/fbb2978e127f2920ab9774076ade2a36.png';
  panelOpenState = false;
  step = 0;

  constructor(private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    ) 
    {
      this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")); 
    }

  ngOnInit() {
  }

}
