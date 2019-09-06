import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:project1/main.dart';



class ForgetPW extends StatelessWidget{
  static const String routeName = "/forgetpw";
  @override
  Widget build(BuildContext context){

    return Scaffold( 
      appBar: AppBar( 
        centerTitle: true,
        title: Text("Password Reset"),
        backgroundColor: Colors.amber[numColor],
      ),

      body: Container( 
        child: Container( 
          child: SingleChildScrollView( 
            child: Column(  
              children: <Widget>[  
                Padding( 
                  padding: EdgeInsets.only(left: 16.0,right: 16.0,top: 50.0),
                  child: Column( 
                  crossAxisAlignment: CrossAxisAlignment.start,
                    children: <Widget>[ 

                      Align(  
                        alignment: Alignment.topLeft,
                        child: Container(  
                          child: Text("   Find your E-mail account", 
                                      style: TextStyle( 
                                        fontSize: 15,
                                        fontFamily: "Poppins-Bold",
                                        fontWeight: FontWeight.bold,
                                        
                                      ),                          
                            ),
                        ),
                        
                      ),
                      SizedBox(height: ScreenUtil.getInstance().setHeight(30),),
                      TextField(
                          decoration: InputDecoration(
                            border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0))
                                                        , borderSide: BorderSide(color: Colors.amber,),
                            ),
                             focusedBorder:OutlineInputBorder(
                                            borderSide: const BorderSide(color: Colors.teal, width: 2.0),
                                            borderRadius: BorderRadius.circular(25.0),
                              ),
                             hintText: 'Enter your email',
                             fillColor: Colors.amber,
                          ),
                      ),

                    ],
                  ),
                ),

                Padding( 
                  padding: EdgeInsets.only(left: 16.0,right: 16.0,top: 16.0),
                  child: Column( 
                    children: <Widget>[ 

                      RaisedButton( 
                        onPressed: (){  },
                        shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(20.0)),
                        padding: const EdgeInsets.all(0.0),
                        child: Container(
                          
                          decoration: const BoxDecoration(
                            gradient: LinearGradient(
                              colors: <Color>[
                                Colors.lightBlueAccent,
                                Colors.tealAccent,
                                // Colors.lightGreenAccent,
                              ],
                            ),
                            borderRadius: BorderRadius.all(Radius.circular(20.0)),              
                          ),
                          padding: const EdgeInsets.all(10.0),
                          child: const Text(
                            'Search',
                            style: TextStyle(fontSize: 20,color: Colors.white),
                          ),
                        )             
                      ),
                      

                    ],
                  ),
                )

              ],
            )
          ),
        ),
      ),
    );
  }


}