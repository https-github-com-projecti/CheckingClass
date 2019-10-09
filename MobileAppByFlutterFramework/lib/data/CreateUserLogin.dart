import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import '../main.dart';
import 'User.dart';

class CreateUserLogin{
  // final int id;
  final String username;
  final String password;

  CreateUserLogin({/*this.id,*/ this.username, this.password});

  factory CreateUserLogin.fromJson(Map<String, dynamic> json){
    return CreateUserLogin(
      // id : json['id'],
      username : json['username'],
      password : json['password'],
    );
  }

  Map toMap(){
    var map = new Map<String, dynamic>();
    map["username"] = username;
    map["password"] = password;
    return map;
  }
 

}


  
Future<CreateUserLogin> createuserLogine(String url, {Map body}) async {
  print(body);
  return http.post(url, body: json.encode(body)).then((http.Response response) {
    final int statusCode = response.statusCode;
    
    print(json.decode(response.body));
    print(".......");
    loginState = json.decode(response.body);    //ได้  json กลับมาจาก post
    Iterable list = json.decode(response.body);
    users = list.map((model) => User.fromJson(model)).toList();
    print(users[0].id);  
    loginerror = statusCode;
    if (statusCode < 200 || statusCode > 400 || json == null) {
      
      throw new Exception("Error while fetching data");
      
      
    }
    
    // return CreateUserLogin.fromJson(json.decode(response.body));
  });
}