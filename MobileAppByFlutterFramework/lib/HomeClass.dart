import 'dart:io';
import 'package:flutter/material.dart';
import 'package:project1/main.dart';
import 'ClassListView.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';




class HomeClass extends StatelessWidget{  
  static const String routeName = "/homeclass";
  final ClassListView classListView =  ClassListView();

  @override 
  Widget build(BuildContext context){ 

    return Scaffold( 
      appBar: AppBar(
        centerTitle: true,
        title: Text('HomeClass',
                    style: TextStyle(  
                      color: Colors.amber
                    ),  
                  ),
        backgroundColor: Colors.amber[numColor2],
        iconTheme: new IconThemeData(color: Colors.grey),
        actions: <Widget>[
        IconButton(icon: Icon(Icons.input), 
                    onPressed: () {
                      showDialog(
                        
                        context: context,
                        builder: (_) => new AlertDialog(
                            title: new Text("Enter Code "),
                            content: Container(
                                height: 180.0, 
                                child: Column(  
                                  children: <Widget>[ 
                                    TextField(
                                            obscureText: true,
                                            decoration: InputDecoration(
                                              border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0))
                                                                                  , borderSide: BorderSide(color: Colors.amber,),
                                              ),
                                              focusedBorder:OutlineInputBorder(
                                                                      borderSide: const BorderSide(color: Colors.teal, width: 2.0),
                                                                      borderRadius: BorderRadius.circular(20.0),
                                                        ),
                                                hintText: "Code",
                                                hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)),
                                      ),
                                      SizedBox(height: ScreenUtil.getInstance().setHeight(30),),
                                      RaisedButton( 
                                        onPressed: (){ Navigator.pop(context); },
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
                                            '  Enter  ',
                                            style: TextStyle(fontSize: 20,color: Colors.white),    
                                          ),
                                          )
                                      ),
                                      SizedBox(height: ScreenUtil.getInstance().setHeight(15),),
                                     // Text("OR"),
                                  ],
                                ),
                              ),                   
                            )
                          );
                        }
          ),       
        ],       
      ),
      //Drawer
      drawer: Drawer(  
        child: Container(
          color: Colors.amber[numColor1],
          child: ListView(
                    // Important: Remove any padding from the ListView.
            padding: EdgeInsets.zero,
            children: <Widget>[
              DrawerHeader(
                
                child: Container( 
                  width: 30.0,
                  height: 30.0,
                  child: Padding( 
                    padding: EdgeInsets.all(30),
                    child: CircleAvatar( 
                      backgroundImage: NetworkImage(''),
                      minRadius: 90,
                      maxRadius: 180,
                      child: Text("My Profile",
                                 style: TextStyle( 
                                   color: Colors.white70
                                 ), 
                      ),
                      backgroundColor: Colors.greenAccent,
                    ), 
                  ),
                ),
                decoration: BoxDecoration(  
                  color: Colors.amber[numColor2],
                ),             
              ),
             
              ListTile(
                leading: Icon(Icons.info),
                title: Text('Student Information'),
                
                onTap: () {
                  Navigator.pop(context);
                },
              ),
              ListTile(
                leading: Icon(Icons.exit_to_app),
                title: Text('LogOut'),
                onTap: () {  Navigator.of(context).pushNamed("/" + login); },
              ),
              ListTile( 
                leading: Icon(Icons.close),
                title: Text("close"),
                onTap: (){exit(0);},
              )
              ],
            ),
          ),
        ),


      // Grid Layout
      body: classListView.build(context),
      backgroundColor: Colors.amber[numColor1], 
     
    );
  }
}