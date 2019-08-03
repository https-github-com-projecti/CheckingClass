import { Component, OnInit } from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";

@Component({
  selector: 'app-login-component',
  templateUrl: './login-component.component.html',
  styleUrls: ['./login-component.component.css']
})
export class LoginComponentComponent implements OnInit {

  constructor(private matIconRegistry: MatIconRegistry,private domSanitizer: DomSanitizer) { 
    this.matIconRegistry.addSvgIcon(
      `icon_label`,
      `path_to_custom_icon.svg`
    );
    this.matIconRegistry.addSvgIcon(
      "outline-vpn_key",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../assets/outline-vpn_key-24px.svg")
    );
  }

  ngOnInit() {
  }

}
