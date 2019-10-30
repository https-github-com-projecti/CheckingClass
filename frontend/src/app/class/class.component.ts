import { ClassService } from './../service/class.service';
import { Component, OnInit, OnDestroy,} from '@angular/core';
import { MatIconRegistry } from "@angular/material/icon";
import { DomSanitizer } from "@angular/platform-browser";
import { DatePipe } from '@angular/common';
import {HomeService} from '../service/home.service';
import { Router } from '@angular/router';
import { SocketService } from "../service/socket.service";
import { DataAuthenService } from '../service/data-authen.service';
import { newTimeSubject } from '../home2/home2.component';

export interface createQrcode {
  time : string;
  user : string;
  passOfCouse : number;
  clientId : string;
}

export interface date{
  ADate : string; 
}

@Component({
  selector: 'app-class',
  templateUrl: './class.component.html',
  styleUrls: ['./class.component.css']
})
export class ClassComponent implements OnInit {
  mypic: string = "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAA9QAAAPcCAMAAACXWQPbAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAGUExURUdwTAgICFc719MAAAABdFJOUwBA5thmAAARMklEQVR42uzXgQ0CMQzFULr/0syAhHI593kDYv/q+HyQ4/yAawGBIRs4EB+zcQPVNds2EJ2zaQPNPVs2UNyzZQPJQVs2ENyzYQPFQRs2UFy0YQO9Qds1EFy0YQO9Qds1EFy0XQO9Rds10Fu0XQO9Rds10Fu0XcOkj1kDFm3XgEXbNWDSZg2YtFnDou0aMGmzBkzarAGLtmvApM0aJm3WgEmbNWDSZg2YtFkDJm3WMGmzBkzarAGTNmvApM0asGmrhknDrGHSZg2YtFkDNm3VgEmbNUwaZg2btmrApM0asGmrBkzarGHSMGvYtFUDJm3WgE1bNWDSZg2bhlXDpq0aMGmzBmzaqmHSMGvYNKwaNg2rhkmbNXDBpmu/ALhu09f/JCDSv18HVKr3K4FM7X4s0MncTwZCffvhQCdtPx8IRe0GVo1Sz+5g1Qil7BpmjVLFLmLVKBXsKlaNUr0uY9Uoles6Vo1StS5k1UgV60hWjVKtLmXWSJXqWFaNUqYOZtVIJepmVo1Un+5m1Si16XROh1SYrud6SFXpgA6IVJJu6IZI9eiMzohUjC7pkkiV6JiOiVSG7umeSDXopE6KVICu6qpQn7u6K9a257ROi1R4ruu6SFXnwA4MyTmxE2NtcG7sylCbO7sz1rbmxFOndmEoza2hM5k5N+6OzH1dHApzczeHvlzd1TFSl/M6PKTl9E6PtWU5ruNDVs7v/BAVAQRgJCm35QB6YoEFqIkHHjATk9MuMeGySlISF9CRjtiAisAHNMQII3goIXflBPoBK1APL7xAO8xAOsqhBsIBOQhl46r8QDNgCHuTcVSKIBiQBLnQRBNmanFTnqAVMAWlcMUVZkJxUrIgE9AFkRBGGCRCGWUKEQhn0AdYQykPB6UN4gBxkAaow0wZ7skdvPYgD7IAffABxx9/mtAEg1AEOEQqCOckEXIAjVADaIQYiCRSC1JgEkoAl9AB2IQMwCZUwCefItAAoZAAKEWpALfkFF51kAr6QSt8qPHKK/fcM8ss9W7JLJgHt9gr3inJhccc7MJbDnpBOsEEc845wQRT7pQUo2XcJTmGVxwkwyMOlsE2eCabbJ55JtshrRo2DarBNKgG0SCbZ57JJptmh6QbLcvuyDe83CAcHm4wDu82KGeYYM45J5hgzsEvWMdivc5IOzzZ4B1ebBAPbkE8tdRSTz2zzFIPYkE+eAX5mPLqiuzDWw364akG/yAV/HPKqQIUQCmlCgCj0AAIhQZAKDTAJ58qUAGddKpABTWdjigDGXiioQN4oSEEcAkhUEmlFKTAJJNSkAKTkAI2m3RDLWjB6wwxwOMMNYBGqIFFGuUgBxZJ1IMeSIQewCF8f8OooQgKKYQiKGRQEpLwLEMTIBCagE8tiII//iAK+vhThSq8yZAF2IMswB5kQR55EAZ33AlDGNxBGOAOwgB1kAZzzEEazDEnDWkkzbmgNrRBHLQBn1gQB3iDOHijDeqgjTZ1qIM2qAP+NUEe8BRDH6SRBn2QRpo+9EEa9AHOoBBQBoVQRhkUQhllClEIZVAI1itzQIlIxDMMjYAwaASEQSOEEQaNEMaXSETCF0QCviAS+LsEldBFF1RCF10qUYl/S5AJ2IJMwBZkwhZbkAlbbMlEJmxBJtgviy2d6MQLDKGAKwgFXEEoXHEFoXDFlVCEwhWEAq4gFHAFoXDFFYTCFVdCEUrRlfspRSneXxg1jBpSAVOQClNMQSpMMSUVqTAFqYApSAVMQSpMMQWpMMWUVKTCFKQCpiAVMAWpMMUUpMIUU1KRClOQCpiCVMAUpAKmIBWmmJKKVJiCVByQKUgFTEEqYApSYYopqUiFKUhFKkxBKmAKUgFTkApTTElFKkxBKlJhClIBU5AKmIJUmGJKKlJhClKRClOQCpiCVMAUpEIVU1AKVRCKUriCUMAVhAKuIBSuuIJQuIJQhMIVhAKuIBRwBaFwxRWEcpcrsnSiE7KgE/isgkzAFmTCFluQCVtsyUQmbEEm2G6LLpWohC6oBHRBJfBvCSLhiy+IhC++RCISf5egERAGjYAwaIQwwqARxgiTiEQYg0Tg2woKAWVQCGWUQSGUUaYQhYSVcSYQfZAGfYA06AOkQR+kkQZ9kMaaPOThKYY6QBvUAdqgDtp4gzh4400c4vCFBW3AawxpgDlIgznmIA3qmFOGMrzHEAa4gzDAHYSBf7gjTxe68CJDFmAPsoDvLKiCPvqgCv7oE4UoPMrQBAiEJkAgfH0zSCEUQSGFilAEhVAEOIQeQCL0QCKJ0AOJLMpBDp5mqAHeZogBPEIM8MUFLRDJpBSkwCSkACohBHAJIYBLCIFMMmUgAzYhA7zSJp0qUIE3GiKARxoaAKHQAKOMKkABlFKqAAUklXIqAHiowT9IBf/w+QX6aaWVfPJ55ZV8EAvqwSyox7RZapmH9xrEw4MN3uHFBu3s0ss66/TSyzr4BecgGJyDYTBOMceEE04yx3zDyw26wTLoBs0gm2eeySabaKKpRt401UzD+w2i4QEHz2AbLNNNN8ss8803x2gKZ5xieMbBMLzjIBik00sv66zTSy/ttJOLpnfiuYXXHNTCcw5mwT2vvJJPPq+8ss8+q2jq559UCACUYncBEmAU3nUQCg87+IQK2GRTBzrgkkshCIFLKAFM4g0paIFIiAE0Qg4gEZM9CIJDeOVBITzzYBCa4I8/VaiCPchCF+RBGKAO0gBxeKANcfAGdYA17M5DH6RBIKAMEiGMMEw2IhK+4OUHXZAJWWRh8oNOKb69YdUgCmKhiSbIhSSS9CIYjqAYxTAEzYAfvCMa1dg0rBrkQDjUUAPpEEOMdtRDC+QjH1IgIFACCRFCCEREBx0qkhEb0JGQqICUQATERAMNeDgnPXGAWlCKogCSIoAALI9KVa4PXbm920NZLu/yGG1LXM6OWl3ycnXk+hKYk0NiDu7gWB6ZylwbOnNrt8by0qTm0BCbMzszluemNzeG4lzYhbG8ueO8zotadsdt3Rax8O5Mz2Fh1e7qrnhTfcdRHRW1AI+LuihqDR7ndM5vO3ZgGzsOBUEQyj9px2AsIA2bVRn4cfq0/6jN8HFLt6S2xMchHZLaGB9XdEVqe3yc0AmpTfJxP/fDLN1O0owv83E4TVMb5+Nqmqa2z/MX6mLYaGqkzoWZlobqVJhqaqvuhLWW9upGqDo1WQdC1aXZOg6yLk3XXVB1ab9ugqpLG3YPVB0aslOg6tKYnQFVhxbtAsg6NOub/3ZUnRv3pX82qk5O/Lo/GFV3l37PX4qsZ/kLobb5X5ff/cvg7O3/M4LYnwPtrEOsDlVrGlStaWSNpFE1mkbVmgZZSxpUrWmQtaRRNZpG1pIGVWsaZC1pULWmkTWSRtWaBllLGlStaZC1pJE1kkbVmgZZSxpkLWmQtaRRNLpG0boGResaJK1rULSsUTS6RtHoGkXrGiQta1C0rpE0skbR6BpJyxoUrWuQtKxRNLpG0sgaScsaFK1rJI2skTSyRtLIGknLGiQtaxSNrpE0skbSyBpJI2skLWskjayRNLJG0sgaTSNrJK1qJI2skTSyRtLIGkkjayQtazR905gdAkmn1+swSDq7WEdC0tWZuheabq7T3ZB0c5buh6SLc3RHbDE5Q+fEBosDdFWMr7g8x8XsgptzYQyuODd3xtKCO3Nurh6Zm7s5pXk5vMMTmpbruz6hUXkDb0BpT97BO1CaksfwGIRm5EW8CKUBeRDPQmk8nsPTUNqN1/A+lDbjLbwRpbl4Cg9Faipe4pzH8hBWYiUejLsm4hk8Gql5eIUzH84jWIZleDz6s/AEHhCLwCOyOgcv0HhID2AKluAtMQO8J9MbcH1vSur9Hd+z4u3xtKw+vNt7XUqv7vJeGA+OV2b1tR3eQ5N6anf31qTe2dk9N94YT87qA7u6Vyf1uo7u4Uk9rZt7ezwr3p/VN3VyEyD1oC5uBVaQek0HNwRD8B9obAH/dcYcPKI3xCK8oAc0CqO4+flc2y7swtthGcy+nGMbh3F4NQwET4aJeDDvhZV4LY9lKIbiqTAVU6k8lEtbi7WkXsmhDcZg/IcXm8F/dTEbj+NxMBxP42WwHe+C9VhP5lWc2YAMyJNgQiY0+yCubEVW5DWwIzuafQtHNiVT8hAYkzHNPoMb25M9eQMsyqK8APeMyomdH7NyfMfHsJze5bEtd8e6rMvVsS/7+uzmDoyJOThGZmTOjZnxyrHdF0tzaWzN1twZa7O2V67svBicE2NyJufAGJ3ROS9m57hui+W5LNieu2J91jd6VpvF/FJHNVgs0EWxQRt0T6zQCl0TO3RLt8QSXRJsMXRHE8UcHREM0gkxSYt0QWzS/dwPq3Q9sEu3A8t0OWzTNl+4m0Finq2rWSP26WZgoS6Gqm3UvbDS9kpdCzt1K01jqbedygIxVYcCY3UmVG2ujoTBxgfrRJhsbLIOhM22Nus+WG1rta6D3bZ26zZYbmu5LoPtxrbrLhhva7zOgvm25usoGHBrwE6CCccm7CCourVh90DVrRW7Bqpu7dgtUHVsyS6BqltTdghU3RqznyyourVmTSNqUWsag3YCMGkHAKNe+PstCqv214Nd+9vBsv3lqFrUokbUx21b06i6tW5No+rYvjWNqlsD1zSqbk1c06g6NnJNo+rWyn2oEXVr5ppG1bGhaxpVt5auaVTd2rqmUXVs7aJG1a21axpRt/auaVQdW7ymUXVr8ppG1bHRixpRt0avaVTdmr2mUXVs+JpG1a3laxpVt7bvxzeijo1f06i6tX5No+rY/jWNqlsB+FAj6lYBmkbVsQY0japbEfhQQ6wCTUMrA01DKwRNQywFUUMrBU1DLAZNQ6sGH2po5aBpiAWhaWgV4UMNsao1Da0mNA2xKjQNrSx8qCHWhaahFYamoZWGH98Qa0PT0IrDhxpiVWsaWnloGmKBiBpagWgaYoloGlqN+FBDLBJNQ6sSTUOsE1FDqxNNQ6wUUUOrFE1DrBVNQysWH2qI1aJpaOXiQw2xXjQNrWA0DbFkRA2tZDQNsWhEDa1oNA2xbEQNrWw0DbFwRA3flaNp8Kk++D834FPtQw0+1T7UUKtH1NCqR9MQ60fU0OpH0xArSNTQKkjTEGtI1NBqSNMQq0jUIGpNw3BGPtQQ60jT0ArJhxpiJWkaWin5UEPsU61piH2qRQ2tqDUNsapFDa2oNQ2xqkUNotY0DBflQw2xT7WmofWp9qGG2Kda0xD7VIsaWlFrGmJVixpaUfvfZDAa9XPgf1BA1aIGUWsa7qpa1CBqTcNwXD7UEPtUixpin2pNQytqH2qI/f7WNLQ+1T7UEPtUaxpin2pRQytqv74h9vtb0xD7VIsaWlFrGmJVixpELWoYjlrTEKta1CBqTcMHVftQw6WfalGDqDUNw7+/fagh9qkWNYha07D8+1vU0Irar2+IVa1pELWoYThqTUOsalGDqEUNw1FrGmJVixpELWpIR+2N4OWqfajhqk+1qEHUmobh398+1BD7VIsaRK1pWP79LWpoRe3XN8R+f4saRK1pWP79LWpoRe3XN8R+f2saYp9qUYOoRQ3DUWsaYlWLGkQtahA18FZ+moZY1aIGUYsaylF7Ffiyah9qyH+qRQ2iFjWUo/YmsPWPah9qiH2qRQ2iFjWUo/YiMPaPalFDK2q/viH2+1vUIGpNw/Lvb1GDqEUNw1H7JzXE/lGtaYh9qkUNohY1iBp4K2r/nwwSUT8+1FD9VIsaRC1qEDXwVtSahljVogZRixpEDYgaRP1J1N4BZqr2oYbkp1rUIGpRg6gBUYOoP4naK8BQ1T7UEPxUixpELWoQNXBK1N4ApqoWNYha1LAdtX9SQ+wf1aIGUYsaRA2cErUXgLGqRQ2iFjWIGngt6j+vOcJU8Jf7dgAAAABJRU5ErkJggg==";
  myClass: Object = null ;
  guillaumeNery = 'Hello my code';
  pipe = new DatePipe('en-US');
  private nows = Date.now();
  myFormattedDates = this.pipe.transform(this.nows, 'medium');
  private Qrcode : any = null;
  private userdata = null;
  public getPic = null; 
  public check : any = null;
  public length : any = null;
  private ShowdDataQr : any = null ;
  private Username : string = null;
  private QrShow : any = null;
  private oneQrcode : any = null;
  private messages: Array<any>;
  private chatBox: string;
  private StudentNow : any = null;  
  today: number = Date.now();
  TypeObject : object;
  public state = false
  private stateClickCreate = 0;
  public ADate : string;
  public ATimeAuthen : string;
  private arr : any = [];
  private arrdata : any = [];
  public timeLimite : any = null; 
  private timeSubject : any = null;

