import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:crypto/crypto.dart';
import 'package:project1/data/CreateUser.dart';
// import 'package:project1/data/CreateUserLogin.dart';
import '../main.dart';


  TextEditingController firstNameControler = new TextEditingController();
  TextEditingController lastNameControler = new TextEditingController();
  TextEditingController idStudentControler = new TextEditingController();
  TextEditingController phoneControler = new TextEditingController();
  TextEditingController emailControler = new TextEditingController();
  TextEditingController usernameControler = new TextEditingController();
  TextEditingController passwordControler = new TextEditingController();
 String logintextstatusSigup = "Sigupping...";
 String logintextstatusNOpass = "Wrong ro Please fill this out !!!";



class SignUp extends StatefulWidget {
  static const String routeName = "/signup";
  //  static final postUserURL = url+'studentuser';
  
  //  static final postUserLoginURL = url+"userlogin"; 

  @override
  _SignUpState createState() => _SignUpState();
}

class _SignUpState extends State<SignUp> {
final GlobalKey<ScaffoldState> _scaffoldKey = new GlobalKey<ScaffoldState>();
  
  void loading(String status){
    _scaffoldKey.currentState.showSnackBar(
      new SnackBar(duration: new Duration(seconds: 3), content:
      new Row(
        children: <Widget>[
          new CircularProgressIndicator(),
          new Text(status)
        ],
      ),
    ));
  } 
  void wrongbar(String status){
    _scaffoldKey.currentState.showSnackBar(
      new SnackBar(duration: new Duration(seconds: 1), content:
      new Row(
        children: <Widget>[
          new Icon(Icons.close),
          new Text(status,
            style: TextStyle(
             color: Colors.red,)
          )
        ],
      ),
    ));
  }

  @override 
  Widget build(BuildContext context){
     ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
     ScreenUtil.instance = ScreenUtil(width: 500, height: 800, allowFontScaling: true);
  


    return Scaffold(
       key: _scaffoldKey,                                                    //loading
       appBar: AppBar(
        // centerTitle: true ,
        title: Text("Sign Up"),
        backgroundColor: Colors.amber[700],
      ),
      backgroundColor: Colors.amber[numColor1],
      body: Container(
        child: Center(
          child: SingleChildScrollView(
                child: Padding(
                  padding: EdgeInsets.only(left: 28.0, right: 28.0, top: 0.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: <Widget>[
                      // FormSignUp(),
                        SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Row(
                        children: <Widget>[
                          Text("Name",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        // textInputAction: TextInputAction.next,
                        controller: firstNameControler,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)),borderSide: BorderSide(color: Colors.amber,),),
                          focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                            hintText: "Name",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      // textform("name", "name"),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(26),
                      ),
                      Row(
                        children: <Widget>[
                          Text("Lastname",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        controller: lastNameControler,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)),borderSide: BorderSide(color: Colors.amber,),),
                          focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                            hintText: "Lastname",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Row(
                        children: <Widget>[
                          Text("Student ID",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        controller: idStudentControler,
                        maxLength: 8,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)),borderSide: BorderSide(color: Colors.amber,),),
                          focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                            hintText: "EX : B59XXXXX",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Row(
                        children: <Widget>[
                          Text("phone",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        controller: phoneControler,
                        maxLength: 10,
                        keyboardType: TextInputType.phone,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)),borderSide: BorderSide(color: Colors.amber,),),
                          focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                          hintText: "Ex: 09-XXXX-XXXX",
                          hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Row(
                        children: <Widget>[
                          Text("E-mail",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        controller: emailControler,
                        keyboardType: TextInputType.emailAddress,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)),borderSide: BorderSide(color: Colors.amber,),),
                          focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                          hintText: "EX : B5907953@g.sut.ac.th",
                          hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                          
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      
                      Row(
                        children: <Widget>[
                          Icon(Icons.portrait),
                          Text("| Username",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                        ],
                      ),
                      TextField(
                        controller: usernameControler,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)), borderSide: BorderSide(color: Colors.amber,),),
                            hintText: "Username",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Row(
                        children: <Widget>[
                          Icon(Icons.vpn_key),
                          Text("| New Password  ",
                              style: TextStyle(
                                  fontFamily: "Poppins-Medium",
                                  fontSize: ScreenUtil.getInstance().setSp(26))),
                          
                        ],
                      ),
                      TextField(
                        controller: passwordControler,
                        obscureText: true,
                        decoration: InputDecoration(
                          border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)), borderSide: BorderSide(color: Colors.amber,),),
                            hintText: "Password",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      SizedBox(height: ScreenUtil.getInstance().setHeight(25),),
                      InkWell(
                        child: Padding(
                          padding: const EdgeInsets.only(bottom: 50),
                          child: Container(
                            //margin: const EdgeInsets.fromLTRB(80.0, 0.0, .0, 0.0),
                            width: ScreenUtil.getInstance().setWidth(200),
                            height: ScreenUtil.getInstance().setHeight(80),
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
                                onTap: () async { 
                                  var bytes  = utf8.encode(passwordControler.text);
                                  var digest = sha256.convert(bytes);
                                  if(firstNameControler.text.toString() != "" && lastNameControler.text != "" && idStudentControler.text != "" && phoneControler.text != ""
                                      && usernameControler.text != "" && passwordControler.text != ""){
                                    
                                    print(firstNameControler.text.toString());
                                    print(lastNameControler.text);
                                    print(digest.toString());
                                    if(sigupState == null ){
                                      print(sigupState);
                                     loading(logintextstatusSigup);
                                      CreateUser newCreateUser = new CreateUser( /*id: 0,*/ firstname: firstNameControler.text, lastname: lastNameControler.text,
                                                              idstudent: idStudentControler.text, phone: phoneControler.text, username: usernameControler.text, password: digest.toString(),
                                                              email: emailControler.text, /*subjects : createSubjects*/);
                                      CreateUser p = await createUsers(postUserURL,body: newCreateUser.toMap());
                                      sigupState = null;
                                      if (p != null){
                                        print("--------------NULL-------------");
                                      }
                                      Navigator.of(context).pushNamed("/" + login);
                                      firstNameControler.clear();
                                      lastNameControler.clear();
                                      idStudentControler.clear();
                                      phoneControler.clear();
                                      usernameControler.clear();
                                      passwordControler.clear();
                                      emailControler.clear();

                                      
                                    }else{
                                     wrongbar(logintextstatusNOpass);
                                    }
                                  }else{
                                    wrongbar(logintextstatusNOpass);
                                  }
                                },
                                child: Padding(
                                  padding: const EdgeInsets.only(bottom: 0.0),
                                  child: Center(
                                    child: Text("SIGN UP",
                                        style: TextStyle(
                                            color: Colors.white,
                                            fontFamily: "Poppins-Bold",
                                            fontSize: 18,
                                            letterSpacing: 1.0)),
                                  ),
                                ),
                              ),
                            ),
                          ),
                        ),
                      )
                    ],
                  ),                 
                )         
        
              ),
            ),
          ),
    );
  }
}



