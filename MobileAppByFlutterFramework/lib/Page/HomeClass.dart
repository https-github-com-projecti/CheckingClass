import 'dart:async';
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

class HomeClass extends StatefulWidget{  
  static const String routeName = "/homeclass";
  final User user ;
  HomeClass({Key key, @required this.user, }) : super(key: key);
  @override
  _HomeClassState createState() => _HomeClassState();
}

class _HomeClassState extends State<HomeClass> {
  final ClassListView classListView =  ClassListView();
  ScrollController controller;
  Timer timer;
  final GlobalKey<RefreshIndicatorState> _refreshIndicatorKey = new GlobalKey<RefreshIndicatorState>();

  @override
  void initState() {
    super.initState();
    // getDataSubjectAllTime();
    // timer = Timer.periodic(Duration(seconds: 15), (Timer t) => getDataSubjectAllTime());
    WidgetsBinding.instance
        .addPostFrameCallback((_) => _refreshIndicatorKey.currentState.show());
    
  }


  @override
  void dispose() {
    timer?.cancel();

    super.dispose();
  }

  Future<void> getDataSubjectAllTime() async {
     API.getSubject(widget.user.idstudent).then((response) {   
        Iterable list = json.decode(response.body);
        setState(() {
          subjectName = list.map((model) => Subjects.fromJson(model)).toList();
        });
        
      });
  }


  @override 
  Widget build(BuildContext context){ 

  
    // final EdgeInsets padding = MediaQuery.of(context) as EdgeInsets;
    return Scaffold(   
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(50.0),
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.only(top: 0.0),
            child: AppBar(
            
            // centerTitle: true,
            title: Text('HomeClass',
                        style: TextStyle(  
                          color: Colors.amber[numColor]
                        ),  
                      ),
            elevation: 1,
            backgroundColor: Colors.amber[numColor],
            iconTheme: new IconThemeData(color: Colors.black),
            actions: <Widget>[
            IconButton(icon: Icon(Icons.add), 
                        onPressed: () {
                          showDialog(
                         
                            context: context,
                            builder: (context) => new AlertDialog(
                                shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
                                title: new Text("Enter Code "),
                                content: Container(
                                    height: 150.0, 
                                    child: Column(  
                                      children: <Widget>[ 
                                        TextField(
                                          controller: codeJoinClassControler,
                                          obscureText: true,
                                          decoration: InputDecoration(
                                            border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)), 
                                            borderSide: BorderSide(color: Colors.amber,),),
                                            focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                                            hintText: "Code",
                                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)
                                            
                                          ),
                                        ),
                                        SizedBox(height: ScreenUtil.getInstance().setHeight(30),),
                                        RaisedButton( 
                                          onPressed: () async { 
                                            
                                            JoinClass newjoinClass = new JoinClass(codeJoinClass: codeJoinClassControler.text, idstudent: widget.user.idstudent,firstname: widget.user.firstname,lastname: widget.user.lastname);
                                            JoinClass joinClass = await joinClasss(joinClassURL,body:newjoinClass.toMap());
                                            if (joinClass != null){
                                              print("--------------NULL-------------");
                                            }
                                            // Navigator.pop(context);
                                            Navigator.pushReplacement(context,MaterialPageRoute(builder: (_) => HomeClass(user: widget.user)));
                                             codeJoinClassControler.clear();
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
                    padding: EdgeInsets.only(top: 30,right: 80.0,left: 80.0),
                    child: CircleAvatar(
                      backgroundImage: NetworkImage("https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/89566ebc-ae66-4173-858c-bdae4ee22eb2/dap48o3-b6df2aed-517b-4feb-96b5-33f0d2677c00.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzg5NTY2ZWJjLWFlNjYtNDE3My04NThjLWJkYWU0ZWUyMmViMlwvZGFwNDhvMy1iNmRmMmFlZC01MTdiLTRmZWItOTZiNS0zM2YwZDI2NzdjMDAucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.5-IneqmHJLinXfaDtuwCdTtA6MNeBzwr-LlYQAx-V9I"),
                      foregroundColor: Colors.black,
                                radius: 30.0,   
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
                  Navigator.push(context,MaterialPageRoute(builder: (context) => UserInFor(user: widget.user,)));
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
      body: RefreshIndicator(
          onRefresh: getDataSubjectAllTime,
          backgroundColor: Colors.amber[50],
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: classListView.build(context,this.widget.user,),
  
        ),
      ),
      backgroundColor: Colors.amber[numColor2], 
     
    );
  }

   


}