  newQr : createQrcode = {
    time : null,
    user : null,
    passOfCouse : null,
    clientId : null,
  }

  select_Class: any = {
    subject_id : null,
    TSName: null,
    TSDescription: null,
    TSID: null,
    TSpassword: null,
  }

  newDate : date = {
    ADate : null,
  }

  constructor(
    private matIconRegistry: MatIconRegistry,
    private domSanitizer: DomSanitizer,
    private homeService : HomeService,
    private classService : ClassService,
    private router : Router,
    private sockets: SocketService,
    private dataAuthenService: DataAuthenService,
  ) {
    this.matIconRegistry.addSvgIcon(
      "add",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/add24px.svg")
    );
    this.matIconRegistry.addSvgIcon(
      "aspectRatio",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/aspect_ratio24px.svg")
    );
    this.matIconRegistry.addSvgIcon(
      "delete",
      this.domSanitizer.bypassSecurityTrustResourceUrl("../../assets/delete24px.svg")
    );
    this.messages = [];
    this.chatBox = "";
   }

  ngOnInit() {
    var showQr = document.getElementById("showQr-contrainer");
    var showQr3 = document.getElementById("qrCode_Container");
    this.Username = localStorage.getItem('isLogin');
    this.classService.getmyClass().subscribe(data =>{
      this.myClass = data;
      this.classService.passofClass(this.myClass[0]["TSpassword"]);
      this.loadData();
    });

    window.onclick = function(event) {
      if (event.target == showQr) {
        showQr.style.transition = "all ease-out 600ms";
        showQr.style.transform = "rotateX(60deg)";
        showQr.style.top = "-100%";
        var panel = document.getElementById("qrCodeBlockPanel");
        panel.style.display = "none";
        clearInterval();
      }
      if (event.target ==  showQr3){
        showQr.style.transition = "all ease-out 600ms";
        showQr.style.transform = "rotateX(60deg)";
        showQr.style.top = "-100%";
        var panel = document.getElementById("qrCodeBlockPanel");
        panel.style.display = "none";
        clearInterval();
      }
    }
    this.sockets.getEventListener().subscribe(event => {
      if(event.type == "message") {
          let data = event.data.content;
          if(event.data.sender) {
              data = event.data.sender + ": " + data;
              if(event.data.content == "Success"){
                this.state = true;
                if(this.stateClickCreate == 1){
                  this.newQr.clientId = event.data.sender;
                  this.classService.ClientID(event.data.sender);
                  this.createAuthenicate()
                }
              } 
              else if ((typeof event.data.content) == "object") {
                var Id = localStorage.getItem('clientID');
                if (Id == event.data.sender){
                  this.StudentNow = event.data.content[0]["Astudent"];
                }
              }
              else {
                alert("Error กรุณาลองใหม่อีกครั้ง");
              }
              
          }
          this.messages.push(data);
      }
      if(event.type == "close") {
          this.messages.push("/The socket connection has been closed");
      }
      if(event.type == "open") {
          this.messages.push("/The socket connection has been established");
      }
    });
  }