// class FormSignUp extends StatelessWidget {

//   @override
//   Widget build(BuildContext context) {
 
//         return new Container(
         
          
//           width: double.infinity,
//           height: ScreenUtil.getInstance().setHeight(500),
//           decoration: BoxDecoration(
//               color: Colors.white,
//               borderRadius: BorderRadius.circular(9.0),
//               boxShadow: [
//                 BoxShadow(
//                     color: Colors.black12,
//                     offset: Offset(0.0, 15.0),
//                     blurRadius: 15.0),
//                 BoxShadow(
//                     color: Colors.black12,
//                     offset: Offset(0.0, -10.0),
//                     blurRadius: 15.0),
//               ]),
//           child: SingleChildScrollView(
//                       child: Padding(
//               padding: EdgeInsets.only(left: 16.0, right: 16.0, top: 10.0),
//               child: Column(
                
//                   crossAxisAlignment: CrossAxisAlignment.start,
//                     children: <Widget>[         
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//                       // Text("Name",
//                       //     style: TextStyle(
//                       //         fontFamily: "Poppins-Medium",
//                       //         fontSize: ScreenUtil.getInstance().setSp(26))),
//                       // TextField(
//                       //   // textInputAction: TextInputAction.next,
//                       //   controller: firstNameControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "Name",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // // textform("name", "name"),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(26),
//                       // ),
//                       // Text("Lastname",
//                       //     style: TextStyle(
//                       //         fontFamily: "Poppins-Medium",
//                       //         fontSize: ScreenUtil.getInstance().setSp(26))),
//                       // TextField(
//                       //   controller: lastNameControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "Lastname",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//                       // Text("Student ID",
//                       //     style: TextStyle(
//                       //         fontFamily: "Poppins-Medium",
//                       //         fontSize: ScreenUtil.getInstance().setSp(26))),
//                       // TextField(
//                       //   controller: idStudentControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "EX : B59XXXXX",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//                       // Text("phone",
//                       //     style: TextStyle(
//                       //         fontFamily: "Poppins-Medium",
//                       //         fontSize: ScreenUtil.getInstance().setSp(26))),
//                       // TextField(
//                       //   controller: phoneControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "Ex: 09-XXXX-XXXX",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//                       // Text("E-mail",
//                       //     style: TextStyle(
//                       //         fontFamily: "Poppins-Medium",
//                       //         fontSize: ScreenUtil.getInstance().setSp(26))),
//                       // TextField(
//                       //   controller: emailControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "EX : B5907953@g.sut.ac.th",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       //       keyboardType: TextInputType.emailAddress,
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
                      
//                       // Row(
//                       //   children: <Widget>[
//                       //     Icon(Icons.portrait),
//                       //     Text("| Username",
//                       //         style: TextStyle(
//                       //             fontFamily: "Poppins-Medium",
//                       //             fontSize: ScreenUtil.getInstance().setSp(26))),
//                       //   ],
//                       // ),
//                       // TextField(
//                       //   controller: usernameControler,
//                       //   decoration: InputDecoration(
//                       //       hintText: "Username",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//                       // Row(
//                       //   children: <Widget>[
//                       //     Icon(Icons.vpn_key),
//                       //     Text("| New Password  ",
//                       //         style: TextStyle(
//                       //             fontFamily: "Poppins-Medium",
//                       //             fontSize: ScreenUtil.getInstance().setSp(26))),
                          
//                       //   ],
//                       // ),
//                       // TextField(
//                       //   controller: passwordControler,
//                       //   obscureText: true,
//                       //   decoration: InputDecoration(
//                       //       hintText: "Password",
//                       //       hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
//                       // ),
//                       // SizedBox(
//                       //   height: ScreenUtil.getInstance().setHeight(20),
//                       // ),
//             ],
//         ),
//               ),

//           ),
//     );
//   }
// }