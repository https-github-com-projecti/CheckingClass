import 'package:http/http.dart' as http;
import 'dart:async';
import '../main.dart';

class API {
   static Future getUsersLogin(String id) {
    var urlUsername = (url2 + 'Student/info/'+ id);
    return http.get(urlUsername);
  }
  static Future getUsersformDB() {
    var urlUsername = (url2 + 'Student/info/'+ xUser);
    return http.get(urlUsername);
  }
  static Future getSubject(String studentID){
    return http.get(url2 +'Student/subject/'+studentID);
  }
  static Future getAttendance(String studentID,String pass){
    return http.get(url2+'Student/Attendance/'+ studentID +'/'+ pass);
  }

 }

