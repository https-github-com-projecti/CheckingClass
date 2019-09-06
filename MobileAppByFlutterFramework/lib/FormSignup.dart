import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';


class FormSignUp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    
        return new Container(
         
          
          width: double.infinity,
          height: ScreenUtil.getInstance().setHeight(500),
          decoration: BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.circular(9.0),
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
          child: SingleChildScrollView(
                      child: Padding(
              padding: EdgeInsets.only(left: 16.0, right: 16.0, top: 10.0),
              child: Column(
                
                  crossAxisAlignment: CrossAxisAlignment.start,
                    children: <Widget>[         
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("Name",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "Name",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(26),
                      ),
                      Text("Lastname",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "Lastname",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("Student ID",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "EX : B59XXXXX",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("Branch",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "สาขาวิชา : Computer Engeneering",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("E-mail",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "EX : B5907953@g.sut.ac.th",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("Username",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        decoration: InputDecoration(
                            hintText: "Username",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
                      Text("New Password",
                          style: TextStyle(
                              fontFamily: "Poppins-Medium",
                              fontSize: ScreenUtil.getInstance().setSp(26))),
                      TextField(
                        obscureText: true,
                        decoration: InputDecoration(
                            hintText: "Password",
                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                      ),
                      SizedBox(
                        height: ScreenUtil.getInstance().setHeight(20),
                      ),
            ],
        ),
              ),

          ),
    );
  }
}
