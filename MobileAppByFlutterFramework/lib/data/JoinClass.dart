import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';


class JoinClass{
  // final int id;
  String codeJoinClass;
  String idstudent;
  String firstname;
  String lastname;

  JoinClass({
    // this.id,
    this.codeJoinClass,
    this.idstudent,
    this.firstname,
    this.lastname
  });

  factory JoinClass.fromJson(Map<String, dynamic> json){
    return JoinClass(
      // id : json['id'],
      codeJoinClass: json['Password'],
      idstudent: json['StudentID'],
      firstname: json['SfirstName'],
      lastname: json['SlastName']
    );
  }

  Map toMap(){
    var map = new Map<String, dynamic>();
    map['Password'] = codeJoinClass;
    map['StudentID'] = idstudent;
    map['SfirstName'] = firstname;
    map['SlastName'] = lastname;
    return map;
  }

}

Future<JoinClass> joinClasss(String url, {Map body}) async {
  print(body);
  return http.post(url, body: json.encode(body))
  .then((http.Response response) {
    final int statusCode = response.statusCode;
 
    if (statusCode < 200 || statusCode > 400 || json == null) {
      throw new Exception("Error while fetching data");
    }
     print(json.decode(response.body));
     
    // return JoinClass.fromJson(json.decode(response.body));
  });
}