  //---------------------------------- start websockget ---------------------------------------------
  startWebsocket(){
    const now = Date.now();
    if (this.Qrcode == null){
      this.sockets.newSocket();
      this.stateClickCreate = 1;
    }
    else {
      this.classService.getmyQr().subscribe(
        data =>{
          this.Qrcode = data;
          var lenNow = this.Qrcode.length - 1;
          this.newDate.ADate = this.Qrcode[lenNow]["ADate"]; 
          // -------------- เวลาQrcode ล่าสุด ----------------------------
          var nowQrcode = this.Qrcode[lenNow]["ADate"];
          nowQrcode = nowQrcode.split(" ");
          var nowQrcodeDate = nowQrcode[0];
          var nowQrcodeTime = nowQrcode[1];
          var nowQrcodeDateSplite = nowQrcodeDate.split("-");
          var nowQrcodeTimeSplite = nowQrcodeTime.split(":");

          // -------------- เวลาสิ้นสุด ----------------------------
          this.classService.getTimeLimite(this.newDate).subscribe(data => {
            this.timeLimite = data;
            var timeLimiteHMS = this.timeLimite
            var timeLimiteHMSSplite = this.timeLimite.split(":");

            // -------------- เวลาปัจจุบัน ----------------------------
            const myFormattedNowDate : string = this.pipe.transform(now, 'd-MMM-y');  
            const myFormattedNowTime : string = this.pipe.transform(now, 'H:mm:ss');
            var myFormattedNowDateSplite = myFormattedNowDate.split("-");
            var myFormattedNowTimeSplite = myFormattedNowTime.split(":");
            
            // Now
            if ((myFormattedNowTimeSplite[0] == timeLimiteHMSSplite[0]) && (nowQrcodeDate == myFormattedNowDate)){
              var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
              var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
              if(SUMTimeNowTime > SUMTimeLimite){
                this.sockets.newSocket();
                this.stateClickCreate = 1;
              }
            }
             // Month & Year 
             else if ((myFormattedNowDateSplite[1] != nowQrcodeDateSplite[1]) || (myFormattedNowDateSplite[2] != nowQrcodeDateSplite[2])){
              this.sockets.newSocket();
              this.stateClickCreate = 1;
            }
            // Day no same
            else if ((myFormattedNowDateSplite[0] != nowQrcodeDateSplite[0]) && (myFormattedNowDateSplite[1] == nowQrcodeDateSplite[1]) && (myFormattedNowDateSplite[2] == nowQrcodeDateSplite[2])){
              var countNowDate = +myFormattedNowDateSplite[0];
              var countnowQrcodeDate = +nowQrcodeDateSplite[0];
              var CountDate = countNowDate - countnowQrcodeDate;
              if(CountDate < 0){
                CountDate = countnowQrcodeDate - countNowDate;
              }
              var countNowTimeH = +myFormattedNowTimeSplite[0];
              var counttimeLimiteH = +timeLimiteHMSSplite[0]; 
              var CountTimeH = countNowTimeH - counttimeLimiteH;
              if(CountTimeH < 0){
                CountTimeH = counttimeLimiteH - countNowTimeH;
              }
              if (CountDate > 1){
                this.sockets.newSocket();
                this.stateClickCreate = 1;
              } 
              else if (CountTimeH != 23) {
                this.sockets.newSocket();
                this.stateClickCreate = 1;
              }
              else{
                var numNowQrcodeTimeSpliteM = +nowQrcodeTimeSplite[2]
                if ((nowQrcodeTimeSplite[0] == "23") && (numNowQrcodeTimeSpliteM > 30) && (CountTimeH == 23)){
                  if (myFormattedNowTimeSplite[0] == "0") {
                    var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
                    var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
                    if(SUMTimeNowTime > SUMTimeLimite){
                      this.sockets.newSocket();
                      this.stateClickCreate = 1;
                    }
                  }
                }
              }
            }
            // Day == 
            else {
              var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
              var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
              if(SUMTimeNowTime > SUMTimeLimite){
                this.sockets.newSocket();
                this.stateClickCreate = 1;
              }
            }
          });
        }
      );
    }
  }

