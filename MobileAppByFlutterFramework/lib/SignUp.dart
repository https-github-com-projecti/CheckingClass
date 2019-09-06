import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'FormSignUp.dart';
import 'main.dart';



class SignUp extends StatelessWidget {
  static const String routeName = "/signup";
  @override 
  Widget build(BuildContext context){
     ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
     ScreenUtil.instance = ScreenUtil(width: 500, height: 800, allowFontScaling: true);

    
      // SystemChrome.setPreferredOrientations([
      //   DeviceOrientation.landscapeRight,
      //   DeviceOrientation.landscapeLeft,
      //   DeviceOrientation.portraitUp,
      //   DeviceOrientation.portraitDown,
      // ]);
     
    

    return Scaffold( 

       appBar: AppBar(
        centerTitle: true ,
        title: Text("Sign Up"),
        backgroundColor: Colors.amber[700],
      ),
      backgroundColor: Colors.amber[numColor],
      body: Container(
        child: Center(
          child: SingleChildScrollView(
                child: Padding(
                  padding: EdgeInsets.only(left: 28.0, right: 28.0, top: 0.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: <Widget>[
                      FormSignUp(),
                      SizedBox(height: ScreenUtil.getInstance().setHeight(25),),
                      InkWell(
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
                              onTap: () { Navigator.of(context).pushNamed("/" + homeclass); },
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

