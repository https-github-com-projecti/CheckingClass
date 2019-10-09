import 'dart:convert';
import 'dart:io';
import 'package:flutter/material.dart';
import 'package:project1/Page/UserInFor.dart';
import 'package:project1/data/API.dart';
import 'package:project1/data/JoinClass.dart';
import 'package:project1/data/Subjects.dart';
import 'package:project1/data/User.dart';
import 'package:project1/main.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'ClassListView.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';


TextEditingController codeJoinClassControler = new TextEditingController();

class HomeClass extends StatelessWidget{  
  static const String routeName = "/homeclass";
  final ClassListView classListView =  ClassListView();
  final User user ;
  
 
  HomeClass({Key key, @required this.user, }) : super(key: key);
 


  @override 
  Widget build(BuildContext context){ 
  // print("HomeClass xIndex==> " + xUser);
  
   
  
    // final EdgeInsets padding = MediaQuery.of(context) as EdgeInsets;
    return Scaffold( 
      

      appBar: PreferredSize(
        preferredSize: Size.fromHeight(80.0),
        child: Padding(
          padding: EdgeInsets.only(top: 20.0),
          child: AppBar(
          
          // centerTitle: true,
          title: Text('HomeClass',
                      style: TextStyle(  
                        color: Colors.amber[numColor]
                      ),  
                    ),
          elevation: 0,
          backgroundColor: Colors.amber[numColor1],
          iconTheme: new IconThemeData(color: Colors.black),
          actions: <Widget>[
          IconButton(icon: Icon(Icons.input), 
                      onPressed: () {
                        showDialog(
                       
                          context: context,
                          builder: (_) => new AlertDialog(
                              shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
                              title: new Text("Enter Code "),
                              content: Container(
                                  height: 180.0, 
                                  child: Column(  
                                    children: <Widget>[ 
                                      TextField(
                                        controller: codeJoinClassControler,
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
                                          onPressed: () async { 
                                            
                                            JoinClass newjoinClass = new JoinClass(codeJoinClass: codeJoinClassControler.text, idstudent: user.idstudent,firstname: user.firstname,lastname: user.lastname);
                                            JoinClass joinClass = await joinClasss(joinClassURL,body:newjoinClass.toMap());
                                            if (joinClass != null){
                                              print("--------------NULL-------------");
                                            }
                                            Navigator.pop(context);
                                          },
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
                                              '  Join Class  ',
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
      ),
      ),
      //Drawer*****
      drawer: Drawer(  
        child: Container(
          color: Colors.amber[numColor1],
          child: ListView(
            padding: EdgeInsets.zero,
            children: <Widget>[
              DrawerHeader(
                
                child: Container( 
                  
                  child: Padding( 
                    padding: EdgeInsets.all(50.0),
                    child: CircleAvatar(
                      backgroundImage: NetworkImage("https://avatars0.githubusercontent.com/u/12619420?s=460&v=4"),
                      foregroundColor: Colors.black,
                                radius: 30.0,
                     
                    ),
                    // child: CircleAvatar( 
                    //   backgroundImage: NetworkImage(''),
                    //   minRadius: 90,
                    //   maxRadius: 180,
                    
                    // ),
                      // child: Text(this.user.firstname,
                      //            style: TextStyle( 
                      //              color: Colors.white70
                      //            ), 
                      // ),
                      
                      // backgroundColor: Colors.greenAccent,
                    // ), 
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
                  Navigator.push(context,MaterialPageRoute(builder: (context) => UserInFor(user: user,)));
                },
              ),
              ListTile(
                leading: Icon(Icons.exit_to_app),
                title: Text('LogOut'),
                onTap: () async {  
                  SharedPreferences prefs = await SharedPreferences.getInstance();
                  prefs.remove('User');
                  print(prefs.getInt('User'));
                  xUser= null;
                  loginState=null;
                  Navigator.of(context).pushNamed("/" + login);
                  users[0]=null;
                   }
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
      body: classListView.build(context,this.user),
      backgroundColor: Colors.amber[numColor1], 
     
    );
  }

    getIndex() async {
   
     API.getSubject(user.idstudent).then((response) {   
        Iterable list = json.decode(response.body);
        subjectName = list.map((model) => Subjects.fromJson(model)).toList();
      });
  }
  void initState() {
    getIndex();
  }
}