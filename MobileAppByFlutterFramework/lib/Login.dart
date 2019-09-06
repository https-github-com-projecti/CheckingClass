import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'data/API.dart';
import 'data/User.dart';
import 'main.dart';
import 'package:flutter/services.dart';
import 'dart:convert';

var xText = " ";

TextEditingController _usernameController = TextEditingController();
TextEditingController _passwordController = TextEditingController();

class Login extends StatefulWidget{
  static const String routeName = "/login";
  @override
  LoginState createState() => new LoginState();
}

class LoginState extends State<Login>{

  var users = new List<User>();
  List data;

  _getUsers() {
    API.getUsers().then((response) {
      setState(() {
        Iterable list = json.decode(response.body);
        users = list.map((model) => User.fromJson(model)).toList();
        data = users;
      });
    });
  }
  // initState() {
  //   super.initState();
  //   _getUsers();
  // }

  // dispose() {
  //   super.dispose();
  // }

  getData(){
    _getUsers();
    print("********************TEST**********************");
    bool login1 = false;
    bool login2 = false;
    for (var i = 0; i < users.length; i++) {
      if(_usernameController.text == users[i].username){
        print("Good1 "  + i.toString() );
        login1 = true;
          // for (var j = 0; j < users.length; j++) {
            if(_passwordController.text == users[i].password){
              print("Good2 " + i.toString());
              login2 = true;
              break;
            }else{
              print("Password wrong " + i.toString());
              break;
            }
          //}
      }else{
        print("ผิด " + i.toString());
      }
    }
    if(login1 == true && login2 == true){
      Navigator.of(context).pushNamed("/" + homeclass);
    }    
    print("******************************************");
  }


  @override 
  Widget build(BuildContext context){
    ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
    ScreenUtil.instance = ScreenUtil(width: 500, height: 1200, allowFontScaling: true);

     SystemChrome.setPreferredOrientations([
        DeviceOrientation.portraitUp,
        DeviceOrientation.portraitDown,
      ]);

    return new Scaffold(
      
      backgroundColor: Colors.amber[numColor] ,
      // resizeToAvoidBottomPadding: false ,
      body: Center( 
        child: SingleChildScrollView(
          child: Column( 
            children: <Widget>[ 
              Column(
                 crossAxisAlignment: CrossAxisAlignment.start,
                children: <Widget>[
                  Padding(
                    padding: EdgeInsets.only(top: 0.0),
                    child: Container(
                      child: Text("SUT",
                          style: TextStyle(
                            fontFamily: "Poppins-Bold",
                            fontSize: 50,
                            fontWeight: FontWeight.bold,
                            color: Colors.white,
                          )
                      )
                    ),
                  ),
                ],
              ),
              SingleChildScrollView(
                  child: Padding( 
                    padding: EdgeInsets.only(left: 28.0, right: 28.0, top: 0.0),
                    child: Column( 
                        crossAxisAlignment: CrossAxisAlignment.center,
                      children: <Widget>[
                        SizedBox(height: ScreenUtil.getInstance().setHeight(50),),
                        FormLogin(),
                        SizedBox(height: ScreenUtil.getInstance().setHeight(40),),
                        InkWell(
                          child: Container(
                           // margin: const EdgeInsets.fromLTRB(0.0, 0.0, .0, 0.0),
                            width: ScreenUtil.getInstance().setWidth(200),
                            height: ScreenUtil.getInstance().setHeight(100),
                            decoration: BoxDecoration(
                                gradient: LinearGradient(colors: [
                                  Color(0xFFFF1744),
                                  Color(0xFF00BCD4)
                                ]),
                                borderRadius: BorderRadius.circular(6.0),
                                boxShadow: [
                                  BoxShadow(
                                      color: Color(0xFF6078ea).withOpacity(.3),
                                      offset: Offset(0.0, 8.0),
                                      blurRadius: 8.0)
                                ]),
                            child: Material(
                              color: Colors.transparent,
                              child: InkWell(
                                //onTap: () { Navigator.of(context).pushNamed("/" + homeclass); },
                                onTap: () { getData(); },
                                child: Center(
                                  child: Text("SIGNIN",
                                          style: TextStyle(
                                              color: Colors.white,
                                              fontFamily: "Poppins-Bold",
                                              fontSize: 18,
                                              letterSpacing: 1.0)
                                    ),
                                ),
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  )    
              )  
            ],
          ),
        )
      ),
    );
  }
}



//----------------------------------------------------------------------------------------------------------------------//

class FormLogin extends StatelessWidget {
  initState() {
    xText = _passwordController.text;
    print(xText);
  }

  @override
  
  Widget build(BuildContext context) {
    //double height = MediaQuery.of(context).size.height;
        return new Container(
         
          width: double.infinity,    
          decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(10.0),
              boxShadow: [
                BoxShadow(
                    color: Colors.black12,
                    offset: Offset(0.0, 15.0),
                    blurRadius: 15.0),
                BoxShadow(
                    color: Colors.black12,
                    offset: Offset(0.0, -10.0),
                    blurRadius: 15.0),
              ]),
          child: Padding(
            padding: EdgeInsets.only(left: 16.0, right: 16.0, top: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                Text("Login",
                    style: TextStyle(
                        fontSize: ScreenUtil.getInstance().setSp(45),
                        fontFamily: "Poppins-Bold",
                        letterSpacing: .6)),
                SizedBox(
                  height: ScreenUtil.getInstance().setHeight(30),
                ),
                //Username
                Text("Username",
                    style: TextStyle(
                        fontFamily: "Poppins-Medium",
                        fontSize: ScreenUtil.getInstance().setSp(26))),
                TextField(
                  controller: _usernameController,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0))
                                                        , borderSide: BorderSide(color: Colors.amber,),
                            ),
                    focusedBorder:OutlineInputBorder(
                                            borderSide: const BorderSide(color: Colors.teal, width: 2.0),
                                            borderRadius: BorderRadius.circular(20.0),
                              ),
                      hintText: "username",
                      hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      
                ),
                SizedBox(
                  height: ScreenUtil.getInstance().setHeight(30),
                ),
                //Password
                Text("Password",
                    style: TextStyle(
                        fontFamily: "Poppins-Medium",
                        fontSize: ScreenUtil.getInstance().setSp(26))),
                TextField(
                  obscureText: true,
                   controller: _passwordController,
                   
                  decoration: InputDecoration(
                    border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0))
                                                        , borderSide: BorderSide(color: Colors.amber,),
                            ),
                    focusedBorder:OutlineInputBorder(
                                            borderSide: const BorderSide(color: Colors.teal, width: 2.0),
                                            borderRadius: BorderRadius.circular(20.0),
                              ),
                      hintText: "Password",
                      hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                ),
                SizedBox(
                  height: ScreenUtil.getInstance().setHeight(20),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: <Widget>[
                  FlatButton( 
                      
                    
                    child: Text("Sign Up",
                      style: TextStyle( 
                        color: Colors.green,
                      ),  
                    ),
                    splashColor: Colors.amber,
                    onPressed:(){
                      Navigator.of(context).pushNamed("/" + signup);
                    },
                ),
                FlatButton(  
                      
                    
                    child: Text("Forgot Password?",
                      style: TextStyle( 
                        color: Colors.blue,
                        fontFamily: "Poppins-Medium",
                        
                      ),  
                    ),
                    onPressed:(){
                      Navigator.of(context).pushNamed("/"+ forgetpw);
                    },
                ),
               
              ],
            )
          ],
        ),
      ),
    );
  }
}
