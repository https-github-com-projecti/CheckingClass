import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';

import 'package:project1/main.dart';


class Selfie{
  String studentId;
  String imageSelfie;
  String dataCheck;
  String firstName;
  String lastName;
  
 

  Selfie({
    this.studentId,
    this.imageSelfie,
    this.dataCheck,
    this.firstName,
    this.lastName
    
  });

  factory Selfie.fromJson(Map<String, dynamic> json){
    return Selfie(
      studentId : json['StudentID'],
      imageSelfie: json['ImageSelfie'],
      dataCheck: json['AQRcode'],
      firstName: json['SfirstName'],
      lastName: json['SlastName']
    );
  }

  Map toMap(){
    var map = new Map<String, dynamic>();
    map['StudentID'] = studentId;
    map['ImageSelfie'] = imageSelfie;
    map['AQRcode'] = dataCheck;
    map['SfirstName'] = firstName;
    map['SlastName'] = lastName;
    return map;
  }

}

Future<Selfie> selfieQR(String url, {Map body}) async {
 print("Send: \\|/");
 print(body);
 print(json.encode(body));
 print(" ");
  return http.post(url, body: json.encode(body))
  .then((http.Response response) {
    final int statusCode = response.statusCode;
 
    if (statusCode < 200 || statusCode > 400 || json == null) {
      throw new Exception("Error while fetching data");
    }
     print("response: >>> \\|/");
     print(json.decode(response.body));
     print("____________________________________________________________");
     selfieState = json.decode(response.body);
    // return JoinClass.fromJson(json.decode(response.body));
  });
}