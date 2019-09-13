import 'package:http/http.dart' as http;
import 'dart:async';

var url = ("https://e18150df-4163-452d-baea-abe25b1b30ca.mock.pstmn.io/");
class API {
  static Future getUsers() {
    var urlname = "users";
    var urlUsername = (url + urlname);
    return http.get(urlUsername);
  }
}