  //---------------------------------- สร้าง Qr code ---------------------------------------------
  createAuthenicate(){
    const now = Date.now();
    const myFormattedDate : string = this.pipe.transform(now, 'd-MMM-y H:mm:ss');
    const user = localStorage.getItem('isLogin');
    this.newQr.time = myFormattedDate;
    this.newQr.user = user;
    this.newQr.passOfCouse = this.myClass[0]["TSpassword"];

    if (this.stateClickCreate == 1){
      this.classService.createQr(this.newQr).subscribe(
        data => {
          this.check = data;
          if (this.check  == "Success"){
            this.loadData();
            this.ADate = null;
            this.ATimeAuthen = null;
            this.showQrCode();
            this.stateClickCreate = 0;
          }
        },
        error  => {
          alert('Error กรุณาลองใหม่' + error);
        }     
      );
    }
  }
  //---------------------------------- dataAuthen ---------------------------------------------
  dataAuthen(x){
    var pass = x.ASpassword;
    var date = x.ADate;
    var time = x.ATimeAuthen;
    this.dataAuthenService.getDataClassAndQrcode(pass,date,time);
    this.router.navigate(['/DataAuthen']);
  }

  //---------------------------------- load date ---------------------------------------------
  loadData() {
    this.homeService.getUserdata().subscribe(data =>{
      this.userdata = data;
      // this.homeService.setID(this.userdata); //สำหรับserver Test
      this.homeService.setID(this.userdata[0]['user_id']); //สำหรับserver DB
      this.homeService.getGetPic().subscribe(data =>{
        this.getPic = data;
        if (this.getPic.trim() === ''){}
        else {  this.mypic = this.getPic }
      });
    });
    this.classService.getmyQr().subscribe(
      data =>{
        this.Qrcode = data;
      }
    );
    this.classService.getTStimesubject().subscribe(data => {
      console.log(data);
      this.timeSubject = data;
    });
    this.StudentNow = null; //refresh UI show student
  }
  //---------------------------------- เช็คค่า null ---------------------------------------------
  isEmptyOrSpaces(str){
    return str === null || str.match(/^ *$/) !== null;
  }

