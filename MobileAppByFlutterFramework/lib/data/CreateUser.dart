import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';

import '../main.dart';


class CreateUser {
  // final int id;
  final String firstname;
  final String lastname;
  final String idstudent;
  final String phone;
  final String email;
  final String username;
  final String password;
  //  final String subjects;///////
  // final List<CreateSubjects>  subjects;/////////
  

 
  CreateUser({
    // this.id,
    this.firstname,
    this.lastname,
    this.idstudent,
    this.phone,
    this.email,
    this.username, 
    this.password
    // this.subjects /////////
  });
 
  factory CreateUser.fromJson(Map<String, dynamic> json) {
    // var list = json['subjects'] as List;////////////////
    // List<CreateSubjects> subjectsList = list.map((i) => CreateSubjects.fromJson(i)).toList(); ///////
    return CreateUser(
        // id : json['id'],
        firstname : json['SfirstName'],
        lastname : json['SlastName'],
        idstudent : json['SID'],
        phone : json['Sphone'],
        email : json['Semail'],
        username : json['Susername'],
        password : json['Spassword'],
        // subjects: json['subjects'] ///////////
        // subjects: subjectsList //////
    );
  }
 
  Map toMap() {
    var map = new Map<String, dynamic >();
    map["SID"] = idstudent;
    map["SfirstName"] = firstname;
    map["SlastName"] = lastname;
    map["Susername"] = username;
    map["Spassword"] = password;
    map["Sphone"] = phone;
    map["Semail"] = email;
    
    //  map["subjects"] = subjects; //////////
 
    return map;
  }
  

}


Future<CreateUser> createUsers(String url, {Map body}) async {
  print(body);
  return http.post(url, body: json.encode(body))
  .then((http.Response response) {
    final int statusCode = response.statusCode;
 
    if (statusCode < 200 || statusCode > 400 || json == null) {
      throw new Exception("Error while fetching data");
    }
    sigupState = json.decode(response.body);
    print(json.decode(response.body));
    print(" ");
       //ลองเปลี่ยนดู
  });
}

Future<CreateUser> editPhoneAndEmail(String url, {Map body}) async {
  print(body);
  return http.put(url, body: json.encode(body))
  .then((http.Response response) {
    final int statusCode = response.statusCode;
 
    if (statusCode < 200 || statusCode > 400 || json == null) {
      throw new Exception("Error while fetching data");
    }  
    print(json.decode(response.body));
    print("Put ลอง ดูนะจ๊ะ");
    
  });
}


