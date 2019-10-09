
import 'dart:async';
import 'dart:convert';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:project1/data/Subjects.dart';
import 'HomeClass.dart';
import 'package:project1/data/API.dart';
import 'package:project1/data/User.dart';
import 'package:project1/main.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';


class SplashScreen extends StatefulWidget {
  @override
  _SplashScreenState createState() => new _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen>  {
  int i;
  var users = new List<User>();
  startTime() async {
    var _duration = new Duration(seconds: 10);
    return new Timer(_duration, navigationPage);
  }

  void navigationPage() {
    // Navigator.of(context).pushNamed("/"+ subjectInFor);
    
    if(xUser == null){
      Navigator.of(context).pushNamed("/" + login);

    }else{
      if(stateCheckLogin==false){
        
        Navigator.of(context).pushNamed("/" + login);
      }else{
        Navigator.push(context,MaterialPageRoute(builder: (context) => HomeClass(user: users[0],)),);
      }
    }
        
    
  }

  getIndex() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
     xUser = prefs.getString('User');
     print("xUser : == "+ xUser.toString());

    API.getUsersformDB().then((response) {   
      Iterable list = json.decode(response.body);
      users = list.map((model) => User.fromJson(model)).toList();
      print(users[0].idstudent);
      API.getSubject(users[0].idstudent).then((response) {   
        Iterable list = json.decode(response.body);
        subjectName = list.map((model) => Subjects.fromJson(model)).toList();
        final int statusCode = response.statusCode;
        if (statusCode < 200 || statusCode > 400 || json == null) {
          throw new Exception("Error while fetching data");
          
        }
      });
    });
     
    
  }
  
  @override
  void initState() {
    super.initState();
    getIndex();
    startTime();
  }
  
  Widget build(BuildContext context) {
     ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
    return  Scaffold(
     body: new Center(
        child:Column( 
          children: <Widget>[ 
            Padding( 
              padding: EdgeInsets.only(top: 150.0),
              child: Column( 
                children: <Widget>[ 
                  Text("Welcome To SUT Attendance",
                    style:  TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 15.0
                    ),
                  ),
                  SizedBox(height: ScreenUtil.getInstance().setHeight(50),),
                  SpinKitDualRing(
                    color: Colors.amber,
                    size: 50.0,
                  ),
                ],
              ),
            )
          ],
        )
      ),
    );
  }


  
}