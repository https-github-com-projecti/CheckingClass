import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import '../main.dart';


class DataCheck{
  // final int id;
  String dataCheck;
  int tspassword;
 

  DataCheck({
    // this.id,
    this.dataCheck,
    this.tspassword
    
  });

  factory DataCheck.fromJson(Map<String, dynamic> json){
    return DataCheck(
      // id : json['id'],
      dataCheck: json['AQRcode'],
      tspassword: json['TSpassword'],
      
    );
  }

  Map toMap(){
    var map = new Map<String, dynamic>();
    map['AQRcode'] = dataCheck;
    map['TSpassword'] = tspassword;
    
    return map;
  }

}

Future<DataCheck> send(String url, {Map body}) async {
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
     print(response.body);
    scanState = response.body;
    // return JoinClass.fromJson(json.decode(response.body));
  });
}