  //----------------------------------  logout ---------------------------------------------
  logout(){
    localStorage.clear();
    this.router.navigate(['/Home']);
  }

  //----------------------------------  หน้าUI โชQrcode ---------------------------------------------
  showQrCode(){
    this.ShowdDataQr == null;
    var panel = document.getElementById("qrCodeBlockPanel");
    panel.style.display = "none";
    this.classService.getmyQr().subscribe(
      data =>{
        this.oneQrcode = data;
        var numQrcode = this.oneQrcode.length - 1;
        this.ShowdDataQr = this.oneQrcode[numQrcode]["PicQRcode"];
        this.ADate = this.oneQrcode[numQrcode]["ADate"];
        this.ATimeAuthen = this.oneQrcode[numQrcode]["ATimeAuthen"];
        this.ShowTime(this.ADate);
      }
    )
    var showQr = document.getElementById("showQr-contrainer");
    showQr.style.display = "flex";
    showQr.style.transition = "all ease-out 600ms";
    showQr.style.top = "0";
    showQr.style.transform = "rotateX(0deg)";
  }

  //----------------------------------  หน้าUI โชQrcode เก่า ---------------------------------------------
  public zoomScreen(x){
    //start web socket 
    if(this.state == false){
      this.sockets.newSocket();
    }
    //load data student
    var pass = x.ASpassword;
    var date = x.ADate;
    var time = x.ATimeAuthen;
    this.loadDataJoinStudent(pass, date, time);
    // set value panel
    this.ShowdDataQr = x.PicQRcode;
    this.classService.ClientID(x.AClientid);
    this.ADate = x.ADate;
    this.ATimeAuthen = x.ATimeAuthen;
    //start time 
    this.ShowTime(x.ADate);
    //check time out
    const now = Date.now();
    this.newDate.ADate = x.ADate;
    // ----------------------- time limite ---------------------------------
    this.classService.getTimeLimite(this.newDate).subscribe(data => {
      this.timeLimite = data;
      var timeLimiteHMS = this.timeLimite;
      var timeLimiteHMSSplite = this.timeLimite.split(":");

      // -------------- เวลาปัจจุบัน ----------------------------
      const myFormattedNowDate : string = this.pipe.transform(now, 'd-MMM-y');  
      const myFormattedNowTime : string = this.pipe.transform(now, 'H:mm:ss');
      var myFormattedNowDateSplite = myFormattedNowDate.split("-");
      var myFormattedNowTimeSplite = myFormattedNowTime.split(":");

      // -------------- เวลาQrcode ล่าสุด ----------------------------
      var nowQrcode = x.ADate; 
      nowQrcode = nowQrcode.split(" ");
      var nowQrcodeDate = nowQrcode[0];
      var nowQrcodeTime = nowQrcode[1];
      var nowQrcodeDateSplite = nowQrcodeDate.split("-");
      var nowQrcodeTimeSplite = nowQrcodeTime.split(":");
      var panel = document.getElementById("qrCodeBlockPanel"); // หน้าต่าง UI
      // Now
      if ((myFormattedNowTimeSplite[0] == timeLimiteHMSSplite[0]) && (nowQrcodeDate == myFormattedNowDate)){
        var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
        var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
        if(SUMTimeNowTime > SUMTimeLimite){
          panel.style.display = "flex";
        }
      }
      // Month & Year 
      else if ((myFormattedNowDateSplite[1] != nowQrcodeDateSplite[1]) || (myFormattedNowDateSplite[2] != nowQrcodeDateSplite[2])){
        panel.style.display = "flex";
      }
      // Day no same
      else if ((myFormattedNowDateSplite[0] != nowQrcodeDateSplite[0]) && (myFormattedNowDateSplite[1] == nowQrcodeDateSplite[1]) && (myFormattedNowDateSplite[2] == nowQrcodeDateSplite[2])){
        var countNowDate = +myFormattedNowDateSplite[0];
        var countnowQrcodeDate = +nowQrcodeDateSplite[0];
        var CountDate = countNowDate - countnowQrcodeDate;
        if(CountDate < 0){
          CountDate = countnowQrcodeDate - countNowDate;
        }
        var countNowTimeH = +myFormattedNowTimeSplite[0];
        var counttimeLimiteH = +timeLimiteHMSSplite[0]; 
        var CountTimeH = countNowTimeH - counttimeLimiteH;
        if(CountTimeH < 0){
          CountTimeH = counttimeLimiteH - countNowTimeH;
        }
        if (CountDate > 1){
          panel.style.display = "flex";
        } 
        else if (CountTimeH != 23) {
          panel.style.display = "flex";
        }
        else{
          var numNowQrcodeTimeSpliteM = +nowQrcodeTimeSplite[2]
          if ((nowQrcodeTimeSplite[0] == "23") && (numNowQrcodeTimeSpliteM > 30) && (CountTimeH == 23)){
            if (myFormattedNowTimeSplite[0] == "23"){}
            else if (myFormattedNowTimeSplite[0] == "0") {
              var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
              var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
              if(SUMTimeNowTime > SUMTimeLimite){
                panel.style.display = "flex";
              }
            }
          }
        }
      }
      // Day == 
      else {
        var SUMTimeNowTime = this.SUMAllArrayTreeValue(myFormattedNowTimeSplite);
        var SUMTimeLimite = this.SUMAllArrayTreeValue(timeLimiteHMSSplite); 
        if(SUMTimeNowTime > SUMTimeLimite){
          panel.style.display = "flex";
        }
      }
    })
    
    //Show Qrcode
    var showQr = document.getElementById("showQr-contrainer");
    showQr.style.display = "flex";
    showQr.style.transition = "all ease-out 600ms";
    showQr.style.top = "0";
    showQr.style.transform = "rotateX(0deg)";
    
  }

