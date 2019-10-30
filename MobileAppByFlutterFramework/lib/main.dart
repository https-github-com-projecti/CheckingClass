import 'package:flutter/material.dart';
import 'package:project1/data/Attendance.dart';
import 'Page//ForgotPW.dart';
import 'Page/SignUp.dart';
import 'Page/Login.dart';
import 'Page/SplashScreen.dart';
import 'Page/SubjectInFor.dart';
import 'data/Subjects.dart';
import 'data/User.dart';

var url = ("https://jsonflutter.herokuapp.com/");
var url2 =("http://5b1d04d3.ap.ngrok.io/");
var postUserURL = url2 +'Student/new';
var urlLogin = url2 + "Student/login";
var joinClassURL = url2 + 'Subject/join';
var dateCheck = url2+'Attendance/CheckQR';
var selfieqr = url2 + 'Attendance/Checknamewithsocket';
var putPhoneURL = url2 + 'Student/Editphone';
var putemailURL = url2 + 'Student/Editemail';
var loginState;
var sigupState;
var scanState;
var selfieState;
bool stateCheckLogin = true;
int loginerror;
Attendance atten;

String signup = "signup";
String forgetpw = "forgetpw";
String homeclass = "homeclass";
String login = "login";
String subject = "subject";
String subjectInFor = "subjectinfor";
String camera = "camera";
int numColor = 500, numColor1 = 50, numColor2 = 100;
String xUser;
var users = new List<User>();
String headerUser;
var attendance = new List<Attendance>();
var subjectName = new List<Subjects>();

void main()  {
 
  runApp(MyApp());
  // runApp(Episode5());
}
class MyApp extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    // List<CameraDescription> cameras;

    return MaterialApp(
      title: 'SUT Attendances',
      home: SplashScreen(),
      debugShowCheckedModeBanner: false,
      routes: <String, WidgetBuilder>{
      // Set named routes
      SignUp.routeName: (BuildContext context) => SignUp(),
      ForgetPW.routeName: (BuildContext context) => ForgetPW(),
      Login.routeName:(BuildContext context) => Login(),
      // Subject.routeName:(BuildContext context) => Subject(),
      SubjectInFor.routeName:(BuildContext context) => SubjectInFor(),
      // Camera.routeName:(BuildContext context) => Camera(cameras),
      },
      // theme: ThemeData(
      //   primarySwatch: Colors.blue,
      //   primaryColor: Colors.white,
      //   canvasColor: Colors.white,
      //   appBarTheme: AppBarTheme(
      //     elevation: 0,
      //     color: Colors.white,
      //   )
      // ),
    );
  }
}

 