  public SUMAllArrayTreeValue(time){
    var sum : number = 0;
    for (var num = 0; num < time.length - 1; num++){
      var intnum = +time[num];
      if (num == 0){
        sum = sum + intnum;
      } else {
        sum = (sum*100) + intnum;
      }
    }
    return sum
  }

  public loadDataJoinStudent(x,y,z){
    this.classService.getClientData(x,y,z).subscribe(data => {
      this.StudentNow = data;
    });
    this.loadData();
  }

  public send() {
    if(this.chatBox) {
        this.sockets.send(this.chatBox);
        this.chatBox = "";
    }
  }

  public isSystemMessage(message: string) {
  return message.startsWith("/") ? "<strong>" + message.substring(1) + "</strong>" : message;
  }

  getRandomColor(index, i) {
    var color = Math.floor(0x1000000 * Math.random()).toString(16);
    var trueColor = '#' + ('000000' + color).slice(-6);
    if (this.arrdata.length < 1 || this.arrdata == undefined){
      this.arrdata.push(index.StudentID);
    }
    if (this.arrdata[i] != index.StudentID){
      this.arrdata.push(index.StudentID);
    }
    if (this.arr.length < this.arrdata.length){
      this.arr.push(trueColor);
    }
    return this.arr
  }

  ShowTime(ADate){
    // console.log("Hello Showtime!");
    var time = new Date();
    var H = time.getHours();
    var M = time.getMinutes();
    var S = time.getSeconds();
    var realTime = H + ":" + M + ":" + S;
    document.getElementById('timeShow').innerText = realTime;
    document.getElementById('timeShow').textContent = realTime;
    this.newDate.ADate = ADate;
    setInterval(this.ShowTime, 1000); 
  }
}
